package components

import "github.com/veandco/go-sdl2/sdl"

const (
	MOTION_IDLE = iota
	MOTION_MOVE
	MOTION_ATTACK
	MOTION_NONE
)

type ActionComponent struct {
	CurDir  sdl.FPoint
	LastDir sdl.FPoint
	Motion  uint32
}

type ActionInterface interface {
	GetActionComponent() *ActionComponent
}

func (c *ActionComponent) GetActionComponent() *ActionComponent {
	return c
}
