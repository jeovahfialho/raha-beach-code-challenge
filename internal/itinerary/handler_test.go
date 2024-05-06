package itinerary

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostItineraries(t *testing.T) {
	// Inicializa o Echo e registra os handlers
	e := echo.New()
	RegisterHandlers(e)

	// JSON input para o teste
	var jsonStr = `[["JFK","LAX"],["LAX","DXB"],["DXB","SFO"],["SFO","SJC"]]`

	// Cria uma nova requisição HTTP POST para o endpoint '/itineraries'
	req := httptest.NewRequest(http.MethodPost, "/itineraries", strings.NewReader(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Cria um novo ResponseRecorder (httptest.ResponseRecorder) para gravar a resposta
	rec := httptest.NewRecorder()

	// O Echo serve a requisição usando o ResponseRecorder
	e.ServeHTTP(rec, req)

	// Expected JSON response
	expectedJSONResponse := `["JFK", "LAX", "DXB", "SFO", "SJC"]`

	// Testa se a resposta HTTP é a esperada
	assert.Equal(t, http.StatusOK, rec.Code, "should return HTTP status 200")

	// Adicione mais testes conforme necessário para verificar o corpo da resposta
	// Por exemplo, você pode querer verificar se a estrutura do itinerário retornado é correta
	assert.JSONEq(t, expectedJSONResponse, rec.Body.String(), "Response body should match expected JSON")
}
