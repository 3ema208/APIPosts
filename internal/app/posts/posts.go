package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/3ema208/APIPosts/internal/app/model"
	"github.com/3ema208/APIPosts/internal/app/store"
	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// New ..
func New(config *Config) *APIPosts {
	return &APIPosts{
		config: config,
		router: mux.NewRouter(),
	}
}

// APIPosts ...
type APIPosts struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

// Start ..
func (p *APIPosts) Start() error {
	if err := p.configStore(); err != nil {
		return err
	}
	p.configRouter()
	log.Info("Start api posts")
	return http.ListenAndServe(p.config.BindAddr, p.router)
}

// configStore ...
func (p *APIPosts) configStore() error {
	log.Info("Configurate Store")
	st := store.New(p.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	p.store = st
	return nil
}

// configRouter ..
func (p *APIPosts) configRouter() {
	log.Info("Configurate Router")
	p.router.HandleFunc("/posts", p.handlePosts()).Methods("GET")
	p.router.HandleFunc("/posts", p.handleCreatePost()).Methods("POST")
	p.router.HandleFunc("/post/{IDPost:[0-9]+}", p.handlePost).Methods("GET")
	p.router.HandleFunc("/post/{IDPost:[0-9]+}", p.handlerPutPost()).Methods("PUT")
	p.router.HandleFunc("/post/{IDPost:[0-9]+}", p.handlerDeletePost).Methods("DELETE")

	p.router.HandleFunc("/post/{IDPost:[0-9]+}/comments", p.handlerComments()).Methods("GET")
	p.router.HandleFunc("/post/{IDPost:[0-9]+}/comments", p.handlerCreateComment()).Methods("POST")
}

func (p *APIPosts) handlePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		posts, err := p.store.Post().Get()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		json.NewEncoder(w).Encode(posts)
	}
}

func (p *APIPosts) handlePost(w http.ResponseWriter, r *http.Request) {
	idPost := mux.Vars(r)["IDPost"]
	post, err := p.store.Post().FindByID(idPost)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (p *APIPosts) handleCreatePost() http.HandlerFunc {
	type Req struct {
		Title  string `json:"title"`
		Link   string `json:"link"`
		Author string `json:"author"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Req{}
		// todo validate fields
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		postObj := &model.Post{
			Title:      req.Title,
			Link:       req.Link,
			AuthorName: req.Author,
		}
		err := p.store.Post().Create(postObj)
		if err != nil {
			// log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(postObj)
	}
}

func (p *APIPosts) handlerPutPost() http.HandlerFunc {
	type Req struct {
		Title  string `json:"title"`
		Link   string `json:"link"`
		Author string `json:"author"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		idPost, _ := strconv.Atoi(mux.Vars(r)["IDPost"])
		req := &Req{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		postObj := &model.Post{
			ID:         idPost,
			Title:      req.Title,
			Link:       req.Link,
			AuthorName: req.Author,
		}
		if err := p.store.Post().Update(postObj); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(postObj)
	}
}

func (p *APIPosts) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	idPost := mux.Vars(r)["IDPost"]
	post, err := p.store.Post().Delete(idPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return 
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(post)
}

func (p *APIPosts) handlerComments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idPost := mux.Vars(r)["IDPost"]
		comments, errComm := p.store.Comments().Get(idPost)
		if errComm != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": errComm.Error()})
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(comments)
	}
}

func (p *APIPosts) handlerCreateComment() http.HandlerFunc {
	type ReqComment struct {
		Author  string `json:"author"`
		Content string `json:"content"`
		PostID  int    `json:"postId"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		reqComm := &ReqComment{}
		if err := json.NewDecoder(r.Body).Decode(reqComm); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		comment := &model.Comment{Author: reqComm.Author, Content: reqComm.Content, PostID: reqComm.PostID}
		if err := p.store.Comments().Create(comment); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(comment)
	}
}
