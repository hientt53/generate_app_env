package main

import (
	"fmt"
	"os"
	"strings"
)

var appEnvPrefix = "APPENV"
var envFileName = "app.env"

type appEnv struct {
	key   string
	value string
}

func main() {
	osEnviron := os.Environ()
	appEnvs := getAppEnv(osEnviron)
	writeDotEnv(appEnvs)

}

func getAppEnv(environs []string) (appEnvs []appEnv) {
	fmt.Println("Read app env ...")
	for _, env := range environs {
		pairKeyValue := strings.Split(env, "=")
		pairKey := strings.Split(env, "_")
		if pairKey[0] == appEnvPrefix {
			appenv := appEnv{
				key:   getAppEnvKey(pairKeyValue[0]),
				value: pairKeyValue[1],
			}
			appEnvs = append(appEnvs, appenv)
		}
	}
	return
}

func getAppEnvKey(osEnv string) string {
	return osEnv[len(appEnvPrefix)+1 : len(osEnv)]
}

func writeDotEnv(appEnvs []appEnv) {
	fmt.Println("Write app env ...")
	f, err := os.Create(envFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, env := range appEnvs {
		_, err := fmt.Fprintln(f, fmt.Sprintf("%s=%s", env.key, env.value))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(fmt.Sprintf("Write to %s successfully!", envFileName))
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
