package git

import (
	"fmt"
	"os"
	"time"

	go_git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/kaibling/apiforge/logging"
)

func Clone(gitURL, localPath, gitToken string) error {
	_, err := go_git.PlainClone(localPath, false, &go_git.CloneOptions{
		URL: gitURL,
		Auth: &http.BasicAuth{
			Username: "not_empty_string",
			Password: gitToken,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}

func Pull(localPath string) error {
	r, err := go_git.PlainOpen(localPath)
	if err != nil {
		return fmt.Errorf("cannot pull local repo: %w", err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	return w.Pull(&go_git.PullOptions{RemoteName: "origin"})
}

func Push(localPath, gitToken string) error {
	r, err := go_git.PlainOpen(localPath)
	if err != nil {
		return fmt.Errorf("cannot push local repo: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	status, err := w.Status()
	if err != nil {
		return err
	}

	if status.IsClean() {
		return nil
	}

	if err := r.Push(&go_git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "not_empty_string",
			Password: gitToken,
		},
	}); err != nil {
		return fmt.Errorf("failed to push repo: %w", err)
	}

	return nil
}

func CommitFiles(localPath string, filenames []string, logger logging.Writer) error {
	r, err := go_git.PlainOpen(localPath)
	if err != nil {
		return fmt.Errorf("cannot pull local repo: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	logger.Debug("staging files..")

	for _, n := range filenames {
		_, err := w.Add(n)
		if err != nil {
			return err
		}
	}

	logger.Debug("check status..")

	status, err := w.Status()
	if err != nil {
		return err
	}

	if status.IsClean() {
		logger.Info("Repo is clean")

		return nil
	}

	logger.Debug(status.String())

	commit, err := w.Commit("workflow export", &go_git.CommitOptions{
		Author: &object.Signature{
			Name:  "iggy",
			Email: "iggy@doe.org",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	obj, err := r.CommitObject(commit)
	if err != nil {
		return err
	}

	logger.Info(obj.Message)

	return nil
}
