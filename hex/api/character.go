package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Character struct {
	Gender    string             `json:"gender"`
	Name      CharacterName      `json:"name"`
	Email     string             `json:"email"`
	Birthdate CharacterBirthdate `json:"dob"`
	Phone     string             `json:"phone"`
	Cell      string             `json:"cell"`
	Nat       string             `json:"nat"`
}

type CharacterName struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type CharacterBirthdate struct {
	Date time.Time `json:"date"`
	Age  int       `json:"age"`
}

type Response struct {
	Results []Character `json:"results"`
}

func GetNewCharacter() (ch *Character, err error) {
	var (
		url  = "https://randomuser.me/api/"
		or   Response
		req  *http.Request
		res  *http.Response
		body []byte
	)

	ch = new(Character)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	req.Header.Add("cache-control", "no-cache")
	res, err = http.DefaultClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return
	}

	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &or); err != nil {
		return nil, err
	}

	ch = &or.Results[0]
	return
}
