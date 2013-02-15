package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
)

// define the mount point (directory) of volumes
const VOLUME_MOUNT_POINT string = "/Volumes/"

type Volume struct {
	Source        string
	Title         string
	KeychainLabel string
}

// Return the path to the config file, which is predefined as .mencfs
// inside the executing users home directory
func GetConfigFilePath() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir + "/.mencfs"
}

// Parse a config file line, which is a tab-seperated list of
// Source, Title and KeychainLabel
func parseConfigLine(line string) (Volume, error) {
	var (
		parts  []string
		volume Volume
	)

	r := regexp.MustCompile(`[a-zA-Z0-9_\-\/~]+`)
	matches := r.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		parts = append(parts, match[0])
	}

	switch {
	case len(parts) == 3:
		volume = Volume{parts[0], parts[1], parts[2]}
	case len(parts) < 3:
		volume = Volume{}
		return volume, errors.New("Not enough parameters.")
	}

	return volume, nil
}

// Read each line of file at path into a slice containing each line as a string
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}

	if err == io.EOF {
		err = nil
	}

	return
}

// Read the instructions in the config file into a slice of Volume's defining
// them
func ReadConfig() (volumes []Volume, err error) {
	config_path := GetConfigFilePath()

	if _, err = os.Open(config_path); err != nil {
		return volumes, errors.New("No config file found, generate a config with '--generate-config'")
	}

	lines, err := readLines(config_path)
	if err != nil {
		return volumes, err
	}

	for i, line := range lines {
		if volume, err := parseConfigLine(line); err != nil {
			fmt.Println("Failed to parse line ", i)
		} else {
			volumes = append(volumes, volume)
		}
	}

	return volumes, nil
}

// Generate a new config file at the path defined in GetConfigFilePath() if
// file doesn't already exist
// if force=true, generate a new file no matter what, overwriting any
// pre-existing configurations
func GenerateConfig(force bool) (err error) {
	var (
		content     string = "~/encrypted		volume_name		password_keychain_label"
		config_path string = GetConfigFilePath()
	)

	if !force {
		_, err = os.Open(config_path)
		if err == nil {
			return errors.New(fmt.Sprintf("%s already exist, use '--force' to regenerate it", config_path))
		}
	}

	content_b := []byte(content)
	err = ioutil.WriteFile(config_path, content_b, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't create %s", config_path))
	}
	fmt.Println("Generated new configuration file", config_path)

	return nil
}
