package daimyo_random_lcg

import (
	"github.com/adarqui/daimyo-go/src/random/lcg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LCG(t *testing.T) {
	lcg := daimyo_random_lcg.New(1, 2, 3, 0)
	assert.Exactly(t, lcg.Get(), 2)
}

func Test_NewBSD(t *testing.T) {
	lcg := daimyo_random_lcg.NewBSD(5)
	assert.Exactly(t, 1222621274, lcg.Get())
}

func Test_NewMS(t *testing.T) {
	lcg := daimyo_random_lcg.NewMS(5)
	assert.Exactly(t, 3601076, lcg.Get())
}
