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
 * literal.go
 *
 *  Created on Apr 25, 2018
 *      Author Massimiliano Ghilardi
 */

package untyped

import (
	"go/constant"
	"go/token"
	"math/big"
	r "reflect"

	"github.com/muazhari/gomacro-custom/base/output"
	"github.com/muazhari/gomacro-custom/base/reflect"
	"github.com/muazhari/gomacro-custom/base/strings"
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

var (
	rtypeOfPtrBigFloat = r.TypeOf((*big.Float)(nil))
	rtypeOfPtrBigInt   = r.TypeOf((*big.Int)(nil))
	rtypeOfPtrBigRat   = r.TypeOf((*big.Rat)(nil))
)

func (untyp *Lit) EqualInt64(i int64) bool {
	val := untyp.Val
	switch val.Kind() {
	case constant.Int:
		m, exact := constant.Int64Val(val)
		return exact && m == i
	case constant.Float:
		m, exact := constant.Float64Val(val)
		return exact && float64(int64(m)) == m && int64(m) == i
	case constant.Complex:
		m, exact := constant.Float64Val(constant.Imag(val))
		if !exact || m != 0.0 {
			return false
		}
		m, exact = constant.Float64Val(constant.Real(val))
		return exact && float64(int64(m)) == m && int64(m) == i
	default:
		return false
	}
}

var constantValZero = constant.MakeInt64(0)

// ================================= Convert =================================

// Convert checks that an untyped.Lit can be converted exactly to the given type.
// performs actual untyped -> typed conversion and subsequent overflow checks.
// returns the constant.Value converted to given type
func (untyp *Lit) Convert(t xr.Type) interface{} {
	val := untyp.Val
	var ret interface{}
again:
	switch t.Kind() {
	case r.Bool:
		if val.Kind() == constant.Bool {
			ret = constant.BoolVal(val)
		}
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64:

		if untyp.Kind == Complex && constant.Compare(constant.Imag(val), token.EQL, constantValZero) {
			// allow conversion from untyped complex to untyped integer or float,
			// provided that untyped complex has zero imaginary part.
			//
			// Required by the example var s uint = complex(1, 0)
			// mentioned at https://golang.org/ref/spec#Complex_numbers
			val = constant.Real(val)
		}
		fallthrough
	case r.Complex64, r.Complex128:

		n := untyp.extractNumber(val, t)
		return ConvertLiteralCheckOverflow(n, t)
	case r.Interface:
		// this can happen too... for example in "var foo interface{} = 7"
		// and it requires to convert the untyped constant to its default type.
		// Obviously, untyped constants can only implement empty interfaces
		// (TODO: unless CTI generics are enabled - they add methods to basic types)
		if t.NumMethod() == 0 {
			t = untyp.DefaultType()
			goto again
		}
	case r.Slice:
		// https://golang.org/ref/spec#String_literals states:
		//
		// 4. Converting a value of a string type to a slice of bytes type
		// yields a slice whose successive elements are the bytes of the string.
		//
		// 5. Converting a value of a string type to a slice of runes type
		// yields a slice containing the individual Unicode code points of the string.
		if val.Kind() == constant.String {
			s := strings.UnescapeString(val.ExactString())
			switch t.Elem().Kind() {
			case r.Uint8:
				ret = []byte(s)
			case r.Int32:
				ret = []rune(s)
			}
		}
	case r.String:
		switch val.Kind() {
		case constant.String:
			// untyped string -> string
			ret = strings.UnescapeString(val.ExactString())
		case constant.Int:
			// https://golang.org/ref/spec#String_literals states:
			//
			// 1. Converting a signed or unsigned integer value to a string type yields
			// a string containing the UTF-8 representation of the integer.
			// Values outside the range of valid Unicode code points are converted to "\uFFFD".

			i, exact := constant.Int64Val(val)
			if exact {
				ret = string(i)
			} else {
				ret = "\uFFFD"
			}
		}
	case r.Ptr:
		ret = untyp.toMathBig(t)
	}
	if ret == nil {
		output.Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil
	}
	v := r.ValueOf(ret)
	if v.Type() != t.ReflectType() {
		ret = v.Convert(t.ReflectType())
	}
	return ret
}

// EXTENSION: conversion from untyped constant to big.Int, bit.Rat, big.Float
func (untyp *Lit) toMathBig(t xr.Type) interface{} {
	var ret interface{}
	if k := untyp.Val.Kind(); k == constant.Int || k == constant.Float {
		switch t.ReflectType() {
		case rtypeOfPtrBigInt:
			if a := untyp.BigInt(); a != nil {
				ret = a
			}
		case rtypeOfPtrBigRat:
			if a := untyp.BigRat(); a != nil {
				ret = a
			}
		case rtypeOfPtrBigFloat:
			if a := untyp.BigFloat(); a != nil {
				ret = a
			}
		}
	}
	return ret
}

func (untyp *Lit) BigInt() *big.Int {
	var b big.Int
	var ret *big.Int

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
	} else if n, exact := untyp.Uint64(); exact {
		ret = b.SetUint64(n)
	} else {
		i, r, f := untyp.rawBignum()
		if i != nil {
			ret = b.Set(i)
		} else if r != nil {
			if !r.IsInt() {
				return nil
			}
			ret = b.Set(r.Num())
		} else if f != nil {
			if !f.IsInt() {
				return nil
			}
			if i, acc := f.Int(&b); acc == big.Exact {
				if i != &b {
					b.Set(i)
				}
				ret = &b
			}
		}
	}
	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		if _, ok := b.SetString(s, 0); ok {
			ret = &b
		}
	}
	return ret
}

