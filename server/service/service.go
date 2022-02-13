package service

import (
	"github.com/pkg/errors"
	"microblog/pkg"
	"microblog/postgres"
	"microblog/types"
)

type Service struct {
	p *postgres.PostgreS
}

func NewService(pg *postgres.PostgreS) (*Service, error) {
	return &Service{
		p: pg,
	}, nil
}

func (s *Service) AddBlog(text *types.Blog) error {
	if text.Text == "" || text.Anous == "" || text.FullText == "" {
		return pkg.ErrorNotData
	}

	err := s.p.AddBlog(text)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with AddBlog in AddBlog"))
		return pkg.ErrorInternalServerError
	}

	return nil
}

func (s *Service) GetBlog(id string) (*types.Blog, error) {

	blog, err := s.p.GetBlogByID(id)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with GetBlogByID in GetBlog"))
		return nil, pkg.ErrorInternalServerError
	}

	return blog, nil
}

func (s *Service) GetAllBlogs() ([]types.Blog, error) {

	posts, err := s.p.GetAllBlogs()
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with GetAllBlogs in GetAllBlogs"))
		return nil, pkg.ErrorInternalServerError
	}

	return posts, nil
}
