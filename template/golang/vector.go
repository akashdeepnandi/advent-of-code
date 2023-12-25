package main

import (
	"math"
	"strconv"
)

type Vector struct {
	x, y int
}

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

func (v Vector) Add(o Vector) Vector {
	return NewVector(v.x+o.x, v.y+o.y)
}

func (v Vector) Sub(o Vector) Vector {
	return NewVector(v.x-o.x, v.y-o.y)
}

func (v Vector) Scale(scalar int) Vector {
	return NewVector(v.x*scalar, v.y*scalar)
}

func (v Vector) Dot(o Vector) int {
	return v.x*o.x + v.y + o.y
}

func (v Vector) Mag() int {
	return int(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v Vector) Norm() Vector {
	mag := v.Mag()
	if mag == 0 {
		return Vector{}
	}
	return NewVector(v.x/mag, v.y/mag)
}

func (v Vector) Key() string {
	return strconv.Itoa(v.x) + "," + strconv.Itoa(v.y)
}
