package robinhood

// import "errors"
import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func margin_call(trades_to_parse [][]string) [][]string {
	return computePortfolio(trades_to_parse)
}

func computePortfolio(trades_to_parse [][]string) [][]string {
	pq := make(PriorityQueue, 0)
	p := Portfolio{
		PortfolioBalances:    map[string]*InstrumentPortfolio{},
		PortfolioBalanceHeap: &pq,
	}
	heap.Init(p.PortfolioBalanceHeap)
	p.PortfolioBalances["CASH"] = &InstrumentPortfolio{
		Instrument:        "CASH",
		Quantity:          1000,
		Price:             1,
		AvailableQuantity: 1000,
		LockedQuantity:    0,
	}
	for _, trade := range trades_to_parse {
		t := MakeTrade(trade)
		if t.Side == "B" {
			p.Buy(t)
		}
		if t.Side == "S" {
			p.Sell(t)
		}
		if p.PortfolioBalances["CASH"].Quantity < 0 {
			err := p.MarginSell(t)
			if err != nil {
				PrettyPrint(p, err)
				return p.Serialize()
			}
		}
	}
	return p.Serialize()
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

type Trade struct {
	Timstamp   int
	Instrument string
	Side       string
	Quantity   int
	Price      int
}

func MakeTrade(t []string) *Trade {
	ts, _ := strconv.Atoi(t[0])
	quantity, _ := strconv.Atoi(t[3])
	price, _ := strconv.Atoi(t[4])
	return &Trade{
		ts, t[1], t[2], quantity, price,
	}
}

type InstrumentPortfolio struct {
	Instrument        string
	Quantity          int
	Price             int
	AvailableQuantity int
	LockedQuantity    int
	LockedBy          []*InstrumentPortfolio

	// for heap
	Index int
}

type Portfolio struct {
	PortfolioBalances    map[string]*InstrumentPortfolio
	PortfolioBalanceHeap *PriorityQueue
}

func (p Portfolio) MarginSell(t *Trade) error {
	// PrettyPrint(p, nil)
	idx := 0
	for p.PortfolioBalances["CASH"].Quantity < 0 {
		idx++
		// PrettyPrint(p, nil)
		cash := p.PortfolioBalances["CASH"].Quantity
		ip := p.PortfolioBalanceHeap.Peek().(*InstrumentPortfolio)
		if ip.AvailableQuantity == 0 {
			heap.Pop(p.PortfolioBalanceHeap)
			continue
		}
		if cash+(ip.AvailableQuantity*ip.Price) < 0 {
			err := p.Sell(&Trade{t.Timstamp, ip.Instrument, "S", ip.AvailableQuantity, ip.Price})
			if err != nil {
				return err
			}
		} else {
			unitsToSell := (-1 * cash) / ip.Price
			if (-1*cash)%ip.Price > 0 {
				unitsToSell++
			}
			err := p.Sell(&Trade{t.Timstamp, ip.Instrument, "S", unitsToSell, ip.Price})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p Portfolio) Serialize() [][]string {
	x := []*InstrumentPortfolio{}
	for key, val := range p.PortfolioBalances {
		if key == "CASH" || val.Quantity == 0 {
			continue
		}
		x = append(x, val)
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].Instrument < x[j].Instrument
	})
	res := [][]string{}
	res = append(res, []string{"CASH", strconv.Itoa(p.PortfolioBalances["CASH"].Quantity)})
	for _, l := range x {
		res = append(res, []string{l.Instrument, strconv.Itoa(l.Quantity)})
	}
	return res
}

func (p Portfolio) CanBuy(t *Trade) bool {
	// cash, ok := p.PortfolioBalances["CASH"]
	// if !ok {
	//     return false
	// }
	// if cash.AvailableQuantity < (t.Price * t.Quantity) {
	//     return false
	// }
	// return true
	if t.Instrument[len(t.Instrument)-1] == 'O' {
		port, ok := p.PortfolioBalances[string(t.Instrument[0:len(t.Instrument)-1])]
		if !ok || port.AvailableQuantity < t.Quantity {
			return false
		}
	}
	return true
}

func (p Portfolio) Buy(t *Trade) error {
	// if !p.CanBuy(t) {
	//     return errors.New("Cant Buy")
	// }
	inst, ok := p.PortfolioBalances[t.Instrument]
	if !ok {
		inst = &InstrumentPortfolio{
			Instrument: t.Instrument,
		}
		p.PortfolioBalances[t.Instrument] = inst
		heap.Push(p.PortfolioBalanceHeap, inst)
	}
	inst.Quantity = inst.Quantity + t.Quantity
	inst.AvailableQuantity = inst.AvailableQuantity + t.Quantity
	inst.Price = t.Price
	if t.Instrument[len(t.Instrument)-1] == 'O' {
		port, _ := p.PortfolioBalances[string(t.Instrument[0:len(t.Instrument)-1])]
		port.AvailableQuantity = port.AvailableQuantity - t.Quantity
		p.PortfolioBalanceHeap.update(port)
	}
	p.PortfolioBalanceHeap.update(inst)

	cash, _ := p.PortfolioBalances["CASH"]
	cash.Quantity = cash.Quantity - (t.Price * t.Quantity)
	cash.AvailableQuantity = cash.AvailableQuantity - (t.Price * t.Quantity)
	return nil
}

func (p Portfolio) CanSell(t *Trade) bool {
	inst, ok := p.PortfolioBalances[t.Instrument]
	if !ok {
		return false
	}
	return !(inst.AvailableQuantity < t.Quantity)
}

func (p Portfolio) Sell(t *Trade) error {
	if !p.CanSell(t) {
		return errors.New("Cant Sell")
	}
	inst, _ := p.PortfolioBalances[t.Instrument]
	inst.Quantity = inst.Quantity - t.Quantity
	inst.AvailableQuantity = inst.AvailableQuantity - t.Quantity
	inst.Price = t.Price
	p.PortfolioBalanceHeap.update(inst)
	if t.Instrument[len(t.Instrument)-1] == 'O' {
		port, _ := p.PortfolioBalances[string(t.Instrument[0:len(t.Instrument)-1])]
		port.AvailableQuantity = port.AvailableQuantity + t.Quantity
		p.PortfolioBalanceHeap.update(port)
	}
	cash, _ := p.PortfolioBalances["CASH"]
	cash.Quantity = cash.Quantity + (t.Price * t.Quantity)
	cash.AvailableQuantity = cash.AvailableQuantity + (t.Price * t.Quantity)
	return nil
}

// Priority Queue
type PriorityQueue []*InstrumentPortfolio

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].AvailableQuantity == 0 || pq[j].AvailableQuantity == 0 {
		return (pq[i].AvailableQuantity) > (pq[j].AvailableQuantity)
	}
	if pq[i].Price != pq[j].Price {
		return pq[i].Price > pq[j].Price
	}
	if pq[i].Instrument[len(pq[i].Instrument)-1] == 'O' && pq[i].Instrument[0:len(pq[i].Instrument)-1] == pq[j].Instrument {
		return true
	}
	if pq[j].Instrument[len(pq[j].Instrument)-1] == 'O' && pq[j].Instrument[0:len(pq[j].Instrument)-1] == pq[i].Instrument {
		return false
	}
	return pq[i].Instrument < pq[j].Instrument
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	ip := x.(*InstrumentPortfolio)
	ip.Index = n
	*pq = append(*pq, ip)
	// PrettyPrint(*pq, nil)
}

func (pq *PriorityQueue) Peek() interface{} {
	return (*pq)[0]
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	ip := old[n-1]
	old[n-1] = nil // avoid memory leak
	ip.Index = -1  // for safety
	*pq = old[0 : n-1]
	return ip
}

// update modifies the priority and value of an PortfolioBalance in the queue.
func (pq *PriorityQueue) update(ip *InstrumentPortfolio) {
	if ip.Index == -1 {
		heap.Push(pq, ip)
		return
	}
	heap.Fix(pq, ip.Index)
}
