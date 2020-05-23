package posts

import (
	"encoding/json"
	"net/http"

	"github.com/3ema208/pythontask/internal/app/model"
	"github.com/3ema208/pythontask/internal/app/store"
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

	p.router.HandleFunc("/posts/comment", p.handlerComments)
}

// handlePosts ...
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
		postObj, err := p.store.Post().Create(&model.Post{
			Title:      req.Title,
			Link:       req.Link,
			AuthorName: req.Author,
		})
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

func (p *APIPosts) handlerComments(w http.ResponseWriter, r *http.Request) {

}
