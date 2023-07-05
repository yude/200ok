package main

type Config struct {
	General   General
	Platforms map[string][]Platform
}

type General struct {
	Target []string
}

type PlatformType string

const (
	Discord  = PlatformType("Discord")
	Mastodon = PlatformType("Mastodon")
)

type Platform struct {
	Type  PlatformType
	Url   string
	Token string
}

type DiscordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type DiscordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type DiscordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  DiscordImg     `json:"image"`
	Thum   DiscordImg     `json:"thumbnail"`
	Author DiscordAuthor  `json:"author"`
	Fields []DiscordField `json:"fields"`
}

type DiscordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []DiscordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}

type Status struct {
	Destination string
	IfUp        bool
}
