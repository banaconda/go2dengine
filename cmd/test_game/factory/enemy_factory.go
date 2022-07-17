package factory

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/entities"
	"go2dengine/pkg/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func CraeteEnemy(target ecs.Identifier, pos sdl.FPoint) entities.Enemy {
	return entities.Enemy{
		Entity:      ecs.NewEntity(),
		AIComponent: &components.AIComponent{Target: target, AIType: components.AI_TYPE_FLEE},
		TransformComponent: &components.TransformComponent{
			Pos:    pos,
			Dim:    sdl.FPoint{X: 100, Y: 100},
			Angle:  0,
			Center: sdl.FPoint{},
			Flip:   0,
			Speed:  0.5,
		},
		DrawComponent: &components.DrawComponent{
			Type:  components.DRAW_TYPE_ENEMY,
			Shape: components.DRAW_SHAPE_RECT,
			R:     255,
			G:     0,
			B:     0,
			A:     255,
		},
	}
}

func CreateTile(pos sdl.FPoint) entities.Tile {
	return entities.Tile{
		Entity: ecs.NewEntity(),
		TransformComponent: &components.TransformComponent{
			Pos:    pos,
			Dim:    sdl.FPoint{X: 100, Y: 100},
			Angle:  0,
			Center: sdl.FPoint{},
			Flip:   0,
			Speed:  0.5,
		},
		DrawComponent: &components.DrawComponent{
			Type:  components.DRAW_TYPE_TILE,
			Shape: components.DRAW_SHAPE_RECT,
			R:     0,
			G:     0,
			B:     255,
			A:     255,
		},
	}
}
