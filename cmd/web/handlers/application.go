package handlers

import "forum/internal/services"

type Application struct {
	Auth         *services.AuthService
	Registration *services.RegistrationService
	Session      *services.SessionService
}
