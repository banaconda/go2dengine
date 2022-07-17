package globals

import (
	"go2dengine/pkg/ecs"
)

var IsRunning bool = false
var World *ecs.World = nil

const (
	CONTROL_SYSTEM = iota
	AI_SYSTEM
	POSITON_SYSTEM
	RENDER_SYSTEM
	SYSTEM_MAX
)
