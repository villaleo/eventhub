//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/magefile/mage/sh"
)

// Install installs all the dependencies required by the application.
func Install() error {
	if err := InstallProtoc(); err != nil {
		return err
	}

	if err := InstallGolangciLint(); err != nil {
		return err
	}

	return nil
}

// InstallProtoc installs the Protocol Buffers compiler (protoc).
func InstallProtoc() error {
	fmt.Println("Checking if protoc is already installed...")

	// Check if `protoc` is already installed
	if _, err := exec.LookPath("protoc"); err == nil {
		fmt.Println("protoc is already installed.")
		return nil
	}

	fmt.Println("protoc not found. Installing...")

	var url string
	switch runtime.GOOS {
	case "darwin":
		url = "https://github.com/protocolbuffers/protobuf/releases/download/v21.3/protoc-21.3-osx-x86_64.zip"
	case "linux":
		url = "https://github.com/protocolbuffers/protobuf/releases/download/v21.3/protoc-21.3-linux-x86_64.zip"
	case "windows":
		url = "https://github.com/protocolbuffers/protobuf/releases/download/v21.3/protoc-21.3-win64.zip"
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// Download the release archive
	if err := sh.Run("curl", "-OL", url); err != nil {
		return fmt.Errorf("failed to download protoc: %w", err)
	}

	// Unzip the downloaded file
	zipFile := url[strings.LastIndex(url, "/")+1:]
	if err := sh.Run("unzip", "-o", zipFile, "-d", "./protoc"); err != nil {
		return fmt.Errorf("failed to unzip protoc: %w", err)
	}

	// Move the binary to /usr/local/bin
	if err := sh.Run("sudo", "mv", "./protoc/bin/protoc", "/usr/local/bin/protoc"); err != nil {
		return fmt.Errorf("failed to move protoc binary: %w", err)
	}

	// Clean up
	os.Remove(zipFile)
	os.RemoveAll("./protoc")

	fmt.Println("protoc installed successfully.")
	return nil
}

// GenerateProtobuf generates the protobuf files.
func GenerateProtobuf() error {
	fmt.Println("Generating protobuf files...")

	protoCommand := `protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			eventhub/eventhub.proto`

	if err := sh.Run("sh", "-c", protoCommand); err != nil {
		return fmt.Errorf("failed to generate protobuf files: %w", err)
	}

	fmt.Println("Protobuf files generated successfully.")
	return nil
}

// InstallGolangciLint installs the golangci-lint CLI tool.
func InstallGolangciLint() error {
	fmt.Println("Checking if golangci-lint is already installed...")

	// Check if `golangci-lint` is already installed
	if _, err := exec.LookPath("golangci-lint"); err == nil {
		fmt.Println("golangci-lint is already installed.")
		return nil
	}

	fmt.Println("golangci-lint not found. Installing...")

	var url string
	switch runtime.GOOS {
	case "darwin":
		url = "https://github.com/golangci/golangci-lint/releases/download/v1.52.2/golangci-lint-1.52.2-darwin-amd64.tar.gz"
	case "linux":
		url = "https://github.com/golangci/golangci-lint/releases/download/v1.52.2/golangci-lint-1.52.2-linux-amd64.tar.gz"
	case "windows":
		url = "https://github.com/golangci/golangci-lint/releases/download/v1.52.2/golangci-lint-1.52.2-windows-amd64.zip"
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// Download the release archive
	if err := sh.Run("curl", "-OL", url); err != nil {
		return fmt.Errorf("failed to download golangci-lint: %w", err)
	}

	// Unpack the downloaded file
	tarFile := url[strings.LastIndex(url, "/")+1:]
	if runtime.GOOS == "windows" {
		if err := sh.Run("unzip", "-o", tarFile); err != nil {
			return fmt.Errorf("failed to unzip golangci-lint: %w", err)
		}
	} else {
		if err := sh.Run("tar", "-xzf", tarFile); err != nil {
			return fmt.Errorf("failed to unpack golangci-lint: %w", err)
		}
	}

	// Move the binary to /usr/local/bin
	if runtime.GOOS == "windows" {
		if err := sh.Run("move", "./golangci-lint-1.52.2-windows-amd64/golangci-lint.exe", "/usr/local/bin/golangci-lint.exe"); err != nil {
			return fmt.Errorf("failed to move golangci-lint binary: %w", err)
		}
	} else {
		if err := sh.Run("sudo", "mv", "./golangci-lint-1.52.2-$(runtime.GOOS)-amd64/golangci-lint", "/usr/local/bin/golangci-lint"); err != nil {
			return fmt.Errorf("failed to move golangci-lint binary: %w", err)
		}
	}

	// Clean up
	os.Remove(tarFile)
	os.RemoveAll("./golangci-lint-1.52.2-$(runtime.GOOS)-amd64")

	fmt.Println("golangci-lint installed successfully.")
	return nil
}
