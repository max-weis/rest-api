package app

func (a *App) routes(){
	a.Router.HandleFunc("/api/v1/book",a.handleBookGet()).Methods("GET")
	a.Router.HandleFunc("/api/v1/book",a.handleBookGetList()).Methods("GET")
	a.Router.HandleFunc("/api/v1/book",a.handleBookCreate()).Methods("POST")
	a.Router.HandleFunc("/api/v1/book",a.handleBookUpdate()).Methods("PUT")
	a.Router.HandleFunc("/api/v1/book",a.handleBookDelete()).Methods("DELETE")
}