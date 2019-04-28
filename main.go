package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type scene struct {
	winW, winH     int
	quad           *Quadtree
	isMousePressed bool
}

func (s *scene) run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, float64(s.winW), float64(s.winH)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			mpos := win.MousePosition()
			s.quad.Insert(int(mpos.X), int(mpos.Y))
		}

		s.quad.BreadthFirst(func(q *Quadtree) {
			x0 := float64(q.TopLeft.X)
			y0 := float64(q.TopLeft.Y)
			x1 := float64(q.BotRight.X)
			y1 := float64(q.BotRight.Y)
			imd.Color = colornames.Grey
			imd.Push(pixel.V(x0, y0), pixel.V(x1, y1))
			imd.Rectangle(1.0)

			if q.Node != nil {
				imd.Color = colornames.Red
				imd.Push(pixel.V(float64(q.Node.Point.X), float64(q.Node.Point.Y)))
				imd.Circle(1.0, 2.0)
			}
		})

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	w, h := 800, 800
	s := scene{
		winW: w,
		winH: h,
		quad: NewQuadtree(w, h),
	}
	pixelgl.Run(s.run)
}
