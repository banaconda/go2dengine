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
	entities []renderEntity
	Renderer *sdl.Renderer
}

type drawable interface {
	ecs.DefaultEntityInterface
	components.TransformInterface
	components.DrawInterface
}

func (s *RenderSystem) Add(id ecs.Identifier) {
	obj := id.(drawable)
	s.entities = append(s.entities, renderEntity{obj.GetEntity(), obj.GetDrawComponent(), obj.GetTransformComponent()})
}

func (s *RenderSystem) Priority() int {
	return globals.RENDER_SYSTEM
}

func (s *RenderSystem) Remove() {
}

func (s *RenderSystem) Update() {
	// globals.Logger.Info("pos system update")
	for _, entity := range s.entities {
		entity.ID()
		vec := entity.GetTransformComponent().Pos
		draw := entity.GetDrawComponent()

		draw.Render(s.Renderer, entity.TransformComponent)
		globals.Logger.Debug("%d, x: %f, y: %f", entity.ID(), vec.X, vec.Y)
	}
}
