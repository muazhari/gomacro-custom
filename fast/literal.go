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
 * literal.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	"go/token"
	"math/big"
	r "reflect"

	"github.com/muazhari/gomacro-custom/base/output"

	"github.com/muazhari/gomacro-custom/base/reflect"
	"github.com/muazhari/gomacro-custom/base/untyped"
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

func (c *Comp) BasicLit(node *ast.BasicLit) *Expr {
	str := node.Value
	var kind untyped.Kind
	var label string
	switch node.Kind {
	case token.INT:
		kind, label = untyped.Int, "integer"
	case token.FLOAT:
		kind, label = untyped.Float, "float"
	case token.IMAG:
		kind, label = untyped.Complex, "complex"
	case token.CHAR:
		kind, label = untyped.Rune, "rune"
	case token.STRING:
		kind, label = untyped.String, "string"
	default:
		c.Errorf("unsupported basic literal: %v", node)
		return nil
	}
	obj := constant.MakeFromLiteral(str, node.Kind, 0)
	if obj.Kind() == constant.Unknown {
		c.Errorf("invalid %s literal: %v", label, str)
		return nil
	}
	return c.exprUntypedLit(kind, obj)
}

func isLiteral(x interface{}) bool {
	if x == nil {
		return true
	}
	rtype := r.TypeOf(x)
	switch reflect.Category(rtype.Kind()) {
	case xr.Bool, r.Int, r.Uint, r.Float64, r.Complex128, r.String:
		return true
	}
	_, ok := x.(UntypedLit)
	return ok
}

func isLiteralNumber(x I, n int64) bool {
	if x == nil {
		return false
	}
	v := xr.ValueOf(x)
	switch reflect.Category(v.Kind()) {
	case xr.Bool:
		return false
	case xr.Int:
		return v.Int() == n
	case xr.Uint:
		// n == -1 means "unsigned integer equals its maximum value"
		// similarly, n == -2 means "unsigned integer equals its maximum value minus 1"
		// and so on...
		return v.Uint() == uint64(n)
	case xr.Float64:
		return v.Float() == float64(n)
	case xr.Complex128:
		return v.Complex() == complex(float64(n), 0)
	case xr.String:
		return false
	}
	// no luck yet... try harder
	switch x := x.(type) {
	case xr.Value:
		return false
	case UntypedLit:
		return x.EqualInt64(n)
	}
	output.Errorf("isLiteralNumber: unexpected literal type %v <%v>", x, r.TypeOf(x))
	return false
}

// ================================= ConstTo =================================

// ConstTo checks that a constant Expr can be used as the given type.
// panics if not constant, or if Expr is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (e *Expr) ConstTo(t xr.Type) I {
	if !e.Const() {
		output.Errorf("internal error: expression is not constant, use Expr.To() instead of Expr.ConstTo() to convert from <%v> to <%v>", e.Type, t)
	}
	val := e.Lit.ConstTo(t)
	fun := makeMathBigFun(val)
	if fun != nil {
		// no longer a constant
		e.Lit.Value = nil
		e.Fun = fun
	} else {
		if e.Fun != nil {
			// e.Fun is no longer valid, recompute it
			e.WithFun()
		}
	}
	return val
}

