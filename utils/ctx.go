package utils

type ContextKey string

const (
	UserIdContextKey    ContextKey = "user_id"
	GuildIdContextKey   ContextKey = "guild_id"
	ChannelIdContextKey ContextKey = "channel_id"
	BotIdContextKey     ContextKey = "bot_id"
)
