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
 * quasiquote.go
 *
 *  Created on Jun 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/muazhari/gomacro-custom/ast2"
	"github.com/muazhari/gomacro-custom/base"
	"github.com/muazhari/gomacro-custom/base/output"
	"github.com/muazhari/gomacro-custom/base/reflect"
	etoken "github.com/muazhari/gomacro-custom/go/etoken"
	mp "github.com/muazhari/gomacro-custom/go/parser"
	xr "github.com/muazhari/gomacro-custom/xreflect"
)

var (
	rtypeOfNode      = r.TypeOf((*ast.Node)(nil)).Elem()
	rtypeOfUnaryExpr = r.TypeOf((*ast.UnaryExpr)(nil))
	rtypeOfBlockStmt = r.TypeOf((*ast.BlockStmt)(nil)).Elem()
)

func (c *Comp) quasiquoteUnary(unary *ast.UnaryExpr) *Expr {
	block := unary.X.(*ast.FuncLit).Body
	node := base.SimplifyNodeForQuote(block, true)

	if block != nil && len(block.List) == 1 {
		unary, ok := base.SimplifyNodeForQuote(block.List[0], false).(*ast.UnaryExpr)
		if ok && (unary.Op == etoken.UNQUOTE || unary.Op == etoken.UNQUOTE_SPLICE) {
			// to support quasiquote{unquote ...} and quasiquote{unquote_splice ...}
			// we invoke SimplifyNodeForQuote() at the end, not at the beginning.

			in := ToAst(block)
			expr := c.quasiquote1(in, 1, true)

			if unary.Op == etoken.UNQUOTE_SPLICE {
				return expr
			}
			fun := expr.AsX1()
			toUnwrap := block != node
			return exprX1(c.Universe.FromReflectType(rtypeOfNode), func(env *Env) xr.Value {
				x := reflect.ValueInterface(fun(env))
				node := AnyToAstWithNode(x, "Quasiquote").Node()
				node = base.SimplifyNodeForQuote(node, toUnwrap)
				return xr.ValueOf(node)
			})
		}
	}
	return c.quasiquote1(ToAst(node), 1, true)
}

// Quasiquote expands and compiles ~quasiquote, if Ast starts with it
func (c *Comp) Quasiquote(in Ast) *Expr {
	switch form := in.(type) {
	case UnaryExpr:
		if form.Op() == etoken.QUASIQUOTE {
			body := form.X.X.(*ast.FuncLit).Body
			return c.quasiquote1(ToAst(body), 1, true)
		}
	}
	return c.Compile(in)
}

func (c *Comp) quasiquote1(in Ast, depth int, canSplice bool) *Expr {
	expr, _ := c.quasiquote(in, depth, canSplice)
	return expr
}

