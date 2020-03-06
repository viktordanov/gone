package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	p := NewPerceptron(0.1, 20)

	// Training data
	pts := []Point{}
	for i := 0; i < 100; i++ {
		pts = append(pts, *NewPoint())
	}

	p.Train(pts)
}