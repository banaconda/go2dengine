package main

import (
	"go2dengine/cmd/test_game/components"
	"go2dengine/cmd/test_game/entities"
	"go2dengine/cmd/test_game/globals"
	"go2dengine/cmd/test_game/systems"
	"go2dengine/pkg/ecs"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	frameCounter  int
	window        *sdl.Window
	renderer      *sdl.Renderer
	width, height int32

	world ecs.World
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

func (g *Game) InitECS() {
	g.world = ecs.NewWorld()

	player := entities.Player{
		Entity: &ecs.Entity{},
		TransformComponent: &components.TransformComponent{
			Pos: sdl.FPoint{X: 100, Y: 110},
			Dim: sdl.FPoint{X: 100, Y: 100},
		},
		DrawComponent: &components.DrawComponent{
			Type:  components.DRAW_TYPE_PLAYER,
			Shape: components.DRAW_SHAPE_RECT,
			R:     255,
			G:     0,
			B:     0,
			A:     255,
		},
		ActionComponent: &components.ActionComponent{
			CurDir:  sdl.FPoint{},
			LastDir: sdl.FPoint{},
			Motion:  0,
		},
	}

	g.world.AddSystem(&systems.PositionSystem{}, systems.Positionable)
	g.world.AddSystem(&systems.RenderSystem{
		Renderer: g.renderer,
	}, systems.Drawable)
	g.world.AddSystem(&systems.ControlSystem{}, systems.Controllable)

	g.world.AddEntity(&player)
	globals.Logger.Info("init done")
}

func (g *Game) Update() {
	if err := g.renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		globals.Logger.Info("draw error: %v", err)
	}

	if err := g.renderer.Clear(); err != nil {
		globals.Logger.Info("render error: %v", err)
	}

	g.world.Update()

	g.renderer.Present()
}

func (g *Game) FrameCounter() int {
	return g.frameCounter
}
