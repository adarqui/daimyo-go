package daimyo_random_lcg

// Linear Congruential Generator
// https://en.wikipedia.org/wiki/Linear_congruential_generator

type LcgBase struct {
	a  int // multiplier
	c  int // increment
	m  int // modulus
	st int // state
}

type LcgI interface {
	New(int) *LcgI
	Get() int
	Seed(int)
}

func New(a, c, m, st int) *LcgBase {
	lcg := LcgBase{a: a, c: c, m: m, st: st}
	return &lcg
}

func (lcg *LcgBase) Get() int {
	lcg.st = (lcg.a*lcg.st + lcg.c) % lcg.m
	return lcg.st
}

func (lcg *LcgBase) Seed(seed int) {
	lcg.st = seed
}
