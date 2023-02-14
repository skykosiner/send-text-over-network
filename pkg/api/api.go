package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skykosiner/text-over-network/pkg/typer"
)

type Msg struct {
	Data string `json:"data"`
}

func SetMsg(w http.ResponseWriter, r *http.Request) {
	var msg Msg
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("There was an error reading the body", err)
	}

	json.Unmarshal(body, &msg)
	fmt.Println(msg.Data)

	typer.TypeMeDaddy(msg.Data)

	r.Body.Close()
}

func Api() {
	router := mux.NewRouter()

	router.HandleFunc("/text", SetMsg).Methods("POST")

	log.Fatal(http.ListenAndServe(":42069", router))
}
