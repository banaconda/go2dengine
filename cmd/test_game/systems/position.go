package systems

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
)

var Positionable *positionable

type positionEntity struct {
	*ecs.Entity
	*components.ActionComponent
	*components.TransformComponent
}

type positionable interface {
	ecs.DefaultEntityInterface
	components.ActionInterface
	components.TransformInterface
}

type PositionSystem struct {
	entities []positionEntity
}

func (s *PositionSystem) Add(id ecs.Identifier) {
	obj := id.(positionable)
	s.entities = append(s.entities,
		positionEntity{
			obj.GetEntity(),
			obj.GetActionComponent(),
			obj.GetTransformComponent(),
		})
}

func (s *PositionSystem) Priority() int {
	return globals.POSITON_SYSTEM
}

func (s *PositionSystem) Remove() {
}

func (s *PositionSystem) Update() {
	for _, entity := range s.entities {
		entity.ID()
		action := entity.GetActionComponent()
		trans := entity.GetTransformComponent()

		trans.Rect.X += action.CurDir.X
		trans.Rect.Y += action.CurDir.Y

		globals.Logger.Debug("%d, x: %f, y: %f, w: %f, h: %f",
			entity.ID(), trans.Rect.X, trans.Rect.Y, trans.Rect.W, trans.Rect.H)
	}
}
