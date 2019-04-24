package main

// #include <curses.h>
// #cgo CFLAGS:
// #cgo LDFLAGS: -lncurses
import "C"

import (
	"fmt"
)

func main() {
	m1, err := newMachine(1, 2, 3)

	if err != nil {
		fmt.Println(err.Error())
	}

	m2, err := newMachine(1, 2, 3)

	if err != nil {
		fmt.Println(err.Error())
	}

	// C.initscr()
	// C.cbreak()
	// C.noecho()

	s := m1.encryptMessage("NELMEZZODELCAMMINDINOSTRAVITA")
	fmt.Println(s)

	fmt.Println(m2.encryptMessage(s))

	// C.echo()

	// C.clear()
}
