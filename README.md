# Go PRD Random

基于Go 参考[PRD伪随机数生成算法](https://liquipedia.net/dota2/Pseudo_Random_Distribution)实现的伪随机分布生成器

## example

``` golang
func main() {
	pr := PrdRandom.New(20, 1000)
    // pr.SetSeed() // 设置干预概率

	for i := 0; i < 100; i++ {
		println(pr.Get()) // 生成下一个
		println(pr.GetSeed()) // seed为当前干预的概率
	}
}
```