package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/simsimler/gokoreanbots"
)


func SearchBot() {
	client := gokoreanbots.NewHTTPClient()


	prompt := promptui.Prompt{
		Label: "Bot Name",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("ERROR: Bot Name cannot be empty")
			}
			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	bot, err := client.SearchBots(result, 1)

	if err != nil {
		fmt.Printf("Koreanbots failed %v\n", err)
		return
	}

	SearchBotList := make(map[string]string)
	SelectbotList := make([]interface{}, len(bot.Data))

	for i, v := range bot.Data {
		SearchBotList[v.Name] = v.ID
		SelectbotList[i] = v.Name
	}


	promptSelect := promptui.Select{
		Label: "Select Bot",
		Items: SelectbotList,
	}

	_, result, err = promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	a, err := client.GetBot(SearchBotList[result])

	if err != nil {
		fmt.Printf("Koreanbots failed %v\n", err)
		return
	}

	println("ID: " + a.ID)
	println("Username: " + a.Name + "#" + a.Tag)
	println("Prefix: " + a.Prefix)
	println("Status: " + a.Status)
	println("Intro: " + a.Intro)
	println("Description: " + a.Description)
	println("Avatar: " + a.Avatar)
	println("Background: " + a.Background)
	println("Banner: " + a.Banner)
	println("Discord: " + a.Discord)
	println("Git: " + a.Git)
	println("Library: " + a.Library)
	println("State: " + a.State)
	println("Url: " + a.Url)
	println("Vanity: " + a.Vanity)
	println("Web: " + a.Web)
}

func GetBot() {
	prompt := promptui.Prompt{
		Label: "Bot ID",
		Validate: func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("ERROR: Invalid number")
			}
			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	a, err := gokoreanbots.NewHTTPClient().GetBot(result)

	if err != nil {
		fmt.Printf("Koreanbots failed %v\n", err)
		return
	}

	println("ID: " + a.ID)
	println("Username: " + a.Name + "#" + a.Tag)
	println("Prefix: " + a.Prefix)
	println("Status: " + a.Status)
	println("Intro: " + a.Intro)
	println("Description: " + a.Description)
	println("Avatar: " + a.Avatar)
	println("Background: " + a.Background)
	println("Banner: " + a.Banner)
	println("Discord: " + a.Discord)
	println("Git: " + a.Git)
	println("Library: " + a.Library)
	println("State: " + a.State)
	println("Url: " + a.Url)
	println("Vanity: " + a.Vanity)
	println("Web: " + a.Web)
}