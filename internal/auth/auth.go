package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string,error){
	val:=headers.Get("Authorization")
	if val==""{
		return "",errors.New("No Authentication header found")
	}
	vals:=strings.Split(val," ")
	if len(vals)!=2{
		return "",errors.New("Malformed Authorization Header")
	}
	if vals[0]!="ApiKey"{
		return "",errors.New("Malformed Authorization Header")
	}
	return vals[1], nil
}