package landing

import "github.com/CaribouBlue/personal-website/internal/pages/landing/api"

type LandingPage struct {
	Api api.LandingApi
}

func NewLandingPage() *LandingPage {
	return &LandingPage{
		Api: api.NewLandingApi(),
	}
}
