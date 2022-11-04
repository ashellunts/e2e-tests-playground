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

	cmd := exec.Command("./answer_app")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Start()
	r.NoError(err)

	cmd2 := exec.Command("./offer_app")
	var outb2, errb2 bytes.Buffer
	cmd2.Stdout = &outb2
	cmd2.Stderr = &errb2
	err2 := cmd2.Start()
	r.NoError(err2)

	exit1 := cmd.Wait()

	exit2 := cmd2.Wait()

	fmt.Println("out offer:", outb2.String(), "err offer:", errb2.String())
	fmt.Println("out answer:", outb.String(), "err answer:", errb.String())

	r.NoError(exit1)
	r.NoError(exit2)
}
