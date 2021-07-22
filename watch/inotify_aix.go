// Copyright (c) 2015 HPE Software Inc. All rights reserved.
// Copyright (c) 2013 ActiveState Software Inc. All rights reserved.

// +build aix

package watch

import (
	"path/filepath"

	"gopkg.in/tomb.v1"
)

// InotifyFileWatcher uses inotify to monitor file changes.
type InotifyFileWatcher struct {
	Filename string
	Size     int64
}

func NewInotifyFileWatcher(filename string) *InotifyFileWatcher {
	fw := &InotifyFileWatcher{filepath.Clean(filename), 0}
	return fw
}

func (fw *InotifyFileWatcher) BlockUntilExists(t *tomb.Tomb) error {
	return nil
}

func (fw *InotifyFileWatcher) ChangeEvents(t *tomb.Tomb, pos int64) (*FileChanges, error) {
	changes := NewFileChanges()
	return changes, nil
}
