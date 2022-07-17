package algorithm

import "github.com/veandco/go-sdl2/sdl"

func AABB(a sdl.FRect, b sdl.FRect) bool {

	if a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y {
		return true
	}
	return false
}
