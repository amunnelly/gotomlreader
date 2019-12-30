package mytomlreader

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type alphabeta struct {
	alpha int
	beta int
}

// Reader reads toml data and returns it as a string
func Reader(uri string) string {
	var details alphabeta

	toml.DecodeFile(uri, &details)

	return fmt.Sprintf(details)
}