package systems

import (
	"go2dengine/cmd/test_game/algorithm"
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
	"math"

	"github.com/veandco/go-sdl2/sdl"
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

		a := sdl.FRect{X: trans.Pos.X, Y: trans.Pos.Y, W: trans.Dim.X, H: trans.Dim.Y}
		b := sdl.FRect{X: targetTrans.Pos.X, Y: targetTrans.Pos.Y, W: targetTrans.Dim.X, H: targetTrans.Dim.Y}
		distance := math.Sqrt(float64(x*x + y*y))
		switch ai.AIType {
		case components.AI_TYPE_CHASE:
			if algorithm.AABB(a, b) {
				globals.Logger.Info("collision")
			} else {
				if distance > float64(targetTrans.Dim.X) {
					trans.Pos.X -= x / float32(distance)
					trans.Pos.Y -= y / float32(distance)
				}
			}
		case components.AI_TYPE_FLEE:
			globals.Logger.Info("distance: %f", distance)
			if distance < float64(targetTrans.Dim.X*2) {
				trans.Pos.X += x / float32(distance) * trans.Speed
				trans.Pos.Y += y / float32(distance) * trans.Speed
			}
		}

		globals.Logger.Debug("%d, x: %f, y: %f", entity.ID(), trans.Pos.X, trans.Pos.Y)
	}
}
