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
