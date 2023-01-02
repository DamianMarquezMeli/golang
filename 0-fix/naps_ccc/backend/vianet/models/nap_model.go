package naps

type Data struct {
	Data []Records `json:"data"`
}

type Nap struct {
	IdRequest     int       `json:"id_request"`
	IdTransaction int       `json:"id_transaction"`
	ResultOk      bool      `json:"result_ok"`
	Errors        string    `json:"errors"`
	Records       []Records `json:"records"`
}

type Records struct {
	IdNodoArbol int    `json:"id_nodo_arbol"`
	Descripcion string `json:"descripcion"`
	Lat         string `json:"latitud"`
	Long        string `json:"longitud"`
	Disponibles int    `json:"disponibles"`
}

// Users lista de usuarios
type Naps []Nap
