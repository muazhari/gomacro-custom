/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * expr_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"fmt"
	r "reflect"
	"testing"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/classic"
	bi "github.com/cosmos72/gomacro/experiments/bytecode_interfaces"
	bv "github.com/cosmos72/gomacro/experiments/bytecode_values"
	ci "github.com/cosmos72/gomacro/experiments/closure_interfaces"
	cm "github.com/cosmos72/gomacro/experiments/closure_maps"
	cv "github.com/cosmos72/gomacro/experiments/closure_values"
	"github.com/cosmos72/gomacro/fast"
)

const (
	collatz_n   = 837799 // sequence climbs to 1487492288, which also fits 32-bit ints
	sum_n       = 1000
	fib_n       = 12
	bigswitch_n = 100
)

var verbose = false

/*
	--------- 2016-05-06: results on Intel Core i7 4770 ---------------

	BenchmarkFibonacciCompiler-8                 	 3000000	       498 ns/op
	BenchmarkFibonacciFastInterpreter-8          	  100000	     14812 ns/op
	BenchmarkFibonacciFastInterpreterBis-8       	  100000	     14446 ns/op
	BenchmarkFibonacciClassicInterpreter-8       	    3000	    575222 ns/op
	BenchmarkFibonacciClassicInterpreterBis-8    	    3000	    575585 ns/op
	BenchmarkFibonacciClosureValues-8            	   10000	    239373 ns/op
	BenchmarkFibonacciClosureInterfaces-8        	   10000	    184985 ns/op
	BenchmarkFibonacciClosureMaps-8              	    5000	    330350 ns/op
	BenchmarkArithCompiler1-8                    	200000000	         8.58 ns/op
	BenchmarkArithCompiler2-8                    	200000000	         8.56 ns/op
	BenchmarkArithFastInterpreter-8              	30000000	        59.2 ns/op
	BenchmarkArithFastInterpreterBis-8           	30000000	        51.4 ns/op
	BenchmarkArithFastInterpreterConst-8         	100000000	        13.9 ns/op
	BenchmarkArithFastInterpreterCompileLoop-8   	  100000	     21307 ns/op
	BenchmarkArithClassicInterpreter-8           	 1000000	      1466 ns/op
	BenchmarkArithClassicInterpreterBis-8        	 1000000	      2378 ns/op
	BenchmarkCollatzCompiler-8                   	 3000000	       426 ns/op
	BenchmarkCollatzFastInterpreter-8            	  100000	     12460 ns/op
	BenchmarkCollatzClassicInterpreter-8         	    3000	    479500 ns/op
	BenchmarkCollatzBytecodeInterfaces-8         	   50000	     29575 ns/op
	BenchmarkCollatzClosureValues-8              	  100000	     16678 ns/op
	BenchmarkSumCompiler-8                       	 3000000	       413 ns/op
	BenchmarkSumFastInterpreter-8                	  100000	     20292 ns/op
	BenchmarkSumFastInterpreterBis-8             	  100000	     20330 ns/op
	BenchmarkSumClassicInterpreter-8             	    2000	    904097 ns/op
	BenchmarkSumBytecodeValues-8                 	   20000	     72740 ns/op
	BenchmarkSumBytecodeInterfaces-8             	   30000	     52509 ns/op
	BenchmarkSumClosureValues-8                  	   30000	     41459 ns/op
	BenchmarkSumClosureInterfaces-8              	   10000	    142466 ns/op
	BenchmarkSumClosureMaps-8                    	   20000	     93106 ns/op
*/

// recursion: fibonacci. fib(n) => if (n <= 2) { return 1 } else { return fib(n-1) + fib(n-2) }

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func BenchmarkFibonacciCompiler(b *testing.B) {
	var total int
	n := fib_n
	for i := 0; i < b.N; i++ {
		total += fibonacci(n)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkFibonacciFastInterpreter(b *testing.B) {
	ce := fast.New()
	c := ce.Comp
	ce.Eval(fibonacci_source_string)

	// compile the call to fibonacci(fib_n)
	fun := c.CompileAst(c.ParseAst(fmt.Sprintf("fibonacci(%d)", fib_n)))
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		retv, _ := fun(env)
		total += int(retv.Int())
	}
}

func BenchmarkFibonacciFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	ce.Eval(fibonacci_source_string)

	// alternative: extract the function fibonacci, and call it ourselves
	//
	// ValueOf is the method to retrieve constants, functions and variables from the classic and fast interpreters
	// (if you set the same interpreter variable repeatedly, use the address returned by AddressOfVar)
	fun := ce.ValueOf("fibonacci").Interface().(func(int) int)
	fun(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_n)
	}
}

func BenchmarkFibonacciClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(fibonacci_source_string))

	// compile the call to fibonacci(fib_n)
	form := env.ParseAst(fmt.Sprintf("fibonacci(%d)", fib_n))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
	}
}

func BenchmarkFibonacciClassicInterpreterBis(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(fibonacci_source_string))

	// alternative: extract the function fibonacci, and call it ourselves
	fun := env.ValueOf("fibonacci").Interface().(func(int) int)
	fun(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_n)
	}
}

func BenchmarkFibonacciClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	fib := cv.DeclFibonacci(env, 0)
	n := r.ValueOf(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureInterfaces(b *testing.B) {
	env := ci.NewEnv(nil)
	fib := ci.DeclFibonacci(env, 0)
	var n interface{} = fib_n

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureMaps(b *testing.B) {
	env := cm.NewEnv(nil)
	fib := cm.DeclFibonacci(env, "fib")
	n := r.ValueOf(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

// ---------------------- bigswitch ------------------------

func bigswitch(n int) int {
	for i := 0; i < 1000; i++ {
		switch n & 15 {
		case 0:
			n++
		case 1:
			n += 2
		case 2:
			n += 3
		case 3:
			n += 4
		case 4:
			n += 5
		case 5:
			n += 6
		case 6:
			n += 7
		case 7:
			n += 8
		case 8:
			n += 9
		case 9:
			n += 10
		case 10:
			n += 11
		case 11:
			n += 12
		case 12:
			n += 13
		case 13:
			n += 14
		case 14:
			n += 15
		case 15:
			n--
		}
	}
	return n
}

func BenchmarkBigSwitchCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += bigswitch(bigswitch_n)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkBigSwitchFastInterpreter(b *testing.B) {
	ce := fast.New()
	ce.Eval(bigswitch_source_string)

	fun := ce.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_n)
	}
}

func BenchmarkBigSwitchClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(bigswitch_source_string))

	fun := env.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_n)
	}
}

func arith(n int) int {
	return ((n*2+3)&4 | 5 ^ 6) / (n | 1)
}

func BenchmarkArithCompiler1(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		n := b.N
		total += ((n*2+3)&4 | 5 ^ 6) / (n | 1)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithCompiler2(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += arith(b.N)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFastInterpreter(b *testing.B) {
	ce := fast.New()
	ce.DeclVar("n", nil, int(0))

	addr := ce.AddressOfVar("n").Interface().(*int)

	fun := ce.Compile("((n*2+3)&4 | 5 ^ 6) / (n|1)")
	env := ce.PrepareEnv()
	fun(env)
	var ret r.Value

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		*addr = b.N
		ret, _ = fun(env)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, n, total int")

	n := ce.AddressOfVar("n").Interface().(*int)
	total := ce.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ce.Compile("for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }")
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*n = b.N
	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastInterpreterConst(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total int")
	// "cheat" a bit and declare n as a constant. checks if constant propagation works :)
	ce.DeclConst("n", nil, int(b.N))
	total := ce.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ce.Compile("for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }")
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastInterpreterCompileLoop(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, n, total int")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ce.Compile("total = 0; for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")
	}
}

func BenchmarkArithClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("n:=0"))

	form := env.ParseAst("((n*2+3)&4 | 5 ^ 6) / (n|1)")

	value := env.ValueOf("n")
	var ret r.Value
	env.EvalAst(form)

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		value.SetInt(int64(b.N))
		ret, _ = env.EvalAst(form)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithClassicInterpreterBis(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.ParseAst("var n, total int"))

	// interpreted code performs iteration and arithmetic
	form := ir.ParseAst("total = 0; for i:= 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")

	value := ir.ValueOf("n")
	ir.EvalAst(form)

	b.ResetTimer()
	value.SetInt(int64(b.N))
	ret, _ := ir.EvalAst(form)

	if verbose {
		println(ret.Int())
	}
}

// collatz conjecture

func collatz(n int) {
	for n > 1 {
		if n&1 != 0 {
			n = ((n * 3) + 1) / 2
		} else {
			n = n / 2
		}
	}
}

func BenchmarkCollatzCompiler(b *testing.B) {
	n := collatz_n
	for i := 0; i < b.N; i++ {
		collatz(n)
	}
}

func BenchmarkCollatzFastInterpreter(b *testing.B) {
	ce := fast.New()
	c := ce.Comp
	ce.DeclVar("n", TypeOfInt, 0)

	addr := ce.AddressOfVar("n").Interface().(*int)

	fun := c.CompileAst(c.ParseAst("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } }"))
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_n
		fun(env)
	}
}

func BenchmarkCollatzClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("var n int"))
	n := env.ValueOf("n")

	form := env.ParseAst("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } }")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.SetInt(collatz_n)
		env.EvalAst(form)
	}
}

func BenchmarkCollatzBytecodeInterfaces(b *testing.B) {
	coll := bi.BytecodeCollatz()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Vars[0] = collatz_n
		coll.Exec(0)
	}
}

func BenchmarkCollatzClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	coll := cv.DeclCollatz(env, 0)
	n := r.ValueOf(collatz_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += coll(n)
	}
}

// looping: sum the integers from 1 to N

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func BenchmarkSumCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_n)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFastInterpreter(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))

	fun := ce.Compile("total = 0; for i = 1; i <= n; i++ { total += i }; total")
	env := ce.PrepareEnv()
	fun(env)

	var total uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret, _ := fun(env)
		total += uint(ret.Uint())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))

	fun := ce.Compile("for i = 1; i <= n; i++ { total += i }")
	env := ce.PrepareEnv()
	fun(env)
	total := ce.AddressOfVar("total").Interface().(*uint)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*total = 0
		fun(env)
	}
	if verbose {
		println(*total)
	}
}

func BenchmarkSumClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("var i, n, total int"))
	env.ValueOf("n").SetInt(sum_n)
	form := env.ParseAst("total = 0; for i = 1; i <= n; i++ { total += i }; total")

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
	}
}

func BenchmarkSumBytecodeValues(b *testing.B) {
	sum := bv.BytecodeSum(sum_n)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(sum.Exec(0)[0].Int())
	}
}

func BenchmarkSumBytecodeInterfaces(b *testing.B) {
	p := bi.BytecodeSum(sum_n)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += p.Exec(0)[0].(int)
	}
}

func BenchmarkSumClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	sum := cv.DeclSum(env, 0)
	n := r.ValueOf(sum_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureInterfaces(b *testing.B) {
	env := ci.NewEnv(nil)
	sum := ci.DeclSum(env, 0)
	var n interface{} = sum_n

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureMaps(b *testing.B) {
	env := cm.NewEnv(nil)
	sum := cm.DeclSum(env, "sum")
	n := r.ValueOf(sum_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}
