package gam

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type Scene interface {
	New(width, height int, window *glfw.Window)
	Render()
	Update(float32)
	Close()
}

func Run(scene Scene, tickRate float32, width, height int) {
	if tickRate == 0 {
		panic("Tickrate must be > 0, try 1./60.")
	}
	runtime.LockOSThread()

	glfw.Init()
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	// glfw window creation
	window, err := glfw.CreateWindow(width, height, "Breakout", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, int32(width*2), int32(height*2))
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	scene.New(width, height, window)

	var frames int
	var dt, accumulator float64
	tickRate64 := float64(tickRate)
	lastFrame := glfw.GetTime()
	lastFPS := lastFrame

	for !window.ShouldClose() {
		currentFrame := glfw.GetTime()
		frames++
		if currentFrame - lastFPS > 1 {
			window.SetTitle(fmt.Sprintf("Breakout | %d FPS", frames))
			frames = 0
			lastFPS = currentFrame
		}
		dt = currentFrame - lastFrame
		lastFrame = currentFrame

		for accumulator += dt; accumulator > tickRate64; accumulator -= tickRate64 {
			scene.Update(tickRate)
		}

		gl.ClearColor(0, 0, 0, 0.5)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		scene.Render()
		window.SwapBuffers()
		glfw.PollEvents()
	}

	scene.Close()
}
