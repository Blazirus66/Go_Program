package main

// test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Créez une nouvelle requête HTTP
	req, err := http.NewRequest("GET", "https://api.kraken.com/0/public/Time", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Envoyez la requête et récupérez la réponse
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Lisez le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Affichez le corps de la réponse
	fmt.Println(string(body))
}

//________________________________________________________________________

// Définissez une structure pour stocker les données de réponse
type AssetPairsResponse struct {
	Error  []string        `json:"error"`
	Result map[string]Pair `json:"result"`
}

type Pair struct {
	Altname           string  `json:"altname"`
	AclassBase        string  `json:"aclass_base"`
	Base              string  `json:"base"`
	AclassQuote       string  `json:"aclass_quote"`
	Quote             string  `json:"quote"`
	Lot               string  `json:"lot"`
	PairDecimals      int     `json:"pair_decimals"`
	LotDecimals       int     `json:"lot_decimals"`
	LotMultiplier     int     `json:"lot_multiplier"`
	LeverageBuy       []int   `json:"leverage_buy"`
	LeverageSell      []int   `json:"leverage_sell"`
	Fees              [][]int `json:"fees"`
	FeesMaker         [][]int `json:"fees_maker"`
	FeeVolumeCurrency string  `json:"fee_volume_currency"`
	MarginCall        int     `json:"margin_call"`
	MarginStop        int     `json:"margin_stop"`
}

func AssetPairs() {
	// Créez une nouvelle requête HTTP
	req, err := http.NewRequest("GET", "https://api.kraken.com/0/public/AssetPairs", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Envoyez la requête et récupérez la réponse
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Lisez le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Décodez le corps de la réponse en un objet de la structure AssetPairsResponse
	var data AssetPairsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Affichez la liste des paires de trading
	for name, pair := range data.Result {
		fmt.Println(name, pair.Altname)
	}
}