// ConstTo checks that a Lit can be used as the given type.
// panics if Lit is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (lit *Lit) ConstTo(t xr.Type) I {
	value := lit.Value
	// output.Debugf("Lit.ConstTo(): converting constant %v <%v> (emulated type <%v>) to <%v>", value, r.TypeOf(value), lit.Type, t)
	if t == nil {
		// only literal nil has type nil
		if value != nil {
			output.Errorf("cannot convert constant %v <%v> to <nil>", value, lit.Type)
		}
		return nil
	}
	// stricter than t == lit.Type
	tfrom := lit.Type
	if tfrom != nil && t != nil && tfrom.IdenticalTo(t) {
		return value
	}
	switch x := value.(type) {
	case UntypedLit:
		val := x.Convert(t)
		lit.Type = t
		lit.Value = val
		// output.Debugf("Lit.ConstTo(): converted untyped constant %v to %v <%v> (emulated type <%v>)", x, val, r.TypeOf(val), t)
		return val
	case nil:
		// literal nil can only be converted to nillable types
		if reflect.IsNillableKind(t.Kind()) {
			lit.Type = t
			return nil
			// lit.Value = xr.Zero(t).Interface()
			// return lit.Value
		}
	}
	if tfrom != nil && t != nil && (tfrom.AssignableTo(t) || t.Kind() == r.Interface && tfrom.Implements(t)) {
		lit.Type = t
		// FIXME: use (*Comp).Converter(), requires a *Comp parameter
		lit.Value = convert(xr.ValueOf(value), t.ReflectType()).Interface()
		return lit.Value
	}
	output.Errorf("cannot convert typed constant %v <%v> to <%v>%s", value, lit.Type, t, interfaceMissingMethod(lit.Type, t))
	return nil
}

// return a closure that duplicates at each invokation any *big.Int, *big.Rat, *big.Float passed as 'val'
func makeMathBigFun(val I) func(*Env) xr.Value {
	switch a := val.(type) {
	case *big.Int:
		return func(*Env) xr.Value {
			var b big.Int
			b.Set(a)
			return xr.ValueOf(&b)
		}
	case *big.Rat:
		return func(*Env) xr.Value {
			var b big.Rat
			b.Set(a)
			return xr.ValueOf(&b)
		}
	case *big.Float:
		return func(*Env) xr.Value {
			var b big.Float
			b.Set(a)
			return xr.ValueOf(&b)
		}
	default:
		return nil
	}
}

// ================================= DefaultType =================================

// DefaultType returns the default type of an expression.
func (e *Expr) DefaultType() xr.Type {
	if e.Untyped() {
		return e.Lit.DefaultType()
	}
	return e.Type
}

// DefaultType returns the default type of a constant.
func (lit *Lit) DefaultType() xr.Type {
	switch x := lit.Value.(type) {
	case UntypedLit:
		return x.DefaultType()
	default:
		return lit.Type
	}
}

// SetTypes sets the expression result types
func (e *Expr) SetTypes(tout []xr.Type) {
	switch len(tout) {
	case 0:
		e.Type = nil
		e.Types = tout
	case 1:
		e.Type = tout[0]
		e.Types = nil
	default:
		e.Type = tout[0]
		e.Types = tout
	}
}

/* used?

// Set sets the expression value to the given (typed or untyped) constant
func (e *Expr) Set(x I) {
	e.Lit.Set(x)
	e.Types = nil
	e.Fun = nil
	e.IsNil = x == nil
}

// Set sets the Lit to the given typed constant
func (lit *Lit) Set(x I) {
	t := TypeOf(x)
	if !isLiteral(x) {
		Errorf("cannot set Lit to non-literal value %v <%v>", x, t)
	}
	lit.Type = t
	lit.Value = x
}
*/

// To checks that an Expr can be used as (i.e. is assignable to) the given type,
// and converts Expr to the given type.
// panics if Expr has an incompatible type.
func (e *Expr) To(c *Comp, t xr.Type) {
	if e.Const() {
		e.ConstTo(t)
		return
	}
	if e.Type.IdenticalTo(t) {
		return
	}
	if !e.Type.AssignableTo(t) {
		c.Errorf("cannot use <%v> as <%v>", e.Type, t)
	}
	k := e.Type.Kind()
	if reflect.IsOptimizedKind(k) {
		if k == t.Kind() {
			// same optimized representation
			e.Type = t
			return
		} else if t.Kind() == r.Interface {
			e.Fun = e.AsX1()
			e.Type = t
			return
		}
		c.Errorf("internal error: cannot use <%v> as <%v> (should not happen, <%v> is assignable to <%v>", e.Type, t, e.Type, t)
	}
	fun := e.AsX1()
	zero := xr.Zero(t)

	if conv := c.Converter(e.Type, t); conv == nil {
		e.Fun = func(env *Env) xr.Value {
			v := fun(env)
			if !v.IsValid() {
				v = zero
			}
			return v
		}
	} else {
		e.Fun = func(env *Env) xr.Value {
			v := fun(env)
			if !v.IsValid() {
				v = zero
			} else {
				v = conv(v)
			}
			return v
		}
	}
	e.Type = t
}

