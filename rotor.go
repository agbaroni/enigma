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

import (
	// "fmt"
	"strconv"
)

type rotor struct {
	configuration string
	rotations     uint
}

type invalidRotorNumber struct {
	n int
}

func (num *invalidRotorNumber) Error() string {
	return "invalid rotor number " + strconv.Itoa(num.n)
}

func newRotor(n int) (r *rotor, err error) {
	if n == 1 {
		return &rotor{
			configuration: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			rotations:     0,
		}, nil
	} else if n == 2 {
		return &rotor{
			configuration: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
			rotations:     0,
		}, nil
	} else if n == 3 {
		return &rotor{
			configuration: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			rotations:     0,
		}, nil
	}

	return nil, &invalidRotorNumber{
		n: n,
	}
}

func (r *rotor) rotate() {
	r.configuration = r.configuration[1:] + r.configuration[0:1]

	r.rotations++
}

func (r *rotor) rotationCompleted() bool {
	return r.rotations%26 == 0
}

func (r *rotor) direct(key rune) rune {
	return rune(r.configuration[int(key-'A')])
}

func (r *rotor) inverse(key rune) rune {
	x := 0

	for i, c := range r.configuration {
		if c == key {
			x = i
			break
		}
	}

	return rune('A' + x)
}
