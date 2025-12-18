package main

import (
	"fmt"
	"net/http"

	"github.com/jnj-97/go-practice/internal/auth"
	"github.com/jnj-97/go-practice/internal/database"
)

type authHnadler func(http.ResponseWriter, *http.Request,database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHnadler ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		apiKey,err:=auth.GetAPIKey(r.Header)
	if err!=nil{
		respondwithError(w,403,fmt.Sprintf("Auth Err: ",err))
		return
	}
	user,err:=apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
	if err!=nil{
		respondwithError(w,403,fmt.Sprintf("Couln't get User: ",err))
		return
	}
	handler(w,r,user)
	}
}