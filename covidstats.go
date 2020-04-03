package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Resp struct {
	Country   string
	Cases     int
	Deaths    int
	Recovered int
}

func getData() []Resp {

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
	return resp

}

func main() {

	resp := getData()

	var country string

	flag.StringVar(&country, "country", "", "specify country for COVID case stats")
	flag.Parse()
	if country != "" {
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
	}
	flag.PrintDefaults()
}
