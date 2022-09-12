# Go PRD Random

基于Go 参考[PRD伪随机数生成算法](https://liquipedia.net/dota2/Pseudo_Random_Distribution)实现的伪随机分布生成器

稳定的输出符合人直觉的概率分布

增加干预因子 将每次的随机结果参考以往输出。

# Example

``` golang
func main() {
	pr := PrdRandom.New(100, 1000)
    // pr.SetSeed() // 设置干预概率
	pr.SetSeedMult(0.5) // 干预概率比例
	for i := 0; i < 1000; i++ {
		println(pr.Get()) // 生成下一个
		fmt.Println("总量 ? 成功量 ? 干预概率 ? 成功概率 ?", pr.GetTotalCount(), pr.GetSuccessCount(), pr.GetSeed(), pr.GetSuccessScale()+pr.GetSeed())
	}
}
```

# 