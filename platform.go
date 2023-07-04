package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var embed_color int
var status_fmt string

func Post(target string, IsUp bool) {
	for _, v := range cfg.Platforms {
		for _, value := range v {
			if value.Type == Discord {
				err := PostToDiscord(target, value.Url, IsUp)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func PostToDiscord(target string, url string, IsUp bool) error {
	if url == "" {
		return nil
	}

	if IsUp {
		embed_color = 0x24FD4D
		status_fmt = "Up"
	} else {
		embed_color = 0xF85A5A
		status_fmt = "Down"
	}

	dw := &DiscordWebhook{
		UserName: "200ok",
	}
	dw.Embeds = []DiscordEmbed{
		DiscordEmbed{
			Title: status_fmt,
			Desc:  target,
			Color: embed_color,
		},
	}

	j, err := json.Marshal(dw)
	if err != nil {
		return errors.New("Failed to create json struct.")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Failed to make our client work.")
	}

	if resp.StatusCode == 204 {
		return nil
	} else {
		return errors.New("Failed to communicate with Discord server.")
	}
}
