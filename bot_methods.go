package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getUpdates(offset int) ([]Update, error) {
	methodName := "/getUpdates"
	url := botAPI + botToken + methodName + "?offset=" + strconv.Itoa(offset)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func sendMessage(update Update) error {
	methodName := "/sendMessage"
	botMessage := BotMessage{
		ChatId: update.Message.Chat.Id,
		Text:   update.Message.Text,
	}

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	url := botAPI + botToken + methodName

	_, err = http.Post(url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}
