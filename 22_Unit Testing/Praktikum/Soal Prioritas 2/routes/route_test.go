package routes

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {

	e := echo.New()

	SetupRoutes(e)

	assert.NotNil(t, e.Routes())
}
