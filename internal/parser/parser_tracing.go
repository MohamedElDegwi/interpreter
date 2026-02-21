package parser

import (
	"fmt"
	"strings"
)

const traceIdentPlaceholder string = "\t"

func (p *Parser) identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, p.traceLevel-1)
}

func (p *Parser) tracePrint(fs string) {
	msg := fmt.Sprintf("%s%s\n", p.identLevel(), fs)

	p.traces = append(p.traces, msg)
}

func (p *Parser) incIdent() { p.traceLevel = p.traceLevel + 1 }
func (p *Parser) decIdent() { p.traceLevel = p.traceLevel - 1 }

func (p *Parser) trace(msg string) string {
	p.incIdent()
	p.tracePrint("BEGIN " + msg)
	return msg
}

func (p *Parser) untrace(msg string) {
	p.tracePrint("END " + msg)
	p.decIdent()
}
