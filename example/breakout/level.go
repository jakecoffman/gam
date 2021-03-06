package breakout

import (
	"bufio"
	"fmt"
	"github.com/jakecoffman/gam"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

type Level struct {
	Bricks       []*Object
	block, solid *gam.Texture2D
}

func NewLevel(block, solid *gam.Texture2D) *Level {
	return &Level{
		Bricks: []*Object{},
		block:  block,
		solid:  solid,
	}
}

func (l *Level) Load(level string, lvlWidth, lvlHeight int) error {
	l.Bricks = l.Bricks[:0]

	var tileData [][]int
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimPrefix(level, "\n")))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		var row []int
		for _, part := range parts {
			i, err := strconv.Atoi(part)
			if err != nil {
				return fmt.Errorf("Failed to parse level: %s", err.Error())
			}
			row = append(row, i)
		}
		tileData = append(tileData, row)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Failed to scan file: %s", err)
	}
	if len(tileData) > 0 {
		return l.init(tileData, lvlWidth, lvlHeight)
	}
	return nil
}

func (l *Level) Draw(renderer *gam.SpriteRenderer) {
	for _, tile := range l.Bricks {
		if !tile.Destroyed {
			tile.Draw(renderer)
		}
	}
}

func (l *Level) IsCompleted() bool {
	for _, tile := range l.Bricks {
		if !tile.IsSolid && !tile.Destroyed {
			return false
		}
	}
	return true
}

func (l *Level) init(tileData [][]int, lvlWidth, lvlHeight int) error {
	height := len(tileData)
	width := len(tileData[0])
	unitWidth := lvlWidth / width
	unitHeight := lvlHeight / height

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if tileData[y][x] == 1 {
				pos := Vec2(unitWidth*x, unitHeight*y)
				size := Vec2(unitWidth, unitHeight)
				obj := NewGameObject(pos, size, l.solid)
				obj.Color = mgl32.Vec3{.8, .8, .7}
				obj.IsSolid = true
				l.Bricks = append(l.Bricks, obj)
			} else if tileData[y][x] > 1 {
				color := mgl32.Vec3{1, 1, 1}
				switch tileData[y][x] {
				case 2:
					color = mgl32.Vec3{.2, .6, 1}
				case 3:
					color = mgl32.Vec3{0, .7, 0}
				case 4:
					color = mgl32.Vec3{.8, .8, .4}
				case 5:
					color = mgl32.Vec3{1, .5, 0}
				}

				pos := Vec2(unitWidth*x, unitHeight*y)
				size := Vec2(unitWidth, unitHeight)
				obj := NewGameObject(pos, size, l.block)
				obj.Color = color
				l.Bricks = append(l.Bricks, obj)
			}
		}
	}

	return nil
}

func Vec2(x, y int) mgl32.Vec2 {
	return mgl32.Vec2{float32(x), float32(y)}
}
