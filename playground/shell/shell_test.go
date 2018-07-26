package main

import (
	"testing"

	"os"

)

func TestExecInput(t *testing.T){
	tests := []struct {
		name string
		givenInput string
		expectedDir string //only test "cd"
		expectedErr error
	}{
		{
			name: "test1 ls",
			givenInput: "ls",
			expectedErr: nil,
		},
		{
			name: "test2 ls -l -a",
			givenInput: "ls -l -a",
			expectedErr: nil,

		},
		{
			name: "test3 cd",
			givenInput: "cd",
			expectedErr: nil,
		},
		{
			name: "test4 cd /",
			givenInput: "cd /",
			expectedDir: "/",
			expectedErr: nil,


		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// test command
			if err := execInput(tt.givenInput); err != tt.expectedErr {
				t.Errorf("execInput() error = %v, wantErr %v", err, tt.expectedErr)
			}

			// test for changed directory only
			if tt.expectedDir != "" {
				curDir, err := os.Getwd()
				if err != nil {
					t.Errorf("Failed to get new directory: %v", err)
				}
				if tt.expectedDir != curDir {
					t.Errorf("Failed to change to desired directory.")
				}
			}
		})
	}

}
