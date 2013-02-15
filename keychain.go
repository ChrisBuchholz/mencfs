// the idea is, that there will exist different platform dependent
// implementation of how to get a Volume's encryption password, using
// the platforms standard keychain-protocol

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// get the Volume's encryption password from Mac OS Xs keychain, using the
// label the Volume has defined in the config file
func GetPassword_mac(label string) string {
	var (
		// this elaborate bash command will look up the Volume's encryption
		// password via the label
		// we do it via bash since the security program is build in to OS X
		// and therefore doesnt add an extra dependency
		bash_cmd string = fmt.Sprintf("security 2>&1 >/dev/null find-generic-password -gl '%s' |grep password|cut -d \\\" -f 2", label)
		out      bytes.Buffer
	)

	cmd := exec.Command(GetBash(), "-c", bash_cmd)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}

	return strings.TrimSpace(out.String())
}
