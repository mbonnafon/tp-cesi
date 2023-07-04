package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080" // Déclaration d'une constante pour le port du serveur

// La fonction principale de votre programme
func main() {
	http.HandleFunc("/", handler) // Cette ligne associe la fonction "handler" à la route racine ("/") de votre serveur

	http.HandleFunc("/health-check", HealthCheckHandler) // Cette ligne associe la fonction "HealthCheckHandler" à la route "/health-check" de votre serveur

	fmt.Println("Server Web started on port", PORT) // Ceci affiche un message à la console indiquant que le serveur a démarré et sur quel port il écoute

	log.Fatal(http.ListenAndServe(":"+PORT, nil)) // Ceci démarre le serveur sur le port spécifié par la constante "PORT". "log.Fatal" est utilisé ici pour que le programme termine son exécution si le serveur ne parvient pas à démarrer
}

// La fonction "handler" est appelée lorsqu'une requête HTTP est faite à la racine de votre serveur
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!") // Ceci envoie la réponse "Hello, World!" au client
}

// La fonction "HealthCheckHandler" est appelée lorsqu'une requête HTTP est faite à la route "/health-check" de votre serveur
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Ceci définit le code de statut de la réponse HTTP à 200 OK

	w.Header().Set("Content-Type", "application/json") // Ceci définit l'en-tête "Content-Type" de la réponse HTTP à "application/json"

	fmt.Fprint(w, `{"alive": true}`) // Ceci envoie la réponse `{"alive": true}` au client
}
