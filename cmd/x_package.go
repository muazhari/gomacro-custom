// this file was generated by gomacro command: import _i "github.com/muazhari/gomacro-custom/cmd"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package cmd

import (
	"github.com/muazhari/gomacro-custom/imports"
	r "reflect"
)

// reflection: allow interpreted code to import "github.com/muazhari/gomacro-custom/cmd"
func init() {
	imports.Packages["github.com/muazhari/gomacro-custom/cmd"] = imports.Package{
		Binds: map[string]r.Value{
			"New": r.ValueOf(New),
		}, Types: map[string]r.Type{
			"Cmd": r.TypeOf((*Cmd)(nil)).Elem(),
		},
	}
}
