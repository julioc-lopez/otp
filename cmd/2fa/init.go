package main

import (
	"fmt"
	"log"
	"os"
)

type initCommand struct{}

func (c initCommand) Name() string {
	return "init"
}

func (c initCommand) Run(args []string) bool {
	path := getCfgPath()
	if f, err := os.Open(path); err == nil {
		_ = f.Close()

		return true
	}

	f, err := os.Create(path)

	if err != nil {
		log.Println("error creating init file:", path, err)

		return false
	}

	defer f.Close()

	_, err = f.WriteString(
		`# 2fa configuration
#
# Example:
#
# [key.label]
# issuer = "The Issuer"
# secret = <Base32 encoded secret key>
`)
	if err != nil {
		log.Println("error writing config file:", path, err)

		return false
	}

	return true
}

func (c initCommand) Usage() {
	usage := "    init        create the user config"
	fmt.Println(usage)
}

func (c initCommand) Help() {
	help := "\n" + c.Name() + " usage:\n\n    2fa " + c.Name() + "\n\n"
	help += "    Creates a configuration file at " + getCfgPath() + " if one does not already exist.\n"
	fmt.Println(help)
}
