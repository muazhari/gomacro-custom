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
 * binary.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	"go/token"

	"github.com/muazhari/gomacro-custom/base/reflect"
	"github.com/muazhari/gomacro-custom/base/untyped"
	etoken "github.com/muazhari/gomacro-custom/go/etoken"
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

func (c *Comp) BinaryExpr(node *ast.BinaryExpr) *Expr {
	x := c.expr1(node.X, nil)
	y := c.expr1(node.Y, nil)
	return c.BinaryExpr1(node, x, y)
}

func (c *Comp) BinaryExpr1(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	if x.Untyped() && y.Untyped() {
		return c.BinaryExprUntyped(node, x.Value.(UntypedLit), y.Value.(UntypedLit))
	}
	bothConst := x.Const() && y.Const()
	var z *Expr

	op := tokenWithoutAssign(node.Op)
	switch op {
	case token.ADD:
		z = c.Add(node, x, y)
	case token.SUB:
		z = c.Sub(node, x, y)
	case token.MUL:
		z = c.Mul(node, x, y)
	case token.QUO:
		z = c.Quo(node, x, y)
	case token.REM:
		z = c.Rem(node, x, y)
	case token.AND:
		z = c.And(node, x, y)
	case token.OR:
		z = c.Or(node, x, y)
	case token.XOR:
		z = c.Xor(node, x, y)
	case token.SHL:
		z = c.Shl(node, x, y)
	case token.SHR:
		z = c.Shr(node, x, y)
	case token.AND_NOT:
		z = c.Andnot(node, x, y)
	case token.LAND:
		z = c.Land(node, x, y)
	case token.LOR:
		z = c.Lor(node, x, y)
	case token.EQL:
		z = c.Eql(node, x, y)
	case token.LSS:
		z = c.Lss(node, x, y)
	case token.GTR:
		z = c.Gtr(node, x, y)
	case token.NEQ:
		z = c.Neq(node, x, y)
	case token.LEQ:
		z = c.Leq(node, x, y)
	case token.GEQ:
		z = c.Geq(node, x, y)
	default:
		return c.unimplementedBinaryExpr(node, x, y)
	}
	if bothConst {
		// constant propagation
		z.EvalConst(COptKeepUntyped)
	}
	return z
}

func (c *Comp) BinaryExprUntyped(node *ast.BinaryExpr, x UntypedLit, y UntypedLit) *Expr {
	op := node.Op
	switch op {
	case token.LAND, token.LOR:
		xb, yb := x.Convert(c.TypeOfBool()).(bool), y.Convert(c.TypeOfBool()).(bool)
		var flag bool
		if op == token.LAND {
			flag = xb && yb
		} else {
			flag = xb || yb
		}
		return c.exprUntypedLit(untyped.Bool, constant.MakeBool(flag))
	case token.EQL, token.LSS, token.GTR, token.NEQ, token.LEQ, token.GEQ:
		// comparison gives an untyped bool
		flag := constant.Compare(x.Val, op, y.Val)
		return c.exprUntypedLit(untyped.Bool, constant.MakeBool(flag))
	case token.SHL, token.SHL_ASSIGN:
		return c.ShiftUntyped(node, token.SHL, x, y)
	case token.SHR, token.SHR_ASSIGN:
		return c.ShiftUntyped(node, token.SHR, x, y)
	default:
		op2 := tokenWithoutAssign(op)
		xint := x.Kind == untyped.Int || x.Kind == untyped.Rune
		yint := y.Kind == untyped.Int || y.Kind == untyped.Rune
		if op2 == token.QUO && xint && yint {
			// untyped integer division
			op2 = token.QUO_ASSIGN
		}
		zobj := constant.BinaryOp(x.Val, op2, y.Val)
		zkind := untyped.MakeKind(zobj.Kind())
		// c.Debugf("untyped binary expression %v %s %v returned {%v %v}", x, op2, y, zkind, zobj)
		// untyped.Rune has precedence over untyped.Int
		if zobj.Kind() == constant.Int {
			if xint && x.Kind != untyped.Int {
				zkind = x.Kind
			} else if yint && y.Kind != untyped.Int {
				zkind = y.Kind
			}
		}
		if zkind == untyped.None {
			c.Errorf("invalid binary operation: %v %v %v", x.Val, op, y.Val)
		}
		return c.exprUntypedLit(zkind, zobj)
	}
}

var tokenRemoveAssign = map[token.Token]token.Token{
	token.ADD_ASSIGN:     token.ADD,
	token.SUB_ASSIGN:     token.SUB,
	token.MUL_ASSIGN:     token.MUL,
	token.QUO_ASSIGN:     token.QUO,
	token.REM_ASSIGN:     token.REM,
	token.AND_ASSIGN:     token.AND,
	token.OR_ASSIGN:      token.OR,
	token.XOR_ASSIGN:     token.XOR,
	token.SHL_ASSIGN:     token.SHL,
	token.SHR_ASSIGN:     token.SHR,
	token.AND_NOT_ASSIGN: token.AND_NOT,
}

