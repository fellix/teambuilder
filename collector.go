package main

import (
	"./collectors"
	"flag"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var moves = flag.Bool("moves", false, "export moves from gen I up to III in the files data/movesGEN.json (e.g. data/moves1.json)")
	var pokemon = flag.Bool("pokemon", false, "download from the available pokemon all the data, each pokemon in his respective file")

	flag.Parse()

	if *moves {
		fmt.Println("Dowloading moves data")
		collectors.CollectMoves()
	}

	if *pokemon {
		fmt.Println("Downloading pokemon data")
		collectors.CollectPokemons()
	}
}
