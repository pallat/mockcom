package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func initConf() {
	_default()
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		fmt.Printf("warning: %s \n", err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

}

func _default() {
	viper.SetDefault("file", "./tests/example.http")
	viper.SetDefault("addr", ":9090")
}

func main() {
	initConf()

	b, err := ioutil.ReadFile(viper.GetString("file"))
	if err != nil {
		log.Fatal(err)
	}

	body := strings.Split(string(b), "\n")
	proto := strings.Split(body[0], " ")

	http.HandleFunc(proto[1], func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, "%s", strings.Join(body[1:], "\n"))
	})

	http.ListenAndServe(viper.GetString("addr"), nil)
}
