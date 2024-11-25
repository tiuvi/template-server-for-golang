package main

import (
	. "dac"
	. "dac/shell"
)

// fields
const (
	countForums string = "countForums"
)

// columns
const (
	id        string = "id"
	nameForum string = "nameForum"
)

// Redeclaramos una estrucura para añadirle funciones personalizadas
type sfForumList struct {
	*SpaceRamSync
}

// Variable global para el archivo
var sfListForum *sfForumList

func init(){

	RoutesDac["initListForum"] = func () {

		fields := NewSpaceList()
		fields = AppendSpaceList(countForums, 8, fields) //int64
	
		lines := NewSpaceList()
		lines = AppendSpaceList(id, 8, lines)         //int64
		lines = AppendSpaceList(nameForum, 32, lines) //string
	
		sf, err := NewSfPermBytes(fields, lines, globalFolderService, "listForum")
		if err != nil {
			ErrorStatusInternalServerError(err.Error())
		}
	
		//Sincronizamos el numero de linea con el nombre del foro
		sfRamSync, err := sf.InitRamSyncString(nameForum)
		if err != nil {
			ErrorStatusInternalServerError(err.Error())
		}
	
		sfListForum = &sfForumList{sfRamSync}
	}
	
}