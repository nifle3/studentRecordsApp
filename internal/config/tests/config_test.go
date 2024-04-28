package tests

import (
	"github.com/stretchr/testify/assert"
	"os"
	"studentRecordsApp/internal/config"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("DB_USER", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("DB_PASSWORD", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("DB_HOST", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("DB_PORT", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("DB_NAME", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("DB_SSL", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("FS_END_POINT", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("FS_USER", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("FS_PASSWORD", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("SERVER_PORT", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("SERVER_HOST", "QWE"); err != nil {
		t.Fatal(err.Error())
	}
	if err := os.Setenv("JWT_SECRET_KEY", "QWE"); err != nil {
		t.Fatal(err.Error())
	}

	result1 := config.GetConfig()
	result2 := config.GetConfig()
	assert.Equal(t, result1, result2)
}
