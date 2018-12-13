package main

import (
	"github.com/Lebonesco/go-compiler/gen"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	tests := []struct {
		src string
		res string
	}{
		{
			src: ``,
			res: `
				#include <string>
				#include <iostream>
				#include "Builtins.cpp"
				int main() { return 0; }
				`},
		{
			src: `5 + 5;`,
			res: `
			#include <string>
			#include <iostream>
			#include "Builtins.cpp"
			int main() {
			Int tmp_1 = Int(5);
			Int tmp_2 = Int(5);
			Int tmp_3 = tmp_1.PLUS(tmp_2);
			tmp_3;
			return 0;
			}`},
		{
			src: `10 < 4;`,
			res: `
			#include <string>
			#include <iostream>
			#include "Builtins.cpp"
			int main() {
			Int tmp_1 = Int(10);
			Int tmp_2 = Int(4);
			Bool tmp_3 = tmp_1.LT(tmp_2);
			tmp_3;
			return 0;
			}`},
		{
			src: `
				let x = "hello ";
				let y = "world!";
				let z = x + y;
				PRINT(z);
				`,
			res: `
				#include <string>
				#include <iostream>
				#include "Builtins.cpp"
				int main() {
				String tmp_1 = String("hello ");
				String x = tmp_1;
				String tmp_2 = String("world!");
				String y = tmp_2;
				String tmp_3 = x.PLUS(y);
				String z = tmp_3;
				Nothing tmp_4 = z.PRINT();
				tmp_4;
				return 0;
				}
				`},
		{
			src: `
				func add(x Int, y Int) Int {
					return x + y;
				}

				let a = add(1, 3);`,
			res: `
				#include <string>
				#include <iostream>
				#include "Builtins.cpp"
				Int add(Int y, Int x) {
					Int tmp_1 = x.PLUS(y);
					return tmp_1;
				}
				int main() {
				Int tmp_2 = Int(3);
				Int tmp_3 = Int(1);
				Int tmp_4 = add(tmp_2, tmp_3);
				Int a = tmp_4;
				return 0;
				}`},
		{
			src: `
				let x = 0;
				if (true) {
					x = 5;
				} else {
					x = 6;
				}`,
			res: `
				#include <string>
				#include <iostream>
				#include "Builtins.cpp"
				int main() {
				Int tmp_1 = Int(0);
				Int x = tmp_1;
				if("true" == Bool("true").val) {
					Int tmp_2 = Int(5);
					x = tmp_2;
				} else {
					Int tmp_3 = Int(6);
					x = tmp_3;
				}
				return 0;
				}`}}

	for i, test := range tests {
		program := Parse(test.src)
		TypeCheck(program)
		code := gen.GenWrapper(program)
		codeString := code.String()
		// remove spaces for comparison
		for _, rep := range []string{" ", "\n", "\t"} {
			codeString = strings.Replace(codeString, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if codeString != test.res {
			t.Log(codeString)
			t.Fatalf("test [%d] failed", i)
		}
	}
}

func TestOutPut(t *testing.T) {
	tests := []struct {
		src string
		out string
	}{
		{
			src: `let x = 5 + 5;
				  PRINT(x);`,
			out: "10"},
		{
			src: `let x = 10 > 4;
				PRINT(x);`,
			out: "true"},
		{
			src: `
				let x = "hello";
				let y = "world!";
				let z = x + y;
				PRINT(z);
				`,
			out: "helloworld!"},
		{
			src: `
				func add(x Int, y Int) Int {
					return x + y;
				}

				let a = add(1, 3);
				PRINT(a);`,
			out: "4"},
		{
			src: `
				let x = 0;
				if (true) {
					x = 5;
				} else {
					x = 6;
				}`,
			out: ""}}

	for i, test := range tests {
		program := Parse(test.src)
		TypeCheck(program)
		code := gen.GenWrapper(program)
		output := Compile(code)

		for _, rep := range []string{" ", "\n", "\t"} {
			output = strings.Replace(output, rep, "", -1)
		}

		if output != test.out {
			t.Fatalf("test [%d] failed wanted '%s', got='%s'", i, test.out, output)
		}
	}
}
