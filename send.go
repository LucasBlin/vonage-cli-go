package main

import (
	"encoding/json"
	"errors"
	"log"
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

func Send(sender string, to string, message string, path string) error {
	err, b := CreatePayload(sender, to, message, path)
	if err != nil {
		log.Printf("Could not build required payload : %s", err)
		return err
	}

	body, err := rest.Post(config.Url, b)
	if err != nil {
		log.Printf("Could not send the request : %s", err)
		return err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Printf("An error occurred unmarshalling the res body : %s", err)
	}

	err = res.Validate()
	if err != nil {
		log.Printf("Validation error: %s", err)
	}

	config.Green("📩  SMS sent successfully !")
	remainingCredits, err := AsEuro(res.Messages[0].RemainingBalance)
	if err != nil {
		log.Printf("An error occurred  parsing the RemainingBalance : %s", err)
	}

	cost, err := AsEuro(res.Messages[0].MessagePrice)
	if err != nil {
		log.Printf("An error occurred  parsing the MessagePrice : %s", err)
	}

	if err != nil {
		log.Printf("An error occurred reading the res body : %s", err)
	} else {
		config.Yellow("💵  Credit remaining : " + remainingCredits)
		config.Yellow("💵  Cost of this sms : " + cost)
	}

	return nil
}

func CreatePayload(sender string, to string, message string, path string) (error, []byte) {
	apiKey, err := config.GetConfig("api", API_KEY, true, path)
	if err != nil {
		return err, nil
	}

	apiSecret, err := config.GetConfig("api", API_SECRET, true, path)
	if err != nil {
		return err, nil
	}

	data := Payload{
		From:      sender,
		Text:      message,
		To:        to,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err, nil
	}
	return err, b
}

func (r *Response) Validate() error {
	if r.MessageCount == "" {
		return errors.New("missing field: MessageCount")
	}
	if len(r.Messages) == 0 {
		return errors.New("missing field: Messages")
	}
	for _, message := range r.Messages {
		if message.To == "" {
			return errors.New("missing field: Message.To")
		}
	}
	return nil
}
