package scanner

import (
	"testing"

	"github.com/nvkp/turtle/assert"
)

var sanitizeTestCases = map[string]struct {
	base     string
	input    string
	token    string
	label    string
	datatype string
	typ      string
}{
	"with-label": {
		input: `"this is an English text"@en`,
		token: `this is an English text`,
		label: `en`,
		typ:   "literal",
	},
	"with-datatype": {
		input:    `"this is an English text"^^xsd:string`,
		token:    `this is an English text`,
		datatype: `xsd:string`,
		typ:      "literal",
	},
	"at-in-literal": {
		input:    `"my email is x@y.com"^^xsd:string`,
		token:    `my email is x@y.com`,
		datatype: `xsd:string`,
		typ:      "literal",
	},
	"multiline-literal-datatype": {
		input:    `"""Note that SI supports only the use of symbols and deprecates the use of any abbreviations for units."""^^qudt:LatexString`,
		token:    `Note that SI supports only the use of symbols and deprecates the use of any abbreviations for units.`,
		datatype: `qudt:LatexString`,
		typ:      "literal",
	},
	"multiline-literal-label": {
		input: `"""Ostrouhej čtyři sta brambor, dokud můžeš ostrou škrabkou!"""@cs`,
		token: `Ostrouhej čtyři sta brambor, dokud můžeš ostrou škrabkou!`,
		label: `cs`,
		typ:   "literal",
	},
	"iri": {
		base:  "http://example.org/",
		input: "</path>",
		token: "http://example.org/path",
		typ:   "iri",
	},
	"escaped-quotes": {
		input: `"a string with escaped \"double quotes\""`,
		token: `a string with escaped \"double quotes\"`,
		typ:   "literal",
	},
	"single-quote": {
		input: `"someone's single quote"`,
		token: `someone's single quote`,
		typ:   "literal",
	},
	"empty-typed-literal": {
		input:    `""^^qudt:LatexString`,
		token:    ``,
		datatype: `qudt:LatexString`,
		typ:      "literal",
	},
	"bare-integer": {
		input:    `42`,
		token:    `42`,
		datatype: `<http://www.w3.org/2001/XMLSchema#integer>`,
		typ:      "literal",
	},
	"bare-decimal": {
		input:    `-48.200`,
		token:    `-48.200`,
		datatype: `<http://www.w3.org/2001/XMLSchema#decimal>`,
		typ:      "literal",
	},
	"bare-double": {
		input:    `1.0E0`,
		token:    `1.0E0`,
		datatype: `<http://www.w3.org/2001/XMLSchema#double>`,
		typ:      "literal",
	},
	"bare-boolean": {
		input:    `false`,
		token:    `false`,
		datatype: `<http://www.w3.org/2001/XMLSchema#boolean>`,
		typ:      "literal",
	},
}

func TestSanitize(t *testing.T) {
	for name, tc := range sanitizeTestCases {
		t.Run(name, func(t *testing.T) {
			s := &Scanner{
				base: tc.base,
			}
			token, label, datatype, typ := s.sanitize(tc.input)
			assert.Equal(t, tc.token, token, "function should have returned correctly sanitized token")
			assert.Equal(t, tc.label, label, "function should have returned correctly extracted label")
			assert.Equal(t, tc.datatype, datatype, "function should have returned correctly extracted datatype")
			assert.Equal(t, tc.typ, typ, "function should have returned the correct object type")
		})
	}
}
