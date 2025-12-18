package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jnj-97/go-practice/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!=nil{
		respondwithError(w,400,fmt.Sprintf("Error parsing JSON: ",err))
		return
	}
	feed,err:=apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		Name:params.Name,
		Url: params.URL,
		UserID:user.ID,
	})
	if err!=nil{
		respondwithError(w,500,fmt.Sprintf("Error creating feed: ",err))
		return
	}

	respondWithJSON(w,201,feed)
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request){


	feeds,err:=apiCfg.DB.GetFeeds(r.Context())
	if err!=nil{
		respondwithError(w,500,fmt.Sprintf("Error Getting Feeds: ",err))
		return
	}

	respondWithJSON(w,201,feeds)
}
	

