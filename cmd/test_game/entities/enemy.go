package entities

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/pkg/ecs"
)

type Enemy struct {
	*ecs.Entity
	*components.AIComponent
	*components.TransformComponent
	*components.DrawComponent
}
