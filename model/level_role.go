package model

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

type LevelRole struct {
	Id        string    `db:"id"`
	GuildId   string    `db:"guildId"`
	Level     int       `db:"level"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (l *LevelRole) AddToMember(memberId string, reason string) error {
	endpoint := discordgo.EndpointGuildMemberRole(l.GuildId, memberId, l.Id)
	req, _ := http.NewRequest(http.MethodPut, endpoint, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN")))
	req.Header.Set("X-Audit-Log-Reason", reason)

	client := &http.Client{}
	if _, err := client.Do(req); err != nil {
		return err
	}

	return nil
}
