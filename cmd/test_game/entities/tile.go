package entities

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/pkg/ecs"
)

type Tile struct {
	*ecs.Entity
	*components.TransformComponent
	*components.DrawComponent
}
