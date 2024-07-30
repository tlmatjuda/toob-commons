package cli

import (
	"github.com/tlmatjuda/toob-commons/logs"
	"github.com/tlmatjuda/toob-commons/text"
	"os"
	"os/exec"
	"strings"
)

const (
	TargetShellInterpreter = "/bin/bash"
	CommandErrorTag        = "Command Error : \n%s\n"
)

func Exec(command string, commandArgs []string, targetPath string, returnOutput bool) string {
	return execCommand(command, commandArgs, targetPath, returnOutput, false)
}

func ExecWithNativeLog(command string, commandArgs []string, targetPath string, returnOutput bool) string {
	return execCommand(command, commandArgs, targetPath, returnOutput, true)
}

func ExecScriptFile(scriptPath string, targetPath string, returnOutput bool) string {
	return Exec(TargetShellInterpreter, []string{scriptPath}, targetPath, returnOutput)
}

func execCommand(command string, commandArgs []string, targetPath string, returnOutput bool, logCommand bool) string {
	var responseOutput string
	if logCommand {
		logNativeCommand(command, commandArgs)
	}
	var cmd = exec.Command(command, commandArgs...)

	// Validate the input and see if there's something there
	if text.StringNotBlank(targetPath) {
		cmd.Dir = targetPath
	}

	// Run the command and return the output./to
	if returnOutput {
		result, err := cmd.CombinedOutput()
		if err != nil {
			logs.Error.Printf("Error for command [ %v ] with options : %v", command, commandArgs)
		}

		responseOutput = string(result)

		// Otherwise just run the command and show in the console as it runs
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			logs.Error.Printf(CommandErrorTag, err)
			return text.EMPTY
		}
	}

	return responseOutput
}

func logNativeCommand(command string, commandArgs []string) {
	result := strings.Join(commandArgs, text.WHITE_SPACE)
	logs.Info.Printf("Running Native Command : %v %v", command, result)
}
