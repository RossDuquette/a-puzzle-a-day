package solver

import (
	"fmt"
	"testing"
)

func TestTileRotationV(t *testing.T) {
	tiles := get_tiles()
	v := tiles["v"]

	tests := []struct {
		rotations uint
		shape string
	}{
		{0, "ddrr"},
		{1, "lldd"},
		{1, "uull"},
		{1, "rruu"},
		{1, "ddrr"},
		{3, "rruu"},
		{3, "uull"},
		{3, "lldd"},
		{3, "ddrr"},
	}

	for i, test := range tests {
		test_name := fmt.Sprintf("%d", i)
		t.Run(test_name, func(t *testing.T) {
			v.rotate_cw(test.rotations)
			if v.shape != test.shape {
				t.Errorf("%d - Expected %s, got %s", i, test.shape, v.shape)
			}
		})
	}
}

func TestTileFlipY(t *testing.T) {
	tiles := get_tiles()
	y := tiles["y"]

	tests := []struct {
		flip bool
		shape string
	}{
		{false, "rrudr"},
		{true, "lludl"},
		{true, "rrudr"},
	}

	for i, test := range tests {
		test_name := fmt.Sprintf("%d", i)
		t.Run(test_name, func(t *testing.T) {
			if test.flip {
				y.flip()
			}
			if y.shape != test.shape {
				t.Errorf("%d - Expected %s, got %s", i, test.shape, y.shape)
			}
		})
	}
}

func TestTileRotationFlipZ(t *testing.T) {
	tiles := get_tiles()
	z := tiles["z"]

	tests := []struct {
		rotations uint
		flip bool
		shape string
	}{
		{0, false, "rddr"},
		{0, true, "lddl"},
		{1, false, "ullu"},
		{0, true, "urru"},
	}

	for i, test := range tests {
		test_name := fmt.Sprintf("%d", i)
		t.Run(test_name, func(t *testing.T) {
			if test.flip {
				z.flip()
			}
			z.rotate_cw(test.rotations)
			if z.shape != test.shape {
				t.Errorf("%d - Expected %s, got %s", i, test.shape, z.shape)
			}
		})
	}
}
