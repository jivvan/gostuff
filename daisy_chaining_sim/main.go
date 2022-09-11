package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	DEVICE_WIDTH  = 50
	DEVICE_HEIGHT = 70
)

type Game struct {
	count   int
	devices [5]device
	proc    processor
}

func (g *Game) Update() error {
	if g.count%30 == 0 {
		for i, device := range g.devices {
			if !device.interrupt {
				// randomly genarate interrupts if a device does not have interrupts
				g.devices[i].interrupt = ProbabilisticInterrupt()
				g.proc.generateInterrupt()
			}
		}
		//if processor is interrupted, send acknowledge
		if g.proc.interrupt {
			g.proc.acknowledgeInterrupt()
		}
		// if acknowledge is sent by processor, higest priority device sends the isr to processor to be executed
		if !g.proc.isBusy {
			for i, device := range g.devices {
				if device.interrupt {
					g.proc.runIsr(device.isr, &g.devices[i])
					break
				}
			}
		}
	}
	g.count++
	if g.count == 60 {
		g.count = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Simulating Daisy Chaining")
	for _, device := range g.devices {
		device.draw(screen)
	}
	g.proc.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowTitle("Daisy Chaining Simulation")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := &Game{}
	for i := range game.devices {
		game.devices[i].priority = i + 1
		game.devices[i].isr = GenericISR
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
