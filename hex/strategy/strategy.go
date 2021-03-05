package strategy

import (
	"errors"
	"math/rand"
	"time"
)

type Strategy interface {
	Change(interface{}) interface{}
}

type Context map[string]interface{}

type InitFunc func(Context) Strategy

const BaseYearCtx = "base-year"

var _strategyMap = make(map[string]InitFunc)

func GetStrategy(context Context, tagname string) (Strategy, error) {
	strategyRef, found := _strategyMap[tagname]
	if !found {
		return nil, errors.New("strategy not found")
	}
	return strategyRef(context), nil
}

func CalcPercentage(percentage int) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) <= percentage
}
