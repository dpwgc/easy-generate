package main

import (
	"easygenerate/generate"
	"fmt"
	"github.com/dpwgc/easierweb"
	"github.com/dpwgc/easierweb/middlewares"
	"github.com/dpwgc/easierweb/plugins"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

// basic usage example
func main() {

	file, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}
	serverConfig := ServerConfig{}
	err = yaml.Unmarshal(file, &serverConfig)
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}

	// create a router and set middleware
	router := easierweb.New(easierweb.RouterOptions{
		RequestHandle: plugins.YAMLRequestHandle(),
	}).Use(middlewares.Logger())

	// set api handle
	router.EasyPOST("/build/:type", build)
	router.Static("/java/*filepath", "web")

	// runs on port 80
	log.Fatal(router.Run(fmt.Sprintf(":%v", serverConfig.Port)))
}

// api handle
func build(ctx *easierweb.Context, config generate.Config) {
	switch ctx.Path.Get("type") {
	case "java_po":
		ctx.WriteString(http.StatusOK, generate.NewJavaPOGenerator(config).Build())
		return
	case "java_dto":
		ctx.WriteString(http.StatusOK, generate.NewJavaDTOGenerator(config).Build())
		return
	case "mysql":
		ctx.WriteString(http.StatusOK, generate.NewMySQLGenerator(config).Build())
		return
	}
	panic("type error")
}

type ServerConfig struct {
	Port int `yaml:"port"`
}
