package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Encrypt(args []string) (err error) {
	if len(args) == 0 {
		return errors.New("No encryption path specified")
	}

	for _, path := range args {
		aPath, _ := Abs(path)

		var (
			// assemble the steps that, via bash, will call encfs and make it
			// mount the encrypted volume
			tmp_mount_location  string = fmt.Sprintf("/tmp/mencfs%d", time.Now().Nanosecond())
			tmp_backup_location string = fmt.Sprintf("/tmp/mencfs%d", time.Now().Nanosecond())
			bash_cmd            string = fmt.Sprintf("encfs %s %s", aPath, tmp_mount_location)
		)

		// create the temporary mount and backup locations
		if err = os.Mkdir(tmp_mount_location, 0700); err != nil {
			fmt.Printf("Could not create %s", path)
		}
		if err = os.Mkdir(tmp_backup_location, 0700); err != nil {
			fmt.Printf("Could not create %s", path)
		}

		// make sure that aPath is ready to be used
		if err = ReadyPath(aPath); err != nil {
			fmt.Println(err)
		} else {
			doEncrypt := true
			// back up existing content and remove it from encryption folder 
			if err = CopyDirContent(aPath, tmp_backup_location); err != nil {
				doEncrypt = false
				fmt.Println("Failed to back up existing content in", aPath)
			} else {
				if err = EmptyDir(aPath); err != nil {
					doEncrypt = false
					fmt.Println("Failed to empty", aPath, "before encryption")
				}
			}

			if doEncrypt {
				cmd := exec.Command(GetBash(), "-c", bash_cmd)
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				if err = cmd.Run(); err != nil {
					fmt.Println("Failed to encrypt", aPath)
				}

				fmt.Printf("\n%s has been encrypted\n", aPath)
				fmt.Println("Add a new entry to your configuration file to be able to manage it with MEncFS")

				if err = CopyDirContent(tmp_backup_location, tmp_mount_location); err != nil {
					fmt.Println("Failed to recover existing content of", aPath, " - it's available in", tmp_backup_location)
				}

				if err = unmountTarget(tmp_mount_location); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return nil
}
