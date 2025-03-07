/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * method.go
 *
 *  Created on Mar 28, 2018
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
	"go/ast"
	r "reflect"

	"github.com/muazhari/gomacro-custom/go/etoken"

	"github.com/muazhari/gomacro-custom/go/types"
	"github.com/muazhari/gomacro-custom/go/typeutil"
)

// return detailed string representation of a method signature, including its receiver if present
func (m Method) String() string {
	if m.GoFun != nil {
		return typeutil.String2(m.Name, m.GoFun.Type())
	} else {
		return m.Name
	}
}

// For interfaces, NumMethod returns *total* number of methods for interface t,
// including wrapper methods for embedded interfaces.
// For all other named types, NumMethod returns the number of explicitly declared methods,
// ignoring wrapper methods for embedded fields.
// Returns 0 for other unnamed types.
func (t *xtype) NumMethod() int {
	num := 0
	if gt, ok := t.gunderlying().(*types.Interface); ok {
		num = gt.NumMethods()
	} else {
		// generics v2 add methods to most types
		num = t.gtype.NumMethods()
	}
	return num
}

// NumExplicitMethod returns the number of explicitly declared methods of named type or interface t.
// Wrapper methods for embedded fields or embedded interfaces are not counted.
func (t *xtype) NumExplicitMethod() int {
	num := 0
	if gt, ok := t.gunderlying().(*types.Interface); ok {
		num = gt.NumExplicitMethods()
	} else {
		// generics v2 add methods to most types
		num = t.gtype.NumMethods()
	}
	return num
}

// NumAllMethod returns the *total* number of methods for interface or named type t,
// including wrapper methods for embedded fields or embedded interfaces.
// Note: it has slightly different semantics from go/types.(*Named).NumMethods(),
//
//	since the latter returns 0 for named interfaces, and callers need to manually invoke
//	goNamedType.Underlying().NumMethods() to retrieve the number of methods
//	of a named interface
func (t *xtype) NumAllMethod() int {
	return goTypeNumAllMethod(t.gtype, make(map[types.Type]struct{}))
}

// recursively count total number of methods for type t,
// including wrapper methods for embedded fields or embedded interfaces
func goTypeNumAllMethod(gt types.Type, visited map[types.Type]struct{}) int {
	count := 0
	for {
		if _, ok := visited[gt]; ok {
			break
		}
		visited[gt] = struct{}{}
		switch t := gt.(type) {
		case *types.Interface:
			count += t.NumMethods()
		case *types.Struct:
			n := t.NumFields()
			for i := 0; i < n; i++ {
				if f := t.Field(i); f.Anonymous() {
					count += goTypeNumAllMethod(f.Type(), visited)
				}
			}
		default:
			// generics v2 add methods to most types
			count += t.NumMethods()
			u := t.Underlying()
			if u != gt {
				gt = u
				continue
			}
		}
		break
	}
	return count
}

// For interfaces, Method returns the i-th method, including methods from embedded interfaces.
// For all other named types, Method returns the i-th explicitly declared method, ignoring wrapper methods for embedded fields.
// It panics if i is outside the range 0 .. NumMethod()-1
func (t *xtype) Method(i int) Method {
	checkMethod(t, i)
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.method(i)
}

func checkMethod(t *xtype, i int) {
	if t.kind == r.Ptr {
		xerrorf(t, "Method of %s type %v. Invoke Method() on type's Elem() instead", t.kind, t)
	}
	if !etoken.GENERICS.V2_CTI() && !t.Named() && t.kind != r.Interface {
		xerrorf(t, "Method of type %v that cannot have methods", t.kind, t)
	}
}

