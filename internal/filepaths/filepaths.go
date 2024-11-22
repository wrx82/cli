package filepaths

import (
	"fmt"
	"path/filepath"
)

type CanonicalisationError struct {
	AttemptedPath   string
	UnderlyingError error
}

func (e CanonicalisationError) Error() string {
	return fmt.Sprintf("could not canonicalise path %q due to %v", e.AttemptedPath, e.UnderlyingError)
}

func Canonicalise(path string) (CanonicalisedPath, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, CanonicalisationError{AttemptedPath: path, UnderlyingError: err}
	}

	symlinkEvaledPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return nil, CanonicalisationError{AttemptedPath: absPath, UnderlyingError: err}
	}

	return canonicalisedPath{symlinkEvaledPath}, nil
}

type CanonicalisedPath interface {
	sealed()
	String() string
}

type canonicalisedPath struct {
	path string
}

func (c canonicalisedPath) sealed() {}

func (c canonicalisedPath) String() string {
	return c.path
}
