package main

import(
	"github.com/google/uuid"
	"github.com/gorillla/mux"
)
type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router 
	user User
}
func NewServer() *Server {
	s:=&Server{
		Router: mux.NewRouter(),
		user: User,
	}
	return s
}


func (s *Server) routes(){
	s.HandleFunc("/adduser", s.createUser().Methods("post"))
	s.HandleFunc("/adduser", s.getUser().Methods("post"))

}
func (s *Server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		var u User
		if err:= json.NewDecoder(r.Body).Decode(&u) ; err!= nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.ID = uuid.New()
		s.user=u
		w.Header().set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
 	}
}

func (s *Server) getUser () http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		w.Header().set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.User); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}



