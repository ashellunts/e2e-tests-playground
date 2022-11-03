package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	r := require.New(t)

	cmd := exec.Command("offer")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	r.Error(err)
}
