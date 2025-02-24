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
 * string.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package strings

import (
	"strconv"
	"strings"

	"github.com/muazhari/gomacro-custom/imports/util"
)

func UnescapeChar(str string) (rune, error) {
	// Debugf("unescapeChar(): parsing CHAR %#v", str)
	n := len(str)
	if n >= 2 && str[0] == '\'' && str[n-1] == '\'' {
		str = str[1 : n-1]
	}
	ret, _, _, err := strconv.UnquoteChar(str, '\'')
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func UnescapeString(str string) string {
	ret, err := strconv.Unquote(str)
	if err != nil {
		panic(err)
	}
	return ret
}

func MaybeUnescapeString(str string) string {
	n := len(str)
	if n >= 2 && (str[0] == '"' || str[0] == '`' || str[0] == '\'') && str[n-1] == str[0] {
		ret, err := strconv.Unquote(str)
		if err != nil {
			panic(err)
		}
		return ret
	}
	return str
}

func FindFirstToken(src []byte) int {
	n := len(src)
	const (
		Normal = iota
		Slash
		LineComment
		MultiLineComment
		MultiLineCommentStar
	)
	mode := Normal
	for i := 0; i < n; i++ {
		ch := src[i]
		switch mode {
		case Normal:
			if ch == '/' {
				mode = Slash
			} else if ch > ' ' {
				return i
			}
		case Slash:
			if ch == '/' {
				mode = LineComment
			} else if ch == '*' {
				mode = MultiLineComment
			} else {
				return i - 1
			}
		case LineComment:
			if ch == '\n' {
				mode = Normal
			}
		case MultiLineComment:
			if ch == '*' {
				mode = MultiLineCommentStar
			}
		case MultiLineCommentStar:
			if ch == '/' {
				mode = Normal
			} else {
				mode = MultiLineComment
			}
		}
	}
	return n
}

// split 's' into a prefix and suffix separated by 'separator'.
// suffix is trimmed with strings.TrimSpace() before returning it
func Split2(s string, separator rune) (string, string) {
	var prefix, suffix string
	if space := strings.IndexByte(s, ' '); space > 0 {
		prefix = s[:space]
		suffix = strings.TrimSpace(s[space+1:])
	} else {
		prefix = s
	}
	return prefix, suffix
}

func TailIdentifier(s string) string {
	return util.TailIdentifier(s)
}

/*
func extractFirstIdentifier(src []byte) []byte {
	n := len(src)
	for i := 0; i < n; i++ {
		ch := src[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') ||
			ch == '_' || ch >= 128 ||
			(i != 0 && (ch >= '0' && ch <= '9')) {
		} else {
			return src[:i]
		}
	}
	return src
}
*/
