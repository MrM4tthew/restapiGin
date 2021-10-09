package service


type LoginService interface {
	LoginUser(email string, password string)bool
}

type LoginInformation struct {
	email string
	password string
}

func (info *LoginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		email: "matthewbennett.mail@gmail.com",
		password:  "Pudge1001!!",
	}
}