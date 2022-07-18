package components

import (
	"go2dengine/pkg/ecs"
)

const (
	AI_TYPE_STOP = iota
	AI_TYPE_CHASE
	AI_TYPE_FLEE
	AI_TYPE_ATTACKER
	AI_TYPE_MAX
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
