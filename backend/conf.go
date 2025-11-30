package main

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	v *viper.Viper
}

func newConfig() *config {
	c := &config{
		v: viper.New(),
	}
	c.init()
	return c
}

func (c *config) init() {
	c.v.SetDefault("PORT", "8080")

	c.v.SetConfigName(".env")
	c.v.SetConfigType("env")
	c.v.AddConfigPath(".")

	err := c.v.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (c *config) port() string {
	return c.v.GetString("PORT")
}