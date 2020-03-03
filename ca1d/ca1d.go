
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

func (ca *CA1D) InitSimple(cs CellState) {
	for x := 0; x < ca.Width; x++ {
		ca.first[x] = cs 
		ca.Current[x] = cs
	}
}

func (ca *CA1D) InitRandom(pct float64) { 
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

func (ca *CA1D) InitCenter(cs CellState, pattern string) {

	ca.InitSimple(cs) // TODO Should probably skip if cs == Dead

        length := len(pattern)
	i := 0

	if length <= ca.Width { 
		i = (ca.Width - length) / 2
	}

	for j := 0; i<ca.Width && j<length;i,j=i+1,j+1 { 
		if pattern[j] == '0' {
			ca.first[i] = Dead
			ca.Current[i] = Dead
		} else { 
			ca.first[i] = Live
			ca.Current[i] = Live
		}
	}
}

func (ca *CA1D) InitRepeat(pattern string) {
	length := len(pattern)
	if length <= 0 {
		return // TODO: real handling of this case.
	}

	j := 0
	for i := 0; i < ca.Width; i++ { 
		if pattern[j] == '0' {
			ca.first[i] = Dead
			ca.Current[i] = Dead
		} else { 
			ca.first[i] = Live
			ca.Current[i] = Live
		}
		j++
		j %= length
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


