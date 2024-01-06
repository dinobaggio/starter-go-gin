package libs

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func CurrentDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}

func RootDir() string {
	findIndex := func(arr []string, word string) int {
		var result int
		for i := range arr {
			if arr[i] == word {
				result = i
				break
			}
		}

		return result
	}

	d := strings.Split(path.Join(CurrentDir()), "/")
	indx := findIndex(d, "starter-go-gin")

	d = d[0 : indx+1]

	return strings.Join(d, "/")
}

func EnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(filepath.Join(RootDir(), ".env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
