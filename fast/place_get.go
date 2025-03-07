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
 * place_get.go
 *
 *  Created on Apr 25, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

// compile a read operation on a place
func (c *Comp) GetPlace(place *Place) *Expr {
	if place.IsVar() {
		return c.Symbol(place.Var.AsSymbol())
	}
	lhs := place.Fun
	mapkey := place.MapKey
	var ret I
	if mapkey == nil {
		switch place.Type.Kind() {
		case xr.Bool:
			ret = func(env *Env) bool {
				return lhs(env).Bool()
			}
		case xr.Int:
			ret = func(env *Env) int {
				return int(lhs(env).Int())
			}
		case xr.Int8:
			ret = func(env *Env) int8 {
				return int8(lhs(env).Int())
			}
		case xr.Int16:
			ret = func(env *Env) int16 {
				return int16(lhs(env).Int())
			}
		case xr.Int32:
			ret = func(env *Env) int32 {
				return int32(lhs(env).Int())
			}
		case xr.Int64:
			ret = func(env *Env) int64 {
				return lhs(env).Int()
			}
		case xr.Uint:
			ret = func(env *Env) uint {
				return uint(lhs(env).Uint())
			}
		case xr.Uint8:
			ret = func(env *Env) uint8 {
				return uint8(lhs(env).Uint())
			}
		case xr.Uint16:
			ret = func(env *Env) uint16 {
				return uint16(lhs(env).Uint())
			}
		case xr.Uint32:
			ret = func(env *Env) uint32 {
				return uint32(lhs(env).Uint())
			}
		case xr.Uint64:
			ret = func(env *Env) uint64 {
				return lhs(env).Uint()
			}
		case xr.Uintptr:
			ret = func(env *Env) uintptr {
				return uintptr(lhs(env).Uint())
			}
		case xr.Float32:
			ret = func(env *Env) float32 {
				return float32(lhs(env).Float())
			}
		case xr.Float64:
			ret = func(env *Env) float64 {
				return lhs(env).Float()
			}
		case xr.Complex64:
			ret = func(env *Env) complex64 {
				return complex64(lhs(env).Complex())
			}
		case xr.Complex128:
			ret = func(env *Env) complex128 {
				return lhs(env).Complex()
			}
		case xr.String:
			ret = func(env *Env) string {
				return lhs(env).String()
			}
		default:
			ret = lhs
		}
	} else {
		switch place.Type.Kind() {
		case xr.Bool:
			ret = func(env *Env) bool {
				return lhs(env).MapIndex(mapkey(env)).Bool()
			}
		case xr.Int:
			ret = func(env *Env) int {
				return int(lhs(env).MapIndex(mapkey(env)).Int())
			}
		case xr.Int8:
			ret = func(env *Env) int8 {
				return int8(lhs(env).MapIndex(mapkey(env)).Int())
			}
		case xr.Int16:
			ret = func(env *Env) int16 {
				return int16(lhs(env).MapIndex(mapkey(env)).Int())
			}
		case xr.Int32:
			ret = func(env *Env) int32 {
				return int32(lhs(env).MapIndex(mapkey(env)).Int())
			}
		case xr.Int64:
			ret = func(env *Env) int64 {
				return lhs(env).MapIndex(mapkey(env)).Int()
			}
		case xr.Uint:
			ret = func(env *Env) uint {
				return uint(lhs(env).MapIndex(mapkey(env)).Uint())
			}
		case xr.Uint8:
			ret = func(env *Env) uint8 {
				return uint8(lhs(env).MapIndex(mapkey(env)).Uint())
			}
		case xr.Uint16:
			ret = func(env *Env) uint16 {
				return uint16(lhs(env).MapIndex(mapkey(env)).Uint())
			}
		case xr.Uint32:
			ret = func(env *Env) uint32 {
				return uint32(lhs(env).MapIndex(mapkey(env)).Uint())
			}
		case xr.Uint64:
			ret = func(env *Env) uint64 {
				return lhs(env).MapIndex(mapkey(env)).Uint()
			}
		case xr.Uintptr:
			ret = func(env *Env) uintptr {
				return uintptr(lhs(env).MapIndex(mapkey(env)).Uint())
			}
		case xr.Float32:
			ret = func(env *Env) float32 {
				return float32(lhs(env).MapIndex(mapkey(env)).Float())
			}
		case xr.Float64:
			ret = func(env *Env) float64 {
				return lhs(env).MapIndex(mapkey(env)).Float()
			}
		case xr.Complex64:
			ret = func(env *Env) complex64 {
				return complex64(lhs(env).MapIndex(mapkey(env)).Complex())
			}
		case xr.Complex128:
			ret = func(env *Env) complex128 {
				return lhs(env).MapIndex(mapkey(env)).Complex()
			}
		case xr.String:
			ret = func(env *Env) string {
				return lhs(env).MapIndex(mapkey(env)).String()
			}
		default:
			ret = func(env *Env) xr.Value {
				return lhs(env).MapIndex(mapkey(env))
			}
		}
	}
	return exprFun(place.Type, ret)
}