// quasiquote expands and compiles the contents of a ~quasiquote
func (c *Comp) quasiquote(in Ast, depth int, canSplice bool) (*Expr, bool) {
	if in == nil || in.Interface() == nil {
		return nil, false
	}
	debug := c.Options&base.OptDebugQuasiquote != 0
	var label string
	if canSplice {
		label = " splice"
	}
	if debug {
		c.Debugf("Quasiquote[%d]%s expanding %s: %v // %T", depth, label, etoken.String(etoken.QUASIQUOTE), in.Interface(), in.Interface())
	}

	switch in := in.(type) {
	case AstWithSlice:
		n := in.Size()
		funs := make([]func(*Env) xr.Value, 0, n)
		splices := make([]bool, 0, n)
		positions := make([]token.Position, 0, n)
		for i := 0; i < n; i++ {
			if form := in.Get(i); form != nil {
				form = base.SimplifyAstForQuote(form, false)
				expr, splice := c.quasiquote(form, depth, true)
				fun := expr.AsX1()
				if fun == nil {
					c.Warnf("Quasiquote[%d]%s: node expanded to nil: %v // %T", depth, label, form.Interface(), form.Interface())
					continue
				}
				funs = append(funs, fun)
				splices = append(splices, splice)
				var position token.Position
				if form, ok := form.(AstWithNode); ok {
					position = c.Fileset.Position(form.Node().Pos())
				}
				positions = append(positions, position)
			}
		}
		form := in.New().(AstWithSlice)

		typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
		rtype := typ.ReflectType()

		return exprX1(typ, func(env *Env) xr.Value {
			out := form.New().(AstWithSlice)
			for i, fun := range funs {
				x := reflect.ValueInterface(fun(env))
				if debug {
					output.Debugf("Quasiquote: env=%p, append to AstWithSlice: <%v> returned %v // %T", env, r.TypeOf(fun), x, x)
				}
				if x == nil {
					continue
				} else if !splices[i] {
					out = out.Append(anyToAst(x, positions[i]))
				} else {
					xs := AnyToAstWithSlice(x, positions[i])
					n := xs.Size()
					for j := 0; j < n; j++ {
						if xj := xs.Get(j); xj != nil {
							out = out.Append(xj)
						}
					}
				}
			}
			return xr.ValueOf(out.Interface()).Convert(rtype)
		}), false
	case UnaryExpr:
		unary := in.X
		switch op := unary.Op; op {
		case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
			inner, unquoteDepth := base.DescendNestedUnquotes(in)
			if debug {
				c.Debugf("Quasiquote[%d]%s deep splice expansion? %v. unquoteDepth = %d, inner.Op() = %s: %v // %T",
					depth, label, unquoteDepth > 1 && unquoteDepth >= depth && inner.Op() == etoken.UNQUOTE_SPLICE,
					unquoteDepth, etoken.String(inner.Op()), inner, inner)
			}
			if unquoteDepth > 1 && unquoteDepth >= depth && inner.Op() == etoken.UNQUOTE_SPLICE {
				// complication: in Common Lisp, the right-most unquote pairs with the left-most comma!
				// we implement the same mechanics, so we must drill down to the last unquote/unquote_splice
				// and, for unquote_splice, create a copy of the unquote/unquote_splice stack for each result.
				// Example:
				//   x:=quote{7; 8}
				//   quasiquote{quasiquote{1; unquote{2}; unquote{unquote_splice{x}}}}
				// must return
				//   quasiquote{1; unquote{2}; unquote{7}; unquote{8}}

				depth -= unquoteDepth
				node := base.SimplifyNodeForQuote(inner.X.X.(*ast.FuncLit).Body, true)
				form := ToAst(node)
				if debug {
					c.Debugf("Quasiquote[%d]%s deep splice compiling %s: %v // %T", depth, label, etoken.String(inner.Op()), node, node)
				}
				fun := c.compileExpr(form).AsX1()
				toks, pos := base.CollectNestedUnquotes(in)
				position := c.Fileset.Position(pos[0])
				pos0 := pos[0]
				end := unary.End()
				toks = toks[:unquoteDepth-1]
				pos = pos[:unquoteDepth-1]

				return exprX1(c.Universe.FromReflectType(rtypeOfBlockStmt), func(env *Env) xr.Value {
					x := reflect.ValueInterface(fun(env))
					// Debugf("Quasiquote: runtime deep expansion returned: %v // %T", x, x)
					form := AnyToAstWithSlice(x, position)
					out := BlockStmt{&ast.BlockStmt{Lbrace: pos0, Rbrace: end}}
					for i, ni := 0, form.Size(); i < ni; i++ {
						// cheat: BlockStmt.Append() does not modify the receiver
						formi := AnyToAstWithNode(form.Get(i), position)
						out.Append(base.MakeNestedQuote(formi, toks, pos))
					}
					return xr.ValueOf(out.X)
				}), true
			}
			fallthrough
		case etoken.QUOTE, etoken.QUASIQUOTE:
			node := base.SimplifyNodeForQuote(unary.X.(*ast.FuncLit).Body, true)
			form := ToAst(node)

			if op == etoken.QUASIQUOTE {
				depth++
			} else if op == etoken.UNQUOTE || op == etoken.UNQUOTE_SPLICE {
				depth--
			}
			if depth <= 0 {
				if debug {
					c.Debugf("Quasiquote[%d]%s compiling %s: %v // %T", depth, label, etoken.String(op), node, node)
				}
				return c.compileExpr(form), op == etoken.UNQUOTE_SPLICE
			}
			fun := c.quasiquote1(form, depth, true).AsX1()
			if fun == nil {
				c.Warnf("Quasiquote[%d]%s: node expanded to nil: %v // %T", depth, label, node, node)
			}
			var pos token.Pos
			var position token.Position
			if node, ok := node.(ast.Node); ok {
				pos = node.Pos()
				position = c.Fileset.Position(pos)
			}
			if op == etoken.UNQUOTE_SPLICE {
				return c.quoteUnquoteSplice(op, pos, position, fun), false
			}
			return exprX1(c.Universe.FromReflectType(rtypeOfUnaryExpr), func(env *Env) xr.Value {
				var node ast.Node
				if fun != nil {
					x := reflect.ValueInterface(fun(env))
					if debug {
						output.Debugf("Quasiquote: env = %p, body of %s: <%v> returned %v <%v>", env, etoken.String(op), r.TypeOf(fun), x, r.TypeOf(x))
					}
					node = AnyToAstWithNode(x, position).Node()
				}
				ret, _ := mp.MakeQuote(nil, op, token.NoPos, node)
				return xr.ValueOf(ret)
			}), false
		}
	}

	// Ast can still be a tree: just not a resizeable one, so support ~unquote but not ~unquote_splice
	in, ok := in.(AstWithNode)
	if !ok {
		x := in.Interface()
		c.Errorf("Quasiquote: unsupported node type, expecting AstWithNode or AstWithSlice: %v <%v>", x, r.TypeOf(x))
		return nil, false
	}
	node := in.Interface()
	if debug {
		c.Debugf("Quasiquote[%d] recursing: %v <%v>", depth, node, r.TypeOf(node))
	}
	if node == nil {
		return nil, false
	}
	form := in.New().(AstWithNode) // we must NOT retain input argument, so clone it
	n := in.Size()
	typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
	rtype := typ.ReflectType()

	if n == 0 {
		return exprX1(typ, func(env *Env) xr.Value {
			return xr.ValueOf(form.New().Interface()).Convert(rtype)
		}), false
	}
	funs := make([]func(*Env) xr.Value, n)
	positions := make([]token.Position, n)
	for i := 0; i < n; i++ {
		if form := in.Get(i); form != nil {
			form = base.SimplifyAstForQuote(form, false)
			fun := c.quasiquote1(form, depth, false).AsX1()
			if fun == nil {
				c.Warnf("Quasiquote[%d]: node expanded to nil: %v", depth, form.Interface())
				continue
			}
			funs[i] = fun
			if form, ok := form.(AstWithNode); ok && form.Node() != nil {
				positions[i] = c.Fileset.Position(form.Node().Pos())
			}
		}
	}

	return exprX1(typ, func(env *Env) xr.Value {
		out := form.New().(AstWithNode)
		for i, fun := range funs {
			if fun != nil {
				x := reflect.ValueInterface(fun(env))
				if debug {
					output.Debugf("Quasiquote: env = %p, <%v> returned %v <%v>", env, r.TypeOf(fun), x, r.TypeOf(x))
				}
				out.Set(i, anyToAst(x, positions[i]))
			}
		}
		return xr.ValueOf(out.Interface()).Convert(rtype)
	}), false
}

