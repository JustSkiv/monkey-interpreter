package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)

	// FEATURE: use rune to work with Unicode
	ch byte // current char under examination

}

// FEATURE: init lexer with io.Reader and filename
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// TODO: use constructor to create tokens (implement it in token package)
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// TODO: use newToken
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// TODO: use newToken
			t = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '/':
		t = newToken(token.SLASH, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case 0:
		// TODO: use general func too
		t.Literal = ""
		t.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			t.Literal = l.readIdentifier()
			t.Type = token.LookupType(t.Literal)
			return t
		case isDigit(l.ch):
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		default:
			t = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()

	return t
}

func (l *Lexer) readChar() {

	l.ch = l.peekChar()

	l.position = l.readPosition

	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// TODO: name
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
