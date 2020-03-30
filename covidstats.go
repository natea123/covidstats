package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	var country string

	flag.StringVar(&country, "country", "", "specify country for COVID case stats")
	flag.Parse()

	type Resp struct {
		Country   string
		Cases     int
		Deaths    int
		Recovered int
	}

	url := "https://corona.lmao.ninja/countries?sort=country"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var resp []Resp
	json.Unmarshal([]byte(body), &resp)

	if country != "" {
		//fmt.Println
		for _, v := range resp {
			if v.Country == country {
				fmt.Println("Country:", v.Country)
				fmt.Println("Cases:", v.Cases)
				fmt.Println("Deaths:", v.Deaths)
				fmt.Println("Recovered:", v.Recovered, "\n")
				return
			}
		}

	}
	for _, v := range resp {
		fmt.Println("Country:", v.Country)
		fmt.Println("Cases:", v.Cases)
		fmt.Println("Deaths:", v.Deaths)
		fmt.Println("Recovered:", v.Recovered, "\n")
	}

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("covidstats -country=USA\n")
		fmt.Printf("list of available countries:\n")
		for i := 0; i < len(resp); i++ {
			fmt.Printf("Country:", resp[i].Country, "\n")
		}
		flag.PrintDefaults()
	}
}
