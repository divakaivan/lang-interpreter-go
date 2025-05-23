package lexer

import "divakaivan/lang-interpreter-go/token"

type Lexer struct {
	input        string
	position     int  // cur pos in input (points to cur char)
	readPosition int  // cur reading pos in input (after cur char)
	ch           byte // cur char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
* Give us the next char and advance our pos in the input str
*
 */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		// 0 is the ASCII char for NUL and signifies
		// either we havent read anything yet or eof
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

/*
* look at cur char under examination
* return a token depending on which char it is
* before return, advance pointers into the input
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
