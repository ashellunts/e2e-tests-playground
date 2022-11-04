package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func RunProcess(proc_name string, r *require.Assertions) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	answer_process := exec.Command(proc_name)

	var answer_stdout, answer_stderr bytes.Buffer
	answer_process.Stdout = &answer_stdout
	answer_process.Stderr = &answer_stderr

	r.NoError(answer_process.Start())

	return answer_process, &answer_stdout, &answer_stderr
}

func Test(t *testing.T) {

	t.Run("ajfkjalks", func(t *testing.T) {
		r := require.New(t)

		answer, answer_stdout, answer_stderr := RunProcess("./answer_app", r)
		time.Sleep(time.Second)
		offer, offer_stdout, offer_stderr := RunProcess("./offer_app", r)

		time.AfterFunc(3*time.Second, func() {
			answer.Process.Kill()
			offer.Process.Kill()
		})

		offer_exit_err := offer.Wait()
		answer_exit_err := answer.Wait()

		PrintOutput(offer_stdout, offer_stderr, answer_stdout, answer_stderr)

		r.NoError(offer_exit_err)
		r.NoError(answer_exit_err)
	})
}

func PrintOutput(offer_stdout *bytes.Buffer, offer_stderr *bytes.Buffer, answer_stdout *bytes.Buffer, answer_stderr *bytes.Buffer) {
	fmt.Println("offer stdout:", offer_stdout.String())
	fmt.Println("offer stderr:", offer_stderr.String())

	fmt.Print("\n\n")
	fmt.Println("answer stdout:", answer_stdout.String())
	fmt.Println("answer stderr:", answer_stderr.String())
}
