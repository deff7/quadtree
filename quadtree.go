package main

import (
	"errors"
)

type Point struct {
	X, Y int
}

type Node struct {
	Point
	Data int
}

type Quadtree struct {
	Node              *Node
	TopLeft, BotRight Point

	TopLeftTree  *Quadtree
	TopRightTree *Quadtree
	BotRightTree *Quadtree
	BotLeftTree  *Quadtree
}

func NewQuadtree(w, h int) *Quadtree {
	return &Quadtree{
		TopLeft:  Point{0, 0},
		BotRight: Point{w, h},
	}
}

func (q *Quadtree) BreadthFirst(fn func(q *Quadtree)) {
	if q == nil {
		return
	}
	fn(q)
	q.TopLeftTree.BreadthFirst(fn)
	q.TopRightTree.BreadthFirst(fn)
	q.BotRightTree.BreadthFirst(fn)
	q.BotLeftTree.BreadthFirst(fn)
}

func (q *Quadtree) Insert(x, y int) error {
	if err := q.checkBounds(x, y); err != nil {
		return err
	}

	if q.Node == nil && q.isNoChildren() {
		q.Node = &Node{
			Point: Point{x, y},
			Data:  1,
		}
		return nil
	}

	dest := q.getQuadrant(x, y)
	dest.Insert(x, y)

	if q.Node != nil && !q.isPoint() {
		x, y = q.Node.Point.X, q.Node.Point.Y
		dest = q.getQuadrant(x, y)
		dest.Insert(x, y)
		q.Node = nil
	}

	return nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (q Quadtree) isPoint() bool {
	return abs(q.TopLeft.X-q.BotRight.X) == 1 && abs(q.TopLeft.Y-q.BotRight.Y) == 1
}

func (q Quadtree) isNoChildren() bool {
	return q.TopLeftTree == nil && q.TopRightTree == nil && q.BotRightTree == nil && q.BotLeftTree == nil
}

func (q *Quadtree) getQuadrant(x, y int) *Quadtree {
	var dest *Quadtree
	midX := (q.BotRight.X + q.TopLeft.X) / 2
	midY := (q.BotRight.Y + q.TopLeft.Y) / 2

	// Top left
	if x < midX && y < midY {
		if q.TopLeftTree == nil {
			q.TopLeftTree = &Quadtree{
				TopLeft:  Point{q.TopLeft.X, q.TopLeft.Y},
				BotRight: Point{midX, midY},
			}
		}
		dest = q.TopLeftTree
	}
	// Top right
	if x >= midX && y < midY {
		if q.TopRightTree == nil {
			q.TopRightTree = &Quadtree{
				TopLeft:  Point{midX, q.TopLeft.Y},
				BotRight: Point{q.BotRight.X, midY},
			}
		}
		dest = q.TopRightTree
	}
	// Bot right
	if x >= midX && y >= midY {
		if q.BotRightTree == nil {
			q.BotRightTree = &Quadtree{
				TopLeft:  Point{midX, midY},
				BotRight: Point{q.BotRight.X, q.BotRight.Y},
			}
		}
		dest = q.BotRightTree
	}
	// Bot left
	if x < midX && y >= midY {
		if q.BotLeftTree == nil {
			q.BotLeftTree = &Quadtree{
				TopLeft:  Point{q.TopLeft.X, midY},
				BotRight: Point{midX, q.BotRight.Y},
			}
		}
		dest = q.BotLeftTree
	}
	return dest
}

func (q Quadtree) checkBounds(x, y int) error {
	if x < q.TopLeft.X {
		return errors.New("x out of bounds: less than top left corner")
	}
	if x > q.BotRight.X {
		return errors.New("x out of bounds: greater than bottom right corner")
	}
	if y < q.TopLeft.Y {
		return errors.New("y out of bounds: less than top left corner")
	}
	if y > q.BotRight.Y {
		return errors.New("y out of bounds: greater than bottom right corner")
	}
	return nil
}
