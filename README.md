This is a simple experiment in writing a programming language using go. 
It's not meant to be practical, just to play with some ideas I have.

The goal is a dynamically typed test-oriented language:

	var square = fn(x) {
		return x*x
	}
	square.before = Test(x) {
		if !isInt(x) {
			fail("x must be an integer")
		}
	}
	square.after = Test(x2) {
		var x = square.inputs[0]
		if x*x != x2 {
			fail("x*x != square(x)! x=${0}, square(x)=${1}", x, x2)
		}
	}
However everything is in flux right now as I'm just setting up a simple interpreter.
