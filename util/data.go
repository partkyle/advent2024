package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"os"
	"path"
	"strings"
)

const YEAR = 2024

func downloadDay(day int) error {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", YEAR, day), nil)
	if err != nil {
		return err
	}

	cookieFile, err := os.Open(path.Join(os.Getenv("HOME"), ".advent"))
	if err != nil {
		return err
	}

	cookie, err := io.ReadAll(cookieFile)
	if err != nil {
		return err
	}

	sCookie := strings.TrimSpace(string(cookie))

	request.Header.Set("User-Agent", "Mozilla/5.0")
	request.Header.Set("Accept", "application/zip")
	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sCookie,
	})

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cachedFile, err := os.OpenFile(dayFileOnDisk("data", day), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cachedFile.Close()

	_, err = io.Copy(cachedFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetData(t string, day int) iter.Seq[string] {
	f, err := os.Open(dayFileOnDisk(t, day))
	if errors.Is(err, os.ErrNotExist) && t == "data" {
		err := downloadDay(day)
		if err != nil {
			panic(err)
		}

		f, err = os.Open(dayFileOnDisk(t, day))
		if err != nil {
			panic(err)
		}
	}

	s := bufio.NewScanner(f)

	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				f.Close()
				return
			}
		}
	}
}

func dayFileOnDisk(t string, day int) string {
	return path.Join(t, fmt.Sprintf("%02d.txt", day))
}

func Data(day int) iter.Seq[string] {
	return GetData("data", day)
}

func Sample(day int) iter.Seq[string] {
	return GetData("sample", day)
}

func DataProcess[V any](day int, transform func(string) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range Data(day) {
			if !yield(transform(i)) {
				return
			}
		}
	}
}

func SampleProcess[V any](day int, transform func(line string) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range Sample(day) {
			if !yield(transform(i)) {
				return
			}
		}
	}
}
