package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateEmbed(opts *discordgo.MessageEmbed, hasError bool) *discordgo.MessageEmbed {
	if opts.Color == 0 {
		if hasError {
			opts.Color = 0xFF0000
		} else {
			opts.Color = 0x2F3136
		}
	}

	if opts.Author == nil {
		opts.Author = &discordgo.MessageEmbedAuthor{
			Name: "Prosperity",
		}
	}

	if opts.Footer == nil {
		opts.Footer = &discordgo.MessageEmbedFooter{
			Text: "Made with ❤️ by Ben#2028",
		}
	}

	return opts
}

func SendResponse(c echo.Context, msg string, ephemeral bool, isError bool) {
	var flag discordgo.MessageFlags
	if ephemeral {
		flag = discordgo.MessageFlagsEphemeral
	}

	SendComplexResponse(c, discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{CreateEmbed(&discordgo.MessageEmbed{Description: msg}, isError)},
		Flags:  flag,
	})
}

func SendComplexResponse(c echo.Context, data discordgo.InteractionResponseData) {
	c.JSON(200, discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &data,
	})
}

func SendChannelMessage(channelID string, msg string) (respMsg *discordgo.Message, err error) {
	endpoint := discordgo.EndpointChannelMessages(channelID)

	body, err := json.Marshal(discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{CreateEmbed(&discordgo.MessageEmbed{Description: msg}, false)},
	})
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(responseBody, &respMsg); err != nil {
		return nil, err
	}

	return
}
