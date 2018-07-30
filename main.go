package main

import (
	"buildbot/minions/git"
	"buildbot/models"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	comp := models.Component{
		ID:     "gittest",
		GitUrl: "https://github.com/mapreducelab/application-images.git",
	}
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	var g git.Giter
	g = git.Git{}
	workDir := func(g *git.Git) {
		g.RootWorkDir = "/tmp/workspace2"
	}
	g = git.NewGit(workDir)
	g = git.LoggingMiddleware{logger, true, g}

	_, err := g.Clone(comp)
	if err != nil {
		panic(err)
	}
}
