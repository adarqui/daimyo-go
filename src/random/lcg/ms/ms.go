package daimyo_random_lcg_ms

import (
	"github.com/adarqui/daimyo-go/src/random/lcg"
	"math"
)

type LcgMS struct {
	lcg *daimyo_random_lcg.LcgBase
}

func New(seed int) *LcgMS {
	lcg := daimyo_random_lcg.New(214013, 2531011, int(math.Pow(2, 31)), seed)
	return &LcgMS{lcg: lcg}
}

func (lcg_ms *LcgMS) Seed(seed int) {
	lcg_ms.lcg.Seed(seed)
}

func (lcg_ms *LcgMS) Get() int {
	return lcg_ms.lcg.Get()
}
