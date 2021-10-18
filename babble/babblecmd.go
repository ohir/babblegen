// (c) 2021 Ohir Ripe. MIT license

package main

import (
	"babblegen"
	"fmt"

	"github.com/ohir/mopt"
)

func main() {
	var cmd mopt.Usage = "\tgenerate [-s number] size babble of [-t type] form a [-r number] seed."

	fmt.Println(babblegen.BabbleStr(
		cmd.OptN('s', 1<<16),
		uint64(cmd.OptN('r', 977)),
		babblegen.AsciiA1), "Thats all!")
}
