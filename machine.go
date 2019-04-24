package main

import (
	"strings"
)

type machine struct {
	plugboard map[rune]rune
	reflector map[rune]rune
	rotor1    rotor
	rotor2    rotor
	rotor3    rotor
}

func newMachine(n1, n2, n3 int) (m *machine, err error) {
	r1, err1 := newRotor(n1)

	if err1 != nil {
		return nil, err
	}

	r2, err2 := newRotor(n2)

	if err2 != nil {
		return nil, err
	}

	r3, err3 := newRotor(n3)

	if err3 != nil {
		return nil, err
	}

	return &machine{
		plugboard: map[rune]rune{
			'A': 'Q',
			'B': 'B',
			'C': 'C',
			'D': 'D',
			'E': 'E',
			'F': 'F',
			'G': 'G',
			'H': 'H',
			'I': 'I',
			'J': 'J',
			'K': 'K',
			'L': 'L',
			'M': 'M',
			'N': 'N',
			'O': 'O',
			'P': 'P',
			'Q': 'A',
			'R': 'R',
			'S': 'S',
			'T': 'T',
			'U': 'U',
			'V': 'V',
			'W': 'W',
			'X': 'X',
			'Y': 'Y',
			'Z': 'Z',
		},
		reflector: map[rune]rune{
			'A': 'Z',
			'B': 'Y',
			'C': 'X',
			'D': 'W',
			'E': 'V',
			'F': 'U',
			'G': 'T',
			'H': 'S',
			'I': 'R',
			'J': 'Q',
			'K': 'P',
			'L': 'O',
			'M': 'N',
			'N': 'M',
			'O': 'L',
			'P': 'K',
			'Q': 'J',
			'R': 'I',
			'S': 'H',
			'T': 'G',
			'U': 'F',
			'V': 'E',
			'W': 'D',
			'X': 'C',
			'Y': 'B',
			'Z': 'A',
		},
		rotor1: *r1,
		rotor2: *r2,
		rotor3: *r3,
	}, nil
}

func (m *machine) encryptKey(key rune) rune {
	c := m.plugboard[key]

	m.rotor3.rotate()
	if m.rotor3.rotationCompleted() {
		m.rotor2.rotate()
		if m.rotor2.rotationCompleted() {
			m.rotor1.rotate()
		}
	}

	c = m.rotor3.direct(c)
	c = m.rotor2.direct(c)
	c = m.rotor1.direct(c)

	c = rune(m.reflector[c])

	c = m.rotor1.inverse(c)
	c = m.rotor2.inverse(c)
	c = m.rotor3.inverse(c)

	c = m.plugboard[c]

	return c
}

func (m *machine) encryptMessage(message string) string {
	var buffer strings.Builder

	for _, c := range message {
		buffer.WriteRune(m.encryptKey(c))
	}

	return buffer.String()
}
