package model

import "time"

type Guild struct {
	Id                  string    `db:"id"`
	Name                string    `db:"name"`
	NotificationType    string    `db:"notificationType"`
	NotificationChannel *string   `db:"notificationChannel"`
	XpDelay             int       `db:"xpDelay"`
	XpRate              float64   `db:"xpRate"`
	RoleAssignType      string    `db:"roleAssignType"`
	Locale              string    `db:"locale"`
	ServerLocaleOnly    bool      `db:"serverLocaleOnly"`
	Active              bool      `db:"active"`
	CreatedAt           time.Time `db:"createdAt"`
	UpdatedAt           time.Time `db:"updatedAt"`
}
