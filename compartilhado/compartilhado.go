package compartilhado

func ContemApenasAlfaNum(str string) bool {
	return ContemApenasDigitos(str) || ContemApenasLetras(str)
}