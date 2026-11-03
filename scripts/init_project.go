//go:build ignore

// Turns the template into your project. setup.sh runs it once, then it deletes itself.
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	templateOwner = "augiraud"
	templateName  = "go-template"
	thisFile      = "scripts/init_project.go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	in := bufio.NewReader(os.Stdin)
	typed := ask(in, "Project name, such as my-project (leave empty to skip): ")
	name := strings.Trim(regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(strings.ToLower(typed), "-"), "-")
	if name == "" {
		fmt.Println("Skipped: run ./scripts/setup.sh again to name the project.")
		return nil
	}
	if name[0] < 'a' || name[0] > 'z' {
		return errors.New("the project name must start with a letter")
	}
	owner := ask(in, "GitHub user or organization: ")
	if owner == "" {
		return errors.New("the GitHub user or organization is required")
	}

	replacer := strings.NewReplacer(
		"github.com/"+templateOwner+"/"+templateName, "github.com/"+owner+"/"+name,
		templateName, name,
	)
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			// Skip hidden folders, except .github
			if path != "." && strings.HasPrefix(entry.Name(), ".") && entry.Name() != ".github" {
				return filepath.SkipDir
			}
			return nil
		}
		if !entry.Type().IsRegular() || path == thisFile {
			return nil
		}
		return rewrite(path, func(text string) string {
			text = replacer.Replace(text)
			if path == "sonar-project.properties" {
				text = setProperty(text, "sonar.organization", strings.ToLower(owner))
				text = setProperty(text, "sonar.projectKey", owner+"_"+name)
			}
			return text
		})
	})
	if err != nil {
		return err
	}

	if err := os.Rename(filepath.Join("cmd", templateName), filepath.Join("cmd", name)); err != nil {
		return err
	}
	return os.Remove(thisFile)
}

func ask(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := in.ReadString('\n')
	return strings.TrimSpace(line)
}

// rewrite applies change to a text file, and leaves binary files alone.
func rewrite(path string, change func(string) string) error {
	data, err := os.ReadFile(path)
	if err != nil || !utf8.Valid(data) || bytes.IndexByte(data, 0) >= 0 {
		return err
	}
	text := change(string(data))
	if text == string(data) {
		return nil
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(text), info.Mode().Perm())
}

func setProperty(text, key, value string) string {
	return regexp.MustCompile(`(?m)^`+regexp.QuoteMeta(key)+`=.*$`).ReplaceAllLiteralString(text, key+"="+value)
}
