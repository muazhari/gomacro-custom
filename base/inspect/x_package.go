// this file was generated by gomacro command: import _i "github.com/muazhari/gomacro-custom/base/inspect"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package inspect

import (
	"github.com/muazhari/gomacro-custom/imports"
	r "reflect"
)

// reflection: allow interpreted code to import "github.com/muazhari/gomacro-custom/base/inspect"
func init() {
	imports.Packages["github.com/muazhari/gomacro-custom/base/inspect"] = imports.Package{
		Types: map[string]r.Type{
			"Inspector": r.TypeOf((*Inspector)(nil)).Elem(),
		},
	}
}
