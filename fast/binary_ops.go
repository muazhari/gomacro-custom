// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * binary_ops.go
 *
 *  Created on Apr 12, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) Add(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) + y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) + y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) + y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) + y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) + y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) + y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) + y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) + y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) + y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) + y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) + y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) + y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) + y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) + y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) + y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) string {
					return x(env) + y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) || y == "" {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) + y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) + y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) + y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) + y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) + y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) + y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) + y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) + y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) + y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) + y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) + y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) + y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) + y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) + y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) + y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := r.ValueOf(y).String()
				fun = func(env *Env) string {
					return x(env) + y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) || x == "" {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x + y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x + y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x + y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x + y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x + y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x + y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x + y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x + y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x + y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x + y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x + y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x + y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x + y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x + y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x + y(env)
				}

			}
		case r.String:

			{
				x := r.ValueOf(x).String()

				y := y.(func(*Env) string)
				fun = func(env *Env) string {
					return x + y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Sub(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) - y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) - y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) - y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) - y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) - y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) - y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) - y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) - y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) - y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) - y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) - y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) - y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) - y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) - y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) - y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) - y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) - y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) - y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) - y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) - y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) - y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) - y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) - y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) - y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) - y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) - y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) - y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) - y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) - y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) - y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x - y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x - y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x - y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x - y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x - y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x - y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x - y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x - y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x - y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x - y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x - y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x - y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x - y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x - y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x - y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Mul(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) * y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) * y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) * y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) * y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) * y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) * y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) * y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) * y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) * y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) * y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) * y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) * y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) * y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) * y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) * y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 1) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) * y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) * y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) * y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) * y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) * y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) * y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) * y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) * y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) * y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) * y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) * y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) * y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) * y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) * y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) * y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 1) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x * y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x * y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x * y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x * y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x * y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x * y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x * y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x * y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x * y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x * y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x * y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x * y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x * y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x * y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x * y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Quo(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) / y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) / y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) / y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) / y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) / y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) / y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) / y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) / y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) / y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) / y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) / y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) / y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) / y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) / y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) / y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			c.Errorf("division by zero")
			return nil
		}
		if isLiteralNumber(y, 1) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) / y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) / y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) / y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) / y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) / y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) / y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) / y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) / y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) / y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) / y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) / y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) / y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) / y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) / y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) / y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x / y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x / y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x / y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x / y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x / y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x / y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x / y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x / y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x / y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x / y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x / y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x / y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x / y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x / y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x / y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Rem(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) % y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) % y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) % y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) % y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) % y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) % y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) % y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) % y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) % y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) % y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) % y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, 0) {
			c.Errorf("division by zero")
			return nil
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) % y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) % y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) % y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) % y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) % y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) % y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) % y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) % y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) % y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) % y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) % y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x % y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x % y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x % y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x % y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x % y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x % y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x % y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x % y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x % y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x % y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x % y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) And(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) & y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) & y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) & y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) & y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) & y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) & y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) & y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) & y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) & y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) & y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) & y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, -1) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) & y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) & y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) & y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) & y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) & y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) & y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) & y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) & y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) & y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) & y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) & y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		if isLiteralNumber(x, -1) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x & y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x & y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x & y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x & y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x & y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x & y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x & y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x & y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x & y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x & y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x & y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Or(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) | y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) | y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) | y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) | y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) | y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) | y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) | y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) | y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) | y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) | y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) | y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) | y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) | y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) | y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) | y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) | y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) | y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) | y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) | y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) | y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) | y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) | y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		if isLiteralNumber(x, 0) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x | y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x | y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x | y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x | y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x | y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x | y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x | y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x | y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x | y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x | y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x | y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Xor(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) ^ y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) ^ y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) ^ y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) ^ y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) ^ y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) ^ y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) ^ y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) ^ y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) ^ y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) ^ y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) ^ y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) ^ y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) ^ y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) ^ y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) ^ y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) ^ y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) ^ y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) ^ y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x ^ y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x ^ y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x ^ y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x ^ y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x ^ y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x ^ y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x ^ y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x ^ y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x ^ y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x ^ y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x ^ y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Andnot(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) &^ y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) &^ y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) &^ y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) &^ y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) &^ y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) &^ y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) &^ y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) &^ y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) &^ y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) &^ y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) &^ y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) &^ y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) &^ y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) &^ y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) &^ y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) &^ y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) &^ y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) &^ y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x &^ y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x &^ y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x &^ y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x &^ y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x &^ y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x &^ y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x &^ y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x &^ y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x &^ y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x &^ y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x &^ y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
