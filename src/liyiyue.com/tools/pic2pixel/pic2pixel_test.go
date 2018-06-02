package main

import (
	"testing"
	"image/color"
	"fmt"
)

func TestTransfColor(t *testing.T) {
	c := transfColor(color.RGBA{
		R: 218,
		G: 210,
		B: 137,
		A: 0xff,
	})
	fmt.Printf("R:0x%x G:0x%x B:0x%x A:0x%x\n", c.R, c.G, c.B, c.A)
	fmt.Printf("R:%d G:%d B:%d A:%d\n", c.R, c.G, c.B, c.A)
}
