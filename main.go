// Copyright 2023 The Markov Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// State is a state
type State [5]uint8

func main() {
	var state State
	rnd, states, points := rand.New(rand.NewSource(1)), make(map[State]bool), make(plotter.XYs, 0, 8)
	for i := 0; i < 8*1024*1024; i++ {
		for i := 1; i < len(state); i++ {
			state[i] = state[i-1]
		}
		state[0] = uint8(math.Abs(rnd.NormFloat64() * 64))
		states[state] = true
		points = append(points, plotter.XY{X: float64(i), Y: float64(len(states))})
	}

	p := plot.New()

	p.Title.Text = "epochs vs size"
	p.X.Label.Text = "epochs"
	p.Y.Label.Text = "size"

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Length(1)
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(scatter)

	err = p.Save(8*vg.Inch, 8*vg.Inch, "markov.png")
	if err != nil {
		panic(err)
	}

}
