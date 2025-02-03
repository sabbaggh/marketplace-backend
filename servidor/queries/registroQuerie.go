package queries

import (
	"github.com/sabbaggh/marketplace-backend/database"
)

func Registrar(username, password, fecha string, tec, ropa, perfumes, vide, colec bool) (int64, error) {
	result, err := database.DB.Exec(
		"INSERT INTO usuarios (username, password, fecha_nacimiento, tecnologia, ropa, perfumes, videojuegos, coleccionables) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		username,
		password,
		fecha,
		tec,
		ropa,
		perfumes,
		vide,
		colec,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
