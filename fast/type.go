/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * type.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/muazhari/gomacro-custom/base"
	"github.com/muazhari/gomacro-custom/base/reflect"
	"github.com/muazhari/gomacro-custom/base/strings"
	"github.com/muazhari/gomacro-custom/base/untyped"
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

// DeclType compiles a type declaration.
func (c *Comp) DeclType(spec ast.Spec) {
	node, ok := spec.(*ast.TypeSpec)
	if !ok {
		c.Errorf("unexpected type declaration, expecting *ast.TypeSpec, found: %v // %T", spec, spec)
	}
	if GENERICS_V1_CXX() || GENERICS_V2_CTI() {
		if lit, _ := node.Type.(*ast.CompositeLit); lit != nil {
			c.DeclGenericType(node)
			return
		}
	}
	name := node.Name.Name
	// support type aliases
	if node.Assign != token.NoPos {
		t := c.Type(node.Type)
		c.DeclTypeAlias(name, t)
		return
	}
	// support self-referencing types, as for example: type List struct { First int; Rest *List }
	oldt := c.Types[name]
	panicking := true
	defer func() {
		// On compile error, restore pre-existing declaration
		if !panicking || c.Types == nil {
			// nothing to do
		} else if oldt != nil {
			c.Types[name] = oldt
		} else {
			delete(c.Types, name)
		}
	}()
	t := c.DeclNamedType(name)
	u := c.Type(node.Type)
	if t != nil { // t == nil means name == "_", discard the result of type declaration
		c.SetUnderlyingType(t, u)
	}
	panicking = false
}

