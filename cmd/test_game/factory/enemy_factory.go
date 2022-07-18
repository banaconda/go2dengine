package factory

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/entities"
	"go2dengine/pkg/ecs"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func CraeteEnemy(target ecs.Identifier, rect sdl.FRect) entities.Enemy {
	return entities.Enemy{
		Entity:             ecs.NewEntity(),
		CameraComponent:    &components.CameraComponent{},
		AIComponent:        &components.AIComponent{Target: target, AIType: rand.Uint32() % components.AI_TYPE_MAX},
		ActionComponent:    &components.ActionComponent{},
		CollisionComponent: &components.CollisionComponent{Rect: rect, Enable: true},
		TransformComponent: &components.TransformComponent{Rect: sdl.FRect{X: 500, Y: 500, W: 100, H: 100}, Angle: 0, Center: sdl.FPoint{}, Flip: 0, Speed: 0.5},
		DrawComponent:      &components.DrawComponent{Type: components.DRAW_TYPE_ENEMY, Shape: components.DRAW_SHAPE_RECT, R: 255, G: 0, B: 0, A: 255},
	}
}

func CreateTile(rect sdl.FRect) entities.Tile {
	return entities.Tile{
		Entity:          ecs.NewEntity(),
		CameraComponent: &components.CameraComponent{},
		CollisionComponent: &components.CollisionComponent{
			Rect:   rect,
			Enable: false,
		},
		TransformComponent: &components.TransformComponent{Rect: rect, Angle: 0, Center: sdl.FPoint{}, Flip: 0, Speed: 0.5},
		DrawComponent:      &components.DrawComponent{Type: components.DRAW_TYPE_TILE, Shape: components.DRAW_SHAPE_RECT, R: 0, G: 0, B: 255, A: 255},
	}
}
