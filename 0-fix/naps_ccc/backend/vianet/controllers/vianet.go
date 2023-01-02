package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	m "vianet/models"
)

type Controller struct{}
type Naps []m.Nap

func (c Controller) Index() http.HandlerFunc {
	/*
		w respuesta del servidor al cliente
		r peticion del cliente al servidor
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenido a mi increible API!")
	}
}

// en funcionamiento
func (c Controller) GetNap() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("Access-Control-Allow-Origin", "*")

		j := `[{"id_action":22015,"operador_ws":1,"descripcion":" "}]`

		var url string
		url = ""
		url = "http://172.30.0.100/spi40/abn/webservice.php?request=" + encBase64(j) //<PETICIÓNENBASE64>

		resp, err := http.Get(url)
		logErrores(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		logErrores(err)

		decoBody := decoBase64(string(body))

		var c m.Naps
		var data m.Data

		err = json.Unmarshal(decoBody, &c)
		logErrores(err)

		data.Data = c[0].Records

		d, err := json.Marshal(&data)
		logErrores(err)

		w.Write(d)
	}
}

// deshabilitado
func (c Controller) GetNaps() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		/*j1 := `[{"id_action":22015,"operador_ws":1,"descripcion":"`
		params := mux.Vars(r)
		j2 := `",}]`
		jFinal := j1 + params["id"] + j2

		fmt.Println("*************jFinal:", jFinal)*/

		j := `[{"id_action":22015,"operador_ws":1,"descripcion":" "}]`

		url := "http://172.30.0.100/spi40/abn/webservice.php?request=" + encBase64(j) //<PETICIÓNENBASE64>

		//fmt.Println(url)

		//resp, err := http.Get("https://mocki.io/v1/44ad4551-06c9-447c-a258-f8f5f30e51bc")

		resp, err := http.Get(url)
		logErrores(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		logErrores(err)

		decoBody := decoBase64(string(body))

		var c m.Naps
		var data m.Data

		err = json.Unmarshal(decoBody, &c)
		logErrores(err)

		data.Data = c[0].Records

		d, err := json.Marshal(&data)
		logErrores(err)

		w.Write(d)
	}
}

func encBase64(j string) string {
	jEnc := b64.StdEncoding.EncodeToString([]byte(j))
	return jEnc
}

func decoBase64(j string) []byte {
	jDec, _ := b64.URLEncoding.DecodeString(j)
	return jDec
}

func logErrores(err error) {
	if err != nil {
		log.Fatal("ERROR!!!", err)
	}
}
