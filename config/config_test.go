package config_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/stretchr/testify/assert"
)

// Change working directory to project root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestNewConfig(t *testing.T) {
	t.Run("Test config without environment variables", func(t *testing.T) {
		cfg, err := config.NewConfig()

		if err != nil {
			t.Errorf("Fail to create new config: %d", err)
		}

		assert.Equal(t, "go-template", cfg.Name)
		assert.Equal(t, "1.0.0", cfg.Version)
		assert.Equal(t, "debug", cfg.Log.Level)

		assert.Equal(t, ":8080", cfg.Address)
		assert.Equal(t, make([]string, 0), cfg.CORSAllowOrigins)
	})

	t.Run("Test config replace with environment variables", func(t *testing.T) {
		os.Setenv("APP_NAME", "test-template")
		os.Setenv("APP_VERSION", "v1.0.0")
		os.Setenv("LOGGER_LOG_LEVEL", "test")

		cfg, err := config.NewConfig()

		if err != nil {
			t.Errorf("Fail to create new config: %d", err)
		}

		assert.Equal(t, "test-template", cfg.Name)
		assert.Equal(t, "v1.0.0", cfg.Version)
		assert.Equal(t, "test", cfg.Log.Level)
		assert.Equal(t, ":8080", cfg.Address)
		assert.Equal(t, make([]string, 0), cfg.CORSAllowOrigins)
	})
}
