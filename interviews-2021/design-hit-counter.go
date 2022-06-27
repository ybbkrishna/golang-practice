package main

import "fmt"

// start: Jun 18 2021 2:06 PM
// end: Jun 18 2021 2:49 PM

type HitCounter struct {
	State []State
}

type State struct {
	Ts         int
	ClickCount int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		State: []State{},
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.resetTo(timestamp - 300)
	if len(this.State) > 0 {
		last := this.State[len(this.State)-1]
		if last.Ts == timestamp {
			last.ClickCount = last.ClickCount + 1
			this.State[len(this.State)-1] = last
			return
		}
	}
	this.State = append(this.State, State{timestamp, 1})
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	this.resetTo(timestamp - 300)
	val := 0
	for i := 0; i < len(this.State); i++ {
		val += this.State[i].ClickCount
	}
	return val
}

func (this *HitCounter) resetTo(timestamp int) {
	for i := 0; i < len(this.State); i++ {
		if this.State[i].Ts > timestamp {
			this.State = this.State[i:]
			return
		}
	}
	this.State = []State{}
}

func main() {
	x := Constructor()
	x.Hit(2)
	x.Hit(3)
	x.Hit(4)
	fmt.Println(x.GetHits(500))
}
