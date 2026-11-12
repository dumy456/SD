package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(other Point) float64 {
	return math.Hypot(p.X-other.X, p.Y-other.Y)
}

func Rendezvous(pos []Point, commRadius float64, tolerance float64, maxiter int) ([]Point, int) {
	robots := make([]Point, len(pos))
	copy(robots, pos)
	iter := 0
	for len(robots) > 1 && iter < maxiter {
		iter++
		nextpositions := make([]Point, len(robots))
		for i, r := range robots {
			sumX := r.X
			sumY := r.Y
			count := 1.0

			for j, other := range robots {
				if i != j && r.Distance(other) <= commRadius {
					sumX += other.X
					sumY += other.Y
					count++
				}
			}

			nextpositions[i] = Point{
				X: sumX / count,
				Y: sumY / count,
			}
		}

		copy(robots, nextpositions)

		var compressed []Point
		for _, r := range robots {
			exists := false
			for _, c := range compressed {
				if r.Distance(c) <= tolerance {
					exists = true
					break
				}
			}
			if !exists {
				compressed = append(compressed, r)
			}
		}

		robots = compressed
		fmt.Printf("iter: %d: %d roboti activi\n", iter, len(robots))
	}

	return robots, iter

}
func main() {
	var n int
	fmt.Scan(&n)

	pos := make([]Point, n)
	for i := range pos {
		fmt.Scan(&pos[i].X, &pos[i].Y)
	}
	commRadius := 15.0
	tolerance := 1e-4
	maxiter := 20

	finalRobots, steps := Rendezvous(pos, commRadius, tolerance, maxiter)

	if len(finalRobots) == 1 {
		fmt.Printf("\nrendezvous atins in %d iter la poz: (%.4f, %.4f)\n", steps, finalRobots[0].X, finalRobots[0].Y)
	} else {
		fmt.Printf("\nalgoritm oprit dupa %d iter. clustere formate: %d \n", steps, len(finalRobots))
		for i, r := range finalRobots {
			fmt.Printf(" Cluster %d: (%.4f, %.4f)\n", i+1, r.X, r.Y)
		}
	}
}
