package main

import (
	"bytes"
	"log"
	"encoding/json"
	"fmt"

)

type cachorro struct {
	Nome string // `json: "nome"`
	Raca string // `json: "raca"`
	Idade uint  // `json: "idade"`
}

func main (){
	c := cachorro{"Rex", "Dalmata", 2}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

}