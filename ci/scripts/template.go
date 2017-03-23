package main

import (
	"fmt"
	"html/template"
	"os"
)

//Formula ...
type Formula struct {
	URL     string
	Version string
	Hash    string
}

func main() {

	msg := `
class GitPhlow < Formula
  desc "pragmatic workflow extension for git"
  homepage "https://github.com/Praqma/git-phlow"
  url "{{.URL}}"
  version "{{.Version}}"
  sha256 "{{.Hash}}"

  def install
    bin.install "git-phlow"
  end

  def caveats; <<-EOS.undent
    Make sure git is installed
    EOS
  end
  test do
    system "true"
  end
end

`

	version := os.Getenv("VERSION")
	url := os.Getenv("URL")
	hash := os.Getenv("HASH")

	substitute := Formula{url, version, hash}

	tmpl, err := template.New("msg").Parse(msg)
	err = tmpl.Execute(os.Stdout, substitute)
	if err != nil {
		fmt.Println(err)

	}
}
