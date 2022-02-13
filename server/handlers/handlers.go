package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"html/template"
	"microblog/pkg"
	"microblog/server/service"
	"microblog/types"
	"net/http"
	"time"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) (*Handler, error) {
	return &Handler{
		s: s,
	}, nil
}

func (h *Handler) Blog(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/blog.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ParseFiles in Blog"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "blog", nil)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ExecuteTemplate in Blog"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}
}

func (h *Handler) SaveParamsFromBlog(w http.ResponseWriter, r *http.Request) {

	loc, err := time.LoadLocation("Europe/Simferopol")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with LoadLocation in SaveParamsFromBlog"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}

	now := time.Now().In(loc).Format("2 Jan 2006 15:04:05")
	text := &types.Blog{
		Text:     r.FormValue("title"),
		Anous:    r.FormValue("anons"),
		FullText: r.FormValue("full_text"),
		Now:      now,
		Username: r.FormValue("username"),
	}

	err = h.s.AddBlog(text)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) AboutUs(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/aboutus.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ParseFiles in AboutUs"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "aboutus", nil)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ExecuteTemplate in AboutUs"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ParseFiles in Index"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}

	posts, err := h.s.GetAllBlogs()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "index", posts)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ExecuteTemplate in Index"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}
}

func (h *Handler) ShowPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ParseFiles in ShowPost"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}

	id := mux.Vars(r)["id"]
	showPost, err := h.s.GetBlog(id)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "show", showPost)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "err with ExecuteTemplate in ShowPost"))
		apiErrorEncode(w, pkg.ErrorInternalServerError)
		return
	}
}

type result struct {
	Err string `json:"error"`
}

func apiErrorEncode(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if customError, ok := err.(*pkg.CustomError); ok {
		w.WriteHeader(customError.Code)
	}

	r := result{Err: err.Error()}

	if err = json.NewEncoder(w).Encode(r); err != nil {
		pkg.LogError(err)
	}
}
