package main

import (
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

const (
	WINDOW_WIDTH  = 600
	WINDOW_HEIGHT = 400
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

	cloth := newCloth(250, 100, 50, 20, 30, false)

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			dt := time.Since(initTime)
			initTime = time.Now()

			//fmt.Println(dt.Seconds())

			cloth.draw(gtx, dt.Seconds()*10)

			op.InvalidateOp{}.Add(gtx.Ops)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
