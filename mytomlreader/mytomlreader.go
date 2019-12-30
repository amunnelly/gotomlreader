package mytomlreader

import "github.com/BurntSushi/toml"

// Alphabeta stores .toml data
type Alphabeta struct {
	Alpha int `toml "alpha"`
	Beta int `toml "beta"`
}

// Reader reads toml data and returns it as a string
func Reader(uri string) Alphabeta {
	var details Alphabeta

	toml.DecodeFile(uri, &details)

	return details
}