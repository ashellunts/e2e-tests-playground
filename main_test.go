package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	r := require.New(t)

	cmd := exec.Command("answer")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Start()
	r.NoError(err)

	cmd2 := exec.Command("offer")
	var outb2, errb2 bytes.Buffer
	cmd2.Stdout = &outb2
	cmd2.Stderr = &errb2
	err2 := cmd2.Start()
	r.NoError(err2)

	r.NoError(cmd.Wait())
	r.NoError(cmd2.Wait())

	fmt.Println("out answer:", outb.String(), "err answer:", errb.String())
	fmt.Println("out offer:", outb.String(), "err offer:", errb.String())
}
