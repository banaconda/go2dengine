package ecs

import "sync/atomic"

var gid uint64

type Entity struct {
	id uint64
}

type Identifier interface {
	ID() uint64
}

type DefaultEntityInterface interface {
	GetEntity() *Entity
}

func NewEntity() *Entity {
	return &Entity{id: atomic.AddUint64(&gid, 1)}
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) GetEntity() *Entity {
	return e
}