func (untyp *Lit) BigRat() *big.Rat {
	var b big.Rat
	var ret *big.Rat

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
	} else {
		i, r, f := untyp.rawBignum()
		if i != nil {
			ret = b.SetInt(i)
		} else if r != nil {
			ret = b.Set(r)
		} else if f != nil {
			if f.IsInt() {
				if i, acc := f.Int(nil); acc == big.Exact {
					ret = b.SetInt(i)
				}
			}
		}
	}
	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		_, ok := b.SetString(s)
		if ok {
			ret = &b
		}
	}
	return ret
}

func (untyp *Lit) BigFloat() *big.Float {
	var b big.Float
	var ret *big.Float

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
		// Debugf("untyped.Lit.BigFloat(): converted int64 %v to *big.Float %v", i, b)
	} else if f, exact := untyp.Float64(); exact {
		ret = b.SetFloat64(f)
		// Debugf("untyped.Lit.BigFloat(): converted float64 %v to *big.Float %v", f, b)
	} else {
		i, r, f := untyp.rawBignum()
		if i != nil {
			ret = b.SetInt(i)
			// Debugf("untyped.Lit.BigFloat(): converted *big.Int %v to *big.Float %v", *i, b)
		} else if r != nil {
			ret = b.SetRat(r)
			// Debugf("untyped.Lit.BigFloat(): converted *big.Rat %v to *big.Float %v", *r, b)
		} else if f != nil {
			ret = b.Set(f)
			// Debugf("untyped.Lit.BigFloat(): converted *big.Float %v to *big.Float %v", *f, b)
		}
	}

	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		snum, sden := strings.Split2(s, '/')
		_, ok := b.SetString(snum)
		if ok && len(sden) != 0 {
			var b2 big.Float
			if _, ok = b2.SetString(sden); ok {
				b.Quo(&b, &b2)
			}
		}
		if ok {
			ret = &b
			// Debugf("untyped.Lit.BigFloat(): converted constant.Value %v %v to *big.Float %v", untyp.Val.Kind(), s, b)
		}
	}
	return ret
}

func (untyp *Lit) Int64() (int64, bool) {
	if c := untyp.Val; c.Kind() == constant.Int {
		return constant.Int64Val(c)
	}
	return 0, false
}

func (untyp *Lit) Uint64() (uint64, bool) {
	if c := untyp.Val; c.Kind() == constant.Int {
		return constant.Uint64Val(c)
	}
	return 0, false
}

func (untyp *Lit) Float64() (float64, bool) {
	if c := untyp.Val; c.Kind() == constant.Float {
		return constant.Float64Val(c)
	}
	return 0, false
}

