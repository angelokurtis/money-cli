package classification

import "strings"

var (
	services = []string{
		"amazon web services",
		"enhancv",
		"spotify",
		"youtube",
		"linkedin",
		"netflix",
		"jetbrainsam",
		"yazio",
		"clubew",
		"webmotors",
		"lehainam219",
		"endomondo",
		"google cloud",
		"claro",
	}
	home = []string{
		"planetaagua",
		"supermercado boa",
		"bittencourt",
		"comgas",
		"guarapari",
		"scalla",
		"cpfl",
		"net ",
		"pjbank",
		"condominio",
		"savegnago",
		"supermercado arco iris",
		"grisi macedo de almeida",
		"bruna batista do carmo",
	}
	transport = []string{
		"sem parar",
		"uber",
		"auto el parana",
		"auto posto campineira",
		"auto posto maga",
		"auto posto petropen",
		"forgerini inouye",
		"josuel luiz de lima ci",
		"marimpa",
		"mustang auto posto",
		"oct revenda de combust",
		"posto b express",
		"posto biquinha gasolin",
		"posto colinas",
		"posto jardim alvorada",
		"posto marimpa",
	}
	overtime = []string{
		"1701 banco de horas a 75%",
		"1809 adicional noturno 30%",
		"3038 dsr s/ adicional noturno",
		"3041 dsr s/horas extras bco",
		"3044 diferença dsr horas",
		"3589 débito banco de horas",
		"3604 dif adicional noturno",
	}
	taxes = []string{
		"5250 ir retido 13 sal",
		"5500 ir retido",
		"5560 inss",
		"5580 inss de ferias",
		"7621 ir férias",
		"9960 inss 13 salário",
	}
	otherIncome = []string{
		"participação lucros ou resultados",
		"dif férias dissídio",
		"13 º salário 1 º parcela",
		"13 º salário  2 ª parcela",
		"desc 13 º salário  adiantamento",
		"13 º salário  média",
		"desc média 13 º salário  adiant ",
		"dif de médias 13 º sal",
		"1/3 adic const fer",
		"media férias",
		"férias no mês",
		"média adto 13 º salário",
		"auxilio indicação",
		"líquido férias",
	}
	otherDebts = []string{
		"convenio unimed  agregado",
		"estacionamento conveniado",
		"dif de assistência médica",
	}
	salary = []string{
		"salário base",
		"dif sal dissidio",
		"adiant quinzenal",
	}
)

func Classify(desc string) string {
	desc = strings.ToLower(desc)
	if contains(services, desc) {
		return "Serviços"
	}
	if contains(home, desc) {
		return "Casa"
	}
	if contains(transport, desc) {
		return "Transporte"
	}
	if contains(overtime, desc) {
		return "Hora Extra"
	}
	if contains(taxes, desc) {
		return "Impostos"
	}
	if contains(otherIncome, desc) {
		return "Outras Receitas"
	}
	if contains(otherDebts, desc) {
		return "Outros Descontos"
	}
	if contains(salary, desc) {
		return "Salário"
	}
	return ""
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}
