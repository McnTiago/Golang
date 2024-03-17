// Servidor Client é relacionado ao que o usuário precisa
// nesse código vamos emitir uma solicitação a um servidor http

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//emita uma solicitação HTTP GET para um servidor http.Get é um atalho conveniente
	// para criar um client serve objeto e chamar seu (Get) método;
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// ele usa o objeto que tem configurações padrão úteis, como predefinição
	// Imprima o status da resposta HTTP.
	fmt.Println("Reponse status:", resp.Status) //O HTTP (Hypertesxt Tranfer Protocol, RFC 2616)
	// é o protocolo responsável por fazer a comunicação entre o cliente e o servidor.
	// Desta forma, a cada "solicitação" feita o HTTP responde se você obteve sucesso,
	//se não, se há algum erro na página, etc. Estas mensagens de erro são os
	// "status HTTP". Por exemplo os erros de HTTP mais comuns são: errp http 404,
	// erro http 500 e erro http 401.

	// Imprima as primeiras 5 linas do corpo de resposta.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
