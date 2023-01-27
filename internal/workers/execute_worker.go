package workers

import (
	"bytes"
	"errors"
	"exec-processor/internal/entity"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

func ProcessExecution(clientData entity.ClientData) (error, *entity.ClientRespondData) {
	var splitedCommand = strings.Split(clientData.Cmd, " ")
	var commandName = splitedCommand[0]

	cmd := exec.Command(commandName)
	cmd.Args = splitedCommand
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err, nil
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if runtime.GOOS != clientData.OsName {
		return err, nil
	}
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Start(); err != nil {
		return err, nil
	}
	_, err = io.WriteString(stdin, clientData.Stdin)
	if err != nil {
		return err, nil
	}

	err = cmd.Wait()
	if err != nil {
		return err, nil
	}
	return err, &entity.ClientRespondData{Stdout: string(stdout.Bytes()), Stderr: string(stderr.Bytes())}
}
