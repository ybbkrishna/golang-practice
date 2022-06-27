package main

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
)

// "1", "AAPL", "B", "10", "10"
type Order struct {
	Ts         int
	Instrument string
	Side       string
	Quantity   int
	Price      int // in cents
}

type Portfolio struct {
	PortfolioBalances    map[string]*PortfolioBalance
	PortfolioBalanceHeap PriorityQueue
}

func (p *Portfolio) CanSell(o *Order) bool {
	instP, ok := p.PortfolioBalances[o.Instrument]
	if !ok {
		return false
	}
	if o.Quantity > instP.AvailableBalance {
		return false
	}
	return true
}

func (p *Portfolio) CanBuy(o *Order) bool {
	cash := p.PortfolioBalances["CASH"]
	if (o.Price * o.Quantity) > cash.AvailableBalance {
		return false
	}
	return true
}

func (p *Portfolio) MarginSell(o *Order) error {
	fulfilled := false
	for !fulfilled && p.PortfolioBalanceHeap.Len() > 0 {
		instPortfolio := p.PortfolioBalanceHeap.Peek().(*PortfolioBalance)
		cash := p.PortfolioBalances["CASH"]
		if ((o.Price*o.Quantity)-cash.AvailableBalance)/instPortfolio.Price > instPortfolio.AvailableBalance {
			p.Sell(&Order{o.Ts, instPortfolio.Instrument, "S", instPortfolio.AvailableBalance, instPortfolio.Price})
			p.PortfolioBalanceHeap.Pop()
		} else {
			qty := ((o.Price * o.Quantity) - cash.AvailableBalance) / instPortfolio.Price
			mod := ((o.Price * o.Quantity) - cash.AvailableBalance) % instPortfolio.Price
			if mod > 0 {
				qty = qty + 1
			}
			p.Sell(&Order{o.Ts, instPortfolio.Instrument, "S", qty, instPortfolio.Price})
			fulfilled = true
		}
	}
	if !fulfilled {
		return errors.New("CANT FULFILL")
	}
	return nil
}

func (p *Portfolio) Buy(o *Order) {
	instPortfolio, ok := p.PortfolioBalances[o.Instrument]
	if !ok {
		instPortfolio = &PortfolioBalance{
			Instrument: o.Instrument,
		}
		p.PortfolioBalances[o.Instrument] = instPortfolio
		p.PortfolioBalanceHeap.Push(instPortfolio)
	}
	instPortfolio.AvailableBalance = instPortfolio.AvailableBalance + o.Quantity
	instPortfolio.Price = o.Price
	instPortfolio.Quantity = instPortfolio.Quantity + o.Quantity
	p.PortfolioBalanceHeap.update(instPortfolio)

	// reduce cash
	cash := p.PortfolioBalances["CASH"]
	cash.AvailableBalance = cash.AvailableBalance - (o.Quantity * o.Price)
	cash.Quantity = cash.AvailableBalance
}

func (p *Portfolio) Sell(o *Order) {
	instPortfolio := p.PortfolioBalances[o.Instrument]
	instPortfolio.AvailableBalance = instPortfolio.AvailableBalance - o.Quantity
	instPortfolio.Price = o.Price
	instPortfolio.Quantity = instPortfolio.Quantity - o.Quantity
	p.PortfolioBalanceHeap.update(instPortfolio)

	// increase cash
	cash := p.PortfolioBalances["CASH"]
	cash.AvailableBalance = cash.AvailableBalance + (o.Quantity * o.Price)
	cash.Quantity = cash.AvailableBalance
}

type PortfolioBalance struct {
	AvailableBalance int
	Quantity         int
	Price            int
	LockedBalance    int
	LockedDueTo      []*PortfolioBalance
	Instrument       string

	// for heap
	index int
}

type TradingService interface {
	executeOrder(o *Order, p *Portfolio) (*Portfolio, error)
}

type tradingService struct {
}

func (t *tradingService) executeOrder(o *Order, p *Portfolio) (*Portfolio, error) {
	if o.Side == "B" {
		if !p.CanBuy(o) {
			if err := p.MarginSell(o); err != nil {
				return nil, err
			}
		}
		// margin sell
		p.Buy(o)
		return p, nil
	}

	if o.Side == "S" {
		if !p.CanSell(o) {
			return nil, errors.New("No Bal")
		}
		p.Sell(o)
		return p, nil
	}
	return p, nil
}

func main() {
	t := &tradingService{}
	p := &Portfolio{
		PortfolioBalances:    map[string]*PortfolioBalance{},
		PortfolioBalanceHeap: []*PortfolioBalance{},
	}
	cash := &PortfolioBalance{
		AvailableBalance: 100000,
		Quantity:         100000,
		Price:            1,
		Instrument:       "CASH",
	}
	p.PortfolioBalances["CASH"] = cash

	var err error

	_, err = t.executeOrder(&Order{1, "AAPL", "B", 10, 10000}, p)
	PrettyPrint(p, err)
	_, err = t.executeOrder(&Order{2, "AAPL", "S", 2, 8000}, p)
	PrettyPrint(p, err)
	_, err = t.executeOrder(&Order{3, "GOOG", "B", 15, 2000}, p)
	PrettyPrint(p, err)
}

func PrettyPrint(x interface{}, err1 error) {
	val, err := json.MarshalIndent(x, " ", "  ")
	if err != nil {
		fmt.Println("Error :", err)
	}
	if err1 != nil {
		fmt.Println("Error1 :", err)
	}
	fmt.Println(string(val))
}

// Priority Queue
type PriorityQueue []*PortfolioBalance

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return (pq[i].AvailableBalance * pq[i].Quantity) > (pq[j].AvailableBalance * pq[j].Quantity)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	PortfolioBalance := x.(*PortfolioBalance)
	PortfolioBalance.index = n
	*pq = append(*pq, PortfolioBalance)
}

func (pq *PriorityQueue) Peek() interface{} {
	old := *pq
	n := len(old)
	PortfolioBalance := old[n-1]
	return PortfolioBalance
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	PortfolioBalance := old[n-1]
	old[n-1] = nil              // avoid memory leak
	PortfolioBalance.index = -1 // for safety
	*pq = old[0 : n-1]
	return PortfolioBalance
}

// update modifies the priority and value of an PortfolioBalance in the queue.
func (pq *PriorityQueue) update(PortfolioBalance *PortfolioBalance) {
	heap.Fix(pq, PortfolioBalance.index)
}