func (t *xtype) method(i int) Method {
	checkMethod(t, i)
	gfunc := t.gmethod(i)
	name := gfunc.Name()
	resizemethodvalues(t, t.NumAllMethod())

	rtype := t.rtype
	var rfunctype r.Type
	rfunc := t.methodvalue[i]
	// debugf("xtype.method(%d): t = %v,\tt.gmethod(%d) = %v,\tt.methodvalue[%d] = %v // %v",
	//	i, t, i, gfunc, i, rfunc, rValueType(rfunc))
	if rfunc.Kind() == r.Func {
		// easy, method is cached already
		// fmt.Printf("DEBUG xtype.method(%d): t = %v,\tt.methodvalue[%d] = %v // %v\n", i, t, i, rfunc, rValueType(rfunc))
		rfunctype = rfunc.Type()
	} else if _, ok := t.gunderlying().(*types.Interface); ok {
		if rtype.Kind() == r.Ptr && isReflectInterfaceStruct(rtype.Elem()) {
			// rtype is our emulated interface type,
			// i.e. a pointer to a struct containing: InterfaceHeader, [0]struct { embeddeds }, methods (without receiver)
			rfield := rtype.Elem().Field(i + 2)
			rfunctype = rAddReceiver(rtype, rfield.Type)
		} else if rtype.Kind() != r.Interface {
			xerrorf(t, "inconsistent interface type <%v>: expecting interface reflect.Type, found <%v>", t, rtype)
		} else if ast.IsExported(name) {
			// rtype is an interface type, and reflect only returns exported methods
			// rtype.MethodByName returns a Method with the following caveats
			// 1) Type == method signature, without a receiver
			// 2) Func == nil.
			rmethod, _ := rtype.MethodByName(name)
			if rmethod.Type == nil {
				xerrorf(t, "interface type <%v>: reflect method %q not found", t, name)
			} else if rmethod.Index != i {
				xerrorf(t, "inconsistent interface type <%v>: method %q has go/types.Func index=%d but reflect.Method index=%d",
					t, name, i, rmethod.Index)
			}
			rfunctype = rAddReceiver(rtype, rmethod.Type)
		}
	} else {
		rmethod, _ := rtype.MethodByName(gfunc.Name())
		rfunc = rmethod.Func
		if rfunc.Kind() != r.Func {
			if rtype.Kind() != r.Ptr {
				// also search in the method set of pointer-to-t
				rmethod, _ = r.PtrTo(rtype).MethodByName(gfunc.Name())
				rfunc = rmethod.Func
			}
		}
		if rfunc.Kind() == r.Func {
			rfunctype = rmethod.Type
			t.methodvalue[i] = rfunc
		}
		// debugf("xtype.method(%d): t = %v,\trmethod(%q) = %v", i, t, gfunc.Name(), rmethod)

		// rfunc and rfunctype will be invalid when bootstrapping Universe
		// and when adding CTI methods to a type
	}
	return t.makemethod(i, gfunc, &t.methodvalue, rfunctype) // lock already held
}

func rValueType(v r.Value) r.Type {
	if v.IsValid() {
		return v.Type()
	}
	return nil
}

// insert recv as the the first parameter of rtype function type
func rAddReceiver(recv r.Type, rtype r.Type) r.Type {
	nin := rtype.NumIn()
	rin := make([]r.Type, nin+1)
	rin[0] = recv
	for i := 0; i < nin; i++ {
		rin[i+1] = rtype.In(i)
	}
	nout := rtype.NumOut()
	rout := make([]r.Type, nout)
	for i := 0; i < nout; i++ {
		rout[i] = rtype.Out(i)
	}
	return r.FuncOf(rin, rout, rtype.IsVariadic())
}

// remove the first parameter of rtype function type
func rRemoveReceiver(rtype r.Type) r.Type {
	nin := rtype.NumIn()
	if nin == 0 {
		return rtype
	}
	rin := make([]r.Type, nin-1)
	for i := 1; i < nin; i++ {
		rin[i-1] = rtype.In(i)
	}
	nout := rtype.NumOut()
	rout := make([]r.Type, nout)
	for i := 0; i < nout; i++ {
		rout[i] = rtype.Out(i)
	}
	return r.FuncOf(rin, rout, nin > 1 && rtype.IsVariadic())
}

