package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

type Polumenta struct {
	Name string `json:"name" xml:"name"`
}

type ContentType int

const (
	UNKNOWN ContentType = iota
	XML
	JSON
)

var ContentTypeHeaders = map[ContentType]string{
	JSON: "application/json",
	XML:  "application/json",
}

func (p *Polumenta) serialize(ct ContentType) ([]byte, error) {
	fmt.Println(ct)
	switch ct {
	case XML:
		tmp := struct {
			*Polumenta
			XMLName struct{} `xml:"polumenta"`
		}{Polumenta: p}
		return xml.Marshal(tmp)
	case JSON:
		return json.Marshal(p)
	default:
		return nil, errors.New("Unsupported content type.")
	}
}

func detectContentType(r *http.Request) ContentType {
	accept := r.Header.Get("Accept")
	fmt.Println("accept: ", accept)
	switch accept {
	case "application/json":
		return JSON
	case "application/xml":
		return XML
	case "":
		return JSON
	default:
		return UNKNOWN
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Polumenta{
			Name: generate(),
		}
		ct := detectContentType(r)

		marshaled, err := p.serialize(ct)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", ContentTypeHeaders[ct])
		w.Write(marshaled)

		// 	var marshaled []byte
		// 	var content_type string
		// 	var marshal_err error
		// 	if rand.Intn(2) == 0 { // just for fun
		// 		marshaled, marshal_err = json.Marshal(p)
		// 		content_type = "application/json"
		// 	} else {
		// 		tmp := struct {
		// 			Polumenta
		// 			XMLName struct{} `xml:"polumenta"`
		// 		}{Polumenta: p}
		// 		marshaled, marshal_err = xml.Marshal(tmp)
		// 		content_type = "application/xml"
		// 	}

		// 	if marshal_err == nil {
		// 		w.Header().Add("Content-Type", content_type)

		// 		w.Write(marshaled)
		// 	} else {
		// 		fmt.Println(marshal_err)
		// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
		// 		return
		// 	}
	})
	http.ListenAndServe(":8000", nil)
}
