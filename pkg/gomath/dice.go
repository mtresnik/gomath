package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/rand"
)

type Rollable interface {
	Roll() int
}

type Die struct {
	Sides int
}

func NewDie(sides ...int) *Die {
	size := 6
	if len(sides) > 0 {
		size = sides[0]
	}
	return &Die{size}
}

func (d *Die) Roll() int {
	return rand.Intn(d.Sides) + 1
}

type Dice struct {
	Dice []*Die
}

func NewDice(num int, size int) *Dice {
	if num <= 0 || size <= 1 {
		return nil
	}
	dice := make([]*Die, num)
	for i := 0; i < num; i++ {
		dice[i] = NewDie(size)
	}
	return &Dice{dice}
}

func (d *Dice) Roll() int {
	return int(goutils.SumBy(d.Dice, func(o *Die) float64 {
		return float64(o.Roll())
	}))
}
