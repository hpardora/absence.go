package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Telegram struct {
	Token    string
	ChatId   string
	ChatName string
}

func New(token string, chatId string, chatName string) *Telegram {
	return &Telegram{
		Token:    token,
		ChatId:   chatId,
		ChatName: chatName,
	}
}

func (t *Telegram) getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", t.Token)
}

func (t *Telegram) SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", t.getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": t.ChatId,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	// Return
	return true, nil
}
