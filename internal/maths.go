package internal

import (
	. "github.com/kingso/brocode/pkg"
)

func DivideBy(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, Ef("Divide by zero")
	}
	return x / y, nil
}
