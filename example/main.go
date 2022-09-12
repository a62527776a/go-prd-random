package main

import (
	PrdRandom "github.com/a62527776a/go-prd-random"
)

func main() {
	pr := PrdRandom.New(20, 1000)

	for i := 0; i < 100; i++ {
		println(pr.Get())
		println(pr.GetSeed())
	}
}
