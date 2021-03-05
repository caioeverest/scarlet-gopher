package api

import (
	_ "embed"
	"encoding/csv"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type Car struct {
	Brand string
	Model string
}

type brandMap map[string]string

const (
	brandFilePath = ".internal-db/car-brand.csv"
	modelFilePath = ".internal-db/car-model.csv"
	UnknownBrand  = "UNKNOWN"
)

func GetCarsFrom(brand string) (cars []Car) {
	for _, car := range loadCarsDB() {
		if car.Brand == brand {
			cars = append(cars, car)
		}
	}
	return
}

func GetCarsRandomCar() (car Car) {
	db := loadCarsDB()
	rand.Seed(time.Now().UnixNano())
	car = db[rand.Intn(len(db)-1)]
	return
}

func GetAnotherCarOfTheSameBrand(model string) (car Car) {
	for _, reg := range loadCarsDB() {
		if reg.Model == model {
			car = reg
		}
	}

	if reflect.ValueOf(car).IsZero() {
		return GetCarsRandomCar()
	}

	carsOfBrand := GetCarsFrom(car.Brand)
	rand.Seed(time.Now().UnixNano())
	car = carsOfBrand[rand.Intn(len(carsOfBrand)-1)]
	return
}

func loadCarsDB() (cars []Car) {
	brands := getCarBrandMap()
	modelRecords, err := readCSV(modelFilePath)
	if err != nil {
		return
	}

	for _, record := range modelRecords {
		brandName, ext := brands[record[0]]
		if !ext {
			brandName = UnknownBrand
		}

		cars = append(cars, Car{
			Brand: brandName,
			Model: record[1],
		})
	}

	return
}

func getCarBrandMap() brandMap {
	var (
		returnMap = make(brandMap)
		records   [][]string
		err       error
	)

	if records, err = readCSV(brandFilePath); err != nil {
		return nil
	}

	for _, record := range records {
		returnMap[record[0]] = record[1]
	}
	return returnMap
}

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
	if _, err = reader.Read(); err != nil {
		return [][]string{}, err
	}

	if records, err = reader.ReadAll(); err != nil {
		return [][]string{}, err
	}
	return records, nil
}