// attempt to unwrap an untyped literal. Returns at most one of *big.Int, *big.Rat, *big.Float
func (untyp *Lit) rawBignum() (*big.Int, *big.Rat, *big.Float) {
	k := untyp.Val.Kind()
	if k != constant.Int && k != constant.Float {
		return nil, nil, nil
	}
	x := constant.Val(untyp.Val) // requires Go 1.13
	switch x := x.(type) {
	case *big.Int:
		return x, nil, nil
	case *big.Rat:
		return nil, x, nil
	case *big.Float:
		return nil, nil, x
	default:
		return nil, nil, nil
	}
}

// ================================= DefaultType =================================

// DefaultType returns the default type of an untyped constant.
func (untyp *Lit) DefaultType() xr.Type {
	switch untyp.Kind {
	case Bool, Rune, Int, Float, Complex, String:
		if basicTypes := untyp.basicTypes; basicTypes == nil {
			output.Errorf("untyped.Lit.DefaultType(): malformed untyped constant %v, has nil BasicTypes!", untyp)
			return nil
		} else {
			return (*basicTypes)[untyp.Kind]
		}

	default:
		output.Errorf("unexpected untyped constant %v, its default type is not known", untyp)
		return nil
	}
}

// ======================= utilities for Convert =======================

// extractNumber converts the untyped constant src to an integer, float or complex.
// panics if src has different kind from constant.Int, constant.Float and constant.Complex
// the receiver (untyp *Lit) and the second argument (t reflect.Type) are only used to pretty-print the panic error message
func (untyp *Lit) extractNumber(src constant.Value, t xr.Type) interface{} {
	var n interface{}
	cat := reflect.Category(t.Kind())
	var exact bool
	switch src.Kind() {
	case constant.Int:
		switch cat {
		case r.Int:
			n, exact = constant.Int64Val(src)
		case r.Uint:
			n, exact = constant.Uint64Val(src)
		default:
			n, exact = constant.Int64Val(src)
			if !exact {
				n, exact = constant.Uint64Val(src)
			}
		}
	case constant.Float:
		n, exact = constant.Float64Val(src)
	case constant.Complex:
		re := untyp.extractNumber(constant.Real(src), t)
		im := untyp.extractNumber(constant.Imag(src), t)
		rfloat := r.ValueOf(re).Convert(reflect.TypeOfFloat64).Float()
		ifloat := r.ValueOf(im).Convert(reflect.TypeOfFloat64).Float()
		n = complex(rfloat, ifloat)
		exact = true
	default:
		output.Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil
	}
	// allow inexact conversions to float64 and complex128:
	// floating point is intrinsically inexact, and Go compiler allows them too
	if !exact && (cat == r.Int || cat == r.Uint) {
		output.Errorf("untyped constant %v overflows <%v>", untyp, t)
		return nil
	}
	return n
}

// ConvertLiteralCheckOverflow converts a literal to type t and returns the converted value.
// panics if the conversion overflows the given type
func ConvertLiteralCheckOverflow(src interface{}, to xr.Type) interface{} {
	v := xr.ValueOf(src)
	rto := to.ReflectType()
	vto := reflect.ConvertValue(v, rto)

	k, kto := v.Kind(), vto.Kind()
	if k == kto {
		return vto.Interface() // no numeric conversion happened
	}
	c, cto := reflect.Category(k), reflect.Category(kto)
	if cto == r.Int || cto == r.Uint {
		if c == r.Float64 || c == r.Complex128 {
			// float-to-integer conversion. check for truncation
			t1 := reflect.ValueType(v)
			vback := reflect.ConvertValue(vto, t1)
			if src != vback.Interface() {
				output.Errorf("constant %v truncated to %v", src, to)
				return nil
			}
		} else {
			// integer-to-integer conversion. convert back and compare the interfaces for overflows
			t1 := reflect.ValueType(v)
			vback := vto.Convert(t1)
			if src != vback.Interface() {
				output.Errorf("constant %v overflows <%v>", src, to)
				return nil
			}
		}
	}
	return vto.Interface()
}
