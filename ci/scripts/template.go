package main

import (
	"fmt"
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
  url "%s"
  version "%s"
  sha256 "%s"

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

	fmt.Fprintf(os.Stdout, msg, url, version, hash)

}
