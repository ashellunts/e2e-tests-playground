package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	cmd, err := RunCommand("echo", "hello")

	r := require.New(t)
	r.NotNil(cmd)
	r.Nil(err)

	cmd_err := cmd.Wait()

	r.Nil(cmd_err)

	// program_output := make([]byte, 0)
	// for {
	// 	tmp := make([]byte, 4)
	// 	n, err := stdoutPipe.Read(tmp)
	// 	if err != nil {
	// 		break
	// 	}
	// 	program_output = append(program_output, tmp[0:n]...)
	// }

	// expected := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x0a}
	// a.Equal(expected, program_output)
}

func RunCommand(name string, arg ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, arg...)

	// stdoutPipe, err := cmd.StdoutPipe()
	// if err != nil {
	// 	return nil, err
	// }

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}
