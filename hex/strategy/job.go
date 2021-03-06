package strategy

import "github.com/caioeverest/scarlet-gopher/hex/api"

func init() {
	_strategyMap[TagJob] = func(b Context) Strategy { return &Job{b} }
}

const (
	TagJob = "job"
)

type Job struct {
	ctx Context
}

func (p *Job) Change(arg interface{}) interface{} {
	if CalcPercentage(84) {
		return api.GetRandomJob()
	}
	return arg
}
