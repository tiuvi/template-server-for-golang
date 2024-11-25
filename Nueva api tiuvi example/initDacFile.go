package main

import (
	. "dac"
	"strconv"

)


// Declaracion de campos
const (

	creatorId string = "creatorId"
)

// Declaracion de estructura para tener funciones personalizadas
type sfForum struct {
	IdMember      int64
	SfMetadata    *PublicSpaceFile
}

//Ejecucion de los campos.
var fieldsMetaData = func() []SpaceList {
	fields := NewSpaceList()
	fields = AppendSpaceList(creatorId, 8, fields) //int64

	return fields
}()

//Ejemplo de iniciar un archivo dac
func initForum(idUserCreator int64, idForum int64, idMember int64) (*sfForum, error) {

	sfMetaData, err := NewSfDeferDiskBytes(fieldsMetaData, nil, globalFolderService,
		strconv.Itoa(int(idUserCreator)), strconv.Itoa(int(idForum)),
		"forumMetaData")
	if err != nil {
		return nil, err
	}

	return &sfForum{
		IdMember:      idMember,
		SfMetadata:    sfMetaData,
	}, nil

}


