package git

import (
	"buildbot/models"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware provides logging for git.
type LoggingMiddleware struct {
	Logger log.Logger
	Debug  bool
	Next   Giter
}

func (mw LoggingMiddleware) Clone(c models.Component) (r Result, err error) {
	defer func(begin time.Time) {
		logger := log.With(mw.Logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
		if mw.Debug {
			comp := fmt.Sprintf("%+v", c)
			res := fmt.Sprintf("%+v", r)
			// get rid of curly braces
			comp = comp[1 : len(comp)-1]
			res = res[1 : len(res)-1]
			logger.Log(
				"method", "Clone",
				"input_component", comp,
				"output_result", res,
				"err", err,
				"took", time.Since(begin),
			)
		}
	}(time.Now())
	r, err = mw.Next.Clone(c)
	return
}
