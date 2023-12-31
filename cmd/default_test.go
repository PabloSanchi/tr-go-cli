package cmd

import (
    "testing"
)

func TestExecuteDefault(t *testing.T) {
    
    testCases := []struct {
        name          string
        args          []string
        input         string
        expectedOutput string
        expectError   bool
    }{
        {
            name:          "Test Case 1",
            args:          []string{"old", "new"},
            input:         "old value",
            expectedOutput: "new vaeue",
            expectError:   false,
        },
		{
            name:          "Test Case 2",
            args:          []string{"old", "new"},
            input:         "",
            expectedOutput: "",
            expectError:   false,
        },
		{
            name:          "Test Case 3",
            args:          []string{"old"},
            input:         "",
            expectedOutput: "",
            expectError:   true,
        },
        {
            name:          "Test Case 4",
            args:          []string{"old", "new"},
            input:         "old value",
            expectedOutput: "new vaeue",
            expectError:   false,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
        
            output, err := ExecuteDefault(tc.input, tc.args)

            if (err != nil) != tc.expectError {
                t.Errorf("Expected error: %v, got: %v", tc.expectError, err)
            }

            if output != tc.expectedOutput {
                t.Errorf("Expected output: %s, got: %s", tc.expectedOutput, output)
            }
        })
    }

}
