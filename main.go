package main

import (
	"fmt"
	"modulo/rotas"
	"net/http"
)

func main() {
	rotas.CarregarRotas()
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)

}