// DeclTypeAlias compiles a typealias declaration, i.e. type Foo = /*...*/
// Returns the second argument.
func (c *Comp) DeclTypeAlias(name string, t xr.Type) xr.Type {
	if name == "_" {
		return t
	}
	if et := c.Types[name]; et != nil {
		// forward-declared types have kind == r.Invalid, see Comp.DeclNamedType() below
		if et.Kind() != r.Invalid {
			c.Warnf("redefined type alias: %v", name)
		}
		c.Universe.InvalidateCache()
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	c.Types[name] = t
	return t
}

// DeclTypeAlias0 declares a type alias
// in Go, types are computed only at compile time - no need for a runtime *Env
func (c *Comp) declTypeAlias(alias string, t xr.Type) xr.Type {
	if alias == "" || alias == "_" {
		// never define bindings for "_"
		return t
	}
	if _, ok := c.Types[alias]; ok {
		c.Warnf("redefined type: %v", alias)
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	c.Types[alias] = t
	return t
}

// DeclNamedType executes a named type forward declaration.
// Returns nil if name == "_"
// Otherwise it must be followed by Comp.SetUnderlyingType(t) where t is the returned type
func (c *Comp) DeclNamedType(name string) xr.Type {
	if name == "_" {
		return nil
	}
	if t := c.Types[name]; t != nil {
		if t.Kind() != r.Invalid {
			c.Warnf("redefined type: %v", name)
		}
		if xr.QName1(t) != xr.QName2(name, c.FileComp().Path) {
			// the current type "name" is an alias, discard it
			c.Universe.InvalidateCache()
		} else {
			// reuse t, change only its underlying type
			return t
		}
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	t := c.Universe.NamedOf(name, c.FileComp().Path)
	c.Types[name] = t
	return t
}

func (c *Comp) SetUnderlyingType(t, underlying xr.Type) {
	t.SetUnderlying(underlying)
}

// DeclType0 declares a type
// in Go, types are computed only at compile time - no need for a runtime *Env
func (c *Comp) DeclType0(t xr.Type) xr.Type {
	if t == nil {
		return nil
	}
	return c.declTypeAlias(t.Name(), t)
}

// Type compiles a type expression.
func (c *Comp) Type(node ast.Expr) xr.Type {
	t, _ := c.compileType2(node, false)
	return t
}

// compileTypeOrNilR compiles a type expression. as a special case used by type switch, compiles *ast.Ident{Name:"nil"} to nil
func (c *Comp) compileTypeOrNilR(node ast.Expr) xr.Type {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ident:
			if expr.Name == "nil" {
				sym := c.TryResolve(expr.Name)
				if sym != nil && sym.Type == nil {
					return nil
				}
			}
		}
		break
	}
	t, _ := c.compileType2(node, false)
	return t
}

// compileType2 compiles a type expression.
// if allowEllipsis is true, it supports the special case &ast.Ellipsis{/*expression*/}
// that represents ellipsis in the last argument of a function declaration.
// The second return value is true both in the case above, and for array types whose length is [...]
func (c *Comp) compileType2(node ast.Expr, allowEllipsis bool) (t xr.Type, ellipsis bool) {
	stars := 0
	for {
		switch expr := node.(type) {
		case *ast.StarExpr:
			stars++
			node = expr.X
			continue
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ellipsis:
			if allowEllipsis {
				node = expr.Elt
				ellipsis = true
				continue
			}
		}
		break
	}
	if node != nil {
		c.Pos = node.Pos()
	}
	universe := c.Universe
	var ellipsisArray bool

	switch node := node.(type) {
	case *ast.ArrayType: // also for slices
		t, ellipsisArray = c.TypeArray(node)
	case *ast.ChanType:
		telem := c.Type(node.Value)
		dir := r.BothDir
		if node.Dir == ast.SEND {
			dir = r.SendDir
		} else if node.Dir == ast.RECV {
			dir = r.RecvDir
		}
		t = universe.ChanOf(dir, telem)
	case *ast.FuncType:
		t, _, _ = c.TypeFunction(node)
	case *ast.Ident:
		t = c.ResolveType(node.Name)
	case *ast.IndexExpr:
		if GENERICS_V1_CXX() || GENERICS_V2_CTI() {
			t = c.GenericType(node)
		} else {
			c.Errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
		}
	case *ast.InterfaceType:
		t = c.TypeInterface(node)
	case *ast.MapType:
		kt := c.Type(node.Key)
		vt := c.Type(node.Value)
		t = universe.MapOf(kt, vt)
	case *ast.SelectorExpr:
		ident, ok := node.X.(*ast.Ident)
		if !ok {
			c.Errorf("invalid qualified type, expecting packagename.identifier, found: %v <%v>", node, r.TypeOf(node))
		}
		// this could be Package.Type, or other non-type expressions: Type.Method, Value.Method, Struct.Field...
		// check for Package.Type
		name := ident.Name
		var bind *Bind
		for o := c; o != nil; o = o.Outer {
			if bind = o.Binds[name]; bind != nil {
				break
			}
		}
		if bind == nil {
			c.Errorf("undefined %q in %v <%v>", name, node, r.TypeOf(node))
		} else if !bind.Const() || bind.Type.ReflectType() != rtypeOfPtrImport {
			c.Errorf("not a package: %q in %v <%v>", name, node, r.TypeOf(node))
		}
		imp, ok := bind.Value.(*Import)
		if !ok {
			c.Errorf("not a package: %q in %v <%v>", name, node, r.TypeOf(node))
		}
		name = node.Sel.Name
		t, ok = imp.Types[name]
		if !ok || t == nil {
			c.Errorf("not a type: %v <%v>", node, r.TypeOf(node))
		}
		if !ast.IsExported(name) {
			c.Errorf("cannot refer to unexported name %v", node)
		}
	case *ast.StructType:
		// c.Debugf("evalType() struct declaration: %v <%v>", node, r.TypeOf(node))
		types, names := c.TypeFields(node.Fields)
		tags := c.fieldsTags(node.Fields)
		// c.Debugf("evalType() struct names = %v types = %v tags = %v", names, types, tags)
		pkg := universe.LoadPackage(c.FileComp().Path)
		fields := c.makeStructFields(pkg, names, types, tags)
		// c.Debugf("compileType2() declaring struct type. fields=%#v", fields)
		t = universe.StructOf(fields)
	case nil:
		// type can be omitted in many case - then we must perform type inference
		break
	default:
		// which types are still missing?
		c.Errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
	}
	if t != nil {
		for i := 0; i < stars; i++ {
			t = universe.PtrTo(t)
		}
		if allowEllipsis && ellipsis {
			// ellipsis in the last argument of a function declaration
			t = universe.SliceOf(t)
		}
	}
	return t, ellipsis || ellipsisArray
}

func (c *Comp) TypeArray(node *ast.ArrayType) (t xr.Type, ellipsis bool) {
	universe := c.Universe
	t = c.Type(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		t = universe.ArrayOf(0, t)
		ellipsis = true
	case nil:
		t = universe.SliceOf(t)
	default:
		// as stated by https://golang.org/ref/spec#Array_types
		// "The length is part of the array's type; it must evaluate to a non-negative constant
		// representable by a value of type int. "
		var count int
		init := c.expr(n, nil)
		if !init.Const() {
			c.Errorf("array length is not a constant: %v", node)
			return
		} else if init.Untyped() {
			count = init.ConstTo(c.TypeOfInt()).(int)
		} else {
			count = untyped.ConvertLiteralCheckOverflow(init.Value, c.TypeOfInt()).(int)
		}
		if count < 0 {
			c.Errorf("array length [%v] is negative: %v", count, node)
		}
		t = universe.ArrayOf(count, t)
	}
	return t, ellipsis
}

func (c *Comp) TypeFunction(node *ast.FuncType) (t xr.Type, paramNames []string, resultNames []string) {
	return c.TypeFunctionOrMethod(nil, node)
}

// TypeFunctionOrMethod compiles a function type corresponding to given receiver and function declaration
// If receiver is not null, the returned tFunc will have it as receiver.
func (c *Comp) TypeFunctionOrMethod(recv *ast.Field, node *ast.FuncType) (t xr.Type, paramNames []string, resultNames []string) {
	paramTypes, paramNames, variadic := c.typeFieldOrParamList(node.Params, true)
	resultTypes, resultNames := c.TypeFields(node.Results)

	var recvType xr.Type
	if recv != nil {
		// methods are functions with receiver. xreflect allows functions to be treated as methods
		// (using the first parameter as receiver), but go/types.Type loaded by go/importer.Default()
		// will have methods as functions with receivers.
		//
		// So be uniform with those.
		//
		// Alas, go/types.Type.String() does *not* print the receiver, making it cumbersome to debug.
		recvTypes, recvNames, _ := c.typeFieldsOrParams([]*ast.Field{recv}, false)
		recvType = recvTypes[0]

		// anyway, return the receiver *name* as first element of paramNames
		paramNames = append(recvNames, paramNames...)
	}
	t = c.Universe.MethodOf(recvType, paramTypes, resultTypes, variadic)
	return t, paramNames, resultNames
}

func (c *Comp) TypeFields(fields *ast.FieldList) (types []xr.Type, names []string) {
	types, names, _ = c.typeFieldOrParamList(fields, false)
	return types, names
}

func (c *Comp) typeFieldOrParamList(fields *ast.FieldList, allowEllipsis bool) (types []xr.Type, names []string, ellipsis bool) {
	var list []*ast.Field
	if fields != nil {
		list = fields.List
	}
	return c.typeFieldsOrParams(list, allowEllipsis)
}

func (c *Comp) typeFieldsOrParams(list []*ast.Field, allowEllipsis bool) (types []xr.Type, names []string, ellipsis bool) {
	types = make([]xr.Type, 0)
	names = ZeroStrings
	n := len(list)
	if n == 0 {
		return types, names, ellipsis
	}
	var t xr.Type
	for i, f := range list {
		t, ellipsis = c.compileType2(f.Type, i == n-1)
		if len(f.Names) == 0 {
			types = append(types, t)
			names = append(names, "")
			// c.Debugf("evalTypeFields() %v -> %v", f.Type, t)
		} else {
			for _, ident := range f.Names {
				types = append(types, t)
				names = append(names, ident.Name)
				// Debugf("evalTypeFields() %v %v -> %v", ident.Name, f.Type, t)
			}
		}
	}
	return types, names, ellipsis
}

func (c *Comp) TryResolveType(name string) xr.Type {
	var t xr.Type
	for ; c != nil; c = c.Outer {
		if t = c.Types[name]; t != nil {
			break
		}
	}
	return t
}

func (c *Comp) ResolveType(name string) xr.Type {
	t := c.TryResolveType(name)
	if t == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	return t
}

func (c *Comp) makeStructFields(pkg *xr.Package, names []string, types []xr.Type, tags []string) []xr.StructField {
	// pkgIdentifier := sanitizeIdentifier(pkgPath)
	fields := make([]xr.StructField, len(names))
	for i, name := range names {
		fields[i] = xr.StructField{
			Name:      name,
			Pkg:       pkg,
			Type:      types[i],
			Tag:       r.StructTag(tags[i]),
			Anonymous: len(name) == 0,
		}
	}
	return fields
}

func (c *Comp) fieldsTags(fields *ast.FieldList) []string {
	var tags []string
	if fields != nil {
		for _, field := range fields.List {
			var tag string
			if lit := field.Tag; lit != nil && lit.Kind == token.STRING {
				tag = strings.MaybeUnescapeString(lit.Value)
			}
			if len(field.Names) == 0 {
				tags = append(tags, tag)
			} else {
				for range field.Names {
					tags = append(tags, tag)
				}
			}
		}
	}
	return tags
}

func rtypeof(v xr.Value, t xr.Type) r.Type {
	if t != nil {
		return t.ReflectType()
	}
	return reflect.ValueType(v)
}

// TypeAssert2 compiles a multi-valued type assertion
func (c *Comp) TypeAssert2(node *ast.TypeAssertExpr) *Expr {
	val := c.Expr1(node.X, nil)
	tin := val.Type
	tout := c.Type(node.Type)
	rtout := tout.ReflectType()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	kout := tout.Kind()
	if kout != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	// extractor to unwrap value from proxy or emulated interface
	extractor := c.extractor(tin)

	fun := val.Fun.(func(*Env) xr.Value) // val returns an interface... must be already wrapped in a reflect.Value

	var ret func(env *Env) (xr.Value, []xr.Value)

	fail := []xr.Value{xr.Zero(tout), False} // returned by type assertion in case of failure
	switch {
	case reflect.IsOptimizedKind(kout):
		ret = func(env *Env) (xr.Value, []xr.Value) {
			v, t := extractor(fun(env))
			if reflect.ValueType(v) != rtout || (t != nil && !t.AssignableTo(tout)) {
				return fail[0], fail
			}
			return v, []xr.Value{v, True}
		}

	case kout == r.Interface:
		if tout.NumMethod() == 0 {
			// type assertion to empty interface.
			// everything, excluding untyped nil, implements an empty interface
			ret = func(env *Env) (xr.Value, []xr.Value) {
				v, _ := extractor(fun(env))
				if !v.IsValid() {
					// type assertion of untyped nil must fail
					return fail[0], fail
				}
				v = convert(v, rtout)
				return v, []xr.Value{v, True}
			}
			break
		}
		if tin.Implements(tout) {
			// type assertion to interface.
			// expression type implements such interface, can only fail if value is untyped nil
			ret = func(env *Env) (xr.Value, []xr.Value) {
				v, _ := extractor(fun(env))
				if !v.IsValid() {
					// type assertion of untyped nil must fail
					return fail[0], fail
				}
				v = convert(v, rtout)
				return v, []xr.Value{v, True}
			}
			break
		}
		// type assertion to interface
		// must check at runtime whether concrete type implements asserted interface
		ret = func(env *Env) (xr.Value, []xr.Value) {
			v, t := extractor(fun(env))
			if !v.IsValid() {
				// type assertion of untyped nil must fail
				return fail[0], fail
			}
			rt := rtypeof(v, t)
			if (rt != rtout && !rt.Implements(rtout)) ||
				(t != nil && !t.IdenticalTo(tout) && !t.Implements(tout)) {
				return fail[0], fail
			}
			v = convert(v, rtout)
			return v, []xr.Value{v, True}
		}

	case reflect.IsNillableKind(kout):
		// type assertion to concrete (nillable) type
		ret = func(env *Env) (xr.Value, []xr.Value) {
			v, t := extractor(fun(env))
			if !v.IsValid() {
				// type assertion of untyped nil must fail
				return fail[0], fail
			}
			rt := rtypeof(v, t)
			if rt != rtout || (t != nil && !t.IdenticalTo(tout)) {
				return fail[0], fail
			}
			return v, []xr.Value{v, True}
		}
	default:
		// type assertion to concrete (non-nillable) type
		ret = func(env *Env) (xr.Value, []xr.Value) {
			v, t := extractor(fun(env))
			if !v.IsValid() {
				// type assertion of untyped nil must fail
				return fail[0], fail
			}
			rt := rtypeof(v, t)
			if rt != rtout || (t != nil && !t.IdenticalTo(tout)) {
				return fail[0], fail
			}
			return v, []xr.Value{v, True}
		}
	}
	e := exprXV([]xr.Type{tout, c.TypeOfBool()}, ret)
	e.EFlags = EIsTypeAssert
	return e
}

// TypeAssert1 compiles a single-valued type assertion
func (c *Comp) TypeAssert1(node *ast.TypeAssertExpr) *Expr {
	if node.Type == nil {
		c.Errorf("invalid type assertion: expecting actual type, found type switch: %v", node)
	}
	val := c.Expr1(node.X, nil)
	tin := val.Type
	tout := c.Type(node.Type)
	kout := tout.Kind()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	if tout.Kind() != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	// extractor to unwrap value from proxy or emulated interface
	extractor := c.extractor(tin)

	fun := val.Fun.(func(*Env) xr.Value) // val returns an interface... must be already wrapped in a reflect.Value

	rtout := tout.ReflectType()
	var ret I
	switch kout {
	case xr.Bool:
		ret = func(env *Env) bool {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return int(v.Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return int8(v.Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return int16(v.Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return int32(v.Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.Int()
		}
	case xr.Uint:
		ret = func(env *Env) uint {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return uint(v.Uint())
		}
	case xr.Uint8:
		ret = func(env *Env) uint8 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return uint8(v.Uint())
		}
	case xr.Uint16:
		ret = func(env *Env) uint16 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return uint16(v.Uint())
		}
	case xr.Uint32:
		ret = func(env *Env) uint32 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return uint32(v.Uint())
		}
	case xr.Uint64:
		ret = func(env *Env) uint64 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.Uint()
		}
	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return uintptr(v.Uint())
		}
	case xr.Float32:
		ret = func(env *Env) float32 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return float32(v.Float())
		}
	case xr.Float64:
		ret = func(env *Env) float64 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.Float()
		}
	case xr.Complex64:
		ret = func(env *Env) complex64 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return complex64(v.Complex())
		}
	case xr.Complex128:
		ret = func(env *Env) complex128 {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.Complex()
		}
	case xr.String:
		ret = func(env *Env) string {
			v, t := extractor(fun(env))
			v = typeassert(v, t, tin, tout)
			return v.String()
		}
	case xr.Interface:
		if tout.NumMethod() == 0 {
			// type assertion to empty interface.
			// everything, excluding untyped nil, implements an empty interface
			ret = func(env *Env) xr.Value {
				v, _ := extractor(fun(env))
				if !v.IsValid() {
					typeassertpanic(nil, nil, tin, tout)
				}
				return convert(v, rtout)
			}
		} else if tin.Implements(tout) {
			// type assertion to interface.
			// expression type implements such interface, can only fail if value is untyped nil
			ret = func(env *Env) xr.Value {
				v, _ := extractor(fun(env))
				// nil is not a valid tout, check for it.
				// IsNil() can be invoked only on nillable types...
				if reflect.IsNillableKind(v.Kind()) && (!v.IsValid() || v.IsNil()) {
					typeassertpanic(nil, nil, tin, tout)
				}
				return convert(v, rtout)
			}
		} else {
			// type assertion to interface.
			// must check at runtime whether concrete type implements asserted interface
			ret = func(env *Env) xr.Value {
				v, t := extractor(fun(env))
				// nil is not a valid tout, check for it.
				// IsNil() can be invoked only on nillable types...
				if reflect.IsNillableKind(v.Kind()) && (!v.IsValid() || v.IsNil()) {
					typeassertpanic(nil, nil, tin, tout)
				}
				rt := rtypeof(v, t)
				if (rt != rtout && !rt.AssignableTo(rtout) && !rt.Implements(rtout)) ||
					(t != nil && !t.AssignableTo(tout) && !t.Implements(tout)) {
					typeassertpanic(rt, t, tin, tout)
				}
				return convert(v, rtout)
			}
		}
	default:
		if reflect.IsNillableKind(kout) {
			// type assertion to concrete (nillable) type
			ret = func(env *Env) xr.Value {
				v, t := extractor(fun(env))
				if !v.IsValid() {
					// untyped nil cannot be converted to concrete type
					typeassertpanic(nil, nil, tin, tout)
				}
				rt := rtypeof(v, t)
				if rt != rtout || (t != nil && !t.IdenticalTo(tout)) {
					panic(&TypeAssertionError{
						Interface:       tin,
						Concrete:        t,
						ReflectConcrete: rt,
						Asserted:        tout,
					})
				}
				return v
			}
		} else {
			// type assertion to concrete (non-nillable) type
			ret = func(env *Env) xr.Value {
				v, t := extractor(fun(env))
				rt := rtypeof(v, t)
				if rt != rtout || (t != nil && !t.IdenticalTo(tout)) {
					panic(&TypeAssertionError{
						Interface:       tin,
						Concrete:        t,
						ReflectConcrete: rt,
						Asserted:        tout,
					})
				}
				return v
			}
		}
	}
	e := exprFun(tout, ret)
	e.EFlags = EIsTypeAssert
	return e
}

