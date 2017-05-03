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
 * global.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/constant"
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/base"
)

// ================================= Untyped =================================

// UntypedLit represents an untyped literal value, i.e. an untyped constant
type UntypedLit struct {
	Kind r.Kind // default type. matches Obj.Kind() except for rune literals, where Kind == reflect.Int32
	Obj  constant.Value
}

var (
	typeOfUntypedLit = r.TypeOf(UntypedLit{})

	UntypedZero = UntypedLit{Kind: r.Int, Obj: constant.MakeInt64(0)}
	UntypedOne  = UntypedLit{Kind: r.Int, Obj: constant.MakeInt64(1)}
)

// pretty-print untyped constants
func (untyp *UntypedLit) String() string {
	obj := untyp.Obj
	var strkind, strobj interface{} = untyp.Kind, nil
	if untyp.Kind == r.Int32 {
		strkind = "rune"
		if obj.Kind() == constant.Int {
			if i, exact := constant.Int64Val(obj); exact {
				if i >= 0 && i <= 0x10FFFF {
					strobj = fmt.Sprintf("%q", i)
				}
			}
		}
	}
	if strobj == nil {
		strobj = obj.ExactString()
	}
	return fmt.Sprintf("{%v %v}", strkind, strobj)
}

// ================================= Lit =================================

// Lit represents a literal value, i.e. a typed or untyped constant
type Lit struct {

	// Type is nil only for literal nils.
	// for all other literals, it is reflect.TypeOf(Lit.Value)
	//
	// when Lit is embedded in other structs that represent non-constant expressions,
	// Type is the first type returned by the expression (nil if returns no values)
	Type r.Type

	// Value is one of:
	//   nil, bool, int, int8, int16, int32, int64,
	//   uint, uint8, uint16, uint32, uint64, uintptr,
	//   float32, float64, complex64, complex128, string,
	//   UntypedLit
	//
	// when Lit is embedded in other structs that represent non-constant expressions,
	// Value is usually nil
	Value I
}

// Untyped returns true if Lit is an untyped constant
func (lit *Lit) Untyped() bool {
	_, ok := lit.Value.(UntypedLit)
	return ok
}

// UntypedKind returns the reflect.Kind of untyped constants,
// i.e. their "default type"
func (lit *Lit) UntypedKind() r.Kind {
	if untyp, ok := lit.Value.(UntypedLit); ok {
		return untyp.Kind
	} else {
		return r.Invalid
	}
}

