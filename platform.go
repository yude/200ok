package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var embed_color int
var status_fmt string

func Post(target string, IsUp bool) {
	for _, v := range cfg.Platforms {
		for _, value := range v {
			if value.Type == Discord {
				err := PostToDiscord(target, value.Url, IsUp)
				if err != nil {
					log.Println("[Discord]", err)
				}
			}
			if value.Type == Mastodon {
				err := PostToMastodon(target, value.Url, value.Token, IsUp)
				if err != nil {
					log.Println("[Mastodon]", err)
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
		return errors.New("failed to create json struct")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("failed to make our client work")
	}

	if resp.StatusCode == 204 {
		return nil
	} else {
		return errors.New("failed to communicate with Discord server")
	}
}

func PostToMastodon(target string, instance_url string, token string, IsUp bool) error {
	status_fmt := ""
	if IsUp {
		status_fmt = target + " is now up."
	} else {
		status_fmt = target + " is down."
	}

	val := url.Values{}
	val.Set("status", status_fmt)
	val.Set("visibility", "unlisted")

	res, err := http.PostForm(instance_url+"/api/v1/statuses?access_token="+token, val)
	if err != nil {
		return err
	}

	log.Println("Bearer " + token)

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.New("failed to retrieve error")
		}
		return fmt.Errorf("http status code: %d, body: %s", res.StatusCode, b)
	}

	return nil
}
