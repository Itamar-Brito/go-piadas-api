package piadas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Piada struct {
	ID     string `json:"id"`
	Titulo string `json:"titulo"`
	Piada  string `json:"piada"`
}

func Teste() string {
	return "testando package"
}

func FetchPiadas()([]Piada, error){
	resp, err := http.Get("http://piada.atwebpages.com/php_action/api/get-piadas.php")
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a solicitação: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
	}

	var piadas []Piada
	err = json.Unmarshal(body, &piadas)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar o JSON: %v", err)
	}

	return piadas, err

}