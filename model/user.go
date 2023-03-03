package model

import "time"

type User struct {
	Id            string    `db:"id"`
	Username      string    `db:"username"`
	AccessLevels  *[]string `db:"access_levels"`
	Discriminator string    `db:"discriminator"`
	PremiumSource *string   `db:"premium_source"`
	Locale        *string   `db:"locale"`
	IsAdmin       bool      `db:"is_admin"`
	IsSupport     bool      `db:"is_support"`
	IsTranslator  bool      `db:"is_translator"`
	PremiumStatus string    `db:"premium_status"`
	CreatedAt     time.Time `db:"createdAt"`
	UpdatedAt     time.Time `db:"updatedAt"`
}
