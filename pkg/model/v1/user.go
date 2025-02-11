package v1

import (
	"fmt"
)

type User struct {
	Id       uint64 `json:"id,string"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Bot      bool   `json:"bot"`
}

func (u *User) AvatarUrl(size int) string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.webp?size=%d", u.Id, u.Avatar, size)
}
