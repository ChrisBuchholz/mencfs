// Get absolute path to binaries of external components like bash

package main

import (
	"log"
	"os/exec"
)

func GetBash() string {
	exe, err := exec.LookPath("bash")
	if err != nil {
		log.Fatal("Looks like you don't have bash installed correctly...")
	}
	return exe
}

func GetMount() string {
	exe, err := exec.LookPath("mount")
	if err != nil {
		log.Fatal("Looks like you don't have mount installed correctly...")
	}
	return exe
}

func GetUnmount() string {
	exe, err := exec.LookPath("diskutil")
	if err != nil {
		log.Fatal("Looks like you don't have diskutil installed correctly...")
	}
	return exe
}

func GetEncFS() string {
	exe, err := exec.LookPath("encfs")
	if err != nil {
		log.Fatal("Looks like you don't have EncFS installed correctly...")
	}
	return exe
}

func GetSecurity() string {
	exe, err := exec.LookPath("security")
	if err != nil {
		log.Fatal("Looks like you don't have security installed correctly...")
	}
	return exe
}
