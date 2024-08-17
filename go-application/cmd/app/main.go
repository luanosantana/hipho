package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", handler)
	http.ListenAndServe(":8087", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Coleta a URL da próxima rota da variável de ambiente
	nextURL := os.Getenv("HOST_REQUEST")
	version := os.Getenv("VERSION")

	fmt.Println("mais 1")

	header_version := r.Header.Get("version")

	var show_version string
	if header_version == "" {
		show_version = version
	} else {
		show_version = header_version
	}

	if nextURL == "" {
		w.Write([]byte("header version: " + show_version))
		return
	}

	if nextURL == "" || version == "" {
		http.Error(w, "Variável de ambiente HOST_REQUEST não definida", http.StatusInternalServerError)
		return
	}

	// Cria uma nova requisição para a próxima rota
	req, err := http.NewRequest(http.MethodGet, nextURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copia todos os headers da requisição atual para a nova requisição
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	// Faz a requisição para a próxima rota
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("version: " + show_version + " body: " + string(body) + "\n"))
}
