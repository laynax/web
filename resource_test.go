package main

import (
	"os"
	"testing"
)

func TestCommit(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	err := rs.commit()
	if err == nil {
		defer deleteResource(rs.Dir, rs.ID)
	}

	fetchedResource, _ := getResource(rs.Dir, rs.ID)
	if fetchedResource.Content != rs.Content {
		t.Error("invalid commited content")
	}
}

func TestDelete(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	rs.commit()

	deleteResource(rs.Dir, rs.ID)

	_, err := getResource(rs.Dir, rs.ID)
	if err != os.ErrNotExist {
		t.Error("error deleting resource")
	}
}

func TestUpdate(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	err := rs.commit()
	if err == nil {
		defer deleteResource(rs.Dir, rs.ID)
	}

	rs.Content = "test2"
	err = rs.update()
	if err != nil {
		t.Fatal(err)
	}

	fetchedResource, err := getResource(rs.Dir, rs.ID)
	if fetchedResource.Content != rs.Content || err != nil {
		t.Error("error updating resource")
	}
}