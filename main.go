package main

import (
	. "github.com/kingso/brocode/utils"
)

func main() {
	var msg Str = "Hello, 世界"
	Pl(msg, msg.Len())
	Pl(msg.Upper())
	Pl(msg.Lower())
	Pl(msg.Reverse())
	Pl(msg.Split(" "))
}