func typeassert(v xr.Value, t xr.Type, tin xr.Type, tout xr.Type) xr.Value {
	rt := rtypeof(v, t)
	if rt != tout.ReflectType() || t != nil && !t.IdenticalTo(tout) {
		panic(&TypeAssertionError{
			Interface:       tin,
			Concrete:        t,
			ReflectConcrete: rt,
			Asserted:        tout,
		})
	}
	return v
}

func typeassertpanic(rt r.Type, t xr.Type, tin xr.Type, tout xr.Type) {
	var missingmethod *xr.Method
	if t != nil && tout.Kind() == r.Interface {
		missingmethod = xr.MissingMethod(t, tout)
	}
	panic(&TypeAssertionError{
		Interface:       tin,
		Concrete:        t,
		ReflectConcrete: rt,
		Asserted:        tout,
		MissingMethod:   missingmethod,
	})
}

func (g *CompGlobals) TypeOfBool() xr.Type {
	return g.Universe.BasicTypes[r.Bool]
}

func (g *CompGlobals) TypeOfInt() xr.Type {
	return g.Universe.BasicTypes[r.Int]
}

func (g *CompGlobals) TypeOfInt8() xr.Type {
	return g.Universe.BasicTypes[r.Int8]
}

func (g *CompGlobals) TypeOfInt16() xr.Type {
	return g.Universe.BasicTypes[r.Int16]
}

