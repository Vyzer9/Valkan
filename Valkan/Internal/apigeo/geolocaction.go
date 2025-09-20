package apigeo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GeoInfo struct {
	Status     string  `json:"status"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	ISP        string  `json:"isp"`
	Org        string  `json:"org"`
	AS         string  `json:"as"`
	Query      string  `json:"query"`
	Message    string  `json:"message"`
}

func IpGeolocation(ip string) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("[!] Erro ao consultar API:", err)
		return
	}
	defer resp.Body.Close()

	var geo GeoInfo
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		fmt.Println("[!] Erro ao decodificar resposta:", err)
		return
	}

	if geo.Status != "success" {
		fmt.Println("[!] API retornou erro:", geo.Message)
		return
	}

	fmt.Printf("IP: %s\nPaís: %s\nRegião: %s\nCidade: %s\nCEP: %s\nLat/Lon: %.4f, %.4f\nISP: %s\nOrg: %s\nAS: %s\n",
		geo.Query, geo.Country, geo.RegionName, geo.City, geo.Zip, geo.Lat, geo.Lon, geo.ISP, geo.Org, geo.AS)
}
