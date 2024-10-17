package models

type RequestBodyUserLogin struct {
	UserNameMobileNo string `json: "userNameMobileNo"`
	Password         string `json: "passWord"`
	LoginPlatform    string `json: "loginPlatform"`
	LoginIPAddress   string `json: "loginIpAddress"`
}
