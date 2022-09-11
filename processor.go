package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type processor struct {
	interrupt           bool
	interrupAcknowledge bool
	isBusy              bool
	servicingDevice     int
}

func (p *processor) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 150, 250, 200, 100, color.RGBA{B: 192, R: 200, A: 255})
	ebitenutil.DebugPrintAt(screen, "Processor", 150, 250)
	if p.isBusy {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Servicing device = %d", p.servicingDevice), 150, 280)
	} else {
		ebitenutil.DebugPrintAt(screen, "Processor is doing other stuff!", 150, 280)
	}
}

func (p *processor) runIsr(isr func() bool, dev *device) {
	p.servicingDevice = dev.priority
	p.isBusy = true
	go func() {
		time.Sleep(time.Second * 3)
		isr()
		p.interrupAcknowledge = false
		dev.interrupt = false
		p.isBusy = false
	}()
}

func (p *processor) generateInterrupt() {
	p.interrupt = true
}

func (p *processor) acknowledgeInterrupt() {
	p.interrupAcknowledge = true
}
