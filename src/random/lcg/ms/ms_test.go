package daimyo_random_lcg_ms

import (
	"github.com/adarqui/daimyo-go/src/random/lcg/ms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LcgMS(t *testing.T) {
	lcg := daimyo_random_lcg_ms.New(5)
	assert.Exactly(t, 3601076, lcg.Get())
}
