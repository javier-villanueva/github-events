package models

type GitHubEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt string    `json:"created_at"` // Si quieres convertirlo a tiempo, usa time.Time
}

type Actor struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID  string `json:"gravatar_id"`
	URL         string `json:"url"`
	AvatarURL   string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Payload struct {
	Action string `json:"action"`
	Forkee *Repo `json:"forkee,omitempty"`
	Commits *[]any `json:"commits,omitempty"`
}

