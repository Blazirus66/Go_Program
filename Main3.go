package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Définissez une structure pour stocker les données de réponse
type TickerResponse struct {
	Error  []string          `json:"error"`
	Result map[string]Ticker `json:"result"`
}

type Ticker struct {
	Ask                        []string `json:"a"`
	Bid                        []string `json:"b"`
	LastTrade                  []string `json:"c"`
	Volume                     []string `json:"v"`
	VolumeWeightedAveragePrice []string `json:"p"`
	NumberOfTrades             []int    `json:"t"`
	Low                        []string `json:"l"`
	High                       []string `json:"h"`
	OpeningPrice               string   `json:"o"`
}

func main() {
	// Définissez la paire de trading que vous souhaitez récupérer les informations
	pair := "1INCHEUR"

	// Créez une nouvelle requête HTTP avec l'endpoint Ticker et la paire de trading en tant que paramètre de query
	req, err := http.NewRequest("GET", "https://api.kraken.com/0/public/Ticker", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	q := req.URL.Query()
	q.Add("pair", pair)
	req.URL.RawQuery = q.Encode()

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

	// Décodez le corps de la réponse en un objet de la structure TickerResponse
	var data TickerResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Récupérez les informations de cotation pour la paire de trading
	ticker := data.Result[pair]

	// Affichez les informations de cotation
	fmt.Println("Prix de demande:", ticker.Ask[0])
	fmt.Println("Prix d'offre:", ticker.Bid[0])
	fmt.Println("Dernier prix d'échange:", ticker.LastTrade[0])
	fmt.Println("Volume:", ticker.Volume[1])
	fmt.Println("Prix moyen pondéré du volume:", ticker.VolumeWeightedAveragePrice[1])
	fmt.Println("Nombre d'échanges:", ticker.NumberOfTrades[1])
	fmt.Println("Prix le plus bas de la journée:", ticker.Low[1])
	fmt.Println("Prix le plus haut de la journée:", ticker.High[1])
	fmt.Println("Prix d'ouverture:", ticker.OpeningPrice[1])

	// Créez un fichier dans le dossier Archive
	file, err := os.Create("./Archive/ticker.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Mettre en forme les données de la structure Ticker
	datas := fmt.Sprintf("Prix de demande: %s\nPrix d'offre: %s\nDernier prix d'échange: %s\nVolume: %s\nPrix moyen pondéré du volume: %s\nNombre d'échanges: %d\nPrix le plus bas de la journée: %s\nPrix le plus haut de la journée: %s\nPrix d'ouverture: %s\n",
		ticker.Ask[0], ticker.Bid[0], ticker.LastTrade[0], ticker.Volume[1], ticker.VolumeWeightedAveragePrice[1], ticker.NumberOfTrades[1], ticker.Low[1], ticker.High[1], ticker.OpeningPrice[1])

	// Écrivez les données dans le fichier
	_, err = file.WriteString(datas)
	if err != nil {
		fmt.Println(err)
		return
	}
}
