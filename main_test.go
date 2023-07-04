package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// La fonction TestHealthCheckHandler est notre test unitaire pour la fonction HealthCheckHandler.
func TestHealthCheckHandler(t *testing.T) {
	// On crée une nouvelle requête HTTP GET pour la route "/health-check".
	req, err := http.NewRequest("GET", "/health-check", nil)

	// Si une erreur se produit lors de la création de la requête, le test se termine avec une erreur.
	if err != nil {
		t.Fatal(err)
	}

	// On crée un nouvel enregistreur de réponse qui va enregistrer la réponse HTTP de la fonction HealthCheckHandler.
	rr := httptest.NewRecorder()

	// On convertit notre fonction HealthCheckHandler en un http.Handler en utilisant http.HandlerFunc.
	handler := http.HandlerFunc(HealthCheckHandler)

	// On appelle la fonction ServeHTTP sur notre handler avec notre requête et notre enregistreur de réponse.
	handler.ServeHTTP(rr, req)

	// On vérifie que le code de statut HTTP de la réponse est bien 200 (OK).
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// On définit ce que nous attendons comme corps de réponse.
	expected := `{"alive": true}`

	// On vérifie que le corps de la réponse est bien celui que nous attendons.
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
