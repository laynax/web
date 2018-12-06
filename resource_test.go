package main

import (
	"testing"
	"os"
)

func TestCommit(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	rs.commit()

	fetchedResource, _ := getResource(rs.Dir)
	if fetchedResource.Content != rs.Content {
		t.Error("invalid commited content")
	}
}

func TestDelete(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	rs.commit()

	deleteResource(rs.Dir)

	_, err := getResource(rs.Dir)
	if err != os.ErrNotExist {
		t.Error("error deleting resource")
	}
}

func TestUpdate(t *testing.T) {
	rs := resource{Dir: "/a/b", Content: "test1"}
	rs.commit()

	rs.Content = "test2"
	rs.update()

	fetchedResource, _ := getResource(rs.Dir)
	if fetchedResource.Content != rs.Content {
		t.Error("error updating resource")
	}
}