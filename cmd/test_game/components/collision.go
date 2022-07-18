package components

import "github.com/veandco/go-sdl2/sdl"

type CollisionComponent struct {
	Rect   sdl.FRect
	Enable bool
}

type CollisionInterface interface {
	GetCollisionComponent() *CollisionComponent
}

func (c *CollisionComponent) GetCollisionComponent() *CollisionComponent {
	return c
}
