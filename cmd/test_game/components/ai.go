package components

import (
	"go2dengine/pkg/ecs"
)

const (
	AI_TYPE_CHASER = iota
	AI_TYPE_FLEE
	AI_TYPE_ATTACKER
)

type AIComponent struct {
	AIType uint32
	Target ecs.Identifier
}

type AIInterface interface {
	GetAIComponent() *AIComponent
}

func (c *AIComponent) GetAIComponent() *AIComponent {
	return c
}
