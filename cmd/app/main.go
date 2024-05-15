package main

import (
	. "github.com/kingso/brocode/internal"
)

func main() {
	var msg Str = "Hello, 世界"
	Pf("Msg: '%s', Length: %d \n", msg, msg.Len())
	Pf("Msg: '%s', Upper: %s \n", msg, msg.Upper())
	Pf("Msg: '%s', Lower: %s \n", msg, msg.Lower())
	Pf("Msg: '%s', Split: %s \n", msg, msg.Split(" "))

}
