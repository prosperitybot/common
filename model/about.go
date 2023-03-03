package model

type AboutStats struct {
	Servers  int64 `json:"servers"`
	Users    int64 `json:"users"`
	Messages int64 `json:"messages"`
}
