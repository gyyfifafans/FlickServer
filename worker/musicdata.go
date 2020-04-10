package worker

type MusicDataParam struct {
	Username string `json:"username" orm:"size(80)"`
	Password string `json:"password" orm:"size(64)"`
}