var tokenAddAssign = map[token.Token]token.Token{
	token.ADD:     token.ADD_ASSIGN,
	token.SUB:     token.SUB_ASSIGN,
	token.MUL:     token.MUL_ASSIGN,
	token.QUO:     token.QUO_ASSIGN,
	token.REM:     token.REM_ASSIGN,
	token.AND:     token.AND_ASSIGN,
	token.OR:      token.OR_ASSIGN,
	token.XOR:     token.XOR_ASSIGN,
	token.SHL:     token.SHL_ASSIGN,
	token.SHR:     token.SHR_ASSIGN,
	token.AND_NOT: token.AND_NOT_ASSIGN,
}

func tokenWithoutAssign(op token.Token) token.Token {
	ret, ok := tokenRemoveAssign[op]
	if !ok {
		ret = op
	}
	return ret
}

func tokenWithAssign(op token.Token) token.Token {
	ret, ok := tokenAddAssign[op]
	if !ok {
		ret = op
	}
	return ret
}

var warnUntypedShift, warnUntypedShift2 = true, true

func (c *Comp) ShiftUntyped(node *ast.BinaryExpr, op token.Token, x UntypedLit, y UntypedLit) *Expr {
	var yn64 uint64
	var exact bool

	switch y.Val.Kind() {
	case constant.Int:
		yn64, exact = constant.Uint64Val(y.Val)
	case constant.Float:
		yf, fexact := constant.Float64Val(y.Val)
		if fexact {
			yn64 = uint64(yf)
			exact = float64(yn64) == yf
		}
		// c.Debugf("ShiftUntyped: %v %v %v, rhs converted to %v <float64> => %v <uint64> (exact = %v)", x.Val, op, y.Val, yf, yn64, exact)
	}
	if !exact {
		c.Errorf("invalid shift: %v %v %v", x.Val.ExactString(), op, y.Val.ExactString())
	}
	yn := uint(yn64)
	if uint64(yn) != yn64 {
		c.Errorf("invalid shift: %v %v %v", x.Val.ExactString(), op, y.Val.ExactString())
	}
	xn := x.Val
	xkind := x.Kind
	switch xkind {
	case untyped.Int, untyped.Rune:
		// nothing to do
	case untyped.Float, untyped.Complex:
		if warnUntypedShift {
			c.Warnf("known limitation (warned only once): untyped floating point constant shifted by untyped constant. returning untyped integer instead of deducing the type from the surrounding context: %v",
				node)
			warnUntypedShift = false
		}
		sign := constant.Sign(xn)
		if xkind == untyped.Complex {
			sign = constant.Sign(constant.Real(xn))
		}
		if sign >= 0 {
			xn = constant.MakeUint64(x.Convert(c.TypeOfUint64()).(uint64))
		} else {
			xn = constant.MakeInt64(x.Convert(c.TypeOfInt64()).(int64))
		}
		xkind = untyped.Int
	default:
		c.Errorf("invalid shift: %v %v %v", x.Val, op, y.Val)
	}
	zobj := constant.Shift(xn, op, yn)
	if zobj.Kind() == constant.Unknown {
		c.Errorf("invalid shift: %v %v %v", x.Val, op, y.Val)
	}
	return c.exprUntypedLit(xkind, zobj)
}

// prepareShift panics if the types of xe and ye are not valid for shifts i.e. << or >>
// returns non-nil expression if it computes the shift operation itself
func (c *Comp) prepareShift(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if xe.Untyped() && ye.Untyped() {
		// untyped << untyped should not happen here, it's handled in Comp.BinaryExpr... but let's be safe
		return c.ShiftUntyped(node, node.Op, xe.Value.(UntypedLit), ye.Value.(UntypedLit))
	}
	xet, yet := xe.DefaultType(), ye.DefaultType()
	if xet == nil || !reflect.IsCategory(xet.Kind(), xr.Int, xr.Uint) {
		return c.invalidBinaryExpr(node, xe, ye)
	}
	if xe.Untyped() {
		xuntyp := xe.Value.(UntypedLit)
		if ye.Const() {
			// untyped << constant
			yuntyp := untyped.MakeLit(untyped.Int, constant.MakeUint64(xr.ValueOf(ye.Value).Uint()), &c.Universe.BasicTypes)
			return c.ShiftUntyped(node, node.Op, xuntyp, yuntyp)
		}
		// untyped << expression
		// BUG! we should deduce left operand type from its context, instead of assuming int
		// see https://golang.org/ref/spec#Operators
		//
		// "If the left operand of a non-constant shift expression is an untyped constant,
		// "it is first converted to the type it would assume if the shift expression
		// "were replaced by its left operand alone."
		if warnUntypedShift2 {
			c.Warnf("known limitation (warned only once): untyped constant shifted by a non-constant expression. returning int instead of deducing the type from the surrounding context: %v",
				node)
			warnUntypedShift2 = false
		}
		xe.ConstTo(c.TypeOfInt())
	}
	if ye.Untyped() {
		// untyped constants do not distinguish between int and uint
		if yet == nil || !reflect.IsCategory(yet.Kind(), xr.Int) {
			return c.invalidBinaryExpr(node, xe, ye)
		}
		ye.ConstTo(c.TypeOfUint64())
	} else {
		// accept shift by signed integer, introduced in Go 1.13
		if yet == nil || !reflect.IsCategory(yet.Kind(), xr.Int, xr.Uint) {
			return c.invalidBinaryExpr(node, xe, ye)
		}
	}
	xe.WithFun()
	ye.WithFun()
	return nil
}

