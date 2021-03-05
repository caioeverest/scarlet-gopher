package strategy

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	_strategyMap[TagYear] = func(b Context) Strategy { return &Year{b} }
}

const (
	TagYear = "year"
)

type Year struct {
	ctx Context
}

func (p *Year) Change(arg interface{}) interface{} {
	currentSTR := fmt.Sprintf("%v", arg)
	currentYear, err := strconv.Atoi(currentSTR)
	if err != nil {
		return arg
	}
	current := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)

	baseYear, ok := p.ctx[BaseYearCtx].(int)
	if !ok {
		return arg
	}

	age := int(time.Now().Sub(current).Hours() / 8760)
	someMonth := time.Month(rand.Intn(11) - 1)
	someDay := rand.Intn(30) - 1
	fakeToday := time.Date(baseYear+1, someMonth, someDay, 0, 0, 0, 0, time.UTC)
	newYear := fakeToday.AddDate(-age, 0, 0).Year()
	return strconv.Itoa(newYear)
}
