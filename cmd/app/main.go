package main

import (
	. "github.com/kingso/brocode/internal"
	. "github.com/kingso/brocode/pkg"
)

func main() {
	var msg Str = "Hello, 世界"
	Pf("Msg: '%s', Length: %d \n", msg, msg.Len())
	Pf("Msg: '%s', Upper: %s \n", msg, msg.Upper())
	Pf("Msg: '%s', Lower: %s \n", msg, msg.Lower())
	Pf("Msg: '%s', Split: %s \n", msg, msg.Split(" "))

	ans, err := DivideBy(1, 10)
	if err != nil {
		Pl(err)
	} else {
		Pl(ans)
	}

}
