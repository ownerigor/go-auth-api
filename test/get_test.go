package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ownerigor/go-api-auth/internal/models"
	"github.com/ownerigor/go-api-auth/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func generateTestToken(userID uint, secretKey string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := jwt.MapClaims{
		"userID": float64(userID),
		"exp":    expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func TestMeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testSecret := "supersecretjwtkey"
	os.Setenv("JWT_SECRET", testSecret)

	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
	}
	defer dbMock.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao inicializar GORM com mock: %v", err)
	}

	testUser := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	testUser.ID = 1

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(testUser.ID, testUser.Name, testUser.Email, testUser.Password)
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(float64(testUser.ID), 1).
		WillReturnRows(rows)

	r := gin.Default()
	routes.SetupRoutes(r, gormDB)

	token, err := generateTestToken(testUser.ID, testSecret)
	if err != nil {
		t.Fatalf("Falha ao gerar token de teste: %v", err)
	}

	req, _ := http.NewRequest("GET", "/me", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado status %d, obtido %d. Corpo da resposta: %s", http.StatusOK, w.Code, w.Body.String())
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Falha ao decodificar a resposta JSON: %v", err)
	}

	if response["id"] != float64(testUser.ID) {
		t.Errorf("Esperado userID %v, obtido %v", testUser.ID, response["id"])
	}

	if response["email"] != testUser.Email {
		t.Errorf("Esperado email %s, obtido %s", testUser.Email, response["email"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectativas do mock não atendidas: %s", err)
	}
}
