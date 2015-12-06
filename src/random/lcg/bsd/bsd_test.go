package daimyo_random_lcg_bsd

import (
	"github.com/adarqui/daimyo-go/src/random/lcg/bsd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LcgBSD(t *testing.T) {
	lcg := daimyo_random_lcg_bsd.New(5)
	assert.Exactly(t, 1222621274, lcg.Get())
}
