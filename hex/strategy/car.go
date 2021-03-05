package strategy

import (
	"math/rand"
	"strings"
	"time"

	"github.com/caioeverest/scarlet-gopher/hex/api"
)

func init() {
	_strategyMap[TagCarBrand] = func(b Context) Strategy { return &CarBrand{b} }
	_strategyMap[TagCarModel] = func(b Context) Strategy { return &CarModel{b} }
}

const (
	TagCarBrand = "car-brand"
	TagCarModel = "car-model"
)

type CarBrand struct {
	ctx Context
}

func (c *CarBrand) Change(arg interface{}) interface{} {
	var (
		brand string
		ok    bool
	)

	if brand, ok = arg.(string); !ok {
		return arg
	}

	if len(api.GetCarsFrom(strings.ToUpper(brand))) == 0 {
		newBrand := api.GetCarsRandomCar().Brand
		c.ctx[TagCarBrand] = newBrand
		return newBrand
	}

	if CalcPercentage(3) {
		newBrand := api.GetCarsRandomCar().Brand
		c.ctx[TagCarBrand] = newBrand
		return newBrand
	}
	return arg
}

type CarModel struct {
	ctx Context
}

func (c *CarModel) Change(arg interface{}) interface{} {
	var (
		cars  []api.Car
		model string
		brand string
		ok    bool
	)

	rand.Seed(time.Now().UnixNano())
	if model, ok = arg.(string); !ok {
		return arg
	}

	if brand, ok = c.ctx[TagCarBrand].(string); !ok {
		brand = api.UnknownBrand
	}

	if cars = api.GetCarsFrom(strings.ToUpper(brand)); len(cars) == 0 {
		return api.GetAnotherCarOfTheSameBrand(strings.ToUpper(model)).Model
	}

	if CalcPercentage(74) {
		return cars[rand.Intn(len(cars)-1)].Model
	}
	return model
}
