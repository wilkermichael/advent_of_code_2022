#!/bin/bash

# Check if a project name was provided
if [ -z "$1" ]; then
  echo "Error: missing project name"
  exit 1
fi

# Check if a module name was provided
if [ -z "$2" ]; then
  echo "Error: missing module name"
  exit 1
fi

# Create a new directory for the project
mkdir "$1"

# Create main.go and add a main function that prints "hello world"
cat > "$1/main.go" <<EOF
package main

import "fmt"

func main() {
  fmt.Println("hello world")
}
EOF

# Initialize go mod with the provided project name
cd "$1"
go mod init "$2"

# Install the testify package for testing
go get github.com/stretchr/testify

# Create a sample_package directory and add sample_package.go and sample_package_test.go
mkdir sample_package

cat > sample_package/sample_package.go <<EOF
package sample_package

func SampleFunction() string {
  return "hello world"
}
EOF

cat > sample_package/sample_package_test.go <<EOF
package sample_package

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestSampleFunction(t *testing.T) {
  tests := []struct {
    name string
    expected string
  }{
    {
      name: "Test Case 1",
      expected: "hello world",
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      assert.Equal(t, test.expected, SampleFunction())
    })
  }
}
EOF
