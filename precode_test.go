package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "github.com/stretchr/testify/assert"
)

//Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be 200")

    assert.NotEmpty(t, responseRecorder.Body.String(), "Expected non-empty body, got empty body")
}

//Город, который передаётся в параметре city, не поддерживается. 
//Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenWrongCity(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code to be 400")

    expectedBody := "wrong city value"
    assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Expected body to be '%s', got '%s'", expectedBody, responseRecorder.Body.String())
    
}

//Если в параметре count указано больше кафе, чем есть на самом деле, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be 200")

    list := strings.Split(responseRecorder.Body.String(), ",")
    assert.Len(t, list, totalCount, "Expected cafe count to be %d, got %d", totalCount, len(list))
}