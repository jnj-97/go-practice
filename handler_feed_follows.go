package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jnj-97/go-practice/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!=nil{
		respondwithError(w,400,fmt.Sprintf("Error parsing JSON: ",err))
		return
	}
	feed_follow,err:=apiCfg.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		FeedID:params.FeedID,
		UserID:user.ID,
	})
	if err!=nil{
		respondwithError(w,500,fmt.Sprintf("Error creating feed follow: ",err))
		return
	}

	respondWithJSON(w,201,feed_follow)
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){


	feed_follows,err:=apiCfg.DB.GetFeedFollows(r.Context(),user.ID)
	if err!=nil{
		respondwithError(w,500,fmt.Sprintf("Error Getting Feed Follows: ",err))
		return
	}

	respondWithJSON(w,201,feed_follows)
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feedFolloWIDStr:= chi.URLParam(r,"feed_id")
	feedFollowID, err:=uuid.Parse(feedFolloWIDStr)
	if err!=nil{
		respondwithError(w,400,fmt.Sprintf("Couldn't parse feed id",err))
		return
	}
	err=apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err!=nil{
		respondwithError(w,400,fmt.Sprintf("Couldn't delete feed follow",err))
		return
	}
	respondWithJSON(w,200,struct{ Deleted bool }{Deleted: true})
}