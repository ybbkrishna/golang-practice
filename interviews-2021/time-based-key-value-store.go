package main

// Start June 19th 2021 3:10PM
// End June 19th 2021 3:55PM
type TimeMap struct {
	Store map[string][]TimeVal
}

type TimeVal struct {
	Ts    int
	Value string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		map[string][]TimeVal{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Store[key] = append(this.Store[key], TimeVal{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vals, ok := this.Store[key]
	if !ok {
		return ""
	}
	idx := this.findTs(vals, timestamp)
	if idx == -1 {
		return ""
	}
	return vals[idx].Value
}

func (this *TimeMap) findTs(vals []TimeVal, ts int) int {
	e := len(vals) - 1
	s := 0
	for e-s > 1 {
		mid := (s + e) / 2
		if ts > vals[mid].Ts {
			s = mid
		} else {
			e = mid
		}
	}
	if ts >= vals[e].Ts {
		return e
	}
	if ts >= vals[s].Ts {
		return s
	}
	return -1
}
