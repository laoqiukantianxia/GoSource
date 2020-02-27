package main

import "flag"

func flag1() {
	var n = flag.Bool("n", false, "default false")
	var name = flag.String("name", "", "input you name")

	flag.Parse()
	println(n, name)
}
