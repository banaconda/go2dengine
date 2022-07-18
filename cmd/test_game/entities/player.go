package entities

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/pkg/ecs"
)

type Player struct {
	*ecs.Entity
	*components.CameraComponent
	*components.ControlComponent
	*components.ActionComponent
	*components.CollisionComponent
	*components.TransformComponent
	*components.DrawComponent
}
