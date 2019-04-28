package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	const w, h = 8, 8
	for _, tc := range []struct {
		name    string
		inserts []Point
		want    *Quadtree
	}{
		{
			name:    "simple one insert in empty field",
			inserts: []Point{{3, 4}},
			want: &Quadtree{
				TopLeft:  Point{0, 0},
				BotRight: Point{w, h},
				Node: &Node{
					Point: Point{3, 4},
					Data:  1,
				},
			},
		},
		{
			name:    "insert to top left quadrant",
			inserts: []Point{{5, 5}, {1, 2}},
			want: &Quadtree{
				TopLeft:  Point{0, 0},
				BotRight: Point{w, h},
				BotRightTree: &Quadtree{
					TopLeft:  Point{w / 2, h / 2},
					BotRight: Point{w, h},
					Node: &Node{
						Point: Point{5, 5},
						Data:  1,
					},
				},
				TopLeftTree: &Quadtree{
					TopLeft:  Point{0, 0},
					BotRight: Point{w / 2, h / 2},
					Node: &Node{
						Point: Point{1, 2},
						Data:  1,
					},
				},
			},
		},
		{
			name:    "insert to top right quadrant",
			inserts: []Point{{5, 5}, {5, 2}},
			want: &Quadtree{
				TopLeft:  Point{0, 0},
				BotRight: Point{w, h},
				BotRightTree: &Quadtree{
					TopLeft:  Point{w / 2, h / 2},
					BotRight: Point{w, h},
					Node: &Node{
						Point: Point{5, 5},
						Data:  1,
					},
				},
				TopRightTree: &Quadtree{
					TopLeft:  Point{w / 2, 0},
					BotRight: Point{w, h / 2},
					Node: &Node{
						Point: Point{5, 2},
						Data:  1,
					},
				},
			},
		},
		{
			name:    "insert to bot right quadrant",
			inserts: []Point{{1, 2}, {5, 5}},
			want: &Quadtree{
				TopLeft:  Point{0, 0},
				BotRight: Point{w, h},
				TopLeftTree: &Quadtree{
					TopLeft:  Point{0, 0},
					BotRight: Point{w / 2, h / 2},
					Node: &Node{
						Point: Point{1, 2},
						Data:  1,
					},
				},
				BotRightTree: &Quadtree{
					TopLeft:  Point{w / 2, h / 2},
					BotRight: Point{w, h},
					Node: &Node{
						Point: Point{5, 5},
						Data:  1,
					},
				},
			},
		},
		{
			name:    "insert to bot left quadrant",
			inserts: []Point{{1, 2}, {3, 5}},
			want: &Quadtree{
				TopLeft:  Point{0, 0},
				BotRight: Point{w, h},
				TopLeftTree: &Quadtree{
					TopLeft:  Point{0, 0},
					BotRight: Point{w / 2, h / 2},
					Node: &Node{
						Point: Point{1, 2},
						Data:  1,
					},
				},
				BotLeftTree: &Quadtree{
					TopLeft:  Point{0, h / 2},
					BotRight: Point{w / 2, h},
					Node: &Node{
						Point: Point{3, 5},
						Data:  1,
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQuadtree(w, h)

			for _, p := range tc.inserts {
				err := q.Insert(p.X, p.Y)
				if err != nil {
					t.Error(err)
				}
			}

			if !reflect.DeepEqual(tc.want, q) {
				t.Errorf("want %v, got %v", tc.want, q)
			}
		})
	}
}
