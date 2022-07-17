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

	globals.World.AddSystem(&systems.PositionSystem{}, systems.Positionable)
	globals.World.AddSystem(&systems.RenderSystem{
		Renderer: g.renderer,
	}, systems.Drawable)
	globals.World.AddSystem(&systems.ControlSystem{}, systems.Controllable)
	globals.World.AddSystem(&systems.AISystem{}, systems.AIable)

	player = entities.Player{
		Entity: ecs.NewEntity(),
		ActionComponent: &components.ActionComponent{
			CurDir:  sdl.FPoint{},
			LastDir: sdl.FPoint{},
			Motion:  0,
		},
		TransformComponent: &components.TransformComponent{
			Pos:    sdl.FPoint{X: 100, Y: 110},
			Dim:    sdl.FPoint{X: 100, Y: 100},
			Angle:  0,
			Center: sdl.FPoint{},
			Flip:   0,
			Speed:  3,
		},
		DrawComponent: &components.DrawComponent{
			Type:  components.DRAW_TYPE_PLAYER,
			Shape: components.DRAW_SHAPE_RECT,
			R:     0,
			G:     255,
			B:     0,
			A:     255,
		},
	}
	globals.World.AddEntity(&player)

	enemy := factory.CraeteEnemy(player, sdl.FPoint{X: 100, Y: 100})
	globals.World.AddEntity(&enemy)

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
		globals.World.AddEntity(factory.CraeteEnemy(
			player,
			sdl.FPoint{
				X: float32(math.Remainder(rand.Float64(), float64(g.width))),
				Y: float32(math.Remainder(rand.Float64(), float64(g.height)))}))
	}
}

func (g *Game) FrameCounter() int {
	return g.frameCounter
}
