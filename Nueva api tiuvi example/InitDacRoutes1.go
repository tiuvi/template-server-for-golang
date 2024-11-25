package main

import (
		. "dac"
	. "dac/http"
	"net/http"
)

func init(){

	Routes["/forum"] = func(res http.ResponseWriter, req *http.Request){

		SK := InitSpeak(res, req)

		err := SK.Write(Txt, []byte("forum"))
		if err != nil {
			println("Error al escribir en la conexion: ", err.Error())
		}
	}
	
}