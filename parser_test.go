package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	const input = `
			/* comment should not be scanned */
			func add() Int {
				return x + y;
			}

			let five = "test";
			let ten = 10;

			let result = 4;  
			5 < 10;

			if (5 < 10) {
				print(10);
			} else {
				print(5);
			}

			10 == 10;
			`

	Parse(input)
}
