package test

var Env EnvironmentInfo

type EnvironmentInfo struct {
	BaseURL     string
	Credentials struct {
		Scope    string
		Type     string
		Username string
		Password string
	}
}
