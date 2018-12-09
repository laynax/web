package main

import (
	"testing"
)

func TestIncludedFolderCount(t *testing.T) {
	rs1 := resource{Dir: "/t/e/s/t/1"}
	rs2 := resource{Dir: "/t/e/s/t/2"}

	err1 := rs1.commit()
	err2 := rs2.commit()
	if err1 != nil || err2 != nil {
		t.Error("invalid committed content")
	} else {
		defer func() {
			deleteResource(rs1.Dir, rs1.ID)
			deleteResource(rs2.Dir, rs2.ID)
		}()
	}

	count, err := includedFolderCount("/t/e/s/t")
	if err != nil || count != 2 {
		t.Error("invalid count number")
	}
}

func TestAverageWorldLength(t *testing.T) {
	rs1 := resource{Dir: "/t/e/s/t/1", Content: "a"}
	rs2 := resource{Dir: "/t/e/s/t/2", Content: "a b"}
	rs3 := resource{Dir: "/t/e/s/t/3", Content: "a b c"}

	err1 := rs1.commit()
	err2 := rs2.commit()
	err3 := rs3.commit()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("invalid committed content")
	} else {
		defer func() {
			deleteResource(rs1.Dir, rs1.ID)
			deleteResource(rs2.Dir, rs2.ID)
			deleteResource(rs3.Dir, rs3.ID)
		}()
	}

	mean, dev, err := averageWorldLength("/t/e/s/t")
	if err != nil || mean != 2 || dev != 1 {
		t.Error("invalid count number")
	}
}

func TestAlphanumericStatic(t *testing.T) {
	rs1 := resource{Dir: "/t/e/s/t/1", Content: "a"}
	rs2 := resource{Dir: "/t/e/s/t/2", Content: "a b"}
	rs3 := resource{Dir: "/t/e/s/t/3", Content: "a b c"}

	err1 := rs1.commit()
	err2 := rs2.commit()
	err3 := rs3.commit()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("invalid committed content")
	} else {
		defer func() {
			deleteResource(rs1.Dir, rs1.ID)
			deleteResource(rs2.Dir, rs2.ID)
			deleteResource(rs3.Dir, rs3.ID)
		}()
	}

	mean, dev, err := alphanumericStatic("/t/e/s/t")
	if err != nil || mean != 2 || dev != 1 {
		t.Error("invalid committed content")
	}
}
