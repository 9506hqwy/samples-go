package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Printf(
		"fint = %v\nfbool = %v\nfstr = %v\n",
		viper.GetInt("fint"),
		viper.GetBool("fbool"),
		viper.GetString("fstr"),
	)
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("fint", viper.GetInt("fint")+1)

	err = viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}
