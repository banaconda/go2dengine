package components

const (
	DRAW_SHAPE_RECT = iota
	DRAW_SHAPE_CIRCLE
)

const (
	DRAW_TYPE_TILE = iota
	DRAW_TYPE_OBJ
	DRAW_TYPE_ENEMY
	DRAW_TYPE_PLAYER
	DRAW_TYPE_EFFECT
	DRAW_TYPE_MAX
)

type DrawComponent struct {
	Type       uint32
	Shape      uint32
	R, G, B, A uint8
}

type DrawInterface interface {
	GetDrawComponent() *DrawComponent
}

func (c *DrawComponent) GetDrawComponent() *DrawComponent {
	return c
}
