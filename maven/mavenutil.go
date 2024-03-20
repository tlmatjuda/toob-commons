package maven

import (
	"github.com/tlmatjuda/toob-commons/cli"
	"github.com/tlmatjuda/toob-commons/logs"
	"github.com/tlmatjuda/toob-commons/text"
	"strings"
)

const (
	Maven                  = "mvn"
	Clean                  = "c"
	Install                = "i"
	Offline                = "o"
	Tests                  = "t"
	SkipTests              = "-DskipTests=true"
	SingleThreadPerCPUCore = "-T 1C"
	PomFile                = "pom.xml"
)

func ExtractMavenArgs(args []string) (string, string) {
	targetPath := text.GetArg(args, "-dir")
	mavenBuildOptions := text.GetArg(args, "-opts")
	return targetPath, mavenBuildOptions
}

func Run(targetPath string, mavenOptions string, captureCmdOutput bool, logFilePath string) string {
	var mavenCliFlags []string

	if strings.Contains(mavenOptions, Clean) {
		mavenCliFlags = append(mavenCliFlags, "clean")
	}

	if strings.Contains(mavenOptions, Install) {
		mavenCliFlags = append(mavenCliFlags, "install")
	}

	if strings.Contains(mavenOptions, Offline) {
		mavenCliFlags = append(mavenCliFlags, "-o")
	}

	if !strings.Contains(mavenOptions, Tests) {
		mavenCliFlags = append(mavenCliFlags, SkipTests)
	}

	return execute(targetPath, mavenCliFlags, captureCmdOutput, logFilePath)
}

func RunWithFlags(targetPath string, mavenCliFlags []string, captureCmdOutput bool, logFilePath string) string {
	return execute(targetPath, mavenCliFlags, captureCmdOutput, logFilePath)
}

func execute(targetPath string, mavenCliFlags []string, captureCmdOutput bool, logFilePath string) string {
	if text.StringNotBlank(logFilePath) {
		mavenCliFlags = append(mavenCliFlags, "-log-file="+logFilePath)
	}

	logs.Info.Printf("Running command : [ %v ] %v", Maven, mavenCliFlags)

	var consoleResponse string
	if captureCmdOutput {
		consoleResponse = cli.Exec(Maven, mavenCliFlags, targetPath, captureCmdOutput)
	} else {
		cli.Exec(Maven, mavenCliFlags, targetPath, captureCmdOutput)
	}

	return consoleResponse
}
