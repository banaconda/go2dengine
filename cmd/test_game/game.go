package main

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/entities"
	"go2dengine/cmd/test_game/factory"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/cmd/test_game/systems"
	"go2dengine/pkg/ecs"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	frameCounter  int
	window        *sdl.Window
	renderer      *sdl.Renderer
	width, height int32
}

func (g *Game) InitSDL(title string, x, y, width, height int32) {
	g.width = width
	g.height = height

	globals.Logger.Info("start init sdl")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		globals.IsRunning = false
		globals.Logger.Error("fail init sdl: %v", err)
		return
	}

	globals.IsRunning = true

	if globals.IsRunning {
		var err error
		if err = ttf.Init(); err != nil {
			globals.IsRunning = false
			globals.Logger.Error("fail init ttf: %v", err)
			return
		}

		if g.window, err = sdl.CreateWindow(title, x, y, width, height, 0); err != nil {
			globals.IsRunning = false
			globals.Logger.Error("fail create window: %v", err)
			return
		}

		if g.renderer, err = sdl.CreateRenderer(g.window, -1, 0); err != nil {
			globals.IsRunning = false
			globals.Logger.Error("fail create renderer: %v", err)
			return
		}
	}
}

var player entities.Player

func (g *Game) InitECS() {
	globals.World = ecs.NewWorld()

	globals.World.AddSystem(&systems.ControlSystem{}, systems.Controllable)
	globals.World.AddSystem(&systems.AISystem{}, systems.AIable)
	globals.World.AddSystem(&systems.CollisionSystem{}, systems.Collidable)
	globals.World.AddSystem(&systems.PositionSystem{}, systems.Positionable)
	globals.World.AddSystem(&systems.RenderSystem{
		Renderer: g.renderer,
	}, systems.Drawable)

	player = entities.Player{
		Entity:           ecs.NewEntity(),
		CameraComponent:  &components.CameraComponent{},
		ControlComponent: &components.ControlComponent{},
		ActionComponent:  &components.ActionComponent{CurDir: sdl.FPoint{}, LastDir: sdl.FPoint{}, Motion: 0},
		CollisionComponent: &components.CollisionComponent{
			Rect:   sdl.FRect{X: 100, Y: 100, W: 100, H: 100},
			Enable: true,
		},
		TransformComponent: &components.TransformComponent{Rect: sdl.FRect{X: 100, Y: 100, W: 100, H: 100}, Angle: 0, Center: sdl.FPoint{}, Flip: 0, Speed: 3},
		DrawComponent:      &components.DrawComponent{Type: components.DRAW_TYPE_PLAYER, Shape: components.DRAW_SHAPE_RECT, R: 0, G: 255, B: 0, A: 255},
	}
	globals.World.AddEntity(&player)

	enemy := factory.CraeteEnemy(player, sdl.FRect{
		X: 80,
		Y: 80,
		W: 80,
		H: 80,
	})
	globals.World.AddEntity(&enemy)

	for x := 0; x < 24; x++ {
		for y := 0; y < 14; y++ {
			tile := factory.CreateTile(sdl.FRect{X: float32(x * 80), Y: float32(y * 80), W: 80, H: 80})
			globals.World.AddEntity(&tile)
		}
	}

	globals.Logger.Info("init done")
}

func (g *Game) Update() {
	if err := g.renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		globals.Logger.Info("draw error: %v", err)
	}

	if err := g.renderer.Clear(); err != nil {
		globals.Logger.Info("render error: %v", err)
	}

	globals.World.Update()

	g.renderer.Present()
}

func (g *Game) Debug() {
	state := sdl.GetKeyboardState()
	if state[sdl.SCANCODE_F4] == 1 {
		for i := 0; i < 1000; i++ {
			globals.World.AddEntity(factory.CraeteEnemy(
				player,
				sdl.FRect{
					X: float32(math.Remainder(rand.Float64(), float64(g.width))),
					Y: float32(math.Remainder(rand.Float64(), float64(g.height))),
					W: 80,
					H: 80,
				},
			))
		}
	}
	if state[sdl.SCANCODE_ESCAPE] == 1 {
		globals.IsRunning = false
	}
}

func (g *Game) FrameCounter() int {
	return g.frameCounter
}
