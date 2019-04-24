/**
 * This file is part of enigma.
 *
 * enigma is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * enigma is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with enigma.  If not, see <http://www.gnu.org/licenses/>.
 */

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
