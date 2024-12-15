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

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	adventCookieFile := path.Join(homeDir, ".advent")
	cookieFile, err := os.Open(adventCookieFile)
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
	filename := dayFileOnDisk(t, day)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		if t == "data" {
			err := downloadDay(day)
			if err != nil {
				panic(err)
			}
		}

		if t == "sample" {
			fmt.Println("No sample data, input it here or cancel:")

			f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}

			//TODO: how am I supposed to end this part? like an email,<cr>.<cr>, lol
			_, err = io.Copy(f, os.Stdin)
			if err != nil {
				panic(err)
			}

			if err := f.Close(); err != nil {
				return nil
			}
		}
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
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
	err := os.MkdirAll(t, 0755)
	if err != nil {
		panic(err)
	}
	return path.Join(t, fmt.Sprintf("%02d.txt", day))
}

func Data(day int) iter.Seq[string] {
	return GetData("data", day)
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

func Sample(day int) iter.Seq[string] {
	return GetData("sample", day)
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

func Enumerate[I int, V any](seq iter.Seq[V]) iter.Seq2[I, V] {
	return func(yield func(I, V) bool) {
		var i I

		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}

	}
}
