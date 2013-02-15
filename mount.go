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
	source, _ = Abs(source)
	target, _ = Abs(target)

	var (
		// get the password to decrypt the volume
		// if no password is found, keychain_password will be empty which
		// will result in encfs decrypting rubbish
		keychain_password string = GetPassword_mac(keychain_label)
		// assemble the steps that, via bash, will call encfs and make it
		// mount the encrypted volume
		encfs_args string = fmt.Sprintf("--extpass=\"%s -c \\\"echo '%s'\\\"\" -ovolname=%s -oallow_root -olocal -ohard_remove -oauto_xattr -onolocalcaches", GetBash(), keychain_password, title)
		bash_cmd   string = fmt.Sprintf("%s %s %s %s", GetEncFS(), source, target, encfs_args)
	)

	// if source doesnt exist, why might as well bail
	_, err = os.Open(source)
	if err != nil {
		return errors.New(fmt.Sprintf("No encrypted folder %s", source))
	}

	// make sure that target is ready to be used
	if err := ReadyPath(target); err != nil {
		return err
	}

	// tell encfs to decrypt the volume
	// we do this by telling bash to execute our assembled bash command
	cmd := exec.Command(GetBash(), "-c", bash_cmd)
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Failed to mount %s", target))
	}

	return nil
}

func unmountTarget(target string) (err error) {
	var out bytes.Buffer

	// mount returns all mounted volumes
	cmd := exec.Command(GetMount())
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return err
	}

	// we use the output from `mount` to check if out Volume is mounted
	// and only unmounts it if it has
	if strings.Contains(out.String(), target) {
		cmd = exec.Command(GetUnmount(), "unmount", target)
		if err = cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to unmount %s", target))
		}
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
		fmt.Println(target, "mounted")
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
		fmt.Println(target, "unmounted")
	}

	return nil
}

func AutoMount() (err error) {
	fmt.Println("automount not implemented yet")
	return nil
}
