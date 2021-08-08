package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func abrirTicket(justificativa string, autoclose bool) string {
	token := tokenGLPI()
	userId := getUser(token)
	chamado := novoTicketGLPI(usuario, justificativa, userId, autoclose)
	return chamado
}

func tokenGLPI() string {

	url := configs["glpi_url"] + "initSession"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", configs["glpi_integracao_authorization"])
	req.Header.Add("App-Token", configs["glpi_integracao_app-token"])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)
	token := result["session_token"]
	return token.(string)
}

func novoTicketGLPI(usuario string, justificativa string, iduser string, autoclose bool) string {
	var codigoTicket string
	if autoclose == true {
		codigoTicket = configs["glpi_categoryid_autoapproved"]
	} else {
		codigoTicket = configs["glpi_categoryid_toapprove"]
	}
	listaString := strings.Join(lista, "\\n")
	payload := strings.NewReader(`{"input": {"entities_id": "TI","name": "Liberação de acesso por autoatendimento - ` + usuario + `","content": "` + justificativa + `\n\nTabelas solicitadas liberação:\n` + listaString + `","itilcategories_id": ` + codigoTicket + `,"type": 2,"status": 6, "requesttypes_id": ` + configs["glpi_requesttypeid"] + `,"users_id_recipient": ` + iduser + `}}`)
	fmt.Println(payload)
	url := configs["glpi_url"] + "Ticket/"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "Erro"
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Session-Token", tokenGLPI())
	req.Header.Add("App-Token", configs["glpi_integracao_app-token"])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "Erro"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "Erro"
	}
	fmt.Println(string(body))
	var result3 map[string]interface{}

	json.Unmarshal([]byte(body), &result3)
	novoChamado := result3["id"]
	return fmt.Sprint(novoChamado)
}

type dataPacote struct {
	Id int `json:"2"`
}

type userPacote struct {
	Data []dataPacote `json:"data"`
}

func getUser(token string) string {

	url := configs["glpi_url"] + "search/User?criteria%5B0%5D%5Bfield%5D=1&criteria%5B0%5D%5Bsearchtype%5D=contains&criteria%5B0%5D%5Bvalue%5D=" + usuario + "&forcedisplay%5B0%5D=2"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "Erro"
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Session-Token", token)
	req.Header.Add("App-Token", configs["glpi_integracao_app-token"])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "Erro"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "Erro"
	}

	result2 := userPacote{}
	json.Unmarshal([]byte(body), &result2)
	return strconv.Itoa(result2.Data[0].Id)
}
