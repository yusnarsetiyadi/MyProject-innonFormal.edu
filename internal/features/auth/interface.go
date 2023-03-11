package auth

/*

	Interface Application

*/

type ServiceInterface interface {
	Login(input UserCore) (data UserCore, token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result UserCore, err error)
}
