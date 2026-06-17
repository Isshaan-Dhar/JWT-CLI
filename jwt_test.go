package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestJWTCLI(t *testing.T) {
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIiLCJleHAiOjE4OTM0NTYwMDB9.dummy_signature"
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIiLCJleHAiOjE1Nzc4MzY4MDB9.dummy_signature"
	malformedToken := "not.a.real.token.at.all.lololololol"

	tests := []struct {
		name           string
		token          string
		expectedOutput string
	}{
		{"ValidToken", validToken, "TOKEN VALID"},
		{"ExpiredToken", expiredToken, "TOKEN EXPIRED"},
		{"MalformedToken", malformedToken, "INVALID JWT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", tt.token)

			outputBytes, _ := cmd.CombinedOutput()
			output := string(outputBytes)

			if !strings.Contains(strings.ToLower(output), strings.ToLower(tt.expectedOutput)) {
				t.Errorf("test '%s' failed.\nExpected output to contain: %q\nGot:\n%s", tt.name, tt.expectedOutput, output)
			}
		})
	}
}
