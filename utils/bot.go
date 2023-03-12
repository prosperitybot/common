package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func GetMainBotId() string {
	return os.Getenv("DISCORD_APPLICATION_ID")
}

func CreateCommands(commands []discordgo.ApplicationCommand, botId string, botToken string, devGuildId string) {
	endpoint := discordgo.EndpointApplicationGlobalCommands(os.Getenv("DISCORD_APPLICATION_ID"))
	if devGuildId != "" && os.Getenv("ENV") == "dev" {
		endpoint = discordgo.EndpointApplicationGuildCommands(os.Getenv("DISCORD_APPLICATION_ID"), devGuildId)
	}

	body, err := json.Marshal(commands)
	if err != nil {
		log.Fatal("Error marshalling commands", err)
	}

	req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error creating commands", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response", err)
	}
}
