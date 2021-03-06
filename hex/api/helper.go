package api

import (
	"encoding/csv"
	"os"
)

func readCSV(fileName string) (records [][]string, err error) {
	var (
		file   *os.File
		reader *csv.Reader
	)

	if file, err = os.Open(fileName); err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader = csv.NewReader(file)
	reader.Comma = ';'
	//reader.LazyQuotes = true
	if _, err = reader.Read(); err != nil {
		return [][]string{}, err
	}

	if records, err = reader.ReadAll(); err != nil {
		return [][]string{}, err
	}
	return records, nil
}
