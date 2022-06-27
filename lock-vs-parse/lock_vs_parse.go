package lockvsparse

import (
	"encoding/json"
	"sync"
)

var (
	cryptoNotionalFilterLock = sync.RWMutex{}
	cryptoNotionalFilterRaw  = "{\"upper_bound\":10,\"route\":1}"
)

var cryptoNotionalFilter *item

type item struct {
	UpperBound int `json:"upper_bound,omitempty"`
	Route      int `json:"route,omitempty"`
}

func parse_always() *item {
	var i item
	json.Unmarshal([]byte(cryptoNotionalFilterRaw), &i)
	return &i
}

func parse_once_with_lock() *item {
	cryptoNotionalFilterLock.RLock()
	currFilter := cryptoNotionalFilter
	cryptoNotionalFilterLock.RUnlock()
	if currFilter != nil {
		return currFilter
	}
	cryptoNotionalFilterLock.Lock()
	cryptoNotionalFilter = parse_always()
	currFilter = cryptoNotionalFilter
	cryptoNotionalFilterLock.Unlock()
	return currFilter
}
