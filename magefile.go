//go:build mage

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/magefile/mage/sh"
)

const tailwindVersion = "v4.1.11"

// InstallTailwind downloads the Tailwind v4 standalone CLI.
func InstallTailwind() error {
	binary := tailwindBinaryPath()
	if _, err := os.Stat(binary); err == nil {
		fmt.Println("Tailwind already installed, skipping.")
		return nil
	}
	if err := os.MkdirAll(".bin", 0o755); err != nil {
		return err
	}
	url := tailwindDownloadURL()
	fmt.Printf("Downloading Tailwind %s from %s\n", tailwindVersion, url)
	if err := sh.Run("curl", "-sLo", binary, url); err != nil {
		return err
	}
	return sh.Run("chmod", "+x", binary)
}

// BuildCSS compiles the showcase stylesheet.
func BuildCSS() error {
	return sh.Run(
		tailwindBinaryPath(),
		"-i", "./styles/flint.css",
		"-o", "./examples/showcase/static/css/showcase.css",
		"--minify",
	)
}

// WatchCSS runs Tailwind in watch mode.
func WatchCSS() error {
	return sh.Run(
		tailwindBinaryPath(),
		"-i", "./styles/flint.css",
		"-o", "./examples/showcase/static/css/showcase.css",
		"--watch",
	)
}

// GenerateTempl runs templ generate across the repo.
func GenerateTempl() error {
	return sh.Run("templ", "generate")
}

// Showcase builds CSS + templ then runs the reference site.
func Showcase() error {
	if err := GenerateTempl(); err != nil {
		return err
	}
	if err := BuildCSS(); err != nil {
		return err
	}
	return sh.Run("go", "run", "./examples/showcase")
}

// Build produces a static showcase binary.
func Build() error {
	if err := GenerateTempl(); err != nil {
		return err
	}
	if err := BuildCSS(); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/showcase", "./examples/showcase")
}

// Test runs the full Go test suite.
func Test() error {
	if err := GenerateTempl(); err != nil {
		return err
	}
	return sh.Run("go", "test", "./...")
}

func tailwindBinaryPath() string {
	if runtime.GOOS == "windows" {
		return "./.bin/tailwindcss.exe"
	}
	return "./.bin/tailwindcss"
}

func tailwindDownloadURL() string {
	osName := map[string]string{
		"darwin":  "macos",
		"linux":   "linux",
		"windows": "windows",
	}[runtime.GOOS]

	archName := map[string]string{
		"amd64": "x64",
		"arm64": "arm64",
	}[runtime.GOARCH]

	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	return fmt.Sprintf(
		"https://github.com/tailwindlabs/tailwindcss/releases/download/%s/tailwindcss-%s-%s%s",
		tailwindVersion, osName, archName, ext,
	)
}
