package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func create(c *gin.Context) {
	var payload resource
	c.Bind(&payload)
	if payload.Dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	if err := payload.commit(); err != nil {
		json(c, nil, err)
		return
	}

	c.JSON(http.StatusOK, payload)
}

func get(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	id := c.Param("id")
	r, err := getResource(dir, id)
	if err != nil {
		json(c, nil, err)
		return
	}

	json(c, r)
}

func deleteHandler(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	id := c.Param("id")
	err := deleteResource(dir, id)
	json(c, nil, err)
}

func update(c *gin.Context) {
	// get id too
	var payload resource
	err := c.Bind(&payload)
	if err != nil {
		json(c, err)
		return
	}

	payload.ID = c.Param("id")
	if payload.ID == "" || payload.Dir == "" {
		json(c, nil, errors.New("dir or id cannot be empty"))
		return
	}

	err = payload.update()
	if err != nil {
		json(c, err)
		return
	}

	c.JSON(http.StatusOK, payload)
}

func totalSubDir(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	count, err := includedFolderCount(dir)
	response := struct {
		TotalFiles int
	}{count}
	json(c, response, err)
}

func alphanumericStatics(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	meanAN, stdDev, err := alphanumericStatic(dir)
	if err != nil {
		json(c, nil, err)
		return
	}

	resp := Stat{meanAN, stdDev}
	json(c, resp, nil)
}

func wordLength(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	mean, dev, err := averageWorldLength(dir)
	if err != nil {
		json(c, nil, err)
		return
	}

	resp := Stat{mean, dev}
	json(c, resp, nil)
}

func size(c *gin.Context) {
	dir := c.Query("dir")
	if dir == "" {
		json(c, nil, errors.New("dir cannot be empty"))
		return
	}

	size, err := directorySize(dir)
	if err != nil {
		json(c, nil, err)
		return
	}

	resp := struct {
		Size int64
	}{size}
	json(c, resp, nil)
}

func json(c *gin.Context, data interface{}, err ...error) {
	var e error
	if len(err) != 0 && err[0] != nil {
		e = err[0]
	}

	// ugly
	statusCode := http.StatusOK
	if e == os.ErrNotExist {
		statusCode = http.StatusNotFound
	} else if e != nil {
		statusCode = http.StatusBadRequest
	}

	if e == nil {
		e = errors.New("")
	}
	c.JSON(http.StatusOK, struct {
		StatusCode   int
		ErrorMessage string
		Data         interface{} `json:",omitempty"`
	}{statusCode, e.Error(), data})
}
