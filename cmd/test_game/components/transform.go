package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TransformComponent struct {
	Pos    sdl.FPoint
	Dim    sdl.FPoint
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
