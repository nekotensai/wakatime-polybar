package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type GrandTotal struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float32 `json:"total_seconds"`
}

type Range struct {
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

type Data struct {
	GrandTotal       GrandTotal    `json:"grand_total"`
	Categories       []interface{} `json:"categories"`
	Dependencies     []interface{} `json:"dependencies"`
	Editors          []interface{} `json:"editors"`
	Languages        []interface{} `json:"languages"`
	Machines         []interface{} `json:"machines"`
	OperatingSystems []interface{} `json:"operating_systems"`
	Projects         []interface{} `json:"projects"`
	Range            Range         `json:"range"`
}

func main() {
	req, err := http.NewRequest("GET", "https://wakatime.com/api/v1/users/current/status_bar/today", nil)
	if err != nil {
		fmt.Println("Error in request ")
	}

	viper.AddConfigPath("wakatime-polybar/")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	token := viper.GetString("token")
	token = base64.RawStdEncoding.EncodeToString([]byte(token))
	req.Header = http.Header{"Authorization": []string{"Basic " + token + "="}}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Error in parsing")
	}
	var parsed struct {
		Data Data `json:"data"`
	}
	if err := json.Unmarshal([]byte(body), &parsed); err != nil {
		log.Fatal(err)
	}
	fmt.Println("</>", parsed.Data.GrandTotal.Hours, "hr", parsed.Data.GrandTotal.Minutes, "min")
}