func (g *CompGlobals) TypeOfInt32() xr.Type {
	return g.Universe.BasicTypes[r.Int32]
}

func (g *CompGlobals) TypeOfInt64() xr.Type {
	return g.Universe.BasicTypes[r.Int64]
}

func (g *CompGlobals) TypeOfUint() xr.Type {
	return g.Universe.BasicTypes[r.Uint]
}

func (g *CompGlobals) TypeOfUint8() xr.Type {
	return g.Universe.BasicTypes[r.Uint8]
}

func (g *CompGlobals) TypeOfUint16() xr.Type {
	return g.Universe.BasicTypes[r.Uint16]
}

func (g *CompGlobals) TypeOfUint32() xr.Type {
	return g.Universe.BasicTypes[r.Uint32]
}

func (g *CompGlobals) TypeOfUint64() xr.Type {
	return g.Universe.BasicTypes[r.Uint64]
}

func (g *CompGlobals) TypeOfUintptr() xr.Type {
	return g.Universe.BasicTypes[r.Uintptr]
}

func (g *CompGlobals) TypeOfFloat32() xr.Type {
	return g.Universe.BasicTypes[r.Float32]
}

func (g *CompGlobals) TypeOfFloat64() xr.Type {
	return g.Universe.BasicTypes[r.Float64]
}

