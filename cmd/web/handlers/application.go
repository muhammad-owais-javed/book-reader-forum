package handlers

import "forum/internal/services"

type Application struct {
	Auth     *services.AuthService
	Register *services.RegistrationService
}
