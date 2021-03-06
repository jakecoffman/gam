package breakout

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/jakecoffman/gam"
)

type Object struct {
	Position, Size, Velocity mgl32.Vec2
	Color                    mgl32.Vec3
	Rotation                 float64

	IsSolid, Destroyed bool

	Sprite *gam.Texture2D
}

func (o Object) String() string {
	return fmt.Sprintf("Object(@ %v - Color: %v)", o.Position, o.Color)
}

var (
	DefaultGameObjectColor = mgl32.Vec3{1, 1, 1}
)

func NewGameObject(pos, size mgl32.Vec2, sprite *gam.Texture2D) *Object {
	return &Object{
		pos,
		size,
		mgl32.Vec2{0, 0},
		DefaultGameObjectColor,
		0,
		false,
		false,
		sprite,
	}
}

func (o *Object) Draw(renderer *gam.SpriteRenderer) {
	renderer.DrawSprite(o.Sprite, o.Position, o.Size, o.Rotation, o.Color)
}
