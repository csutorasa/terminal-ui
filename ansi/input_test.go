package ansi_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/csutorasa/terminal-ui/ansi"
)

func TestInputDecoderSingleByteKey(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte{'a'}))

	inputs := readAll(d)

	if len(inputs) != 1 {
		t.Logf("expected length: %d got: %d", 1, len(inputs))
		t.FailNow()
	}
	if 'a' != inputs[0] {
		t.Logf("expected: %X got: %X", 'a', inputs[0])
		t.Fail()
	}
}

func TestInputDecoderMultipleByteKey(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte("µ")))

	inputs := readAll(d)

	if len(inputs) != 1 {
		t.Logf("expected length: %d got: %d", 1, len(inputs))
		t.FailNow()
	}
	if 'µ' != inputs[0] {
		t.Logf("expected: %X got: %X", 'µ', inputs[0])
		t.Fail()
	}
}

func TestInputDecoderEscape(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte{0x1B, 0x5B, 0x41}))

	inputs := readAll(d)

	if len(inputs) != 1 {
		t.Logf("expected length: %d got: %d", 1, len(inputs))
		t.FailNow()
	}
	if 0x1B5B41 != inputs[0] {
		t.Logf("expected: %X got: %X", 0x1B5B41, inputs[0])
		t.Fail()
	}
}

func TestInputDecoderPasted(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte{'a', 'b', 'c'}))

	inputs := readAll(d)

	if len(inputs) != 3 {
		t.Logf("expected length: %d got: %d", 3, len(inputs))
		t.FailNow()
	}
	if 'a' != inputs[0] {
		t.Logf("expected: %X got: %X", 'a', inputs[0])
		t.Fail()
	}
	if 'b' != inputs[1] {
		t.Logf("expected: %X got: %X", 'b', inputs[1])
		t.Fail()
	}
	if 'c' != inputs[2] {
		t.Logf("expected: %X got: %X", 'c', inputs[2])
		t.Fail()
	}
}

func TestInputDecoderMultipleEscapes(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte{
		0x1B, 0x5B, 0x41,
		0x1B, 0x5B, 0x41,
	}))

	inputs := readAll(d)

	if len(inputs) != 2 {
		t.Logf("expected length: %d got: %d", 2, len(inputs))
		t.FailNow()
	}
	if 0x1B5B41 != inputs[0] {
		t.Logf("expected: %X got: %X", 0x1B5B41, inputs[0])
		t.Fail()
	}
	if 0x1B5B41 != inputs[1] {
		t.Logf("expected: %X got: %X", 0x1B5B41, inputs[1])
		t.Fail()
	}
}

func TestInputDecoderEscapeAndCharacters(t *testing.T) {
	d := ansi.NewInputDecoder(bytes.NewReader([]byte{0x1B, 0x5B, 0x41, 'a', 'b', 'c'}))

	inputs := readAll(d)

	if len(inputs) != 4 {
		t.Logf("expected length: %d got: %d", 4, len(inputs))
		t.FailNow()
	}
	if 0x1B5B41 != inputs[0] {
		t.Logf("expected: %X got: %X", 0x1B5B41, inputs[0])
		t.Fail()
	}
	if 'a' != inputs[1] {
		t.Logf("expected: %X got: %X", 'a', inputs[1])
		t.Fail()
	}
	if 'b' != inputs[2] {
		t.Logf("expected: %X got: %X", 'b', inputs[2])
		t.Fail()
	}
	if 'c' != inputs[3] {
		t.Logf("expected: %X got: %X", 'c', inputs[3])
		t.Fail()
	}
}

func readAll(d *ansi.InputDecoder) []ansi.Input {
	inputs := []ansi.Input{}
	for {
		i, err := d.Decode()
		if err == io.EOF {
			return inputs
		}
		inputs = append(inputs, i)
	}
}
