module exemplo.com/externos

go 1.21.7

require github.com/natanfeitosa/portuscript v0.5.0

require github.com/rivo/uniseg v0.4.4 // indirect

// Apenas faz com que o Go nao precise baixar as dependecias, ele usa do nosso módulo principal
replace github.com/natanfeitosa/portuscript => ../..