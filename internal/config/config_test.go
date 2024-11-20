package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DecodeConfig(t *testing.T) {
	prop := properties.LoadMap(map[string]string{
		"scope":             "local",
		"database.cluster":  "db",
		"database.name":     "testlocal",
		"database.username": "root",
		"database.password": "root",
	})

	t.Run("should return success when decode config", func(t *testing.T) {
		expectedResult := &Configuration{
			Scope:    "local",
			Database: buildDatabase(),
		}

		result, err := decodeConfig(prop)
		require.NoError(t, err)

		isEqual := reflect.DeepEqual(expectedResult, result)
		assert.True(t, isEqual)
	})

	t.Run("should return error when decode config", func(t *testing.T) {
		prop := properties.LoadMap(map[string]string{
			"scope": "local",
		})
		_, err := decodeConfig(prop)
		require.Error(t, err)
	})
}

func Test_GetEnv(t *testing.T) {
	defaultEnv := setupTestGetEnv(t)
	defer defaultEnv(t)

	t.Run("should given an loaded env return the value", func(t *testing.T) {
		expectedResult := "any env value"
		result := getEnv("TEST_ENV_KEY", "")
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should given an unloaded env return the value", func(t *testing.T) {
		expectedResult := "default value"
		result := getEnv("ANY_ENV_KEY", "default value")
		assert.Equal(t, expectedResult, result)
	})
}

func TestConfiguration_overrideConfigurations(t *testing.T) {
	defaultEnv := setupTestOverrideConfigurations(t)
	defer defaultEnv(t)

	t.Run("should given a valid configuration and env with other appPath value should override appPath property", func(t *testing.T) {
		expectedResult := Configuration{
			AppPath:  "PATH",
			Scope:    "local",
			Database: Database{},
		}

		config := Configuration{}
		config.overrideConfigurations()

		isEqual := reflect.DeepEqual(expectedResult, config)
		assert.True(t, isEqual)
	})
}

func Test_loadProperties(t *testing.T) {
	t.Run("should return service properties when scope is not local", func(t *testing.T) {
		cleanEnvsNewConfig(t)
		setupTestLoadApplicationProperties(t)

		result, err := loadProperties()
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when mandatory envs not provided", func(t *testing.T) {
		cleanEnvsNewConfig(t)
		expectedResult := "environment variable APP_PATH not set"
		_, err := loadProperties()
		require.EqualError(t, err, expectedResult)
	})

	t.Run("should return local properties when scope is local", func(t *testing.T) {
		cleanEnvsNewConfig(t)
		setupTestLoadLocalProperties(t)

		result, err := loadProperties()
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestNewConfig(t *testing.T) {
	t.Run("given local empty scope and appPath should return error", func(t *testing.T) {
		cleanEnvsNewConfig(t)

		expectedError := fmt.Errorf("environment variable APP_PATH not set")
		os.Setenv("LOCAL_CONFIG_FILE_NAME", "testdata/valid.properties")
		result, err := NewConfig()

		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("given a valid local properties should load properties", func(t *testing.T) {
		cleanEnvsNewConfig(t)
		setupNewConfigEnv(t)
		fmt.Println(os.Getenv("APP_PATH"))

		appPath, err := getProjectPath()
		require.NoError(t, err)

		expectedResult := &Configuration{
			AppPath: appPath,
			Scope:   "",
			Database: Database{
				Cluster:  "db",
				Name:     "testlocal",
				Username: "root",
				Password: "root",
			},
		}

		result, err := NewConfig()

		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})
}

func buildDatabase() Database {
	return Database{
		Cluster:  "db",
		Name:     "testlocal",
		Username: "root",
		Password: "root",
	}
}

func setupTestGetEnv(t *testing.T) func(t *testing.T) {
	t.Log("setup test get env")
	_ = os.Setenv("TEST_ENV_KEY", "any env value")

	return func(t *testing.T) {
		t.Log("teardown test get env")
		_ = os.Unsetenv("TEST_ENV_KEY")
	}
}

func setupTestOverrideConfigurations(t *testing.T) func(t *testing.T) {
	t.Log("setup test override configurations")
	os.Setenv("APP_PATH", "PATH")
	os.Setenv("SCOPE", "local")

	return func(t *testing.T) {
		t.Log("teardown test override configurations")
		os.Unsetenv("APP_PATH")
		os.Unsetenv("SCOPE")
	}
}

func cleanEnvsNewConfig(t *testing.T) {
	t.Log("clean envs new config")
	_ = os.Unsetenv("LOCAL_CONFIG_FILE_NAME")
	_ = os.Unsetenv("SCOPE")
	_ = os.Unsetenv("APP_PATH")
	_ = os.Unsetenv("DB_CLUSTER")
	_ = os.Unsetenv("DB_NAME")
	_ = os.Unsetenv("DB_USERNAME")
	_ = os.Unsetenv("DB_PASSWORD")
}

func setupTestLoadLocalProperties(t *testing.T) {
	t.Log("setup test load local properties")
	appPath, err := getProjectPath()
	require.NoError(t, err)

	_ = os.Setenv("SCOPE", "local")
	_ = os.Setenv("APP_PATH", appPath)

	t.Cleanup(func() {
		t.Log("teardown test load local properties")
		os.Unsetenv("SCOPE")
		os.Unsetenv("APP_PATH")
	})
}

func setupTestLoadApplicationProperties(t *testing.T) func(t *testing.T) {
	t.Log("setup test load local properties")
	appPath, err := getProjectPath()
	require.NoError(t, err)

	_ = os.Setenv("SCOPE", "")
	_ = os.Setenv("APP_PATH", appPath)

	return func(t *testing.T) {
		t.Log("teardown test load local properties")
		_ = os.Setenv("configFileName", applicationConfigScope)
		_ = os.Unsetenv("SCOPE")
		_ = os.Unsetenv("APP_PATH")
	}
}

func setupNewConfigEnv(t *testing.T) {
	t.Log("setup new config env")
	appPath, _ := getProjectPath()

	_ = os.Setenv("configFileName", applicationConfigScope)
	_ = os.Setenv("SCOPE", "")
	_ = os.Setenv("APP_PATH", appPath)
	_ = os.Setenv("DB_USER", "root")
	_ = os.Setenv("DB_PASSWORD", "root")
	_ = os.Setenv("DB_PORT", "1234")
	_ = os.Setenv("DB_NAME", "testlocal")
	_ = os.Setenv("DB_CLUSTER", "db")
	_ = os.Setenv("DB_HOST", "testlocal")
}
