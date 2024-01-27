package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

/*func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}*/

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", config.Porta),
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			)(r)))
}
