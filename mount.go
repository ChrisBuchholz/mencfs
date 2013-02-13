package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func mountTarget(source string, target string, title string, keychain_label string) (err error) {
	var (
		doesNotExist bool = true
		// get the password to decrypt the volume
		// if no password is found, keychain_password will be empty which
		// will result in encfs decrypting rubbish
		keychain_password string = GetPassword_mac(keychain_label)
		// assemble the steps that, via bash, will call encfs and make it
		// mount the encrypted volume
		encfs    string = "encfs"
		extpass  string = fmt.Sprintf("--extpass=\"/bin/bash -c \\\"echo '%s'\\\" \"", keychain_password)
		rargs    string = fmt.Sprintf("-ovolname=%s -oallow_root -olocal -ohard_remove -oauto_xattr -onolocalcaches", title)
		bash_cmd string = fmt.Sprintf("%s %s %s %s %s", encfs, source, target, extpass, rargs)
	)

	// find out if the mountpoint already exist and if it does but it
	// is a file, not a directory, bail
	file, err := os.Open(target)
	if err == nil {
		doesNotExist = false
		fi, err := file.Stat()
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return errors.New(fmt.Sprintf("Mountpoint %s is not a folder", target))
		}
	}

	// if it doesnt exist, create the mountpoint
	if doesNotExist {
		if err = os.Mkdir(target, 0700); err != nil {
			return errors.New(fmt.Sprintf("Couldn't create target %s", target))
		}
		fmt.Println("Created new mountpoint", target)
	}

	// tell encfs to decrypt the volume
	// we do this by telling bash to execute our assembled bash command
	//
	// we assume that bash is installed to /bin/bash, this is a bad thing
	// and should be done otherwise
	cmd := exec.Command("/bin/bash", "-c", bash_cmd)
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Failed to mount %s", target))
	}
	fmt.Println(target, "mounted")

	return nil
}

func unmountTarget(target string) (err error) {
	var (
		out bytes.Buffer
		// assuming the location of diskutil is bad and locating diskutil
		// should be done otherwise
		diskutil string = "/usr/sbin/diskutil"
	)

	// mount returns all mounted volumes
	cmd := exec.Command("mount")
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return err
	}

	// we use the output from `mount` to check if out Volume is mounted
	// and only unmounts it if it has
	if strings.Contains(out.String(), target) {
		cmd = exec.Command(diskutil, "unmount", target)
		if err = cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to unmount %s", target))
		}
		fmt.Println(target, "unmounted")
	}

	return nil
}

// Mount any volumes defined in the config file
func Mount() (err error) {
	volumes, err := ReadConfig()

	for _, volume := range volumes {
		target := VOLUME_MOUNT_POINT + volume.Title
		err = mountTarget(volume.Source, target, volume.Title, volume.KeychainLabel)
		if err != nil {
			return err
		}
	}

	return nil
}

// Unmount any mounted volumes defined in the config file
func UMount() (err error) {
	volumes, err := ReadConfig()

	for _, volume := range volumes {
		target := VOLUME_MOUNT_POINT + volume.Title
		err = unmountTarget(target)
		if err != nil {
			return err
		}
	}

	return nil
}

func AutoMount() (err error) {
	fmt.Println("automount")
	return nil
}
