package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateViper(t *testing.T) {
	var vpConfig *viper.Viper = viper.New()
	fmt.Println("test")
	assert.NotNil(t, vpConfig)
}

func TestReadConfigViper(t *testing.T) {
	var config *viper.Viper = viper.New()

	// set config file
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./")

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)
}

// load config file (from json)
func TestGetValueFromEnv(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}

	// get value
	appName := config.GetString("app.name")
	appVersion := config.GetString("app.version")
	databaseShow := config.GetBool("database.show_sql")
	databasePort := config.GetInt("database.port")

	assert.Equal(t, "belajar viper", appName)
	assert.Equal(t, "1.0", appVersion)
	assert.True(t, databaseShow)
	assert.Equal(t, 3306, databasePort)
}

// read config from json file
func TestGetValueFromJson(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.json")
	config.AddConfigPath("./")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}

	// get value from json
	appName := config.GetString("app.name")
	appAuthor := config.GetString("app.author")
	databaseShow := config.GetBool("database.show_sql")
	databaseHost := config.GetString("database.host")
	databasePort := config.GetInt("database.port")

	assert.Equal(t, "belajar viper", appName)
	assert.Equal(t, "reo sahobby", appAuthor)
	assert.True(t, databaseShow)
	assert.Equal(t, "localhost", databaseHost)
	assert.Equal(t, 3306, databasePort)
}

// read config from yaml file
func TestGetValueFromYaml(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath("./")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}

	// read data from yaml
	appName := config.GetString("app.name")
	databaseShow := config.GetBool("database.parse_date")
	databasePort := config.GetInt("database.port")

	assert.Equal(t, "belajar golang viper", appName)
	assert.True(t, databaseShow)
	assert.Equal(t, 3306, databasePort)
}
