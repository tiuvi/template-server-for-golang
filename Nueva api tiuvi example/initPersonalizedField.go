package main



const (
	Admin     int64 = 4
	Moderator int64 = 3
	User      int64 = 2
	Visitor   int64 = 1
)

var roles = map[int64]string{
	Visitor:   "Visitante",
	User:      "Usuario",
	Moderator: "Moderador",
	Admin:     "Administrador",
}