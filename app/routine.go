package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)

var isRunning bool

func StartRoutine() {
	isRunning = false

	c := cron.New()
	c.AddFunc("@every 5s", func() { Check() })
	c.Start()
}

func Check() {
	if isRunning {
		return
	}

	isRunning = true

	cfg := GetConfig()
	log.Println("Checking targets...")

	for _, target := range cfg.General.Target {
		log.Println("Trying to check `" + target + "`.")
		resp, err := http.Get(target)

		if err != nil {
			if IsTargetUp(target) {
				log.Println("Target `" + target + "` is not returning 200 OK.")
				DownTarget(target)
				Post(target, false)
			} else {
				log.Println("`" + target + "` is still down.")
			}
		} else {
			if resp.StatusCode != 200 && IsTargetUp(target) {
				log.Println("Target `" + target + "` is not working properly.")
				DownTarget(target)
				Post(target, false)
			} else {
				if IsTargetUp(target) {
					log.Println("Target `" + target + "` is working fine.")
				} else {
					// If the node becomes up status
					log.Println("Target `" + target + "` is recovered from disaster.")
					IsTargetUp(target)
					Post(target, true)
					UpTarget(target)
				}
			}
		}
	}

	isRunning = false
}
