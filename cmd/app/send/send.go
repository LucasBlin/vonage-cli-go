package send

import (
	"encoding/json"
	"log"
	"lucas.blin/v2/cmd"
	"lucas.blin/v2/config"
	"lucas.blin/v2/transport/rest"
)

const API_SECRET = "API-SECRET"
const API_KEY = "API-KEY"

type Payload struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type Message struct {
	To               string `json:"to"`
	MessageID        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	ClientRef        string `json:"client-ref"`
	AccountRef       string `json:"account-ref"`
}

type Response struct {
	MessageCount string    `json:"message-count"`
	Messages     []Message `json:"messages"`
}

func Send(sender string, to string, message string) error {
	apiKey, err := config.GetConfig("api", API_KEY, true)
	if err != nil {
		return err
	}

	apiSecret, err := config.GetConfig("api", API_SECRET, true)
	if err != nil {
		return err
	}

	data := Payload{
		From:      sender,
		Text:      message,
		To:        to,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	body, err := rest.Post("https://rest.nexmo.com/sms/json", bytes)
	if err != nil {
		return err
	}

	config.Green("📩  SMS sent successfully !")

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("An error occurred unmarshalling the response body : %s", err)
	}

	remainingCredits, err := cmd.AsEuro(response.Messages[0].RemainingBalance)
	if err != nil {
		log.Printf("An error occurred  parsing the RemainingBalance : %s", err)
	}
	cost, err := cmd.AsEuro(response.Messages[0].MessagePrice)
	if err != nil {
		log.Printf("An error occurred  parsing the MessagePrice : %s", err)
	}

	if err != nil {
		log.Printf("An error occurred reading the response body : %s", err)
	} else {
		config.Yellow("💵  Credit remaining : " + remainingCredits)
		config.Yellow("💵  Cost of this sms : " + cost)
	}

	return nil
}
