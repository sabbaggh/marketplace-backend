package models

type User struct {
	Username              string  `json:"username"`
	Password              string  `json:"password"`
	Fecha                 string  `json:"fecha"`
	Tecnologia            bool    `json:"tec"`
	Ropa                  bool    `json:"ropa"`
	Perfumes              bool    `json:"perf"`
	Videojuegos           bool    `json:"vid"`
	Coleccionables        bool    `json:"colec"`
	CalificacionComprador float32 `json:"califComp"`
	CalificacionVendedor  float32 `json:"califVen"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
