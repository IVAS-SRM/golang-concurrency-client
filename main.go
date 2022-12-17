package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const baseUrl = "http://numbersapi.com/random/year?json"

type Fact struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
	Found  bool
	Type   string
}

func getFact() (fact *Fact, err error) {
	resonse, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resonse.Body).Decode(&fact)
	if err != nil {
		return nil, err
	}

	return fact, err
}

func getFactsObautRundomYears(n int) {
	var factsMap sync.Map

	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			fact, err := getFact()
			if err != nil {
				panic(err)
			}
			factsMap.Store(idx, fact)
			fmt.Printf("Fact obout number %d\n Fact:%s\n", fact.Number, fact.Text)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func timeCount(t time.Time) {
	logrus.Println("DONE!! timeOut:", time.Since(t))
}

func main() {
	startTime := time.Now()
	defer timeCount(startTime)
	getFactsObautRundomYears(100)
}
