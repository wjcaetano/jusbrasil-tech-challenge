package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
)

const (
	localConfigScope       = "resources/config/local.properties"
	applicationConfigScope = "resources/config/application.properties"
	localConfigFileEnv     = "variables.env"
	scopeEnv               = "SCOPE"
	appPathEnv             = "APP_PATH"
	localScope             = "local"
)

type (
	Database struct {
		Cluster  string `properties:"cluster"`
		Name     string `properties:"name"`
		Username string `properties:"username"`
		Password string `properties:"password"`
	}

	Configuration struct {
		AppPath  string   `properties:"app_path,default="`
		Scope    string   `properties:"scope,default="`
		Database Database `properties:"database"`
	}
)

func NewConfig() (*Configuration, error) {
	prop, err := loadProperties()
	if err != nil {
		return nil, err
	}

	conf, err := decodeConfig(prop)
	if err != nil {
		return nil, err
	}

	conf.overrideConfigurations()

	return conf, nil
}

func loadProperties() (*properties.Properties, error) {
	if err := checkMandatoryEnvs(); err != nil {
		return nil, err
	}

	if getEnv(scopeEnv, "SCOPE") == localScope {
		return loadLocalProperties(), nil
	}

	return loadServiceProperties()
}

func checkMandatoryEnvs() error {
	mandatoryEnvs := [...]string{appPathEnv, scopeEnv}

	for _, env := range mandatoryEnvs {
		if _, ok := os.LookupEnv(env); !ok {
			return fmt.Errorf("environment variable %s not set", env)
		}
	}
	return nil
}

func (c *Configuration) overrideConfigurations() {
	c.AppPath = getEnv(appPathEnv, "")
	c.Scope = getEnv(scopeEnv, "")
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func loadLocalProperties() *properties.Properties {
	appPath, err := getProjectPath()
	if err != nil {
		return nil
	}
	configFile := filepath.Join(appPath, localConfigScope)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil
	}

	return properties.MustLoadFile(configFile, properties.UTF8)
}

func loadServiceProperties() (*properties.Properties, error) {
	inputConfig := os.Getenv("configFileName")
	if inputConfig == "" {
		inputConfig = applicationConfigScope
	}

	appPath, err := getProjectPath()
	if err != nil {
		return nil, fmt.Errorf("could not get project path: %w", err)
	}

	configFile := filepath.Join(appPath, inputConfig)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find configuration file: %s", inputConfig)
	}

	prop, _ := properties.LoadFile(configFile, properties.UTF8)

	return prop, nil
}

func decodeConfig(prop *properties.Properties) (*Configuration, error) {
	var cfg Configuration
	if err := prop.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func IsLocalScope() bool {
	return getEnv(scopeEnv, "") == localScope
}

func getProjectPath() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get working directory: %w", err)
	}

	workingDir = filepath.Dir(filepath.Dir(workingDir))

	return workingDir, nil
}