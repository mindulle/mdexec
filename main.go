/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mindulle/mdexec/cmd"
	_ "github.com/mindulle/mdexec/cmd/bulk"
	_ "github.com/mindulle/mdexec/cmd/csv"
	_ "github.com/mindulle/mdexec/cmd/explore"
	_ "github.com/mindulle/mdexec/cmd/http"
	_ "github.com/mindulle/mdexec/cmd/json"
	_ "github.com/mindulle/mdexec/cmd/markdown"
	_ "github.com/mindulle/mdexec/cmd/scrap"
	_ "github.com/mindulle/mdexec/cmd/text"
	_ "github.com/mindulle/mdexec/cmd/text/format"
	_ "github.com/mindulle/mdexec/cmd/text/process"
	_ "github.com/mindulle/mdexec/cmd/text/regexp"
	_ "github.com/mindulle/mdexec/cmd/text/validate"
	_ "github.com/mindulle/mdexec/cmd/text/web"
	_ "github.com/mindulle/mdexec/cmd/url"
)

func main() {
	cmd.Execute()
}
