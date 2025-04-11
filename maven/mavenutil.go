package maven

import (
	"github.com/tlmatjuda/toob-commons/cli"
	"github.com/tlmatjuda/toob-commons/text"
	"strings"
)

const (
	Maven                  = "mvn"
	CleanId                = "c"
	InstallId              = "i"
	Offline                = "o"
	Tests                  = "t"
	SkipTests              = "-DskipTests=true"
	SingleThreadPerCPUCore = "-T 1C"
	PomFile                = "pom.xml"
)

var (
	mavenCleanInstallArgs           = []string{"clean", "install"}
	mavenCleanInstallSkipTestskArgs = append(mavenCleanInstallArgs, SkipTests)
	mavenSpotlessApplyArgs          = []string{"spotless:apply"}
)

func ExtractMavenArgs(args []string) (string, string) {
	targetPath := text.GetArg(args, "-dir")
	mavenBuildOptions := text.GetArg(args, "-opts")
	return targetPath, mavenBuildOptions
}

func Run(targetPath string, mavenOptions string, captureCmdOutput bool, logFilePath string) string {
	var mavenCliFlags []string

	if strings.Contains(mavenOptions, CleanId) {
		mavenCliFlags = append(mavenCliFlags, "clean")
	}

	if strings.Contains(mavenOptions, InstallId) {
		mavenCliFlags = append(mavenCliFlags, "install")
	}

	if strings.Contains(mavenOptions, Offline) {
		mavenCliFlags = append(mavenCliFlags, "-o")
	}

	if !strings.Contains(mavenOptions, Tests) {
		mavenCliFlags = append(mavenCliFlags, SkipTests)
	}

	if text.StringNotBlank(logFilePath) {
		mavenCliFlags = append(mavenCliFlags, "-log-file="+logFilePath)
	}

	return execute(targetPath, mavenCliFlags, captureCmdOutput)
}

func RunWithFlags(targetPath string, mavenCliFlags []string, captureCmdOutput bool) string {
	return execute(targetPath, mavenCliFlags, captureCmdOutput)
}

func CleanInstall(targetPath string, captureCmdOutput bool) string {
	return execute(targetPath, mavenCleanInstallArgs, captureCmdOutput)
}

func CleanInstallSkipTests(targetPath string, captureCmdOutput bool) string {
	return execute(targetPath, mavenCleanInstallSkipTestskArgs, captureCmdOutput)
}

func CleanInstallLogFile(targetPath string, logFile string, captureCmdOutput bool) string {
	mavenCleanInstallArgs = append(mavenCleanInstallArgs, "-log-file="+logFile)
	return execute(targetPath, mavenCleanInstallArgs, captureCmdOutput)
}

func CleanInstallSkipTestsLogFile(targetPath string, logFile string, captureCmdOutput bool) string {
	mavenCleanInstallSkipTestskArgs = append(mavenCleanInstallSkipTestskArgs, "-log-file="+logFile)
	return execute(targetPath, mavenCleanInstallSkipTestskArgs, captureCmdOutput)
}

// SPOTLESS:APPLICY SECTION
func SpotlessApply(targetPath string, captureCmdOutput bool) string {
	return execute(targetPath, mavenSpotlessApplyArgs, captureCmdOutput)
}

func SpotlessApplyLogFile(targetPath string, logFile string, captureCmdOutput bool) string {
	mavenSpotlessApplyArgs = append(mavenSpotlessApplyArgs, "-log-file="+logFile)
	return execute(targetPath, mavenSpotlessApplyArgs, captureCmdOutput)
}

func BuildFailed(buildResponse string) bool {
	return text.StringNotBlank(buildResponse) &&
		strings.Contains(buildResponse, "BUILD FAILURE")
}

func execute(targetPath string, mavenCliFlags []string, captureCmdOutput bool) string {
	var consoleResponse string
	if captureCmdOutput {
		consoleResponse = cli.Exec(Maven, mavenCliFlags, targetPath, captureCmdOutput)
	} else {
		cli.Exec(Maven, mavenCliFlags, targetPath, captureCmdOutput)
	}

	return consoleResponse
}
