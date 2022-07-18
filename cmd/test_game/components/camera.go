package components

import "github.com/veandco/go-sdl2/sdl"

type CameraComponent struct {
	Rect  sdl.FRect
	Scale float32
}

type CameraInterface interface {
	GetCameraComponent() *CameraComponent
}

func (c *CameraComponent) GetCameraComponent() *CameraComponent {
	return c
}
