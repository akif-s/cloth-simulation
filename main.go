package main

import (
	"image"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 500
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Cloth"),
			app.Size(unit.Dp(WINDOW_WIDTH), unit.Dp(WINDOW_HEIGHT)),
		)

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var initTime = time.Now()

func loop(w *app.Window) error {
	var ops op.Ops

	cX := 80
	cY := 20
	cGap := 5
	cloth := newCloth(WINDOW_WIDTH-cX/2*cGap, 100, cX, cY, cGap, false)

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			dt := time.Since(initTime)
			initTime = time.Now()

			//fmt.Println(dt.Seconds() * 10)

			// Pointer Events
			pointer.InputOp{
				Tag:   w,
				Types: pointer.Drag,
				ScrollBounds: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: -30,
					},
					Max: image.Point{
						X: 0,
						Y: 30,
					},
				},
			}.Add(gtx.Ops)

			for _, ev := range gtx.Queue.Events(w) {
				switch ev := ev.(type) {
				case pointer.Event:
					switch ev.Type {
					case pointer.Drag:
						mPos := ev.Position
						for _, p := range cloth.points {
							dx := p.x - float64(mPos.X)
							dy := p.y - float64(mPos.Y)

							dst := math.Sqrt(dx*dx + dy*dy)

							if dst <= 40 {
								p.isActive = false
							}
						}
					}
				}
			}

			//Draw Cloth
			cloth.draw(gtx, dt.Seconds()*10)

			op.InvalidateOp{}.Add(gtx.Ops)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
