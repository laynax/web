package main

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
)

type Stat struct {
	Mean              float64
	StandardDeviation float64
}

func includedFolderCount(dir string) (int, error) {
	dir = mainDir + dir

	var counter int
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return os.ErrNotExist
		}

		if !info.IsDir() {
			counter++
		}
		return nil
	})
	if err == os.ErrNotExist {
		return 0, err
	}
	assert(err)

	return counter, nil
}

func averageWorldLength(dir string) (float64, float64, error) {
	dir = mainDir + dir

	var arr []float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return os.ErrNotExist
		}

		if !info.IsDir() {
			f, err := os.Open(path)
			assert(err)

			data, err := ioutil.ReadAll(f)
			assert(err)
			defer f.Close()

			words := regexp.MustCompile(`[[:alnum:]]+`).FindAllString(string(data), -1)
			arr = append(arr, float64(len(words)))
		}
		return nil
	})
	if err == os.ErrNotExist {
		return 0, 0, os.ErrNotExist
	}
	assert(err)

	if len(arr) == 0 {
		return 0, 0, errors.New("no resource was found in given path")
	} else if len(arr) == 1 {
		return arr[0], 0, nil
	}

	m := mean(arr...)
	sd := stdDeviation(m, arr...)

	return m, sd, nil
}

func alphanumericStatic(dir string) (float64, float64, error) {
	dir = mainDir + dir

	var arr []float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return os.ErrNotExist
		}

		if !info.IsDir() {
			f, err := os.Open(path)
			assert(err)

			data, err := ioutil.ReadAll(f)
			assert(err)
			defer f.Close()

			arr = append(arr, countAlphanumeric(data))
		}
		return nil
	})
	if err == os.ErrNotExist {
		return 0, 0, err
	}
	assert(err)

	m := mean(arr...)
	return m, stdDeviation(m, arr...), nil
}

func directorySize(dir string) (int64, error) {
	dir = mainDir + dir

	info, err := os.Stat(dir)
	if err == os.ErrNotExist {
		return 0, err
	}
	assert(err)

	return info.Size(), nil
}

func countAlphanumeric(data []byte) float64 {
	var counter int
	for i := range data {
		b := data[i]

		if b > 47 && b < 58 ||
			b > 64 && b < 91 ||
			b > 96 && b < 123 {
			counter++
		}
	}

	return float64(counter)
}

func stdDeviation(mean float64, indexes ...float64) float64 {
	total := 0.0
	for _, number := range indexes {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(indexes)-1)
	return math.Sqrt(variance)
}

func mean(indexes ...float64) float64 {
	var sum float64
	for i := range indexes {
		sum += indexes[i]
	}
	return sum / float64(len(indexes))
}
