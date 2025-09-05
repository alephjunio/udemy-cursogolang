package main

import "runtime/debug"

func function3() {
	debug.PrintStack()
}

func function2() {
	function3()
}

func function1() {
	function2()
}

func main() {
	function1()
}
