// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/fast"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package fast

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/fast"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/fast"] = imports.Package{
		Binds: map[string]r.Value{
			"AnyDepth":            r.ValueOf(AnyDepth),
			"ConstBind":           r.ValueOf(ConstBind),
			"ConstBindDescriptor": r.ValueOf(ConstBindDescriptor),
			"FileDepth":           r.ValueOf(FileDepth),
			"FuncBind":            r.ValueOf(FuncBind),
			"IntBind":             r.ValueOf(IntBind),
			"Interrupt":           r.ValueOf(&Interrupt).Elem(),
			"MakeBindDescriptor":  r.ValueOf(MakeBindDescriptor),
			"New":                 r.ValueOf(New),
			"NewComp":             r.ValueOf(NewComp),
			"NewEnv":              r.ValueOf(NewEnv),
			"NewInnerInterp":      r.ValueOf(NewInnerInterp),
			"NewThreadGlobals":    r.ValueOf(NewThreadGlobals),
			"NoIndex":             r.ValueOf(NoIndex),
			"OptDefaults":         r.ValueOf(OptDefaults),
			"OptKeepUntyped":      r.ValueOf(OptKeepUntyped),
			"OptIsCompiled":       r.ValueOf(OptIsCompiled),
			"PlaceAddress":        r.ValueOf(PlaceAddress),
			"PlaceSettable":       r.ValueOf(PlaceSettable),
			"PoolCapacity":        r.ValueOf(PoolCapacity),
			"SigDefer":            r.ValueOf(SigDefer),
			"SigNone":             r.ValueOf(SigNone),
			"SigReturn":           r.ValueOf(SigReturn),
			"TopDepth":            r.ValueOf(TopDepth),
			"VarBind":             r.ValueOf(VarBind),
		}, Types: map[string]r.Type{
			"Assign":             r.TypeOf((*Assign)(nil)).Elem(),
			"Bind":               r.TypeOf((*Bind)(nil)).Elem(),
			"BindClass":          r.TypeOf((*BindClass)(nil)).Elem(),
			"BindDescriptor":     r.TypeOf((*BindDescriptor)(nil)).Elem(),
			"Builtin":            r.TypeOf((*Builtin)(nil)).Elem(),
			"Call":               r.TypeOf((*Call)(nil)).Elem(),
			"Code":               r.TypeOf((*Code)(nil)).Elem(),
			"Comp":               r.TypeOf((*Comp)(nil)).Elem(),
			"CompGlobals":        r.TypeOf((*CompGlobals)(nil)).Elem(),
			"CompileOptions":     r.TypeOf((*CompileOptions)(nil)).Elem(),
			"Env":                r.TypeOf((*Env)(nil)).Elem(),
			"Expr":               r.TypeOf((*Expr)(nil)).Elem(),
			"FuncInfo":           r.TypeOf((*FuncInfo)(nil)).Elem(),
			"Function":           r.TypeOf((*Function)(nil)).Elem(),
			"I":                  r.TypeOf((*I)(nil)).Elem(),
			"Import":             r.TypeOf((*Import)(nil)).Elem(),
			"Interp":             r.TypeOf((*Interp)(nil)).Elem(),
			"Lit":                r.TypeOf((*Lit)(nil)).Elem(),
			"LoopInfo":           r.TypeOf((*LoopInfo)(nil)).Elem(),
			"Macro":              r.TypeOf((*Macro)(nil)).Elem(),
			"NamedType":          r.TypeOf((*NamedType)(nil)).Elem(),
			"Place":              r.TypeOf((*Place)(nil)).Elem(),
			"PlaceOption":        r.TypeOf((*PlaceOption)(nil)).Elem(),
			"Signal":             r.TypeOf((*Signal)(nil)).Elem(),
			"Stmt":               r.TypeOf((*Stmt)(nil)).Elem(),
			"Symbol":             r.TypeOf((*Symbol)(nil)).Elem(),
			"ThreadGlobals":      r.TypeOf((*ThreadGlobals)(nil)).Elem(),
			"TypeAssertionError": r.TypeOf((*TypeAssertionError)(nil)).Elem(),
			"UntypedLit":         r.TypeOf((*UntypedLit)(nil)).Elem(),
			"Var":                r.TypeOf((*Var)(nil)).Elem(),
		}, Proxies: map[string]r.Type{
			"I": r.TypeOf((*I_github_com_cosmos72_gomacro_fast)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"AnyDepth":     "int:-1",
			"FileDepth":    "int:-2",
			"PoolCapacity": "int:32",
			"TopDepth":     "int:-3",
		}, Wrappers: map[string][]string{
			"Bind":              []string{"ConstTo", "DefaultType", "ReflectValue", "Untyped", "UntypedKind"},
			"Comp":              []string{"CollectAst", "CollectNode", "CollectPackageImports", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymEmbedded", "GensymPrivate", "IncLine", "IncLineBytes", "Init", "LookupPackage", "ParseBytes", "Position", "Sprintf", "ToString", "TypeOfBool", "TypeOfBuiltin", "TypeOfComplex128", "TypeOfComplex64", "TypeOfError", "TypeOfFloat32", "TypeOfFloat64", "TypeOfFunction", "TypeOfImport", "TypeOfInt", "TypeOfInt16", "TypeOfInt32", "TypeOfInt64", "TypeOfInt8", "TypeOfInterface", "TypeOfMacro", "TypeOfString", "TypeOfUint", "TypeOfUint16", "TypeOfUint32", "TypeOfUint64", "TypeOfUint8", "TypeOfUintptr", "TypeOfUntypedLit", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"CompThreadGlobals": []string{"CollectAst", "CollectNode", "CollectPackageImports", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymEmbedded", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "Init", "LookupPackage", "ParseBytes", "Position", "Sprintf", "ToString", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"Expr":              []string{"ReflectValue", "Untyped", "UntypedKind"},
			"Place":             []string{"Address", "AsPlace", "AsSymbol"},
			"Symbol":            []string{"AsSymbol", "Const", "ConstTo", "ConstValue", "DefaultType", "ReflectValue", "String", "Untyped", "UntypedKind"},
			"ThreadGlobals":     []string{"CollectAst", "CollectNode", "CollectPackageImports", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymEmbedded", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "Init", "LookupPackage", "ParseBytes", "Position", "Sprintf", "ToString", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
		},
	}
}

// --------------- proxy for github.com/cosmos72/gomacro/fast.I ---------------
type I_github_com_cosmos72_gomacro_fast struct {
	Object interface{}
}
