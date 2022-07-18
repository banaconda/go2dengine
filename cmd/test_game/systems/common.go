package systems

import (
	"go2dengine/cmd/test_game/globals"
	"reflect"
)

func GetPositionSystem() *PositionSystem {
	for _, system := range globals.World.Systems() {
		if reflect.TypeOf(system).Elem() == reflect.TypeOf(PositionSystem{}) {
			return system.(*PositionSystem)
		}
	}
	return nil
}
