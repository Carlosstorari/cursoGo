package samples

//paramentro variadico x ...int deve ser sempre o ultimo paramentro
func SomaMultiplosRetornos(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia de soma!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde de soma!"
	} else {
		oi = "Oi, boa noite de soma!"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi
}
