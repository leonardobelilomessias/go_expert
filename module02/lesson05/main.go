//featch cep

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Endereco struct {
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
	 for _, url := range os.Args[1:]{
		req , err := http.Get(url)

		if err != nil{
			fmt.Fprintf(os.Stderr,"Error to get requisition: %v \n",err)
		}
		defer req.Body.Close()
		res, err:= io.ReadAll(req.Body)
		if err!= nil{
			fmt.Fprintf(os.Stderr,"Error to get response: %v \n",err)
		}
		var data Endereco
		err= json.Unmarshal(res,&data)
		if err!= nil{
			fmt.Fprintf(os.Stderr,"Error to get parser response: %v \n",err)
		}
		fmt.Println(data)

		file, err := os.Create("city.txt")
		if err != nil{
			fmt.Fprintf(os.Stderr,"Error to create file: %v \n",err)
		}
		defer file.Close()
		_, err= file.WriteString(fmt.Sprintf("Cep:%s , Localidade: %s, UF: %s ",data.Cep, data.Localidade, data.Uf ))
	 
	}
}