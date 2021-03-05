package strategy

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/caioeverest/scarlet-gopher/hex/api"
)

func init() {
	_strategyMap[TagPersonName] = func(c Context) Strategy { return &PersonName{c} }
	_strategyMap[TagBirthdate] = func(c Context) Strategy { return &PersonBirthdate{c} }
}

const (
	TagPersonName = "subject-name"
	TagBirthdate  = "birthdate"
)

type PersonName struct {
	ctx Context
}

func (p *PersonName) Change(arg interface{}) interface{} {
	current, ok := arg.(string)
	if !ok {
		return arg
	}

	ch, _ := api.GetNewCharacter()
	if ch != nil {
		if CalcPercentage(50) {
			return fmt.Sprintf("%s %s", ch.Name.First, ch.Name.Last)
		} else {
			return fmt.Sprintf("%s", ch.Name.First)
		}
	}
	return current
}

type PersonBirthdate struct {
	ctx Context
}

func (p *PersonBirthdate) Change(arg interface{}) interface{} {
	current, ok := arg.(time.Time)
	if !ok {
		return arg
	}

	baseYear, ok := p.ctx[BaseYearCtx].(int)
	if !ok {
		return arg
	}

	age := int(time.Now().Sub(current).Hours() / 8760)
	someMonth := time.Month(rand.Intn(11) - 1)
	someDay := rand.Intn(30) - 1
	fakeToday := time.Date(baseYear+1, someMonth, someDay, 0, 0, 0, 0, time.UTC)
	return fakeToday.AddDate(-age, 0, 0)
}
