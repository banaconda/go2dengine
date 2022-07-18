// author: Jacky Boen

package main

import (
	"go2dengine/cmd/test_game/globals"
	"go2dengine/pkg/ecs"
	"time"

	nblogger "github.com/banaconda/nb-logger"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
}

func run() int {
	g := Game{}
	g.InitSDL("test", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1920, 1080)
	if !globals.IsRunning {
		globals.Logger.Error("sdl initialization fail")
		return -1
	}

	g.InitECS()

	frameRate := 120
	count := int64(0)
	before1s := time.Now()
	frameDelay := int64(1000000000 / frameRate)
	for globals.IsRunning {
		tick := time.Now()
		if tick.Sub(before1s).Nanoseconds() > 1000000000 {
			globals.Logger.Info("fps: %d, num of entities: %d", count, ecs.GetGid())
			before1s = tick
			count = 0
		}
		count++

		g.Update()
		g.Debug()

		totalDelay := frameDelay * count
		diff := totalDelay - time.Since(before1s).Nanoseconds()
		if diff > 0 {
			time.Sleep(time.Duration(diff))
		}
	}
	return 0
}

func main() {
	globals.InitLogger("main.log", nblogger.Info, 1000,
		nblogger.Lshortfile|nblogger.LstdFlags|nblogger.Lmicroseconds|nblogger.Lstdout)
	defer globals.Logger.Close()

	run()
}
