package blog

import "github.com/CaribouBlue/personal-website/internal/pages/blog/api"

type BlogPage struct {
	Api api.BlogApi
}

func NewBlogPage() *BlogPage {
	return &BlogPage{
		Api: api.NewBlogApi(),
	}
}
