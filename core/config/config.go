package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config contém as informações de conexãocom o banco de dados
type Config struct {
	Host      string `json:"host"`
	Password  string `json:"password"`
	User      string `json:"user"`
	DbName    string `json:"dbName"`
	Port      string `json:"port"`
	LogConfig string `json:"logConfig"`
}

// ToString gera a string de conexão com o banco de dados de acordo com o 'objeto' Config responável pela chamada do método
func (config *Config) ToString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	return psqlInfo
}

// LoadFromJSON carrega as informações de conexão com o banco de dados de um arquivo do formato json.
func (config *Config) LoadFromJSON(fileName string) (err error) {
	var file *os.File
	var decoder *json.Decoder
	if file, err = os.Open(fileName); err != nil {
		return err
	}
	decoder = json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return err
	}
	return nil
}
