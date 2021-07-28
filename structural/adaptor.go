package structural

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

// Given interface is Vector type
type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangele(height, width int) *VectorImage {
	width -= 1
	height -= 1

	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// Our interface is Raster type
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// Adapter to adapt from VectorImage to RasterImage
type vectorToRasterAdapter struct {
	points []Point
}

func VectorToRaster(v *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range v.Lines {
		adapter.addLine(line)
	}
	return adapter
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

var pointCache = map[[16]byte][]Point{}

func hash(obj interface{}) [16]byte {
	bytes, _ := json.Marshal(obj)
	return md5.Sum(bytes)
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		a.points = append(a.points, pts...)
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	pointCache[h] = a.points
	fmt.Println("we have generated", len(a.points), "points")
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
