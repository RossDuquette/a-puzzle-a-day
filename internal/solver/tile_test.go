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
		shape     string
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
		flip  bool
		shape string
	}{
		{false, "rrdur"},
		{true, "lldul"},
		{true, "rrdur"},
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
		flip      bool
		shape     string
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

func TestTilePointsP(t *testing.T) {
	tiles := get_tiles()
	p := tiles["p"]
	points := p.get_points()

	expected_points := []Point{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
		{0, 2},
	}

	if len(points) != 5 {
		t.Errorf("Shape has %d points", len(points))
	}

	for i := range points {
		if points[i] != expected_points[i] {
			t.Errorf("Got point (%d, %d), expected (%d, %d)",
				points[i].x, points[i].y, expected_points[i].x, expected_points[i].y)
		}
	}
}

func TestTilePointsPRotated(t *testing.T) {
	tiles := get_tiles()
	p := tiles["p"]
	p.rotate_cw(2)
	points := p.get_points()

	expected_points := []Point{
		{0, 2},
		{-1, 2},
		{-1, 1},
		{0, 1},
		{0, 0},
	}

	if len(points) != 5 {
		t.Errorf("Shape has %d points", len(points))
	}

	for i := range points {
		if points[i] != expected_points[i] {
			t.Errorf("Got point (%d, %d), expected (%d, %d)",
				points[i].x, points[i].y, expected_points[i].x, expected_points[i].y)
		}
	}
}
