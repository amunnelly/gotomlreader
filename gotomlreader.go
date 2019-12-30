package main

import ("github.com/amunnelly/gotomlreader/mytomlreader"
		"fmt")


func main() {
	test := mytomlreader.Reader("alphabeta.toml")
	fmt.Printf("%T\n", test)
}