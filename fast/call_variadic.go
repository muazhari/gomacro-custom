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
 * call_variadic.go
 *
 *  Created on May 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

// call a variadic function, when arguments DO NOT contain '...'
func call_variadic_ret0(c *Call, maxdepth int) func(env *Env) {
	exprfun := c.Fun.AsX1()
	argfunsX1 := c.MakeArgfunsX1()
	var call func(*Env)
	switch len(argfunsX1) {
	case 1:
		argfun := argfunsX1[0]
		call = func(env *Env) {
			funv := exprfun(env)
			argv := []xr.Value{
				argfun(env),
			}
			funv.Call(argv)
		}
	case 2:
		argfuns := [2]func(env *Env) xr.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		call = func(env *Env) {
			funv := exprfun(env)
			argv := []xr.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			funv.Call(argv)
		}
	case 3:
		argfuns := [3]func(env *Env) xr.Value{
			argfunsX1[0],
			argfunsX1[1],
			argfunsX1[2],
		}
		call = func(env *Env) {
			funv := exprfun(env)
			argv := []xr.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			funv.Call(argv)
		}
	default:
		call = func(env *Env) {
			funv := exprfun(env)
			argv := make([]xr.Value, len(argfunsX1))
			for i, argfun := range argfunsX1 {
				argv[i] = argfun(env)
			}
			funv.Call(argv)
		}
	}
	return call
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func call_variadic_ret1(c *Call, maxdepth int) I {
	exprfun := c.Fun.AsX1()
	argfunsX1 := c.MakeArgfunsX1()
	kout := c.Fun.Type.Out(0).Kind()
	var call I
	switch len(argfunsX1) {
	case 1:
		argfun := argfunsX1[0]
		switch kout {
		case xr.Bool:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.Bool()
			}
		case xr.Int:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return int(retv.Int())
			}
		case xr.Int8:
			call = func(env *Env) int8 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return int8(retv.Int())
			}
		case xr.Int16:
			call = func(env *Env) int16 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return int16(retv.Int())
			}
		case xr.Int32:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return int32(retv.Int())
			}
		case xr.Int64:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.Int()
			}
		case xr.Uint:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return uint(retv.Uint())
			}
		case xr.Uint8:
			call = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return uint8(retv.Uint())
			}
		case xr.Uint16:
			call = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return uint16(retv.Uint())
			}
		case xr.Uint32:
			call = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return uint32(retv.Uint())
			}
		case xr.Uint64:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.Uint()
			}
		case xr.Uintptr:
			call = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return uintptr(retv.Uint())
			}
		case xr.Float32:
			call = func(env *Env) float32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return float32(retv.Float())
			}
		case xr.Float64:
			call = func(env *Env) float64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.Float()
			}
		case xr.Complex64:
			call = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return complex64(retv.Complex())
			}
		case xr.Complex128:
			call = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.Complex()
			}
		case xr.String:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				retv := funv.Call(argv)[0]
				return retv.String()
			}
		default:
			call = func(env *Env) xr.Value {
				funv := exprfun(env)
				argv := []xr.Value{
					argfun(env),
				}
				return funv.Call(argv)[0]
			}
		}
	case 2:
		argfuns := [2]func(*Env) xr.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		switch kout {
		case xr.Bool:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.Bool()
			}
		case xr.Int:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return int(retv.Int())
			}
		case xr.Int8:
			call = func(env *Env) int8 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return int8(retv.Int())
			}
		case xr.Int16:
			call = func(env *Env) int16 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return int16(retv.Int())
			}
		case xr.Int32:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return int32(retv.Int())
			}
		case xr.Int64:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.Int()
			}
		case xr.Uint:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return uint(retv.Uint())
			}
		case xr.Uint8:
			call = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return uint8(retv.Uint())
			}
		case xr.Uint16:
			call = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return uint16(retv.Uint())
			}
		case xr.Uint32:
			call = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return uint32(retv.Uint())
			}
		case xr.Uint64:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.Uint()
			}
		case xr.Uintptr:
			call = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return uintptr(retv.Uint())
			}
		case xr.Float32:
			call = func(env *Env) float32 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return float32(retv.Float())
			}
		case xr.Float64:
			call = func(env *Env) float64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.Float()
			}
		case xr.Complex64:
			call = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return complex64(retv.Complex())
			}
		case xr.Complex128:
			call = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.Complex()
			}
		case xr.String:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				retv := funv.Call(argv)[0]
				return retv.String()
			}
		default:
			call = func(env *Env) xr.Value {
				funv := exprfun(env)
				argv := []xr.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				return funv.Call(argv)[0]
			}
		}
	default:
		switch kout {
		case xr.Bool:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.Bool()
			}
		case xr.Int:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return int(retv.Int())
			}
		case xr.Int8:
			call = func(env *Env) int8 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return int8(retv.Int())
			}
		case xr.Int16:
			call = func(env *Env) int16 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return int16(retv.Int())
			}
		case xr.Int32:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return int32(retv.Int())
			}
		case xr.Int64:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.Int()
			}
		case xr.Uint:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return uint(retv.Uint())
			}
		case xr.Uint8:
			call = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return uint8(retv.Uint())
			}
		case xr.Uint16:
			call = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return uint16(retv.Uint())
			}
		case xr.Uint32:
			call = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return uint32(retv.Uint())
			}
		case xr.Uint64:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.Uint()
			}
		case xr.Uintptr:
			call = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return uintptr(retv.Uint())
			}
		case xr.Float32:
			call = func(env *Env) float32 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return float32(retv.Float())
			}
		case xr.Float64:
			call = func(env *Env) float64 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.Float()
			}
		case xr.Complex64:
			call = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return complex64(retv.Complex())
			}
		case xr.Complex128:
			call = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.Complex()
			}
		case xr.String:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				retv := funv.Call(argv)[0]
				return retv.String()
			}
		default:
			call = func(env *Env) xr.Value {
				funv := exprfun(env)
				argv := make([]xr.Value, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					argv[i] = argfun(env)
				}
				return funv.Call(argv)[0]
			}
		}

	}
	return call
}