// remove the first parameter of t function type
func removeReceiver(t Type) Type {
	nin := t.NumIn()
	if nin == 0 {
		return t
	}
	tin := make([]Type, nin-1)
	for i := 1; i < nin; i++ {
		tin[i-1] = t.In(i)
	}
	nout := t.NumOut()
	tout := make([]Type, nout)
	for i := 0; i < nout; i++ {
		tout[i] = t.Out(i)
	}
	return t.Universe().FuncOf(tin, tout, nin > 1 && t.IsVariadic())
}

// assumes xreflect.Type methods and go/types.Type methods are in the same order!
func (t *xtype) gmethod(i int) *types.Func {
	var gfun *types.Func
	if gtype, ok := t.gunderlying().(*types.Interface); ok {
		gfun = gtype.Method(i)
	} else {
		gfun = t.gtype.Method(i)
	}
	return gfun
}

func (t *xtype) makemethod(index int, gfun *types.Func, rfuns *[]r.Value, rfunctype r.Type) Method {
	// sanity checks
	name := gfun.Name()
	gsig := gfun.Type().Underlying().(*types.Signature)
	if rfunctype != nil {
		nparams := 0
		if gsig.Params() != nil {
			nparams = gsig.Params().Len()
		}
		if gsig.Recv() != nil {
			if nparams+1 != rfunctype.NumIn() {
				xerrorf(t, `type <%v>: inconsistent %d-th method %q signature:
	go/types.Type has receiver <%v> and %d parameters: %v
	reflect.Type has %d parameters: %v`, t, index, name, gsig.Recv(), nparams, gsig, rfunctype.NumIn(), rfunctype)
			}
		} else if nparams != rfunctype.NumIn() {
			xerrorf(t, `type <%v>: inconsistent %d-th method %q signature:
	go/types.Type has no receiver and %d parameters: %v
	reflect.Type has %d parameters: %v`, t, index, name, nparams, gsig, rfunctype.NumIn(), rfunctype)
		}
	}
	var tmethod Type
	if rfunctype != nil {
		rsig := ReflectUnderlying(rfunctype)
		tmethod = t.universe.maketype(gsig, rsig, OptDefault) // lock already held
	}
	return Method{
		Name:  name,
		Pkg:   (*Package)(gfun.Pkg()),
		Type:  tmethod,
		Funs:  rfuns,
		Index: index,
		GoFun: gfun,
	}
}

func resizemethodvalues(t *xtype, n int) {
	if cap(t.methodvalue) < n {
		slice := make([]r.Value, n, n+n/2+4)
		copy(t.methodvalue, slice)
		t.methodvalue = slice
	} else if len(t.methodvalue) < n {
		t.methodvalue = t.methodvalue[0:n]
	}
}

// return one of the methods defined by interface tinterf but missing from t
func MissingMethod(t, tinterf Type) *Method {
	n := tinterf.NumMethod()
	var mtdinterf Method
	if t == nil && n > 0 {
		mtdinterf = tinterf.Method(0)
		return &mtdinterf
	}
	xt := unwrap(t)
	xtinterf := unwrap(tinterf)
	for i := 0; i < n; i++ {
		mtdinterf = tinterf.Method(i)
		mtd, count := t.MethodByName(mtdinterf.Name, mtdinterf.Pkg.Name())
		if count == 1 {
			tfunc := mtd.Type
			if t.Kind() != r.Interface {
				tfunc = removeReceiver(tfunc)
			}
			fmt.Printf("MissingMethod: comparing %v against expected interface method %v\n", tfunc, mtdinterf)
			if mtdinterf.Type.IdenticalTo(tfunc) && matchReceiverType(xt, xtinterf) {
				continue
			}
		}
		return &mtdinterf
	}
	return nil
}
