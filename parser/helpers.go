package parser

import "encoding/json"

var palavrasChave = map[string]bool{
	"se":       true,
	"senao":    true,
	"enquanto": true,
	"const":    true,
	"var":      true,
	"func":     true,
	"pare":     true,
	"continue": true,
	"para":     true,
	"nova":     true,
	// "Verdadeiro": true,
	// "Falso":      true,
	// "Nulo":       true,
	"assegura": true,
}

func IsKeyword(s string) bool {
	_, ok := palavrasChave[s]
	return ok
}

func Ast2string(ast BaseNode) ([]byte, error) {
	return json.MarshalIndent(ast, "", "    ")
}
