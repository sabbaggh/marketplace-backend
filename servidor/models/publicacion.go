package models

type Publicacion struct {
	Titulo      string  `json:"titulo"`
	Descripcion string  `json:"descripcion"`
	TipoEntrega string  `json:"entrega"`
	Categoria   string  `json:"categoria"`
	Precio      float32 `json:"precio"`
	Ubicacion   string  `json:"ubicacion"`
}
