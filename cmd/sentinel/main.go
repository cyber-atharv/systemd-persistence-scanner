/*
©cyber-atharv | 2026
main.go

Entry point for the sentinel CLI
*/

package main

import (
	"github.com/cyber-atharv/sentinel/internal/cli"
	_ "github.com/cyber-atharv/sentinel/internal/scanner"
)

func main() {
	cli.Execute()
}
