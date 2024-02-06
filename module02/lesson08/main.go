// Creating function to featch cet
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main(){
	http.HandleFunc("/", FeatchCepHandler)
	http.ListenAndServe(":8080", nil)
}

func  FeatchCepHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path!= "/"{
		res.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := req.URL.Query().Get("cep")
	if cepParam == ""{
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("hello world!"))
}

func FeatchCep(cep string)(*ViaCep, error){
	resp , error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error!= nil{
		return nil, error
	}
	defer resp.Body.Close()
	body , error:= io.ReadAll(resp.Body)
	if error!= nil{
		return nil, error
	}
	var c ViaCep
	error = json.Unmarshal(body,&c)
	if error != nil{
		return nil, error
	}
	return &c, nil
}