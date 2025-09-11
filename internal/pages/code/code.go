package code

import "github.com/CaribouBlue/personal-website/internal/pages/code/api"

type CodePage struct {
	Api api.CodeApi
}

func NewCodePage() *CodePage {
	return &CodePage{
		Api: api.NewCodeApi(),
	}
}
