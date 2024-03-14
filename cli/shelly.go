package cli

import (
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
	"os"
	"os/exec"
)

const CommandErrorTag = "Command Error : \n%s\n"

func Exec(command string, commandArgs []string, targetPath string, returnOutput bool) string {
	var responseOutput string
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
			logs.Error.Fatalf(CommandErrorTag, err)
		}
	}

	return responseOutput
}
