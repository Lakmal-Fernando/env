package env

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

const (
	InvalidVariable = "invalid variable in env file"
	PanicError      = "error occured in setting environment variables"
)

type EnvVariable struct {
	Key   string
	Value string
}

func SetEnvVariables(filePath string) (err error) {
	defer func() {
		if recoveredError := recover(); recoveredError != nil {
			log.Printf("%s, error : %v", PanicError, err)
			err = errors.New(PanicError)
			return
		}
	}()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var envVariables []EnvVariable
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#") || strings.HasPrefix(txt, "//") { //ignoring comments
			continue
		}
		variable := strings.Split(txt, "=")
		if len(variable) != 2 { //Checking for invalid equal signs and missing key-value pairs
			return errors.New(InvalidVariable)
		}
		envVariables = append(envVariables, EnvVariable{
			Key:   strings.Trim(variable[0], " "),
			Value: strings.Trim(variable[1], " "),
		})
	}
	if err = scanner.Err(); err != nil {
		return err
	}
	for _, envVariable := range envVariables {
		os.Setenv(envVariable.Key, envVariable.Value)
	}
	return nil
}
