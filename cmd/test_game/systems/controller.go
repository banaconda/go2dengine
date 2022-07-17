package systems

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var Controllable *controllable

type controlEntity struct {
	*ecs.Entity
	*components.ActionComponent
	*components.TransformComponent
}

type controllable interface {
	ecs.DefaultEntityInterface
	components.ActionInterface
	components.TransformInterface
}

type ControlSystem struct {
	entities []controlEntity
}

func (s *ControlSystem) Add(id ecs.Identifier) {
	obj := id.(controllable)
	s.entities = append(s.entities, controlEntity{obj.GetEntity(), obj.GetActionComponent(), obj.GetTransformComponent()})
}

func (s *ControlSystem) Priority() int {
	return globals.CONTROL_SYSTEM
}

func (s *ControlSystem) Remove() {
}

func (s *ControlSystem) Update() {
	if event := sdl.PollEvent(); event != nil {
		switch event.(type) {
		case *sdl.QuitEvent:
			globals.IsRunning = false
			globals.Logger.Info("quit game")
		}
	}

	state := sdl.GetKeyboardState()
	curDir := sdl.FPoint{X: 0, Y: 0}
	if state[sdl.SCANCODE_LEFT] == 1 {
		curDir.X += -1
	}

	if state[sdl.SCANCODE_RIGHT] == 1 {
		curDir.X += 1
	}

	if state[sdl.SCANCODE_UP] == 1 {
		curDir.Y += -1
	}

	if state[sdl.SCANCODE_DOWN] == 1 {
		curDir.Y += 1
	}

	for _, entity := range s.entities {
		entity.ID()
		action := entity.GetActionComponent()

		// TODO: change to check collision
		if curDir.X != 0 || curDir.Y != 0 {
			distance := math.Sqrt(float64(curDir.X*curDir.X + curDir.Y*curDir.Y))

			trans := entity.GetTransformComponent()
			trans.Pos.X += curDir.X / float32(distance) * trans.Speed
			trans.Pos.Y += curDir.Y / float32(distance) * trans.Speed
		}

		globals.Logger.Debug("%d, x: %f, y: %f", entity.ID(), action.CurDir.X, action.CurDir.Y)
	}
}
