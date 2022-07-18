package systems

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
	"math"
)

var AIable *aiable

type aiEntity struct {
	*ecs.Entity
	*components.AIComponent
	*components.ActionComponent
	*components.TransformComponent
}

type aiable interface {
	ecs.DefaultEntityInterface
	components.AIInterface
	components.ActionInterface
	components.TransformInterface
}

type AISystem struct {
	entities []aiEntity
}

func (s *AISystem) Add(id ecs.Identifier) {
	obj := id.(aiable)
	s.entities = append(s.entities,
		aiEntity{
			obj.GetEntity(),
			obj.GetAIComponent(),
			obj.GetActionComponent(),
			obj.GetTransformComponent(),
		})
}

func (s *AISystem) Priority() int {
	return globals.AI_SYSTEM
}

func (s *AISystem) Remove() {
}

func (s *AISystem) Update() {
	for _, entity := range s.entities {
		entity.ID()
		ai := entity.GetAIComponent()

		target := ai.Target.(positionable)
		targetTrans := target.GetTransformComponent()

		trans := entity.GetTransformComponent()

		x := trans.Rect.X - targetTrans.Rect.X
		y := trans.Rect.Y - targetTrans.Rect.Y

		distance := math.Sqrt(float64(x*x + y*y))

		action := entity.GetActionComponent()
		action.LastDir = action.CurDir
		switch ai.AIType {
		case components.AI_TYPE_CHASE:
			action.CurDir.X = -x / float32(distance)
			action.CurDir.Y = -y / float32(distance)
		case components.AI_TYPE_FLEE:
			action.CurDir.X = +x / float32(distance)
			action.CurDir.Y = +y / float32(distance)
		}

		globals.Logger.Debug("%d, x: %f, y: %f, w: %f, h: %f",
			entity.ID(), trans.Rect.X, trans.Rect.Y, trans.Rect.W, trans.Rect.H)
	}
}