func (c *Comp) Land(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	xval, xfun, xerr := x.TryAsPred()
	yval, yfun, yerr := y.TryAsPred()
	if xerr || yerr {
		return c.invalidBinaryExpr(node, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return y
		}
		return c.exprValue(nil, false)
	}
	if yfun == nil {
		if yval {
			return x
		}
		return c.exprBool(func(env *Env) bool {
			return xfun(env) && false
		})
	}
	return c.exprBool(func(env *Env) bool {
		return xfun(env) && yfun(env)
	})
}

func (c *Comp) Lor(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	xval, xfun, xerr := x.TryAsPred()
	yval, yfun, yerr := y.TryAsPred()
	if xerr || yerr {
		return c.invalidBinaryExpr(node, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return c.exprValue(nil, true)
		}
		return y
	}
	if yfun == nil {
		if yval {
			return c.exprBool(func(env *Env) bool {
				return xfun(env) || true
			})
		}
		return x
	}
	return c.exprBool(func(env *Env) bool {
		return xfun(env) || yfun(env)
	})
}

func (c *Comp) invalidBinaryExpr(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.badBinaryExpr("invalid", node, x, y)
}

func (c *Comp) unimplementedBinaryExpr(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.badBinaryExpr("unimplemented", node, x, y)
}

func (c *Comp) badBinaryExpr(reason string, node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	opstr := etoken.String(node.Op)
	var xstr, ystr string
	if x.Const() {
		xstr = x.String() + " "
	}
	if y.Const() {
		ystr = y.String() + " "
	}
	c.Errorf("%s binary operation %s between %s<%v> and %s<%v>: %v %s %v",
		reason, opstr, xstr, x.Type, ystr, y.Type, node.X, opstr, node.Y)
	return nil
}

// convert x and y to the same single-valued expression type. needed to convert untyped constants to regular Go types
func (c *Comp) toSameFuncType(node ast.Expr, xe *Expr, ye *Expr) {
	xe.CheckX1()
	ye.CheckX1()
	xconst, yconst := xe.Const(), ye.Const()
	if yconst {
		if xconst {
			c.constsToSameType(node, xe, ye)
			xe.WithFun()
			ye.WithFun()
		} else {
			ye.ConstTo(xe.Type)
		}
	} else if xconst {
		xe.ConstTo(ye.Type)
	} else if !xe.Type.IdenticalTo(ye.Type) {
		c.mismatchedTypes(node, xe, ye)
	}
}

func (c *Comp) constsToSameType(node ast.Expr, xe *Expr, ye *Expr) {
	x, y := xe.Value, ye.Value
	if x == nil {
		if y == nil {
			return
		} else {
			switch node := node.(type) {
			case *ast.BinaryExpr:
				c.invalidBinaryExpr(node, xe, ye)
			default:
				c.Errorf("invalid operation between %v <%v> and %v <%v>: %v", x, xe.Type, y, ye.Type, node)
			}
		}
	}
	xu, yu := xe.Untyped(), ye.Untyped()
	if xu && yu {
		c.Errorf("internal error, operation between untyped constants %v and %v not optimized away: %v",
			xe.Lit, ye.Lit, node)
	} else if xu {
		xe.ConstTo(ye.Type)
	} else if yu {
		ye.ConstTo(xe.Type)
	} else if xe.Type.ReflectType() != ye.Type.ReflectType() {
		c.mismatchedTypes(node, xe, ye)
	}
}

func (c *Comp) mismatchedTypes(node ast.Expr, xe *Expr, ye *Expr) {
	switch node := node.(type) {
	case *ast.BinaryExpr:
		c.badBinaryExpr("mismatched types in", node, xe, ye)
	default:
		c.Errorf("mismatched types %v and %v in: %v", xe.Type, ye.Type, node)
	}
}
