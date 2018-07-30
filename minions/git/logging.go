package git

import (
	"buildbot/models"
	"log"
)

// LoggingMiddleware provides logging for git.
type LoggingMiddleware struct {
	Logger log.Logger
	Next   Giter
}

func (mw LoggingMiddleware) Clone(c models.Component) (r Result, err error) {
	mw.Logger.Println("workDir")
	r, err = mw.Next.Clone(c)
	return
}
