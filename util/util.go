package util

import (
	"encoding/hex"
	"errors"
	"image/color"
)

func ParseColor(str string) (color.RGBA, error) {
	byteArr, err := hex.DecodeString(str)
	if err != nil {
		return color.RGBA{}, err
	}

	if len(byteArr) != 4 {
		return color.RGBA{}, errors.New("invalid color length")
	}

	return color.RGBA{byteArr[0], byteArr[1], byteArr[2], byteArr[3]}, nil
}
