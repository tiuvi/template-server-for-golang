package main

import (
	."dac"
	."dac/shell"
	. "dac/buildspace"
	. "dac/dan"
	"net/http"
	"strconv"
)




// Variables de configuracion dac, 
// startDac.go


//Variables de configuracion dab, 
// main.go, getArgs.go
var dab *Dab
var globalDacRoute string
var globalServeDomain string
var globalPort int64

var globalEmail string
var globalFolderService string

//Estos son los nombres de los argumentos que se necesitan para ejecutar el proceso
const (
	//Ruta principal de la carpeta y carpeta donde funcionara la app
	ArgDacReq   = "dac"
	ArgFolderService   = "folder"	
	ArgPort   = "port"	

	//Requisitos para certbot
	ArgDomainReq = "domain"
	ArgEmailReq = "email"

	//Inicia el servidor
	RouteStart = "start"
)

/*
	Funcion principal de la aplicacion
*/
func main() {
	
	dab = InitDabSpace([]Flags{
		
		// Se crean las banderas y se ponde si son requeridos o no.
		BuildFlagReq(ArgDacReq, "Obtiene la ruta dac"),
		BuildFlagReq(ArgFolderService, "Obtiene la carpeta principal"),
		BuildFlagReq(ArgPort, "Obtiene el puerto de operacion"),
		
		BuildFlagReq(ArgDomainReq, "Dominio donde se ejecutara el servicio de contenido"),
		BuildFlagReq(ArgEmailReq, "Email para generar el certificado ssl"),

		BuildNewChainMidleware(RouteStart, "Inicia el servicio", []DabMidleFunc{
			//Primero se obtienen los argumentos por linea de comandos
			getArgs,
			//Luego se inicia la base de datos
			startDac,
			//Por ultimo se inicia el servidor
		}, start),
	})

}

//Funcion para obtener los argumentos del sistema
func getArgs(handle DabHandlerFunc) DabHandlerFunc {
	return func(dab *Dab) {

		globalDacRoute = dab.GetFlagReq(ArgDacReq)

		globalFolderService = dab.GetFlagReq(ArgFolderService)

		globalPort = dab.GetFlagInt64Req(ArgPort)

		globalServeDomain = dab.GetFlagReq(ArgDomainReq)
		
		globalEmail = dab.GetFlagReq(ArgEmailReq)

		handle(dab)
	}
}

var RoutesDac = make(map[string]func())

//Funcion para iniciar la carpeta y los archivos permanentes
func startDac(handle DabHandlerFunc) DabHandlerFunc {
	return func(dab *Dab) {
		NewBasicDac(globalDacRoute)

		for name , fileDac := range RoutesDac {
			println("iniciando: " , name)
			fileDac()
		}
		
		handle(dab)
	}
}



var Routes = make(map[string]func(http.ResponseWriter, *http.Request))

func start(*Dab) {

	dan := NewDanSpace(uint16(globalPort))

	urlTempFullchain, urlTempPrivate, err := GetCerts(globalServeDomain, globalServeDomain, globalEmail)
	if err != nil {
		ErrorStatusInternalServerError(err.Error())
	}

	//Cargamos todas las rutas.
	for path , route := range Routes {
		dan.NewRoute(path , route)
	}

	//Dejamos la direccion de enlace donde podemos obtener todos los datos
	println(`Reiniciar:` + globalServeDomain +`:`+strconv.Itoa(int(globalPort)) + `/?terminal`  )

	dan.NewDan(urlTempFullchain, urlTempPrivate)
}

