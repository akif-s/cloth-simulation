package main

import (
	"image"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
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

	cX := 50
	cY := 30
	cGap := 10
	cloth := newCloth(WINDOW_WIDTH-cX/2*cGap, 100, cX, cY, cGap, false)

	var keyTag struct{}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			dt := time.Since(initTime)
			initTime = time.Now()

			//fmt.Println(dt.Seconds() * 10)

			//Keyboard Inputs
			key.InputOp{
				Tag:  &keyTag,
				Keys: key.NameEscape,
			}.Add(gtx.Ops)

			for _, ev := range gtx.Queue.Events(&keyTag) {
				if e, ok := ev.(key.Event); ok {
					if e.State == key.Press {
						if e.Name == key.NameEscape {
							w.Perform(system.ActionClose)
						}
					}
				}
			}

			// Pointer Inputs
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
						mPos := NewVector2(float64(ev.Position.X), float64(ev.Position.Y))
						for _, p := range cloth.points {
							dx := p.pos.Substract(mPos)

							dst := dx.Magnitude()

							if dst <= 20 {
								p.isActive = false
							}
						}
					}
				}
			}

			//Draw Cloth
			cloth.draw(gtx, dt.Seconds()*15)

			op.InvalidateOp{}.Add(gtx.Ops)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
