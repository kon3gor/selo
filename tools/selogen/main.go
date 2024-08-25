package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var t string
	flag.StringVar(&t, "type", "", "")

	flag.Parse()

	f := readLockFile()
	defer f.Close()

	cnt := parseNumFromFile(f)

	createKeyfileForType(t, cnt+1)

	err := f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(strconv.Itoa(cnt + 1))
	if err != nil {
		log.Fatal(err)
	}

}

func parseNumFromFile(f *os.File) int {
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if len(b) == 0 {
		return -1
	}

	cnt, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}

	return cnt
}

const lockFile = ".selogen"

func readLockFile() *os.File {
	f, err := os.OpenFile(lockFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	lifetime := lockFileLifetime()
	stat, err := os.Stat(lockFile)
	if err != nil {
		log.Fatal(err)
	}

	delta := time.Since(stat.ModTime())

	if delta > lifetime {
		err := f.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	return f
}

func lockFileLifetime() time.Duration {
	return 1 * time.Second
}

const keyFileTemplate = `// GENERATED USING SELOGEN PLEASE DO NOT EDIT BY HAND
package selogen

import "gihub.com/kon3gor/selo"

const %s = %d
`

func createKeyfileForType(t string, key int) {
	content := fmt.Sprintf(keyFileTemplate, t, key)

	_, err := os.Stat("selogen")
	if os.IsNotExist(err) {
		err = os.Mkdir("selogen", 0777)
	}

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("selogen/%s.generated.go", strings.ToLower(t)))
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}
