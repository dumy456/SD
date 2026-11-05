package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// Returneaza true daca a precede b in ordinea lamport
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func main() {
	var mu sync.Mutex
	var events []Event

	recordEvent := func(e Event) {
		mu.Lock()
		events = append(events, e)
		mu.Unlock()
	}

	// Canale pentru topologia inel: P1 -> P2 -> P3 -> P1
	ch12 := make(chan Message, 10)
	ch23 := make(chan Message, 10)
	ch31 := make(chan Message, 10)

	var wg sync.WaitGroup
	wg.Add(3)

	runProcess := func(id int, in <-chan Message, out chan<- Message, rounds int) {
		defer wg.Done()
		clock := LamportClock{counter: 0}

		for r := 0; r < rounds; r++ {
			ticks := rand.Intn(3) + 1
			for t := 0; t < ticks; t++ {
				clock.Tick()
			}
			recordEvent(Event{process: id, internal_counter: clock.counter, kind: "internal"})

			if id == 1 {
				ts := clock.Send()
				recordEvent(Event{process: id, internal_counter: clock.counter, kind: "send"})
				out <- Message{sender: id, ts: ts}

				msg := <-in
				clock.Receive(msg.ts)
				recordEvent(Event{process: id, internal_counter: clock.counter, kind: "recv"})
			} else {
				msg := <-in
				clock.Receive(msg.ts)
				recordEvent(Event{process: id, internal_counter: clock.counter, kind: "recv"})

				ts := clock.Send()
				recordEvent(Event{process: id, internal_counter: clock.counter, kind: "send"})
				out <- Message{sender: id, ts: ts}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	go runProcess(1, ch31, ch12, 2)
	go runProcess(2, ch12, ch23, 2)
	go runProcess(3, ch23, ch31, 2)

	wg.Wait()

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, e := range events {
		fmt.Printf("P%d | %d | %s\n", e.process, e.internal_counter, e.kind)
	}
}
