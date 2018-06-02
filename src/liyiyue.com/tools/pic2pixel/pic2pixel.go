package main

import (
	"os"
	"image/png"
	"image"
	"image/color"
	"math"
)

const (
	FILE_NAME_1     = "1.png"
	OUT_FILE_NAME_1 = "1_out.png"
	FILE_NAME_2     = "2.png"
	OUT_FILE_NAME_2 = "2_out.png"
	FILE_NAME_3     = "3.png"
	OUT_FILE_NAME_3 = "3_out.png"
	FILE_NAME_4     = "4.png"
	OUT_FILE_NAME_4 = "4_out.png"
)

const (
	LEN_OUT = 64
)

type rgba struct {
	R, G, B, A int
}

var colorList []rgba

func init() {
	colorList = make([]rgba, 16, 16)
	colorList[0] = rgba{R: 240, G: 190, B: 55, A: 0xff}
	colorList[1] = rgba{R: 0x23, G: 0xff, B: 0xc2, A: 0xff}
	colorList[2] = rgba{R: 0x33, G: 0x33, B: 0xff, A: 0xff}
	colorList[3] = rgba{R: 0x33, G: 0xff, B: 0x33, A: 0xff}
	colorList[4] = rgba{R: 0x55, G: 0xd3, B: 0xff, A: 0xff}
	colorList[5] = rgba{R: 0x66, G: 0x66, B: 0x66, A: 0xff}
	colorList[6] = rgba{R: 0x7f, G: 0x00, B: 0x80, A: 0xff}
	colorList[7] = rgba{R: 0x93, G: 0xff, B: 0x6f, A: 0xff}
	colorList[8] = rgba{R: 0x9e, G: 0x7c, B: 0xfe, A: 0xff}
	colorList[9] = rgba{R: 0xd3, G: 0xff, B: 0x73, A: 0xff}
	colorList[10] = rgba{R: 0xf3, G: 0x80, B: 0x04, A: 0xff}
	colorList[11] = rgba{R: 0xf6, G: 0x99, B: 0xc0, A: 0xff}
	colorList[12] = rgba{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	colorList[13] = rgba{R: 0xff, G: 0x33, B: 0xff, A: 0xff}
	colorList[14] = rgba{R: 0xff, G: 0xff, B: 0x33, A: 0xff}
	colorList[15] = rgba{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	//return

	//2 8
	//3 27
	//4 64
	//5 125
	//6 216
	//7 343
	//8 512
	dp := 2
	d := 255 / (dp - 1)
	colorList = make([]rgba, dp*dp*dp, dp*dp*dp)
	for r := 0; r < dp; r++ {
		for g := 0; g < dp; g++ {
			for b := 0; b < dp; b++ {
				colorList = append(colorList, rgba{
					R: r * d,
					G: g * d,
					B: b * d,
					A: 0xff,
				})
			}
		}
	}
}

func main() {
	Pic2Pixel(FILE_NAME_1, OUT_FILE_NAME_1)
	Pic2Pixel(FILE_NAME_2, OUT_FILE_NAME_2)
	Pic2Pixel(FILE_NAME_3, OUT_FILE_NAME_3)
	Pic2Pixel(FILE_NAME_4, OUT_FILE_NAME_4)
}

func Pic2Pixel(fileName, outName string) {
	file, err := os.Open(fileName)
	check(err)
	originImg, err := png.Decode(file)
	check(err)
	outImg := scale(originImg, LEN_OUT)

	outFile, err := os.Create(outName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	png.Encode(outFile, outImg)
}

func scale(originImg image.Image, len int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, len, len))

	//缩放
	maxX := originImg.Bounds().Dx()
	maxY := originImg.Bounds().Dy()
	scale := float64(0)
	dX := 0
	dY := 0
	if maxX >= maxY {
		scale = float64(len) / float64(maxX)
		dY = int(float64(maxX-maxY) * scale / 2)
	} else {
		scale = float64(len) / float64(maxY)
		dX = int(float64(maxY-maxX) * scale / 2)
	}

	if scale <= 1 {
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				tX := int(float64(x)*scale) + dX
				tY := int(float64(y)*scale) + dY
				newImg.Set(tX, tY, originImg.At(x, y))
			}
		}
	} else {
		for x := 0; x < len; x++ {
			for y := 0; y < len; y++ {
				tX := int(float64(x-dX) / scale)
				tY := int(float64(y-dY) / scale)
				if tX < 0 || tX >= maxX || tY < 0 || tY >= maxY {
					continue
				}
				newImg.Set(x, y, originImg.At(tX, tY))
			}
		}
	}

	//转索引色
	for x := 0; x < len; x++ {
		for y := 0; y < len; y++ {
			originColor := newImg.At(x, y)
			newColor := transfColor(originColor)
			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}

//单一颜色转为索引色
func transfColor(originColor color.Color) color.RGBA {
	var r int
	var g int
	var b int
	var a uint8
	switch originColor.(type) {
	case color.RGBA:
		originRGBA := originColor.(color.RGBA)
		r = int(originRGBA.R)
		g = int(originRGBA.G)
		b = int(originRGBA.B)
		a = originRGBA.A
	case color.NRGBA:
		originRGBA := originColor.(color.NRGBA)
		r = int(originRGBA.R)
		g = int(originRGBA.G)
		b = int(originRGBA.B)
		a = originRGBA.A
	}

	i := 0
	m := float64(math.MaxInt32)
	p := 2
	for idx, c := range colorList {
		d := math.Sqrt(float64(pow(abs(c.R-r), p) + pow(abs(c.G-g), p) + pow(abs(c.B-b), p)))
		if d < m {
			m = d
			i = idx
		}
	}
	newColor := colorList[i]

	return color.RGBA{
		R: uint8(newColor.R),
		G: uint8(newColor.G),
		B: uint8(newColor.B),
		A: a,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		r *= a
		b--
	}
	return r
}

func check(err error) {
	if nil != err {
		panic(err)
	}
}
