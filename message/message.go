package message

const (
	MESSAGE_TYPE_LOGIN = "login_message"
)

type Message struct {
	Type  string `json:"type"`
	Data string `json:"data"`
}

// 登录信息
type LoginMessage struct {
	UserId int64 `json:"user_id"`
	Password string `json:"password"`
}