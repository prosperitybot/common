package model

import "time"

type GuildUser struct {
	GuildId           string    `db:"guildId"`
	Username          string    `db:"username"`
	UserId            string    `db:"userId"`
	Level             int       `db:"level"`
	Xp                int64     `db:"xp"`
	LastXpMessageSent time.Time `db:"lastXpMessageSent"`
	MessageCount      int       `db:"messageCount"`
	CreatedAt         time.Time `db:"createdAt"`
	UpdatedAt         time.Time `db:"updatedAt"`
}
