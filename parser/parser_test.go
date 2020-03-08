package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Incorrect number of statements. Received %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected 'let' got %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s isn't a LetStatement. Received: %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Let statement name not '%s'. Received '%s'", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("LetStatement Name TokenLiteral() not expected value %s. Received %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("Program should have 3 statements")
	}

	for _, stmt := range program.Statements {
		returnStatement, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("Statement not of ReturnStatement type. Got %T", stmt)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("Return token literal type is not return. Instead got %q",
				returnStatement.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Error: %q", msg)
	}
	t.FailNow()
}