func (c *Comp) quoteUnquoteSplice(op token.Token, pos token.Pos, position token.Position, fun func(*Env) xr.Value) *Expr {
	return exprX1(c.Universe.FromReflectType(rtypeOfUnaryExpr), func(env *Env) xr.Value {
		var node ast.Node
		if fun != nil {
			x := reflect.ValueInterface(fun(env))
			form := anyToAst(x, position)
			switch form := form.(type) {
			case AstWithNode:
				node = form.Node()
			case AstWithSlice:
				block := BlockStmt{&ast.BlockStmt{Lbrace: pos}}
				n := form.Size()
				for i := 0; i < n; i++ {
					if formi := form.Get(i); formi != nil {
						/*block =*/ block.Append(formi)
					}
				}
				node = block.X
			default:
				var prefix string
				if pos != token.NoPos {
					prefix = fmt.Sprintf("%s: ", position)
				}
				output.Errorf("%s%s returned invalid type, expecting AstWithNode or AstWithSlice: %v, <%v>",
					prefix, etoken.String(etoken.UNQUOTE_SPLICE), form, r.TypeOf(form))
				return xr.Value{}
			}
		}
		ret, _ := mp.MakeQuote(nil, op, token.NoPos, node)
		return xr.ValueOf(ret)
	})
}
