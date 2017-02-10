package textfixture

import (
	"os"
	"archive/zip"
	"path/filepath"
	"io"
	"fmt"
	"errors"
)

var (
	GoPathNotSet error = errors.New("GOPATH is empty")
	goPath       string
	phlowPath    = goPath + "/src/github.com/praqma/git-phlow"
)

//init
//Runs before functions to setup variable gopath
func init() {
	goPath = os.Getenv("GOPATH")
	if len(goPath) == 0 {
		fmt.Fprintln(os.Stdout, GoPathNotSet)
		os.Exit(1)
	}
}

//unzip
//unzips archive to target directory
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		fmt.Println(archive)
		return err
	}
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

//SetupTestRepo
//Creates git test repository from a zip file in /testfixture
func SetupTestRepo() {

	target := phlowPath + "/build"
	archive := phlowPath + "/testfixture/phlow-test-pkg.zip"
	err := unzip(archive, target)

	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	os.Chdir(target)

	fmt.Fprintln(os.Stdout, "Local test repository created from 'zip'")
}

//TearDownTestRepo
//removes the unzipped test repository is it exists
func TearDownTestRepo() {

	var target = phlowPath + "/build"

	os.Chdir(phlowPath)
	err := os.RemoveAll(target)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Deleted local test repository")

}
