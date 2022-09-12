package PrdRandom

import (
	"math/rand"
	"time"
)

type randomScale struct {
	SuccessScale float32
	TotalScale   float32
}

type PrdRandom struct {
	// 干预概率的种子
	seed float32
	// 成功的占比总数
	successCount int
	// 总数
	totalCount int
}

func New(successCount int, totalCount int) *PrdRandom {
	return &PrdRandom{
		seed:         0,
		successCount: successCount,
		totalCount:   totalCount,
	}
}

// 获取成功或者失败
func (prdRandom *PrdRandom) Get() bool {
	if prdRandom.totalCount == 0 {
		return false
	}

	rand.Seed(time.Now().Unix())

	random := float32(rand.Intn(100))

	isSuccess := prdRandom.bingo(random)

	if isSuccess {
		prdRandom.subSuccessCount()
	} else {
		prdRandom.subTotalCount()
	}

	return isSuccess
}

// 获取成功或者失败
func (prdRandom *PrdRandom) bingo(n float32) bool {
	if n >= 0 && (n <= prdRandom.getRandomScale().SuccessScale) {
		return true
	}
	return false
}

// 获取当前的随机比例
func (prdRandom *PrdRandom) getRandomScale() *randomScale {
	successScale := (float32(prdRandom.successCount) / float32(prdRandom.totalCount) * 100) + prdRandom.seed

	if successScale < 0 {
		successScale = 0
	}
	var totalScale float32 = 100.00
	if totalScale < successScale {
		totalScale = successScale
	}

	return &randomScale{
		SuccessScale: successScale,
		TotalScale:   totalScale,
	}
}

func (prdRandom *PrdRandom) subSuccessCount() {
	prdRandom.successCount -= 1
	prdRandom.calcSeed(true)
}

func (prdRandom *PrdRandom) subTotalCount() {
	prdRandom.totalCount -= 1
	prdRandom.calcSeed(false)
}

func (prdRandom *PrdRandom) calcSeed(success bool) {
	successScale := float32(prdRandom.successCount) / float32(prdRandom.totalCount) * 100
	if success {
		prdRandom.seed -= (100 - successScale)
	} else {
		prdRandom.seed += successScale
	}
}

func (prdRandom *PrdRandom) GetSeed() float32 {
	return prdRandom.seed
}

func (prdRandom *PrdRandom) SetSeed(s float32) {
	prdRandom.seed = s
}
