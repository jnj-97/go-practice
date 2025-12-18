package main

import (
	"net/http"

	"github.com/jnj-97/go-practice/internal/database"
)

type authHnadler func(http.ResponseWriter, *http.Request,database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHnadler ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){}
}