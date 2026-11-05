package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = ((t2.Sub(t1)) + (t3.Sub(t4))) / 2

	delta = (t4.Sub(t1)) - (t3.Sub(t2))
	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	// Procesare între 5 și 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	sClock := Clock{offset: 0}
	initialOffset := 250 * time.Millisecond
	cClock := Clock{offset: initialOffset}

	fmt.Println("A)")
	thetaA, deltaA := measure(cClock, sClock)
	adjustedClockA := Clock{offset: cClock.offset - thetaA}
	remainingOffsetA := adjustedClockA.offset
	if remainingOffsetA < 0 {
		remainingOffsetA = -remainingOffsetA
	}

	fmt.Printf("offset initial %v\n", initialOffset)
	fmt.Printf("theta = %v, delta = %v\n", thetaA, deltaA)
	fmt.Printf("eroare %v\n\n", remainingOffsetA)

	fmt.Println("B)")
	minDelta := time.Duration(1<<62 - 1)
	var bestTheta time.Duration

	for i := 1; i <= 10; i++ {
		theta, delta := measure(cClock, sClock)
		fmt.Printf("%2d: theta = %12v delta = %12v\n", i, theta, delta)
		if delta < minDelta {
			minDelta = delta
			bestTheta = theta
		}
	}

	adjustedClockB := Clock{offset: cClock.offset - bestTheta}
	remainingOffsetB := adjustedClockB.offset
	if remainingOffsetB < 0 {
		remainingOffsetB = -remainingOffsetB
	}

	fmt.Printf("min delta: %v theta: %v)\n", minDelta, bestTheta)
	fmt.Printf("eroare: %v\n", remainingOffsetB)
}
