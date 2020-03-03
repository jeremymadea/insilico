
package ca1d

import (
	"math/rand"
)

type CellState uint8

type CA1D struct {
	Width int
        ruleset uint32 
	first []CellState
        last []CellState
        Current []CellState
}

var Dead CellState = 0
var Live CellState = 1

func NewCA1D(w int, rules uint32) *CA1D { 
	return &CA1D{
		Width: w,
		ruleset: rules,
		first: make([]CellState, w),
		last: make([]CellState, w),
		Current: make([]CellState, w) } 
}

func (ca *CA1D) InitRandomly(pct float64) { 
	for x := 0; x < ca.Width; x++ {
		if rand.Float64() < pct {
			ca.first[x] = Live
			ca.Current[x] = Live
		} else { 
			ca.first[x] = Dead
			ca.Current[x] = Dead
		}
	}
}

func (ca *CA1D) Generate() { 
        var rule uint8
	ca.last, ca.Current = ca.Current, ca.last
        for x := 0; x < ca.Width; x++ { 
		rule = 0
		rule += uint8(ca.last[(x + ca.Width - 2) % ca.Width])
                rule <<= 1
                rule += uint8(ca.last[(x + ca.Width - 1) % ca.Width])
		rule <<= 1
                rule += uint8(ca.last[x])
		rule <<= 1
		rule += uint8(ca.last[(x + 1) % ca.Width])
		rule <<= 1
		rule += uint8(ca.last[(x + 2) % ca.Width])
                ca.Current[x] = CellState((ca.ruleset >> rule) & 1)
	}
}


