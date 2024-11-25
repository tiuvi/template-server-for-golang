package main

import (
	. "dac"
	. "dac/http"
	"net/http"
	"strconv"
	"syscall"
)

func init() {

	Routes["/"] = func(res http.ResponseWriter, req *http.Request) {

		SK := InitSpeak(res, req)

		if SK.ExistMsgUrl("terminal") {

			SK.Write(Html, []byte(
				//Mata el proceso

				`kill `+strconv.Itoa(syscall.Getpid())+` && \<br>`+
					//Limpia la consola
					`clear`+` && \<br>`+

					//Navega a la carpeta
					`cd /media/franky/tiuviweb/eukaryote/`+globalFolderService+` && \<br>`+

					//Copila el proyecto go
					`go build -o `+globalFolderService+` && \<br>`+

					//Inicia el proyecto desde el ejecutable copilado
					`./`+globalFolderService+` \<br>`+
					`-dac `+globalDacRoute+` \<br>`+
					`-folder `+globalFolderService+` \<br>`+
					`-port `+strconv.Itoa(int(globalPort))+` \<br>`+
					`-domain `+globalServeDomain+` \<br>`+
					`-email `+globalEmail+` \<br>`+
					` -start &`))
			return
		}

		err := SK.Write(Txt, []byte("profileuser"))
		if err != nil {
			println("Error al escribir en la conexion: ", err.Error())
		}

	}

}