func (g *CompGlobals) TypeOfComplex64() xr.Type {
	return g.Universe.BasicTypes[r.Complex64]
}

func (g *CompGlobals) TypeOfComplex128() xr.Type {
	return g.Universe.BasicTypes[r.Complex128]
}

func (g *CompGlobals) TypeOfString() xr.Type {
	return g.Universe.BasicTypes[r.String]
}

func (g *CompGlobals) TypeOfError() xr.Type {
	return g.Universe.TypeOfError
}

func (g *CompGlobals) TypeOfInterface() xr.Type {
	return g.Universe.TypeOfInterface
}

const (
	MaxInt = base.MaxInt
)

var (
	nilInterface = xr.ZeroR(base.TypeOfInterface)

	rtypeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
	rtypeOfForward   = r.TypeOf((*xr.Forward)(nil)).Elem()

	rtypeOfBuiltin        = r.TypeOf(Builtin{})
	rtypeOfFunction       = r.TypeOf(Function{})
	rtypeOfMacro          = r.TypeOf(Macro{})
	rtypeOfPtrImport      = r.TypeOf((*Import)(nil))
	rtypeOfPtrGenericFunc = r.TypeOf((*GenericFunc)(nil))
	rtypeOfPtrGenericType = r.TypeOf((*GenericType)(nil))
	rtypeOfReflectType    = r.TypeOf((*r.Type)(nil)).Elem()
	rtypeOfUntypedLit     = r.TypeOf((*UntypedLit)(nil)).Elem()

	zeroOfReflectType = xr.ZeroR(rtypeOfReflectType)

	None  = reflect.None // indicates "no value"
	True  = xr.ValueOf(true)
	False = xr.ValueOf(false)

	ZeroStrings = []string{}
	ZeroValues  = []xr.Value{}
)

