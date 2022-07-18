package components

type ControlComponent struct {
}

type ControlInterface interface {
	GetControlComponent() *ControlComponent
}

func (c *ControlComponent) GetControlComponent() *ControlComponent {
	return c
}
