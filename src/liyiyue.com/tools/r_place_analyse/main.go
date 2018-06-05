package main

import (
	"os"
	"image/png"
	"image/color"
	"sort"
	"fmt"
)

func main() {
	file, err := os.Open("/Users/liyiyue/Documents/go/reddit_pixel/src/liyiyue.com/tools/test/r_place.png")
	check(err)
	img, err := png.Decode(file)
	check(err)

	m := make(map[rgb]int)
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			clr := img.At(x, y).(color.RGBA)
			c := rgb{
				r: clr.R,
				g: clr.G,
				b: clr.B,
			}
			m[c]++
		}
	}

	s := &rgb_i{
		s: make([]*rgb_s, 0, len(m)),
	}

	for k, v := range m {
		s.s = append(s.s, &rgb_s{
			r: k.r,
			g: k.g,
			b: k.b,
			n: v,
		})
	}

	sort.Sort(s)

	for i := 0; i < 16; i++ {
		fmt.Println(s.s[i])
	}
}

type rgb struct {
	r, g, b uint8
}

type rgb_s struct {
	r, g, b uint8
	n       int
}

type rgb_i struct {
	s []*rgb_s
}

func (c *rgb_i) Len() int {
	return len(c.s)
}

func (c *rgb_i) Less(i, j int) bool {
	return c.s[i].n > c.s[j].n
}

func (c *rgb_i) Swap(i, j int) {
	c.s[i], c.s[j] = c.s[j], c.s[i]
}

func check(err error) {
	if nil != err {
		panic(err)
	}
}
