package getbalance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	url := ""

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	rst, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(rst))
}

type address struct {
	Addrs []string `json:"addresses"`
}

type account struct {
	Addr      string  `json:"addr"`
	Available float64 `json:"available_balance"`
	Frozen    float64 `json:"frozen_assets"`
	Total     float64 `json:"total_balance"`
}

type result struct {
	TotalCount int       `json:"total_count"`
	GdxAccount []account `json:"gdx_account"`
}

func post(addrs []string) float64 {
	url := ""
	data := address{
		Addrs: addrs,
	}
	s, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(s))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(s))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//rst, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(rst))

	var rst2 result
	err = json.NewDecoder(resp.Body).Decode(&rst2)
	if err != nil {
		log.Fatalln(err)
	}

	var sum float64
	for _, value := range rst2.GdxAccount {
		sum += value.Total
	}
	return sum

}

func GetAddrTotalBal(addrs []string) float64 {
	return post(addrs)
}

func main() {
	//get()
}
