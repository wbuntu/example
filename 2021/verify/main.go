package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		logStdoutf("invalid args: %v", os.Args)
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		logStdoutf("open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logStdoutf("scan file: %s", err)
		os.Exit(1)
	}
	if len(params) != 2 {
		logStdoutf("invalid params: %v", params)
		os.Exit(1)
	}
	name := params[0]
	password := params[1]

	fileName := "verify-ccd/" + getMD5(name)
	userData, err := ioutil.ReadFile(fileName)
	if err != nil {
		logStdoutf("ioutil.ReadFile: %s", err)
		os.Exit(1)
	}

	expectPassword := strings.TrimSuffix(string(userData), "\n")
	if expectPassword != password {
		logStdoutf("password mismatch: %s -> %s", expectPassword, password)
		os.Exit(1)
	}
}

func getMD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func logStdout(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func logStdoutf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format, a...)
	fmt.Println()
}

func logStderr(a ...interface{}) {
	fmt.Fprintln(os.Stderr, append([]interface{}{"Error:"}, a...)...)
}

func logStderrf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format, a...)
	fmt.Println()
}
