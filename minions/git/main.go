package git

import (
	"buildbot/models"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

// A Giter clones remote repository.
type Giter interface {
	Clone(c models.Component) (r Result, err error)
}

// A Git struct to represent git driver.
type Git struct {
	RootWorkDir string
}

// A Result provides information of completed operation.
type Result struct {
	WorkDir string
}

// A NewGit function instantiate Git driver with default values.
func NewGit(options ...func(*Git)) *Git {
	git := Git{RootWorkDir: "/tmp/workspace"}
	for _, option := range options {
		option(&git)
	}
	return &git
}

// Clone remote repository to working directory.
func (g Git) Clone(c models.Component) (Result, error) {
	if c.ID == "" {
		return Result{}, errors.New("component passed to Clone is empty")
	}
	workDir := fmt.Sprintf("%s/%s", g.RootWorkDir, c.ID)

	_, err := git.PlainClone(workDir, false, &git.CloneOptions{
		URL: c.GitUrl,
	})
	if err != nil {
		return Result{}, errors.Wrap(err, "git clone failed")
	}
	res := Result{WorkDir: workDir}

	return res, nil
}
