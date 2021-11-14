package logic

import "github.com/dembygenesis/blog_exam/pkg/store"

// Logic describes all of our modules
type Logic interface {
	ArticleAdmin
}

// ArticleAdmin is our article's module
type ArticleAdmin interface {
	Create() error
	Read(id int) error
}

// Handler is our logic implementation
type Handler struct {
	store store.Store
}

// NewLogicHandler instantiates a new handler
func NewLogicHandler(store store.Store) Logic {
	return &Handler{
		store,
	}
}
