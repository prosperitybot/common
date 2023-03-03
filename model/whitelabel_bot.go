package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	WhitelabelBotActionNone    = "none"
	WhitelabelBotActionStart   = "start"
	WhitelabelBotActionStop    = "stop"
	WhitelabelBotActionRestart = "restart"
	WhitelabelBotActionDelete  = "delete"
)

type WhitelabelBot struct {
	// Ids
	Id     string  `db:"botId"`
	OldId  *string `db:"oldBotId"`
	UserId *string `db:"userId"`
	// Connection Info
	Token     string  `db:"token"`
	PublicKey *string `db:"publicKey"`
	// Actions
	Action     *string `db:"action"`
	LastAction string  `db:"lastAction"`
	// Bot Info
	Name          *string `db:"botName"`
	Discriminator *string `db:"botDiscrim"`
	AvatarHash    *string `db:"botAvatarHash"`
	// Bot Status
	StatusType    string `db:"statusType"`
	StatusContent string `db:"statusContent"`
	// Bot Settings
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (b *WhitelabelBot) FillInfoByToken() error {
	endpoint := discordgo.EndpointUser("@me")
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", b.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code was %d, expected 200", resp.StatusCode)
	}

	defer resp.Body.Close()

	var botUser discordgo.User

	// decode body of response to botUser
	if err := json.NewDecoder(resp.Body).Decode(&botUser); err != nil {
		return err
	}

	b.Id = botUser.ID
	b.Name = &botUser.Username
	b.Discriminator = &botUser.Discriminator
	b.AvatarHash = &botUser.Avatar

	return nil
}
