package zinterpreter

import (
	"testing"
)

func TestReadSchema(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{input: `definition monsujet { }`, expectError: false},
		{input: `definition monsujet { relation marelation1: monsujet2 }`, expectError: false},
		{input: `definition monsujet { relation marelation1: monsujet2 | monsujet3 }`, expectError: false},
		{input: `definition monsujet { relation marelation1: monsujet2 | monsujet3  relation marelation2: monsujet2 }`, expectError: false},
		{input: `definition monsujet { } definition monsujet2 { } definition maressource { relation marelation: monsujet | monsujet2   }`, expectError: false},
		{input: `definition monsujet { } definition monsujet2 { } definition maressource { relation marelation: monsujet | monsujet2  relation marelation2: monsujet | monsujet2 | monsujet3  }`, expectError: false},
		{input: `definition monsujet {`, expectError: true}, // syntax error
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		lexer.NextToken()
		_, err := lexer.ReadZSchema()

		if tt.expectError && err == nil {
			t.Errorf("expected an error but got none for input: %s", tt.input)
		}
		if !tt.expectError && err != nil {
			t.Errorf("did not expect an error but got one for input: %s, error: %v", tt.input, err)
		}

	}
}
