package main

import (
		. "dac"
	. "dac/http"
	"net/http"
)

func init(){

	Routes["/forumuser"] = func(res http.ResponseWriter, req *http.Request){

		SK := InitSpeak(res, req)

		err := SK.Write(Txt, []byte("forumuser"))
		if err != nil {
			println("Error al escribir en la conexion: ", err.Error())
		}
	}
	
}