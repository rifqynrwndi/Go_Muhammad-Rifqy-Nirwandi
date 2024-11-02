package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
    router := gin.Default()
    router.GET("/todos", GetTodos)

    req, _ := http.NewRequest("GET", "/todos", nil)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    // Tambahkan assert sesuai response body jika diperlukan
}

func TestCreateTodo(t *testing.T) {
    router := gin.Default()
    router.POST("/todos", CreateTodo)

    req, _ := http.NewRequest("POST", "/todos", nil)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
}