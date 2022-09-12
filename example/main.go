package main

import (
	"fmt"

	PrdRandom "github.com/a62527776a/go-prd-random"
)

func main() {
	pr := PrdRandom.New(10, 100)

	pr.SetSeedMult(0.5)

	for i := 0; i < 100; i++ {
		pr.Get()
		fmt.Println("总量 ? 成功量 ? 干预概率 ? 成功概率 ?", pr.GetTotalCount(), pr.GetSuccessCount(), pr.GetSeed(), pr.GetSuccessScale()+pr.GetSeed())
	}
}
