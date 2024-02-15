package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	width        = 800
	height       = 600
	maxFrames    = 200
	numParticles = 100
	minFallSpeed = 1
	maxFallSpeed = 5
	fadeRate     = 0.03
	trailLength  = 50
	outputDir    = "output"
)

var (
	characters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-=_+[]{}|;':,.<>/?`~ "
	defaultColors = []float64{0, 1, 0}
)

type CodeParticle struct {
	X, Y         float64
	Character    rune
	FallSpeed    float64
	Opacity      float64
	FadingFactor float64
	Trail        []TrailPoint
}

type TrailPoint struct {
	X, Y      float64
	Opacity   float64
	Character rune 
}

func main() {
	codeParticles := initializeCodeParticles()

	context := gg.NewContext(width, height)

	stop := make(chan bool)

	go func() {
		fmt.Println("Press 'Enter' to stop the simulation.")
		fmt.Scanln()
		stop <- true
	}()

	frameCount := 0
	for frameCount < maxFrames {
		select {
		case <-stop:
			fmt.Println("Simulation stopped.")
			return
		default:
			updateCodeParticles(codeParticles)
			drawCodeParticles(context, codeParticles, frameCount)
			frameCount++
			writeFrameToFile(context, frameCount)
			time.Sleep(time.Second / 30)
		}
	}

	context.SavePNG("output/final_frame.png")
	fmt.Println("Frames saved to:", outputDir)
}

func initializeCodeParticles() []CodeParticle {
	codeParticles := make([]CodeParticle, numParticles)

	for i := range codeParticles {
		codeParticles[i] = CodeParticle{
			X:            float64(rand.Intn(width)),
			Y:            -float64(rand.Intn(height)),
			Character:    rune(characters[rand.Intn(len(characters))]),
			FallSpeed:    rand.Float64()*(maxFallSpeed-minFallSpeed) + minFallSpeed,
			Opacity:      1.0,
			FadingFactor: fadeRate,
		}
	}
	return codeParticles
}

func updateCodeParticles(codeParticles []CodeParticle) {
	for i := range codeParticles {
		codeParticles[i].Y += codeParticles[i].FallSpeed
		codeParticles[i].Opacity -= codeParticles[i].FadingFactor

		// Add current position to the trail
		trailPoint := TrailPoint{X: codeParticles[i].X, Y: codeParticles[i].Y, Opacity: codeParticles[i].Opacity, Character: codeParticles[i].Character}
		codeParticles[i].Trail = append(codeParticles[i].Trail, trailPoint)

		// Remove old trail points
		if len(codeParticles[i].Trail) > trailLength {
			codeParticles[i].Trail = codeParticles[i].Trail[1:]
		}

		// Reset particle if it goes below the screen
		if codeParticles[i].Y > height {
			codeParticles[i] = CodeParticle{
				X:            float64(rand.Intn(width)),
				Y:            -float64(rand.Intn(height)),
				Character:    rune(characters[rand.Intn(len(characters))]),
				FallSpeed:    rand.Float64()*(maxFallSpeed-minFallSpeed) + minFallSpeed,
				Opacity:      1.0,
				FadingFactor: fadeRate,
			}
		}
	}
}

func drawCodeParticles(context *gg.Context, codeParticles []CodeParticle, frameCount int) {
	context.SetRGB(0, 0, 0)
	context.Clear()

	trailSpacing := 10.0 // Adjust the trail spacing

	for _, p := range codeParticles {
		context.SetRGBA(defaultColors[0], defaultColors[1], defaultColors[2], p.Opacity)

		for i, trailPoint := range p.Trail {
			if i%10 == frameCount%10 {
				trailPoint.Character = rune(characters[rand.Intn(len(characters))])
			}

			speedVariation := rand.Float64() * 2.0   // Adjust the speed variation
			opacityVariation := rand.Float64() * 0.5 // Adjust the opacity variation

			trailPoint.Y += speedVariation
			trailPoint.Opacity -= opacityVariation

			context.SetRGBA(defaultColors[0], defaultColors[1], defaultColors[2], trailPoint.Opacity)
			context.DrawStringAnchored(string(trailPoint.Character), trailPoint.X, trailPoint.Y+(float64(i)*trailSpacing), 0.5, 0.5)
		}

		context.DrawStringAnchored(string(p.Character), p.X, p.Y, 0.5, 0.5)
	}
}

func writeFrameToFile(context *gg.Context, frameCount int) {
	filename := fmt.Sprintf("%s/frame_%03d.png", outputDir, frameCount)
	context.SavePNG(filename)
}
