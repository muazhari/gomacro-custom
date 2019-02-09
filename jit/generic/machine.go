/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * machine.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

const SUPPORTED = false
const Name = "generic"

const (
	NoRegId RegId = iota
	RLo           = NoRegId
	RHi           = NoRegId
)

func (r RegId) Valid() bool {
	return false
}

var alwaysLiveRegIds RegIds // empty

type Op0 uint8
type Op1 uint8
type Op2 uint8
type Op3 uint8
type Op4 uint8

func (asm *Asm) Op0(op Op0) *Asm {
	return asm
}

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Load(src Mem, dst Reg) *Asm {
	return asm
}

func (asm *Asm) Store(src Reg, dst Mem) *Asm {
	return asm
}

func (asm *Asm) Prologue() *Asm {
	return asm
}

func (asm *Asm) Epilogue() *Asm {
	return asm
}

func (s *Save) ArchInit(start SaveSlot, end SaveSlot) {
}
