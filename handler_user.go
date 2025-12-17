package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jnj-97/go-practice/internal/database"
)

func (apiCfg *apiConfig) handleUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!=nil{
		respondwithError(w,400,fmt.Sprintf("Error parsing JSON: ",err))
		return
	}
	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		Name:params.Name,
	})
	if err!=nil{
		respondwithError(w,500,fmt.Sprintf("Error creating user: ",err))
		return
	}

	respondWithJSON(w,200,user)
}