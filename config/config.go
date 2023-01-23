package Config

var Port = ":8080"
var BasePath = "/api"
var PathDocs = BasePath + "/docs"
var PathCreateUser = BasePath + "/user/{nickname}/create"
var PathProfile = BasePath + "/user/{nickname}/profile"

var Headers = map[string]string{
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
