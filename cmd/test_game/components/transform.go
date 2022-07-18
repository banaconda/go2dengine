package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TransformComponent struct {
	Rect   sdl.FRect
	Angle  float64
	Center sdl.FPoint
	Flip   sdl.RendererFlip
	Speed  float32
}

type TransformInterface interface {
	GetTransformComponent() *TransformComponent
}

func (p *TransformComponent) GetTransformComponent() *TransformComponent {
	return p
}