func (g *CompGlobals) TypeOfBuiltin() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfBuiltin]
}

func (g *CompGlobals) TypeOfFunction() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfFunction]
}

func (g *CompGlobals) TypeOfMacro() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfMacro]
}

func (g *CompGlobals) TypeOfPtrImport() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfPtrImport]
}

func (g *CompGlobals) TypeOfPtrGenericFunc() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfPtrGenericFunc]
}

func (g *CompGlobals) TypeOfPtrGenericType() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfPtrGenericType]
}

func (g *CompGlobals) TypeOfUntypedLit() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfUntypedLit]
}

// A TypeAssertionError explains a failed type assertion.
type TypeAssertionError struct {
	Interface       xr.Type
	Concrete        xr.Type
	ReflectConcrete r.Type // in case Concrete is not available
	Asserted        xr.Type
	MissingMethod   *xr.Method // one method needed by Interface, missing from Concrete
}

func (*TypeAssertionError) RuntimeError() {}

func (e *TypeAssertionError) Error() string {
	in := e.Interface
	var concr interface{}
	if e.Concrete != nil {
		concr = e.Concrete
	} else if e.ReflectConcrete != nil {
		concr = e.ReflectConcrete
	}
	if concr == nil {
		return fmt.Sprintf("interface conversion: <%v> is nil, not <%v>", in, e.Asserted)
	}
	if e.MissingMethod == nil {
		return fmt.Sprintf("interface conversion: <%v> is <%v>, not <%v>", in, concr, e.Asserted)
	}
	return fmt.Sprintf("interface conversion: <%v> does not implement <%v>: missing method %s", concr, e.Asserted, e.MissingMethod.String())
}
