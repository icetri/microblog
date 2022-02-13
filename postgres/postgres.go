package postgres

import (
	"database/sql"
	"github.com/pkg/errors"
	"microblog/types"
)

type PostgreS struct {
	db *sql.DB
}

func NewSQL(psqlInfo string) (*PostgreS, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "err with Open DB")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "err with ping DB")
	}

	return &PostgreS{db}, nil
}

func (p *PostgreS) Close() error {
	return p.db.Close()
}

func (p *PostgreS) GetBlogByID(id string) (*types.Blog, error) {

	var showPost types.Blog
	res, err := p.db.Query("SELECT id, title, anous, full_text, datenow, username FROM blog WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err = res.Scan(&showPost.ID, &showPost.Text, &showPost.Anous, &showPost.FullText, &showPost.Now, &showPost.Username)
		if err != nil {
			return nil, err
		}
	}

	return &showPost, nil
}

func (p *PostgreS) GetAllBlogs() ([]types.Blog, error) {

	posts := make([]types.Blog, 0)
	res, err := p.db.Query("SELECT id, title, anous, full_text, datenow, username FROM blog")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var text types.Blog
		err = res.Scan(&text.ID, &text.Text, &text.Anous, &text.FullText, &text.Now, &text.Username)
		if err != nil {
			return nil, err
		}

		posts = append(posts, text)
	}

	return posts, nil
}

func (p *PostgreS) AddBlog(blog *types.Blog) error {

	_, err := p.db.Exec("INSERT INTO blog (title, anous, full_text, datenow, username) VALUES ($1, $2, $3,$4,$5)", blog.Text, blog.Anous, blog.FullText, blog.Now, blog.Username)
	if err != nil {
		return err
	}

	return nil
}
