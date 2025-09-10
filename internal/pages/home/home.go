package home

import "github.com/CaribouBlue/personal-website/internal/pages/home/api"

type HomePage struct {
	Api api.HomeApi
}

func NewHomePage() *HomePage {
	return &HomePage{
		Api: api.NewHomeApi(),
	}
}
