package ecs

import (
	"reflect"
	"sort"
)

type World struct {
	systems
	systemTypeToAbleSlice
}

type systemType reflect.Type
type ableType reflect.Type

type systemTypeToAbleSlice map[systemType][]ableType

func NewWorld() *World {
	world := World{}
	world.systemTypeToAbleSlice = make(systemTypeToAbleSlice)
	return &world
}

func (w *World) AddSystem(system System, ables interface{}) {
	w.systems = append(w.systems, system)
	sysT := reflect.TypeOf(system)

	if !reflect.TypeOf(ables).AssignableTo(reflect.TypeOf([]interface{}{})) {
		ables = []interface{}{ables}
	}

	for _, able := range ables.([]interface{}) {
		ableT := reflect.TypeOf(able).Elem()
		w.systemTypeToAbleSlice[sysT] = append(
			w.systemTypeToAbleSlice[sysT],
			ableT)
	}
	sort.Sort(w.systems)
}

func (w *World) Systems() []System {
	return w.systems
}

func (w *World) Update() {
	for _, system := range w.systems {
		system.Update()
	}
}

func (w *World) AddEntity(entityId Identifier) {
	search := func(entityId Identifier, ableTs []ableType) bool {
		for _, ableT := range ableTs {
			if reflect.TypeOf(entityId).Implements(ableT) {
				return true
			}
		}
		return false
	}

	for _, system := range w.systems {
		t := reflect.TypeOf(system)
		if ableTypeSlice, err := w.systemTypeToAbleSlice[t]; err {
			if search(entityId, ableTypeSlice) {
				system.Add(entityId)
			}
		}
	}
}
