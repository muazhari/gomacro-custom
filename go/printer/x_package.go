// this file was generated by gomacro command: import _i "github.com/muazhari/gomacro-custom/go/printer"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package printer

import (
	r "reflect"

	"github.com/muazhari/gomacro-custom/imports"
)

// reflection: allow interpreted code to import "github.com/muazhari/gomacro-custom/go/printer"
func init() {
	imports.Packages["github.com/muazhari/gomacro-custom/go/printer"] = imports.Package{
		Binds: map[string]r.Value{
			"Fprint":    r.ValueOf(Fprint),
			"RawFormat": r.ValueOf(RawFormat),
			"SourcePos": r.ValueOf(SourcePos),
			"TabIndent": r.ValueOf(TabIndent),
			"UseSpaces": r.ValueOf(UseSpaces),
		}, Types: map[string]r.Type{
			"CommentedNode": r.TypeOf((*CommentedNode)(nil)).Elem(),
			"Config":        r.TypeOf((*Config)(nil)).Elem(),
			"Mode":          r.TypeOf((*Mode)(nil)).Elem(),
		},
	}
}
