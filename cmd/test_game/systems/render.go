package systems

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

var Drawable *drawable

type renderEntity struct {
	*ecs.Entity
	*components.DrawComponent
	*components.TransformComponent
}

type RenderSystem struct {
	entities [][]renderEntity
	Renderer *sdl.Renderer
}

type drawable interface {
	ecs.DefaultEntityInterface
	components.TransformInterface
	components.DrawInterface
}

func (s *RenderSystem) Add(id ecs.Identifier) {
	obj := id.(drawable)
	t := obj.GetDrawComponent().Type
	if s.entities == nil {
		s.entities = make([][]renderEntity, components.DRAW_TYPE_MAX)
	}

	s.entities[t] = append(s.entities[t], renderEntity{obj.GetEntity(), obj.GetDrawComponent(), obj.GetTransformComponent()})
}

func (s *RenderSystem) Priority() int {
	return globals.RENDER_SYSTEM
}

func (s *RenderSystem) Remove() {
}

func (s *RenderSystem) Update() {
	for _, entities := range s.entities {
		for _, entity := range entities {
			entity.ID()
			trans := entity.GetTransformComponent()
			vec := &trans.Pos

			draw := entity.GetDrawComponent()

			s.RenderDraw(draw, trans)
			globals.Logger.Debug("%d, x: %f, y: %f", entity.ID(), vec.X, vec.Y)
		}
	}
}

func (s *RenderSystem) RenderDraw(d *components.DrawComponent, t *components.TransformComponent) {
	if err := s.Renderer.SetDrawColor(d.R, d.G, d.B, d.A); err != nil {
		globals.Logger.Info("draw error: %v", err)
	}

	switch d.Shape {
	case components.DRAW_SHAPE_RECT:
		rect := sdl.FRect{
			X: t.Pos.X,
			Y: t.Pos.Y,
			W: t.Dim.X,
			H: t.Dim.Y,
		}
		s.Renderer.DrawRectF(&rect)

	case components.DRAW_SHAPE_CIRCLE:
	}
}
