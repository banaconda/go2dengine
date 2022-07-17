package systems

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
)

var Positionable *positionable

type positionEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.ActionComponent
}

type positionable interface {
	ecs.DefaultEntityInterface
	components.TransformInterface
	components.ActionInterface
}

type PositionSystem struct {
	entities []positionEntity
}

func (s *PositionSystem) Add(id ecs.Identifier) {
	obj := id.(positionable)
	s.entities = append(s.entities,
		positionEntity{
			obj.GetEntity(),
			obj.GetTransformComponent(),
			obj.GetActionComponent(),
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
		trans := entity.GetTransformComponent()
		action := entity.GetActionComponent()

		trans.Pos.X += action.CurDir.X
		trans.Pos.Y += action.CurDir.Y

		globals.Logger.Debug("%d, x: %f, y: %f", entity.ID(), trans.Pos.X, trans.Pos.Y)
	}
}
