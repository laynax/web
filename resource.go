package main

import (
	"os"
	"io/ioutil"
	"errors"
	"syscall"
	"time"
	"strconv"
)

// TODO first char / validation
// TODO create duplication
// TODO .. security check

type resource struct {
	ID      string
	Dir     string
	Content string
}

func (r *resource) commit() error {
	err := os.MkdirAll(mainDir+r.Dir, 0777)
	if err == os.ErrExist {
		return errors.New("path already exists")
	}
	assert(err)

	id := strconv.Itoa(time.Now().Nanosecond())
	f, err := os.Create(mainDir + r.Dir + "/" + id)
	assert(err)

	r.ID = id
	_, err = f.Write([]byte(r.Content))
	assert(err)

	return nil
}

func (r *resource) update() error {
	//check err
	dir := mainDir + r.Dir + "/" + r.ID
	err := os.Truncate(dir, 0)
	if err == os.ErrNotExist {
		return errors.New("no such directory")
	}
	assert(err)

	f, err := os.OpenFile(dir, os.O_RDWR, 0777)
	assert(err)

	_, err = f.Write([]byte(r.Content))
	assert(err)

	return nil
}

func deleteResource(dir, id string) error {
	err := os.Remove(mainDir + dir + "/" + id)
	if pathError, ok := err.(*os.PathError); ok && pathError.Err == syscall.ENOENT {
		return os.ErrNotExist
	}
	assert(err)

	return nil
}

func getResource(dir, id string) (resource, error) {
	f, err := os.Open(mainDir + dir + "/" +id)
	if pathError, ok := err.(*os.PathError); ok && pathError.Err == syscall.ENOENT {
		return resource{}, os.ErrNotExist
	}
	assert(err)

	data, err := ioutil.ReadAll(f)
	assert(err)

	return resource{Content: string(data), Dir: dir, ID:id}, nil
}
