package entities

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/pkg/ecs"
)

type Enemy struct {
	*ecs.Entity
	*components.CameraComponent
	*components.AIComponent
	*components.ActionComponent
	*components.CollisionComponent
	*components.TransformComponent
	*components.DrawComponent
}
