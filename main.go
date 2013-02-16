package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	NAME        string = "MEncFS"
	DESCRIPTION string = "Manage EncFS"
	VERSION     string = "0.8.1"
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s - %s %s\n\n", NAME, DESCRIPTION, VERSION)
	fmt.Fprintf(os.Stderr, "usage: %s [arguments ..] [action] [path ..]\n\n", NAME)
	fmt.Fprintf(os.Stderr, "arguments:\n")
	fmt.Fprintf(os.Stderr, "  --force\tForce action\n")
	fmt.Fprintf(os.Stderr, "\nactions:\n")
	fmt.Fprintf(os.Stderr, "  encrypt\tEncrypt a folder\n")
	fmt.Fprintf(os.Stderr, "  mount\t\tMount Volumes as defined in the config file\n")
	fmt.Fprintf(os.Stderr, "  umount\tUnmount Volumes defined in the config file\n")
	fmt.Fprintf(os.Stderr, "  generate\tGenerate config file %s\n", GetConfigFilePath())
	fmt.Fprintf(os.Stderr, "  automount\tToggle automounting of Volumes when the computer starts\n")
	os.Exit(2)
}

func main() {
	var (
		action string = ""
		err    error  = nil
	)

	Greete()

	force := flag.Bool("force", false, "Force action")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		action = args[0]
	}

	switch action {
	case "encrypt":
		rest_args := args[1:]
		err = Encrypt(rest_args)
	case "generate":
		err = GenerateConfig(*force)
	case "mount":
		err = Mount()
	case "umount":
		err = UMount()
	case "automount":
		err = AutoMount()
	default:
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
