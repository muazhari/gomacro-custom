// this file was generated by gomacro command: import _i "github.com/muazhari/gomacro-custom/base"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package base

import (
	"reflect"
	r "reflect"

	"github.com/muazhari/gomacro-custom/imports"
	"github.com/muazhari/gomacro-custom/xreflect"
)

// reflection: allow interpreted code to import "github.com/muazhari/gomacro-custom/base"
func init() {
	imports.Packages["github.com/muazhari/gomacro-custom/base"] = imports.Package{
		Binds: map[string]r.Value{
			"CMacroExpand":               r.ValueOf(CMacroExpand),
			"CMacroExpand1":              r.ValueOf(CMacroExpand1),
			"CMacroExpandCodewalk":       r.ValueOf(CMacroExpandCodewalk),
			"CmdOptForceEval":            r.ValueOf(CmdOptForceEval),
			"CmdOptQuit":                 r.ValueOf(CmdOptQuit),
			"CollectNestedUnquotes":      r.ValueOf(CollectNestedUnquotes),
			"DescendNestedUnquotes":      r.ValueOf(DescendNestedUnquotes),
			"DuplicateNestedUnquotes":    r.ValueOf(DuplicateNestedUnquotes),
			"False":                      r.ValueOf(&False).Elem(),
			"IsGensym":                   r.ValueOf(IsGensym),
			"IsGensymAnonymous":          r.ValueOf(IsGensymAnonymous),
			"IsGensymInterface":          r.ValueOf(IsGensymInterface),
			"IsGensymPrivate":            r.ValueOf(IsGensymPrivate),
			"MakeBufReadline":            r.ValueOf(MakeBufReadline),
			"MakeNestedQuote":            r.ValueOf(MakeNestedQuote),
			"MakeQuote":                  r.ValueOf(MakeQuote),
			"MakeQuote2":                 r.ValueOf(MakeQuote2),
			"MakeTtyReadline":            r.ValueOf(MakeTtyReadline),
			"MaxInt":                     r.ValueOf(MaxInt),
			"MaxUint":                    r.ValueOf(MaxUint),
			"MaxUint16":                  r.ValueOf(MaxUint16),
			"MinInt":                     r.ValueOf(MinInt),
			"NewGlobals":                 r.ValueOf(NewGlobals),
			"NilR":                       r.ValueOf(&NilR).Elem(),
			"NoneR":                      r.ValueOf(&NoneR).Elem(),
			"One":                        r.ValueOf(&One).Elem(),
			"OptCollectDeclarations":     r.ValueOf(OptCollectDeclarations),
			"OptCollectStatements":       r.ValueOf(OptCollectStatements),
			"OptCtrlCEnterDebugger":      r.ValueOf(OptCtrlCEnterDebugger),
			"OptDebugCallStack":          r.ValueOf(OptDebugCallStack),
			"OptDebugDebugger":           r.ValueOf(OptDebugDebugger),
			"OptDebugField":              r.ValueOf(OptDebugField),
			"OptDebugFromReflect":        r.ValueOf(OptDebugFromReflect),
			"OptDebugGenerics":           r.ValueOf(OptDebugGenerics),
			"OptDebugMacroExpand":        r.ValueOf(OptDebugMacroExpand),
			"OptDebugMethod":             r.ValueOf(OptDebugMethod),
			"OptDebugParse":              r.ValueOf(OptDebugParse),
			"OptDebugQuasiquote":         r.ValueOf(OptDebugQuasiquote),
			"OptDebugRecover":            r.ValueOf(OptDebugRecover),
			"OptDebugSleepOnSwitch":      r.ValueOf(OptDebugSleepOnSwitch),
			"OptDebugger":                r.ValueOf(OptDebugger),
			"OptKeepUntyped":             r.ValueOf(OptKeepUntyped),
			"OptMacroExpandOnly":         r.ValueOf(OptMacroExpandOnly),
			"OptPanicStackTrace":         r.ValueOf(OptPanicStackTrace),
			"OptShowCompile":             r.ValueOf(OptShowCompile),
			"OptShowEval":                r.ValueOf(OptShowEval),
			"OptShowEvalType":            r.ValueOf(OptShowEvalType),
			"OptShowMacroExpand":         r.ValueOf(OptShowMacroExpand),
			"OptShowParse":               r.ValueOf(OptShowParse),
			"OptShowPrompt":              r.ValueOf(OptShowPrompt),
			"OptShowTime":                r.ValueOf(OptShowTime),
			"OptTrapPanic":               r.ValueOf(OptTrapPanic),
			"ParseOptions":               r.ValueOf(ParseOptions),
			"ReadBytes":                  r.ValueOf(ReadBytes),
			"ReadMultiline":              r.ValueOf(ReadMultiline),
			"ReadOptCollectAllComments":  r.ValueOf(ReadOptCollectAllComments),
			"ReadOptShowPrompt":          r.ValueOf(ReadOptShowPrompt),
			"ReadString":                 r.ValueOf(ReadString),
			"SigAll":                     r.ValueOf(SigAll),
			"SigDebug":                   r.ValueOf(SigDebug),
			"SigDefer":                   r.ValueOf(SigDefer),
			"SigInterrupt":               r.ValueOf(SigInterrupt),
			"SigNone":                    r.ValueOf(SigNone),
			"SigReturn":                  r.ValueOf(SigReturn),
			"SimplifyAstForQuote":        r.ValueOf(SimplifyAstForQuote),
			"SimplifyNodeForQuote":       r.ValueOf(SimplifyNodeForQuote),
			"StartSignalHandler":         r.ValueOf(StartSignalHandler),
			"StopSignalHandler":          r.ValueOf(StopSignalHandler),
			"StrGensym":                  r.ValueOf(StrGensym),
			"StrGensymAnonymous":         r.ValueOf(StrGensymAnonymous),
			"StrGensymInterface":         r.ValueOf(StrGensymInterface),
			"StrGensymPrivate":           r.ValueOf(StrGensymPrivate),
			"True":                       r.ValueOf(&True).Elem(),
			"TypeOfBool":                 r.ValueOf(&TypeOfBool).Elem(),
			"TypeOfByte":                 r.ValueOf(&TypeOfByte).Elem(),
			"TypeOfComplex128":           r.ValueOf(&TypeOfComplex128).Elem(),
			"TypeOfComplex64":            r.ValueOf(&TypeOfComplex64).Elem(),
			"TypeOfDeferFunc":            r.ValueOf(&TypeOfDeferFunc).Elem(),
			"TypeOfError":                r.ValueOf(&TypeOfError).Elem(),
			"TypeOfFloat32":              r.ValueOf(&TypeOfFloat32).Elem(),
			"TypeOfFloat64":              r.ValueOf(&TypeOfFloat64).Elem(),
			"TypeOfInt":                  r.ValueOf(&TypeOfInt).Elem(),
			"TypeOfInt16":                r.ValueOf(&TypeOfInt16).Elem(),
			"TypeOfInt32":                r.ValueOf(&TypeOfInt32).Elem(),
			"TypeOfInt64":                r.ValueOf(&TypeOfInt64).Elem(),
			"TypeOfInt8":                 r.ValueOf(&TypeOfInt8).Elem(),
			"TypeOfInterface":            r.ValueOf(&TypeOfInterface).Elem(),
			"TypeOfPtrBool":              r.ValueOf(&TypeOfPtrBool).Elem(),
			"TypeOfPtrComplex128":        r.ValueOf(&TypeOfPtrComplex128).Elem(),
			"TypeOfPtrComplex64":         r.ValueOf(&TypeOfPtrComplex64).Elem(),
			"TypeOfPtrFloat32":           r.ValueOf(&TypeOfPtrFloat32).Elem(),
			"TypeOfPtrFloat64":           r.ValueOf(&TypeOfPtrFloat64).Elem(),
			"TypeOfPtrInt":               r.ValueOf(&TypeOfPtrInt).Elem(),
			"TypeOfPtrInt16":             r.ValueOf(&TypeOfPtrInt16).Elem(),
			"TypeOfPtrInt32":             r.ValueOf(&TypeOfPtrInt32).Elem(),
			"TypeOfPtrInt64":             r.ValueOf(&TypeOfPtrInt64).Elem(),
			"TypeOfPtrInt8":              r.ValueOf(&TypeOfPtrInt8).Elem(),
			"TypeOfPtrString":            r.ValueOf(&TypeOfPtrString).Elem(),
			"TypeOfPtrUint":              r.ValueOf(&TypeOfPtrUint).Elem(),
			"TypeOfPtrUint16":            r.ValueOf(&TypeOfPtrUint16).Elem(),
			"TypeOfPtrUint32":            r.ValueOf(&TypeOfPtrUint32).Elem(),
			"TypeOfPtrUint64":            r.ValueOf(&TypeOfPtrUint64).Elem(),
			"TypeOfPtrUint8":             r.ValueOf(&TypeOfPtrUint8).Elem(),
			"TypeOfPtrUintptr":           r.ValueOf(&TypeOfPtrUintptr).Elem(),
			"TypeOfReflectType":          r.ValueOf(&TypeOfReflectType).Elem(),
			"TypeOfRune":                 r.ValueOf(&TypeOfRune).Elem(),
			"TypeOfString":               r.ValueOf(&TypeOfString).Elem(),
			"TypeOfUint":                 r.ValueOf(&TypeOfUint).Elem(),
			"TypeOfUint16":               r.ValueOf(&TypeOfUint16).Elem(),
			"TypeOfUint32":               r.ValueOf(&TypeOfUint32).Elem(),
			"TypeOfUint64":               r.ValueOf(&TypeOfUint64).Elem(),
			"TypeOfUint8":                r.ValueOf(&TypeOfUint8).Elem(),
			"TypeOfUintptr":              r.ValueOf(&TypeOfUintptr).Elem(),
			"UnwrapTrivialAst":           r.ValueOf(UnwrapTrivialAst),
			"UnwrapTrivialAstKeepBlocks": r.ValueOf(UnwrapTrivialAstKeepBlocks),
			"UnwrapTrivialNode":          r.ValueOf(UnwrapTrivialNode),
			"ZeroStrings":                r.ValueOf(&ZeroStrings).Elem(),
			"ZeroTypes":                  r.ValueOf(&ZeroTypes).Elem(),
			"ZeroValues":                 r.ValueOf(&ZeroValues).Elem(),
		}, Types: map[string]r.Type{
			"BufReadline":      r.TypeOf((*BufReadline)(nil)).Elem(),
			"CmdOpt":           r.TypeOf((*CmdOpt)(nil)).Elem(),
			"Globals":          r.TypeOf((*Globals)(nil)).Elem(),
			"Inspector":        r.TypeOf((*Inspector)(nil)).Elem(),
			"Options":          r.TypeOf((*Options)(nil)).Elem(),
			"Output":           r.TypeOf((*Output)(nil)).Elem(),
			"ReadOptions":      r.TypeOf((*ReadOptions)(nil)).Elem(),
			"Readline":         r.TypeOf((*Readline)(nil)).Elem(),
			"Signal":           r.TypeOf((*Signal)(nil)).Elem(),
			"Signals":          r.TypeOf((*Signals)(nil)).Elem(),
			"TtyReadline":      r.TypeOf((*TtyReadline)(nil)).Elem(),
			"WhichMacroExpand": r.TypeOf((*WhichMacroExpand)(nil)).Elem(),
		}, Proxies: map[string]r.Type{
			"Inspector": r.TypeOf((*P_github_com_muazhari_gomacro_base_Inspector)(nil)).Elem(),
			"Readline":  r.TypeOf((*P_github_com_muazhari_gomacro_base_Readline)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"CmdOptForceEval": "int:2",
			"CmdOptQuit":      "int:1",
		}, Wrappers: map[string][]string{
			"Globals": []string{"Copy", "Debugf", "Error", "ErrorAt", "Errorf", "Fprintf", "IncLine", "IncLineBytes", "MakeRuntimeError", "Position", "Sprintf", "ToString", "WarnExtraValues", "Warnf"},
			"Output":  []string{"Copy", "ErrorAt", "Errorf", "Fprintf", "IncLine", "IncLineBytes", "MakeRuntimeError", "Position", "Sprintf", "ToString"},
		},
	}
}

// --------------- proxy for github.com/muazhari/gomacro-custom/base.Inspector ---------------
type P_github_com_muazhari_gomacro_base_Inspector struct {
	Object   interface{}
	Inspect_ func(_proxy_obj_ interface{}, name string, val reflect.Value, typ reflect.Type, xtyp xreflect.Type, globals *Globals)
}

func (P *P_github_com_muazhari_gomacro_base_Inspector) Inspect(name string, val reflect.Value, typ reflect.Type, xtyp xreflect.Type, globals *Globals) {
	P.Inspect_(P.Object, name, val, typ, xtyp, globals)
}

// --------------- proxy for github.com/muazhari/gomacro-custom/base.Readline ---------------
type P_github_com_muazhari_gomacro_base_Readline struct {
	Object interface{}
	Read_  func(_proxy_obj_ interface{}, prompt string) ([]byte, error)
}

func (P *P_github_com_muazhari_gomacro_base_Readline) Read(prompt string) ([]byte, error) {
	return P.Read_(P.Object, prompt)
}
