package api

import (
	"math/rand"
	"time"
)

type Job struct {
	BusinessTitle    string
	CivilService     string
	TitleJobCategory string
}

const (
	jobsFilePath = ".internal-db/jobs.csv"
	Unemployed   = "UNEMPLOYED"
)

func GetRandomJob() string {
	records, err := readCSV(jobsFilePath)
	if err != nil {
		return Unemployed
	}

	rand.Seed(time.Now().UnixNano())
	return records[rand.Intn(len(records)-1)][0]
}
