package main

import (
	"fmt"
	"log"
)

func Greeting() {
	cfg := GetConfig()

	logo := `
	 ___   ___   ___          _    
	|__ \ / _ \ / _ \        | |   
	   ) | | | | | | |   ___ | | __
	  / /| | | | | | |  / _ \| |/ /
	 / /_| |_| | |_| | | (_) |   < 
	|____|\___/ \___/   \___/|_|\_\
	                       Made by yude.

    GitHub: https://github.com/yude/200ok
	Licensed under the MIT license.
   `

	fmt.Println(logo)
	log.Println("Targets:", cfg.General.Target)
}
