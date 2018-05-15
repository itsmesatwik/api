package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

const RegistrationURL = config.RegistrationURL

const RegistrationFormat string = "JSON"

var RegistrationRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentRegistrationInfo",
		"GET",
		"/user/",
		alice.New(middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetCurrentRegistrationInfo).ServeHTTP,
	},
	arbor.Route{
		"CreateCurrentRegistrationInfo",
		"POST",
		"/user/",
		alice.New(middleware.AuthMiddleware([]string{"User"})).ThenFunc(CreateCurrentRegistrationInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentRegistrationInfo",
		"PUT",
		"/user/",
		alice.New(middleware.AuthMiddleware([]string{"Attendee"})).ThenFunc(UpdateCurrentRegistrationInfo).ServeHTTP,
	},
	arbor.Route{
		"GetRegistrationInfo",
		"GET",
		"/user/{id}/",
		alice.New(middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetRegistrationInfo).ServeHTTP,
	},
}

func GetCurrentRegistrationInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}

func CreateCurrentRegistrationInfo(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}

func UpdateCurrentRegistrationInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}

func GetRegistrationInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, RegistrationURL+r.URL.String(), RegistrationFormat, "", r)
}
