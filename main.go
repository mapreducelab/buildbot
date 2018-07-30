package main

import (
	"buildbot/minions/git"
	"buildbot/models"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	comp := models.Component{
		ID:     "gittest",
		GitUrl: "https://github.com/mapreducelab/application-images.git",
	}
	logger := log.New(os.Stderr, "logger_test: ", 2)
	log2 := *logger

	var g git.Giter
	g = git.Git{}
	workDir := func(g *git.Git) {
		g.RootWorkDir = "/tmp/workspace2"
	}
	g = git.NewGit(workDir)
	g = git.LoggingMiddleware{log2, g}

	res, err := g.Clone(comp)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(g))
	fmt.Printf("%+v\n", g)
	fmt.Printf("%+v\n", res)
}
