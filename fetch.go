package main

import (
	"github.com/ztrue/tracerr"
	"gopkg.in/src-d/go-git.v4"
)

func fetch(repoPath string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return tracerr.Wrap(err)
	}

	remotes, err := r.Remotes()
	if err != nil {
		return tracerr.Wrap(err)
	}

	for _, remote := range remotes {
		remoteName := remote.Config().Name

		err, _, errB := GitCommand(repoPath, []string{"fetch", remoteName})
		if err != nil {
			return tracerr.Errorf("Remote: %s %s %w", remoteName, errB.Bytes(), err)
		}
	}

	return nil
}
