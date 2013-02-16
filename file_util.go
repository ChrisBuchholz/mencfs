package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
)

// ReadyFolder makes sure that a path is ready to be used as a
// folder
// this includes checking if the path exist and if it does, checking
// that its not a file, and that if it does not exist, creating it
func ReadyPath(path string) (err error) {
	var doesNotExist = true
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		doesNotExist = false
		fi, err := file.Stat()
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return errors.New(fmt.Sprintf("%s is not a folder", path))
		}
	}
	if doesNotExist {
		if err = os.Mkdir(path, 0700); err != nil {
			return errors.New(fmt.Sprintf("Could not create %s", path))
		}
		fmt.Println("Created", path)
	}
	return nil
}

// get absolute path from possibly relative path
func Abs(name string) (string, error) {
	if name[0:1] == "~" {
		if u, err := user.Current(); err == nil {
			return strings.Replace(name, "~", u.HomeDir, 1), nil
		}
	} else {
		if path.IsAbs(name) {
			return name, nil
		}
		wd, err := os.Getwd()
		return path.Join(wd, name), err
	}
	return name, nil
}

// Copy the content of a directory to another
// This include hidden (.dot) files
func CopyDirContent(from string, to string) (err error) {
	from, _ = Abs(from)
	to, _ = Abs(to)
	// make sure that from ends with a slash so that it doesnt copy
	// the folder but the content of it
	if from[len(from)-1:len(from)] != "/" {
		from = from + "/"
	}
	cp_cmd := fmt.Sprintf("cp -r %s %s", from, to)
	cmd := exec.Command(GetBash(), "-c", cp_cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Failed to copy content from %s to %s", from, to))
	}
	return nil
}

// EmptyDir will empty a dir by deleting every file and folder inside the
// directory
// This includes hidden (.dot) files
func EmptyDir(target string) (err error) {
	target, _ = Abs(target)
	if target[len(target)-1:len(target)] != "/" {
		target = target + "/"
	}
	// do not allow deletion of / or anything in cwd
	if target != "/" || target != "" {
		rm_cmd := fmt.Sprintf("rm -rf %s* %s.[!.]*", target, target)
		cmd := exec.Command(GetBash(), "-c", rm_cmd)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err = cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to empty directry %s", target))
		}
	}
	return nil
}
