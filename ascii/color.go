package ascii

import (
	"fmt"
	"image/color"
)

type ColorExtent struct {
	Start  int
	Length int
	ColorPair
}

type ColorPair struct {
	Foreground color.Color
	Background color.Color
}

func (c ColorPair) Code(p *Palette) []byte {
	if c.Foreground == nil && c.Background == nil {
		return []byte{'\x0f'}
	}
	if c.Foreground == nil {
		panic("bg but no fg")
	}
	fgc := lookupIndex(c.Foreground)
	if fgc == -1 {
		panic("invalid fg")
	}
	ret := []byte{'\x03'}
	ret = append(ret, []byte(fmt.Sprintf("%d", fgc))...)
	if c.Background == nil {
		return ret
	}
	bgc := lookupIndex(c.Background)
	if bgc == -1 {
		panic("invalid bg")
	}
	ret = append(ret, []byte(fmt.Sprintf(",%d", bgc))...)
	return ret
}
