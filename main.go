package main

import (
	"context"
	_ "dbproject/docs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"

	deliv "dbproject/delivery"
	rep "dbproject/repository"
	usecase "dbproject/usecase"

	conf "dbproject/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)

		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	myRouter := mux.NewRouter()
	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	config, _ := pgxpool.ParseConfig(urlDB)
	config.MaxConns = 120
	db, err := pgxpool.New(context.Background(), config.ConnString())

	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	store := rep.NewStore(db)

	usecase := usecase.NewUsecase(store)

	handler := deliv.NewHandler(usecase)

	myRouter.HandleFunc(conf.PathCreateUser, handler.CreateUser).Methods(http.MethodPost)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)

	myRouter.Use(loggingMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println("cant serve", err)
	}
}
