package types

type UserData struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialData struct {
	CredentialId string `json:"credential_id"`
	Title        string `json:"title"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Optional     string `json:"optional"`
	UserId       string `json:"-"`
}

type EmailConfigData struct {
	Email       string
	Password    string
	SmtpHost    string
	SmtpPort    string
	MimeHeaders string
	UIUrl       string
}
