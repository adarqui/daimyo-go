package daimyo_random_lcg_bsd

import (
	"github.com/adarqui/daimyo-go/src/random/lcg"
	"math"
)

type LcgBSD struct {
	lcg *daimyo_random_lcg.LcgBase
}

func New(seed int) *LcgBSD {
	lcg := daimyo_random_lcg.New(1103515245, 12345, int(math.Pow(2, 31)), seed)
	return &LcgBSD{lcg: lcg}
}

func (lcg_bsd *LcgBSD) Seed(seed int) {
	lcg_bsd.lcg.Seed(seed)
}

func (lcg_bsd *LcgBSD) Get() int {
	return lcg_bsd.lcg.Get()
}
