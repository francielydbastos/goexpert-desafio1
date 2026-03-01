package entities

type Cotacao struct {
	USDBRL struct {
		Code            string `json:"code"`
		Codein          string `json:"codein"`
		Nome            string `json:"name"`
		Maximo          string `json:"high"`
		Minimo          string `json:"low"`
		Variacao        string `json:"varBid"`
		PorcentVariacao string `json:"pctChange"`
		Compra          string `json:"bid"`
		Venda           string `json:"ask"`
		Timestamp       string `json:"timestamp"`
		DataCriacao     string `json:"create_date"`
	} `json:"USDBRL"`
}
