package components

import (
	"go2dengine/cmd/test_game/globals"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DRAW_SHAPE_RECT = iota
	DRAW_SHAPE_CIRCLE
)

const (
	DRAW_TYPE_TILE = iota
	DRAW_TYPE_OBJ
	DRAW_TYPE_ENEMY
	DRAW_TYPE_PLAYER
	DRAW_TYPE_EFFECT
)

type DrawComponent struct {
	Type       uint32
	Shape      uint32
	R, G, B, A uint8
}

type DrawInterface interface {
	GetDrawComponent() *DrawComponent
	Render(r *sdl.Renderer, t *TransformComponent)
}

func (c *DrawComponent) GetDrawComponent() *DrawComponent {
	return c
}

func (c *DrawComponent) Render(r *sdl.Renderer, t *TransformComponent) {
	if err := r.SetDrawColor(255, 0, 0, 255); err != nil {
		globals.Logger.Info("draw error: %v", err)
	}

	switch c.Shape {
	case DRAW_SHAPE_RECT:
		rect := sdl.FRect{
			X: t.Pos.X,
			Y: t.Pos.Y,
			W: t.Dim.X,
			H: t.Dim.Y,
		}
		r.DrawRectF(&rect)
	case DRAW_SHAPE_CIRCLE:

	}
}
