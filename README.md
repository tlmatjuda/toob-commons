# TOOB - ( Go Lang ) Commons

This Go module is a collection of various operations. <br/>
The idea is to have a centralized module that has all the commons operations you may need when working on a Go project _(  CLI tool, Web Application ... )_ <br/><br/>
This includes operations for :

* File IO operations.
* Command Line Runner
* Maven
* Text / String
<br/>
<br/>

## TECH STACK

The technology stack used in the project is :

* Golang 1.22.x
* You can use any IDE of your choice.

<br/>
<br/>

## SETTING UP

You have a choice of checking this repo out and using the Go Workspaces feature. <br/>
Or you can pull this module directly into your projects using :

```go
go get github.com/tlmatjuda/toob-commons
```

<br/>
<br/>

## USAGE / EXAPMLES

### CLI

When you are inside your Go module and you need to execute a native commands on the respective Windows or Unix based OS. <br/>
The `bool` is for when you want to capture the command output into a string when set to `true` otherwise it will log out to console as normal.

```go
// Defintion
cli.Exec(command string, commandArgs []string, targetPath string, returnOutput bool) string {}

// Example
cli.Exec("mvn", "clean install", "/project/folder/with/pom.xml", false)
```

<br/>

### File IO

The File IO Operations include :

```go
// Defintion
// Exists
// Checks if a given file Exists on the file system
fileio.Exists(path string) (bool, error) {}


// Example
fileio.Exists("/path/of/the/file.txt")
```

<br/>

### Text

The Text / String Operations include :

```go
// StringBlank
// Strempty checks whether string contains only whitespace or not
text.StringBlank(s string) bool {}


// Example
text.StringBlank("Thabo Matjuda")

// OR

text.StringBlank(" ")

```

<br/>

## CONCLUSION

That's all. These are the operations I found important when building a few CLI tools with Go lang. <br/>
I will be adding more as we go.

<br/>
</br
><br/>
