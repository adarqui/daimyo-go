package daimyo_random_lcg

import (
	"math"
)

func NewBSD(seed int) *LcgBase {
	return New(1103515245, 12345, int(math.Pow(2, 31)), seed)
}

func NewMS(seed int) *LcgBase {
	return New(214013, 2531011, int(math.Pow(2, 31)), seed)
}
