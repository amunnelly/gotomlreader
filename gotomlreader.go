package main

import ("github.com/amunnelly/gotomlreader/mytomlreader"
"fmt")

function main() {
	test := mytomlreader.Reader("alphabeta.toml")
	fmt.Println(test)
}