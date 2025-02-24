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
 * proxy.go
 *
 *  Created on Mar 06, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"fmt"
	"go/types"

	"github.com/muazhari/gomacro-custom/base/paths"
)

type writeTypeOpts int

const (
	writeMethodsAsFields writeTypeOpts = 1 << iota
	writeForceParamNames
	writeIncludeParamTypes
	writeLastParamIsVariadic
)

func lastParamIsVariadic(variadic bool) writeTypeOpts {
	if variadic {
		return writeLastParamIsVariadic
	}
	return 0
}

func (gen *genimport) writeInterfaceProxy(pkgPath string, name string, t *types.Interface) {
	fmt.Fprintf(gen.out, "\n// --------------- proxy for %s.%s ---------------\ntype %s%s struct {", pkgPath, name, gen.proxyprefix, name)
	gen.writeInterfaceMethods(name, t, writeMethodsAsFields)
	gen.out.WriteString("\n}\n")
	gen.writeInterfaceMethods(name, t, writeForceParamNames)
}

func (gen *genimport) writeInterfaceMethods(name string, t *types.Interface, opts writeTypeOpts) {
	if opts&writeMethodsAsFields != 0 {
		fmt.Fprint(gen.out, "\n\tObject\tinterface{}") // will be used to retrieve object wrapped in the proxy
	}
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		gen.writeInterfaceMethod(name, t.Method(i), opts)
	}
}

func (gen *genimport) writeInterfaceMethod(interfaceName string, method *types.Func, opts writeTypeOpts) {
	if !method.Exported() {
		return
	}
	sig, ok := method.Type().(*types.Signature)
	if !ok {
		return
	}
	out := gen.out
	params := sig.Params()
	if opts&writeMethodsAsFields != 0 {
		var param0 string
		if opts&writeForceParamNames != 0 || isNamedTypeTuple(params) {
			param0 = "_proxy_obj_ "
		}
		fmt.Fprintf(out, "\n\t%s_\tfunc(%sinterface{}", method.Name(), param0)
		if params != nil && params.Len() != 0 {
			out.WriteString(", ")
		}
	} else {
		fmt.Fprintf(out, "func (P *%s%s) %s(", gen.proxyprefix, interfaceName, method.Name())
	}
	variadic := lastParamIsVariadic(sig.Variadic())
	gen.writeTypeTuple(params, opts|variadic|writeIncludeParamTypes)
	out.WriteString(") ")
	results := sig.Results()
	gen.writeTypeTupleOut(results)
	if opts&writeMethodsAsFields != 0 {
		return
	}
	out.WriteString(" {\n\t")
	if results != nil && results.Len() > 0 {
		out.WriteString("return ")
	}
	fmt.Fprintf(out, "P.%s_(P.Object", method.Name())
	if params != nil && params.Len() != 0 {
		out.WriteString(", ")
	}
	gen.writeTypeTuple(params, variadic|writeForceParamNames)
	out.WriteString(")\n}\n")
}

func isNamedTypeTuple(tuple *types.Tuple) bool {
	if tuple == nil || tuple.Len() == 0 {
		return false
	}
	for i, n := 0, tuple.Len(); i < n; i++ {
		if len(tuple.At(i).Name()) != 0 {
			return true
		}
	}
	return false
}

func (gen *genimport) writeTypeTupleOut(tuple *types.Tuple) {
	if tuple == nil || tuple.Len() == 0 {
		return
	}
	out := gen.out
	ret0 := tuple.At(0)
	if tuple.Len() > 1 || len(ret0.Name()) > 0 {
		out.WriteString("(")
		gen.writeTypeTuple(tuple, writeIncludeParamTypes)
		out.WriteString(")")
	} else {
		types.WriteType(out, ret0.Type(), gen.packageNameQualifier)
	}
}

func (gen *genimport) writeTypeTuple(tuple *types.Tuple, opts writeTypeOpts) {
	opt := opts &^ writeLastParamIsVariadic
	n := tuple.Len()
	for i := 0; i < n; i++ {
		if i != 0 {
			gen.out.WriteString(", ")
		}
		if i == n-1 && opt != opts {
			opt = opts // restore 'variadic' flag
		}
		gen.writeTypeVar(tuple.At(i), i, opt)
	}
}

func (gen *genimport) writeTypeVar(v *types.Var, index int, opts writeTypeOpts) {
	name := v.Name()
	if len(name) == 0 && opts&writeForceParamNames != 0 {
		name = fmt.Sprintf("unnamed%d", index)
	}
	out := gen.out
	out.WriteString(name)
	if opts&writeIncludeParamTypes != 0 {
		if len(name) != 0 {
			out.WriteString(" ")
		}
		t := v.Type()
		if opts&writeLastParamIsVariadic != 0 {
			out.WriteString("...")
			t = t.(*types.Slice).Elem()
		}
		types.WriteType(out, t, gen.packageNameQualifier)
	} else if len(name) != 0 && opts&writeLastParamIsVariadic != 0 {
		out.WriteString("...")
	}
}

func (gen *genimport) packageNameQualifier(pkg *types.Package) string {
	path := pkg.Path()
	name, ok := gen.pkgrenames[path]
	if !ok {
		name = paths.FileName(path)
	}
	return name
}
