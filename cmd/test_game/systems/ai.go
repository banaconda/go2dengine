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
	*components.TransformComponent
}

type aiable interface {
	ecs.DefaultEntityInterface
	components.AIInterface
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

		x := trans.Pos.X - targetTrans.Pos.X
		y := trans.Pos.Y - targetTrans.Pos.Y

		// TODO: change to check collision
		if x != 0 || y != 0 {
			distance := math.Sqrt(float64(x*x + y*y))

			switch ai.AIType {
			case components.AI_TYPE_CHASER:
				if distance > float64(targetTrans.Dim.X) {
					trans.Pos.X -= x / float32(distance)
					trans.Pos.Y -= y / float32(distance)
				}
			case components.AI_TYPE_FLEE:
				globals.Logger.Info("distance: %f", distance)
				if distance < float64(targetTrans.Dim.X*2) {
					trans.Pos.X += x / float32(distance) * .5
					trans.Pos.Y += y / float32(distance) * .5
				}
			}
		}

		globals.Logger.Info("%d, x: %f, y: %f", entity.ID(), trans.Pos.X, trans.Pos.Y)
	}
}