// WithFun ensures that Expr.Fun is a closure that will return the expression result:
//
// if Expr is an untyped constant, WithFun converts the constant to its default type (panics on overflows),
//
//	then sets Expr.Fun to a closure that will return such constant.
//
// if Expr is a typed constant, WithFun sets Expr.Fun to a closure that will return such constant.
// if Expr is not a constant, WithFun does nothing (Expr.Fun must be set already)
func (e *Expr) WithFun() I {
	if !e.Const() {
		return e.Fun
	}
	var fun I
again:
	value := e.Value
	v := xr.ValueOf(value)
	t := e.Type
	if t == nil {
		e.Fun = eNilR
		return eNilR
	}
	if value == nil {
		if !reflect.IsNillableKind(t.Kind()) {
			output.Errorf("internal error: constant of type <%v> cannot be nil", t)
		}
		zero := xr.Zero(t)
		fun = func(*Env) xr.Value {
			return zero
		}
		e.Fun = fun
		return fun
	}
	rtactual := r.TypeOf(value)
	rtexpected := t.ReflectType()
	if rtexpected != rtactual {
		if rtexpected.Kind() == r.Interface && rtactual.Implements(rtexpected) {
			v = convert(v, rtexpected)
		} else {
			output.Errorf("internal error: constant %v <%v> was assumed to have type <%v>", value, rtactual, rtexpected)
		}
	}
	switch v.Kind() {
	case xr.Invalid:
		fun = eNilR
	case xr.Bool:
		if v.Bool() {
			fun = eTrue
		} else {
			fun = eFalse
		}
	case xr.Int:
		x := int(v.Int())
		fun = func(env *Env) int {
			return x
		}
	case xr.Int8:
		x := int8(v.Int())
		fun = func(env *Env) int8 {
			return x
		}
	case xr.Int16:
		x := int16(v.Int())
		fun = func(env *Env) int16 {
			return x
		}
	case xr.Int32:
		x := int32(v.Int())
		fun = func(env *Env) int32 {
			return x
		}
	case xr.Int64:
		x := v.Int()
		fun = func(env *Env) int64 {
			return x
		}
	case xr.Uint:
		x := uint(v.Uint())
		fun = func(env *Env) uint {
			return x
		}
	case xr.Uint8:
		x := uint8(v.Uint())
		fun = func(env *Env) uint8 {
			return x
		}
	case xr.Uint16:
		x := uint16(v.Uint())
		fun = func(env *Env) uint16 {
			return x
		}
	case xr.Uint32:
		x := uint32(v.Uint())
		fun = func(env *Env) uint32 {
			return x
		}
	case xr.Uint64:
		x := v.Uint()
		fun = func(env *Env) uint64 {
			return x
		}
	case xr.Uintptr:
		x := uintptr(v.Uint())
		fun = func(env *Env) uintptr {
			return x
		}
	case xr.Float32:
		x := float32(v.Float())
		fun = func(env *Env) float32 {
			return x
		}
	case xr.Float64:
		x := v.Float()
		fun = func(env *Env) float64 {
			return x
		}
	case xr.Complex64:
		x := complex64(v.Complex())
		fun = func(env *Env) complex64 {
			return x
		}
	case xr.Complex128:
		x := v.Complex()
		fun = func(env *Env) complex128 {
			return x
		}
	case xr.String:
		x := v.String()
		fun = func(env *Env) string {
			return x
		}
	default:
		if t.ReflectType() == rtypeOfUntypedLit {
			e.ConstTo(e.DefaultType())
			goto again
		}
		fun = func(env *Env) xr.Value {
			return v
		}
	}
	e.Fun = fun
	return fun
}
