package helper

import "math"

type Vec2D[T Number] struct {
	X, Y T
}

func NewVec2D[T Number](x, y T) Vec2D[T] {
	return Vec2D[T]{X: x, Y: y}
}

func (p Vec2D[T]) Add(p2 Vec2D[T]) Vec2D[T] {
	return Vec2D[T]{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Vec2D[T]) Sub(p2 Vec2D[T]) Vec2D[T] {
	return Vec2D[T]{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Vec2D[T]) Neg() Vec2D[T] {
	return Vec2D[T]{X: -p.X, Y: -p.Y}
}

func (p Vec2D[T]) Mul(factor T) Vec2D[T] {
	return Vec2D[T]{X: p.X * factor, Y: p.Y * factor}
}

func (p Vec2D[T]) Div(divisor T) Vec2D[T] {
	return Vec2D[T]{X: p.X / divisor, Y: p.Y / divisor}
}

func (p Vec2D[T]) RotCW() Vec2D[T] {
	return Vec2D[T]{X: -p.Y, Y: p.X}
}

func (p Vec2D[T]) RotCCW() Vec2D[T] {
	return Vec2D[T]{X: p.Y, Y: -p.X}
}

func (p Vec2D[T]) Cross(p2 Vec2D[T]) float64 {
	return float64(p.X)*float64(p2.Y) - float64(p.Y)*float64(p2.X)
}

func (p Vec2D[T]) Len() float64 {
	return math.Sqrt(float64(p.X*p.X) + float64(p.Y*p.Y))
}

func (p Vec2D[T]) Dist(p2 Vec2D[T]) float64 {
	return p.Sub(p2).Len()
}

func (p Vec2D[T]) InBounds(min, max Vec2D[T]) bool {
	return p.X >= min.X && p.Y >= min.Y && p.X <= max.X && p.Y <= max.Y
}

func ConvertVec2D[FROM, TO Number](from Vec2D[FROM]) Vec2D[TO] {
	return Vec2D[TO]{X: TO(from.X), Y: TO(from.Y)}
}

type Vec3D[T Number] struct {
	X, Y, Z T
}

func (p Vec3D[T]) Add(p2 Vec3D[T]) Vec3D[T] {
	return Vec3D[T]{X: p.X + p2.X, Y: p.Y + p2.Y, Z: p.Z + p2.Z}
}

func (p Vec3D[T]) Sub(p2 Vec3D[T]) Vec3D[T] {
	return Vec3D[T]{X: p.X - p2.X, Y: p.Y - p2.Y, Z: p.Z - p2.Z}
}

func (p Vec3D[T]) Neg() Vec3D[T] {
	return Vec3D[T]{X: -p.X, Y: -p.Y, Z: -p.Z}
}

func (p Vec3D[T]) Mul(factor T) Vec3D[T] {
	return Vec3D[T]{X: p.X * factor, Y: p.Y * factor, Z: p.Z * factor}
}

func (p Vec3D[T]) XY() Vec2D[T] {
	return Vec2D[T]{X: p.X, Y: p.Y}
}
