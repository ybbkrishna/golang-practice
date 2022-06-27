package main

type Trade struct {
	Instrument string
	Side       string
	Quantity   string
	Id         string
}

func (t *Trade) hash() string {
	return t.Id + t.Instrument + t.Quantity + t.Side
}

func RemoveIntersectionTrades(houseTrades []*Trade, streetTrades []*Trade) []*Trade {
	// preprocess one with less trades
	store := map[string][]*Trade{}
	for _, t := range houseTrades {
		store[t.hash()] = append(store[t.hash()], t)
	}

	var matching []*Trade

	var onlyStreetTrades []*Trade

	for _, t := range streetTrades {
		if x, ok := store[t.hash()]; ok && len(x) > 0 {
			store[t.hash()] = x[0 : len(x)-1]
			matching = append(matching, t)
		} else {
			onlyStreetTrades = append(onlyStreetTrades, t)
		}
	}

	var onlyHouseTrades []*Trade
	for _, t := range store {
		onlyHouseTrades = append(onlyHouseTrades, t...)
	}
	return append(onlyHouseTrades, onlyStreetTrades...)
}
