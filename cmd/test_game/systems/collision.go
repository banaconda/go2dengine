package systems

import (
	"go2dengine/cmd/test_game/algorithm"
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
)

var Collidable *collidable

type collisionEntity struct {
	*ecs.Entity
	*components.ActionComponent
	*components.CollisionComponent
	*components.TransformComponent
}

type collidable interface {
	ecs.DefaultEntityInterface
	components.ActionInterface
	components.CollisionInterface
	components.TransformInterface
}

type CollisionSystem struct {
	entities []collisionEntity
}

func (s *CollisionSystem) Add(id ecs.Identifier) {
	obj := id.(collidable)
	s.entities = append(s.entities,
		collisionEntity{
			obj.GetEntity(),
			obj.GetActionComponent(),
			obj.GetCollisionComponent(),
			obj.GetTransformComponent(),
		})
}

func (s *CollisionSystem) Priority() int {
	return globals.POSITON_SYSTEM
}

func (s *CollisionSystem) Remove() {
}

func (s *CollisionSystem) Update() {
	for _, entity := range s.entities {
		trans := entity.GetTransformComponent()
		action := entity.GetActionComponent()
		collision := entity.GetCollisionComponent()
		collision.Rect = trans.Rect
		action.LastDir = action.CurDir
		if !collision.Enable {
			continue
		}

		for _, otherEntity := range s.entities {
			if entity.ID() == otherEntity.ID() || !otherEntity.CollisionComponent.Enable ||
				(action.CurDir.X == 0 && action.CurDir.Y == 0) {
				continue
			}

			rect := collision.Rect
			rect.X += action.CurDir.X
			rect.Y += action.CurDir.Y

			if algorithm.AABB(rect, otherEntity.GetCollisionComponent().Rect) {
				if algorithm.AABBX(rect, otherEntity.GetCollisionComponent().Rect) {
					action.CurDir.X = 0
				}

				if algorithm.AABBY(rect, otherEntity.GetCollisionComponent().Rect) {
					action.CurDir.Y = 0
				}
			}
		}
	}
}