func (lit Lit) String() string {
	switch val := lit.Value.(type) {
	case string, nil:
		return fmt.Sprintf("%#v", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ================================= Expr =================================

// Expr represents an expression in the "compiler"
type Expr struct {
	Lit
	Types []r.Type // in case the expression produces multiple values. if nil, use Lit.Type.
	Fun   I        // function that evaluates the expression at runtime.
	Sym   *Symbol  // in case the expression is a symbol
	IsNil bool
}

func (e *Expr) Const() bool {
	return e.Value != nil || e.IsNil
}

// NumOut returns the number of values that an expression will produce when evaluated
func (e *Expr) NumOut() int {
	if e.Types == nil {
		return 1
	}
	return len(e.Types)
}

// Out returns the i-th type that an expression will produce when evaluated
func (e *Expr) Out(i int) r.Type {
	if i == 0 && e.Types == nil {
		return e.Type
	}
	return e.Types[i]
}

// Outs returns the types that an expression will produce when evaluated
func (e *Expr) Outs() []r.Type {
	if e.Types == nil {
		return []r.Type{e.Type}
	}
	return e.Types
}

func (e *Expr) String() string {
	if e == nil {
		return "nil"
	}
	var str string
	if e.Const() {
		str = e.Lit.String()
	} else {
		str = fmt.Sprintf("%#v", e.Fun)
	}
	return str
}

// ================================= Stmt =================================

// Stmt represents a statement in the fast interpreter
type Stmt func(*Env) (Stmt, *Env)

// ================================= Builtin =================================

// Builtin represents a builtin function in the fast interpreter
type Builtin struct {
	// interpreted code should not access "compile": not exported.
	// compile usually needs to modify Symbol: pass it by value.
	compile func(c *Comp, sym Symbol, node *ast.CallExpr) *Call
	ArgMin  uint16
	ArgMax  uint16
}

var TypeOfBuiltin = r.TypeOf(Builtin{})

// ================================= EnvFunction =================================

// Function represents a function that accesses *CompEnv in the fast interpreter
type Function struct {
	Fun  I
	Type r.Type
}

var TypeOfFunction = r.TypeOf(Function{})

// ================================= BindClass =================================

type BindClass int

const (
	ConstBind = BindClass(iota)
	FuncBind
	VarBind
	IntBind
)

func (class BindClass) String() string {
	switch class {
	case ConstBind:
		return "constant"
	case FuncBind:
		return "function"
	default:
		return "variable"
	}
}

// ================================== BindDescriptor =================================

// the zero value of BindDescriptor is a valid descriptor for all constants
type BindDescriptor BindClass

const (
	bindClassMask  = BindClass(0x3)
	bindIndexShift = 2

	NoIndex             = int(0)                    // index of constants, functions and variables named "_"
	ConstBindDescriptor = BindDescriptor(ConstBind) // bind descriptor for all constants
)

func MakeBindDescriptor(class BindClass, index int) BindDescriptor {
	class &= bindClassMask
	return BindDescriptor(index<<bindIndexShift | int(class))
}

// IntBind returns true if BindIndex refers to a slot in Env.IntBinds (the default is a slot in Env.Binds)
func (desc BindDescriptor) Class() BindClass {
	return BindClass(desc) & bindClassMask
}

// Index returns the slice index to use in Env.Binds or Env.IntBinds to access a variable or function.
// returns NoIndex for variables and functions named "_"
func (desc BindDescriptor) Index() int {
	index := int(desc >> bindIndexShift)
	// debugf("BindDescriptor=%v, class=%v, index=%v", desc, desc.Class(), index)
	return index
}

func (desc BindDescriptor) Settable() bool {
	class := desc.Class()
	return class == IntBind || class == VarBind
}

func (desc BindDescriptor) String() string {
	return fmt.Sprintf("%s index=%d", desc.Class(), desc.Index())
}

// ================================== Bind =================================

// Bind represents a constant, variable, function or builtin in the "compiler"
type Bind struct {
	Lit
	Desc BindDescriptor
	Name string
}

func (bind *Bind) String() string {
	return fmt.Sprintf("{%s %q %#v <%v>}", bind.Desc, bind.Name, bind.Lit.Value, bind.Lit.Type)
}

func (bind *Bind) Const() bool {
	return bind.Desc.Class() == ConstBind
}

func BindConst(value I) *Bind {
	return &Bind{Lit: Lit{Type: r.TypeOf(value), Value: value}, Desc: ConstBindDescriptor}
}

func (bind *Bind) AsVar(upn int, opt PlaceOption) *Var {
	class := bind.Desc.Class()
	switch class {
	case VarBind, IntBind:
		return &Var{Upn: upn, Desc: bind.Desc, Type: bind.Type, Name: bind.Name}
	default:
		base.Errorf("%s a %s: %s <%v>", opt, class, bind.Name, bind.Type)
		return nil
	}
}

func (bind *Bind) AsSymbol(upn int) *Symbol {
	return &Symbol{Bind: *bind, Upn: upn}
}

type NamedType struct {
	Name, Path string
}

// ================================== Symbol, Var, Place =================================

// Symbol represents a resolved constant, function, variable or builtin
type Symbol struct {
	Bind
	Upn int
}

func (sym *Symbol) AsVar(opt PlaceOption) *Var {
	return sym.Bind.AsVar(sym.Upn, opt)
}

// Var represents a settable variable
type Var struct {
	// when Var is embedded in other structs that represent non-identifiers,
	// Upn and Desc are usually the zero values
	Upn  int
	Desc BindDescriptor
	Type r.Type
	Name string
}

// Place represents a settable place or, equivalently, its address
type Place struct {
	Var
	// Fun is nil for variables.
	// For non-variables, returns a settable and addressable reflect.Value: the place itself.
	// For map[key], Fun returns the map itself (which may NOT be settable).
	// Call Fun only once, it may have side effects!
	Fun func(*Env) r.Value
	// Fddr is nil for variables.
	// For non-variables, it will return the address of the place.
	// For map[key], it is nil since map[key] is not addressable
	// Call Addr only once, it may have side effects!
	Addr func(*Env) r.Value
	// used only for map[key], returns key. call it only once, it may have side effects!
	MapKey func(*Env) r.Value
}

func (place *Place) IsVar() bool {
	return place.Fun == nil
}

type PlaceOption bool // the reason why we want a place: either to write into it, or to take its address

const (
	PlaceSettable PlaceOption = false
	PlaceAddress  PlaceOption = true
)

func (opt PlaceOption) String() string {
	if opt == PlaceAddress {
		return "cannot take the address of"
	} else {
		return "cannot assign to"
	}
}

// ================================== Comp, Env =================================

type CompileOptions int

const (
	CompileKeepUntyped CompileOptions = 1 << iota // if set, Compile() on expressions will keep all untyped constants as such (in expressions where Go compiler would compute an untyped constant too)
	CompileDefaults    CompileOptions = 0
)

type Code struct {
	List []Stmt
}

type LoopInfo struct {
	Break      *int
	Continue   *int
	Labels     map[string]*int
	ThisLabels []string // sorted. for labeled "switch" and "for"
}

func (l *LoopInfo) HasLabel(label string) bool {
	i := sort.SearchStrings(l.ThisLabels, label)
	return i >= 0 && i < len(l.ThisLabels) && l.ThisLabels[i] == label
}

type FuncInfo struct {
	Params       []*Bind
	Results      []*Bind
	NamedResults bool
}

const (
	PoolCapacity = 32
)

// ThreadGlobals contains per-goroutine interpreter bookeeping information
type ThreadGlobals struct {
	FileEnv   *Env
	TopEnv    *Env
	Interrupt Stmt
	Signal    Signal // set by interrupts: Return, Defer...
	PoolSize  int
	Pool      [PoolCapacity]*Env
	*base.Globals
}

// Comp is a tree-of-closures builder: it transforms ast.Nodes into closures
// for faster execution. Consider it a poor man's compiler (hence the name)
type Comp struct {
	Binds      map[string]*Bind
	BindNum    int // len(Binds) == BindNum + IntBindNum + # of constants
	IntBindNum int
	// UpCost is the number of *Env.Outer hops to perform at runtime to reach the *Env corresponding to *Comp.Outer
	// usually equals one. will be zero if this *Comp defines no local variables/functions.
	UpCost         int
	Depth          int
	Types          map[string]r.Type
	NamedTypes     map[r.Type]NamedType
	Code           Code      // "compiled" code
	Loop           *LoopInfo // != nil when compiling a for or switch
	Func           *FuncInfo // != nil when compiling a function
	Outer          *Comp
	Name           string // set by "package" directive
	Path           string
	CompileOptions CompileOptions
	*base.Globals
}

const (
	// conventional values
	AnyDepth  = -1
	FileDepth = -2
	TopDepth  = -3
)

type Signal int

const (
	SigNone Signal = iota
	SigReturn
	SigInstallDeferHandler
)

// Env is the interpreter's runtime environment
type Env struct {
	Binds         []r.Value
	IntBinds      []uint64
	Outer         *Env
	IP            int
	Code          []Stmt
	ThreadGlobals *ThreadGlobals
	UsedByClosure bool // a bitfield would introduce more races among goroutines
	AddressTaken  bool // true if &Env.IntBinds[index] was executed... then we cannot reuse IntBinds
}

// CompEnv is the composition of both the tree-of-closures builder Comp
// and the interpreter's runtime environment Env
type CompEnv struct {
	Comp *Comp
	env  *Env // not exported. to access it, call CompEnv.PrepareEnv()
}

var typeOfCompEnv = r.TypeOf((*CompEnv)(nil))

type (
	I interface{}
	X func(*Env)
	/*
		XBool func(*Env) bool
		XInt        func(*Env) int
		XInt8       func(*Env) int8
		XInt16      func(*Env) int16
		XInt32      func(*Env) int32
		XInt64      func(*Env) int64
		XUint       func(*Env) uint
		XUint8      func(*Env) uint8
		XUint16     func(*Env) uint16
		XUint32     func(*Env) uint32
		XUint64     func(*Env) uint64
		XUintptr    func(*Env) uintptr
		XFloat32    func(*Env) float32
		XFloat64    func(*Env) float64
		XComplex64  func(*Env) complex64
		XComplex128 func(*Env) complex128
		XString     func(*Env) string
		XV          func(*Env) (r.Value, []r.Value)
	*/
)
