// this file was generated by gomacro command: import _i "github.com/muazhari/gomacro-custom/base/dep"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package dep

import (
	"github.com/muazhari/gomacro-custom/imports"
	r "reflect"
)

// reflection: allow interpreted code to import "github.com/muazhari/gomacro-custom/base/dep"
func init() {
	imports.Packages["github.com/muazhari/gomacro-custom/base/dep"] = imports.Package{
		Binds: map[string]r.Value{
			"Const":           r.ValueOf(Const),
			"DEBUG_GRAPH":     r.ValueOf(DEBUG_GRAPH),
			"Expr":            r.ValueOf(Expr),
			"Func":            r.ValueOf(Func),
			"Import":          r.ValueOf(Import),
			"Macro":           r.ValueOf(Macro),
			"Method":          r.ValueOf(Method),
			"NewDecl":         r.ValueOf(NewDecl),
			"NewDeclExpr":     r.ValueOf(NewDeclExpr),
			"NewDeclFunc":     r.ValueOf(NewDeclFunc),
			"NewDeclImport":   r.ValueOf(NewDeclImport),
			"NewDeclPackage":  r.ValueOf(NewDeclPackage),
			"NewDeclStmt":     r.ValueOf(NewDeclStmt),
			"NewDeclType":     r.ValueOf(NewDeclType),
			"NewDeclVar":      r.ValueOf(NewDeclVar),
			"NewDeclVarMulti": r.ValueOf(NewDeclVarMulti),
			"NewScope":        r.ValueOf(NewScope),
			"NewSorter":       r.ValueOf(NewSorter),
			"Package":         r.ValueOf(Package),
			"Stmt":            r.ValueOf(Stmt),
			"Type":            r.ValueOf(Type),
			"TypeFwd":         r.ValueOf(TypeFwd),
			"Unknown":         r.ValueOf(Unknown),
			"Var":             r.ValueOf(Var),
			"VarMulti":        r.ValueOf(VarMulti),
		}, Types: map[string]r.Type{
			"ConstDeps": r.TypeOf((*ConstDeps)(nil)).Elem(),
			"Decl":      r.TypeOf((*Decl)(nil)).Elem(),
			"DeclList":  r.TypeOf((*DeclList)(nil)).Elem(),
			"DeclMap":   r.TypeOf((*DeclMap)(nil)).Elem(),
			"Extra":     r.TypeOf((*Extra)(nil)).Elem(),
			"Kind":      r.TypeOf((*Kind)(nil)).Elem(),
			"Scope":     r.TypeOf((*Scope)(nil)).Elem(),
			"Sorter":    r.TypeOf((*Sorter)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"DEBUG_GRAPH": "bool:false",
		},
	}
}
