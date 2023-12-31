package database

type ICredentials interface {
	GetAllCredentials()
	GetCredential(uint64)
	AddCredential()
	UpdateCredential(uint64)
	DeleteCredential(uint64)
	DeleteAllCredentials()
}

type IUser interface {
	AddUser()
	GetUser(string, string)
	DeleteUser(string, string)
	UpdateUser(string, string)
}

type ISession interface {
	UpdateSession(string) //userId and sessionToken
	ValidateSession(string)
}
