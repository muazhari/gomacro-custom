// this file was generated by gomacro command: import _b "net/netip"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	netip "net/netip"
)

// reflection: allow interpreted code to import "net/netip"
func init() {
	Packages["net/netip"] = Package{
		Name: "netip",
		Binds: map[string]Value{
			"AddrFrom16":	ValueOf(netip.AddrFrom16),
			"AddrFrom4":	ValueOf(netip.AddrFrom4),
			"AddrFromSlice":	ValueOf(netip.AddrFromSlice),
			"AddrPortFrom":	ValueOf(netip.AddrPortFrom),
			"IPv4Unspecified":	ValueOf(netip.IPv4Unspecified),
			"IPv6LinkLocalAllNodes":	ValueOf(netip.IPv6LinkLocalAllNodes),
			"IPv6Unspecified":	ValueOf(netip.IPv6Unspecified),
			"MustParseAddr":	ValueOf(netip.MustParseAddr),
			"MustParseAddrPort":	ValueOf(netip.MustParseAddrPort),
			"MustParsePrefix":	ValueOf(netip.MustParsePrefix),
			"ParseAddr":	ValueOf(netip.ParseAddr),
			"ParseAddrPort":	ValueOf(netip.ParseAddrPort),
			"ParsePrefix":	ValueOf(netip.ParsePrefix),
			"PrefixFrom":	ValueOf(netip.PrefixFrom),
		}, Types: map[string]Type{
			"Addr":	TypeOf((*netip.Addr)(nil)).Elem(),
			"AddrPort":	TypeOf((*netip.AddrPort)(nil)).Elem(),
			"Prefix":	TypeOf((*netip.Prefix)(nil)).Elem(),
		}, 
	}
}
