package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const PROBABILITY_OF_INTERRUPT = 25

func init() {
	rand.Seed(time.Now().Unix())
}

type device struct {
	interrupt bool
	isr       func() bool
	priority  int
}

func (d *device) draw(screen *ebiten.Image) {
	var fillColor color.Color
	if d.interrupt {
		fillColor = color.RGBA{R: 255, A: 255}
	} else {
		fillColor = color.RGBA{B: 255, A: 255}
	}
	ebitenutil.DrawRect(screen, 1.3*float64((d.priority-1)*DEVICE_WIDTH+50), 50, DEVICE_WIDTH, DEVICE_HEIGHT, fillColor)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Device-%d", d.priority), int(1.3*float64((d.priority-1)*DEVICE_WIDTH+50)), 50)
}

func GenericISR() bool {
	fmt.Println("A interrupt was serviced")
	return true
}

func ProbabilisticInterrupt() bool {
	return rand.Intn(PROBABILITY_OF_INTERRUPT) == 2
}
