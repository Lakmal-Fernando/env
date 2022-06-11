package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type EnvVariable struct {
	Key   string
	Value string
}

func SetEnvVariables(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var envVariables []EnvVariable
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(scanner.Text())
		variable := strings.Split(txt, "=")
		envVariables = append(envVariables, EnvVariable{
			Key:   strings.Trim(variable[0], " "),
			Value: strings.Trim(variable[1], " "),
		})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, envVariable := range envVariables {
		os.Setenv(envVariable.Key, envVariable.Value)
	}
}
