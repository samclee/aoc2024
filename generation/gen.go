package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const rootCobraFolder string = "cmd"
const rootDaysFolder string = "days"

type PackageAndIndex struct {
	Index       int
	PackageName string
}

func main() {
	// make folders + files
	var dayPackages []PackageAndIndex
	for i := 1; i < 26; i++ {
		packageName, err := createDayFolder(i)
		if err != nil {
			fmt.Printf("Error occurred on day [%d]: %s", i, err.Error())
		}
		dayPackages = append(dayPackages, PackageAndIndex{i, packageName})
	}
	// make solver function
	fullpath := filepath.Join(rootCobraFolder, "solver.go")
	f, err := os.Create(fullpath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = solverFileTemplate.Execute(f, struct {
		DayPackages []PackageAndIndex
	}{
		DayPackages: dayPackages,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

}

func createDayFolder(day int) (string, error) {
	folderName := fmt.Sprintf("day%d", day)
	err := createSolFile(folderName, 1)
	if err != nil {
		return "", err
	}
	err = createSolFile(folderName, 2)
	if err != nil {
		return "", err
	}

	return folderName, nil
}

func createSolFile(folderName string, part int) error {
	fname := fmt.Sprintf("sol%d.go", part)
	fullpath := filepath.Join(rootDaysFolder, folderName, fname)
	err := os.MkdirAll(filepath.Dir(fullpath), 0777)
	if err != nil {
		return err
	}
	f, err := os.Create(fullpath)
	if err != nil {
		return err
	}

	err = solFileTemplate.Execute(f, struct {
		PackageName string
		Part        int
	}{
		PackageName: folderName,
		Part:        part,
	})
	if err != nil {
		return err
	}

	return nil
}

var solFileTemplate = template.Must(template.New("").Parse(`package {{ .PackageName }}

import "os"

func Sol{{ .Part }}(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	return ""
}`))

var solverFileTemplate = template.Must(template.New("").Parse(`package cmd

import(
{{- range .DayPackages }}
	"aoc2024/days/{{.PackageName}}"
{{- end }}
)

func Solve(day int, part int, inputFname string) (string, error) {
	switch day{
	{{- range .DayPackages }}
	case {{.Index}}:
		if part == 1 {
			return {{.PackageName}}.Sol1(inputFname), nil
		} else {
			return {{.PackageName}}.Sol2(inputFname), nil
		}
	{{- end }}
	default:
		return "", nil
	}
}

`))
