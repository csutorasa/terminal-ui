package ansi

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

// ANSI terminal input
type Input uint64

// Gets if the input is an escape sequence.
func (i Input) IsEscape() bool {
	v := i
	for {
		if v>>8 == 0 {
			return v == 0x1B
		}
		v >>= 8
	}
}

// ANSI [Input] decoder.
type InputDecoder struct {
	r bufio.Reader
	p []byte
}

func NewInputDecoder(r io.Reader) *InputDecoder {
	return &InputDecoder{
		r: *bufio.NewReader(r),
		p: make([]byte, 8),
	}
}

func (d *InputDecoder) Decode() (Input, error) {
	for i := range 8 {
		b, err := d.r.ReadByte()
		if err != nil {
			return 0, err
		}
		d.p[i] = b
		if FullInput(d.p[:i+1]) {
			input, n, err := DecodeInput(d.p[:i+1])
			if n != i+1 {
				return 0, fmt.Errorf("invalid input length %X", input)
			}
			return input, err
		}
	}
	return 0, fmt.Errorf("unrecognized input %X", d.p)
}

// Checks if the bytes represent a known input.
func FullInput(p []byte) bool {
	if len(p) == 0 {
		return false
	}
	if p[0] == byte(KeyEsc) {
		n := 1
		input := Input(p[0])
		for _, b := range p[1:] {
			if b == byte(KeyEsc) {
				return false
			}
			input = (input << 8) + Input(b)
			n++
			if isKnownKeyEscape(input) {
				return true
			}
			if n > 7 {
				return false
			}
		}
		return false
	}
	return utf8.FullRune(p)
}

// Reads an ANSI input from the bytes.
func DecodeInput(p []byte) (Input, int, error) {
	if len(p) == 0 {
		return 0, 0, io.EOF
	}
	if p[0] == byte(KeyEsc) {
		n := 1
		input := Input(p[0])
		for _, b := range p[1:] {
			if b == byte(KeyEsc) {
				break
			}
			input = (input << 8) + Input(b)
			n++
			if isKnownKeyEscape(input) {
				return input, n, nil
			}
			if n > 7 {
				return 0, n, fmt.Errorf("unrecognized input %X", p[:n])
			}
		}
		return 0, n, fmt.Errorf("unrecognized input %X", p[:n])
	}
	rune, n := utf8.DecodeRune(p)
	if rune == utf8.RuneError {
		return 0, n, fmt.Errorf("invalid rune %X", p[:n])
	}
	return Input(rune), n, nil
}

// Writes an ANSI input to the bytes.
func EncodeInput(p []byte, i Input) (int, error) {
	if len(p) == 0 {
		return 0, io.EOF
	}
	n := 0
	for i := range 8 {
		b := byte(i >> (7 - i))
		if n > 0 || b > 0 {
			if len(p) < n+1 {
				return n, io.EOF
			}
			p[n] = b
			n++
		}
	}
	return n, nil
}
