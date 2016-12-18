/*
Copyright Pascal Limeux. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hyperledger

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"github.com/pascallimeux/ocms/utils"
	"github.com/pascallimeux/ocms/utils/log"
	"io/ioutil"
	"net/http"
	"time"
)

var client http.Client = http.Client{Timeout: time.Duration(10 * time.Second)}

type HP_Helper struct {
	HttpHyperledger string
}

const (
	JSONRPC     = "2.0"
	CHAINCODE   = "/chaincode"
	REGISTAR    = "/registrar"
	TRANSACTION = "/transactions"
	CONTENTTYPE = "application/json"
)

/*
func Display_json(jsonbytes []byte) {
	var out bytes.Buffer
	json.Indent(&out, jsonbytes, "", "  ")
	fmt.Println("Json object: ", out.String())
}*/

func (h *HP_Helper) Registar(enrollId, enrollSecret string) (SimpleResponse, error) {
	log.Trace(log.Here(), "Registar() : calling method -")
	response := SimpleResponse{}
	url := h.HttpHyperledger + REGISTAR
	log.Trace(log.Here(), "URL: ", url)
	contentBytes, err1 := Build_registar_body(enrollId, enrollSecret)
	if err1 != nil {
		return response, err1
	}
	log.Trace(log.Here(), "BODY: ", string(contentBytes))
	resp, err2 := client.Post(url, CONTENTTYPE, bytes.NewBuffer(contentBytes))
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func (h *HP_Helper) IsRegistar(enrollId string) (SimpleResponse, error) {
	log.Trace(log.Here(), "Registar() : calling method -")
	response := SimpleResponse{}
	url := h.HttpHyperledger + REGISTAR + "/" + enrollId
	log.Trace(log.Here(), "URL: ", url)
	resp, err2 := client.Get(url)
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func (h *HP_Helper) DeployChainCode(smartcontract_path, hp_account, function string, args []string) (Response, error) {
	log.Trace(log.Here(), "deployChainCode() : calling method -")
	response := Response{}
	url := h.HttpHyperledger + CHAINCODE
	log.Trace(log.Here(), "URL: ", url)
	contentBytes, err1 := Build_deploy_body(smartcontract_path, hp_account, function, args)
	if err1 != nil {
		return response, err1
	}
	log.Trace(log.Here(), "BODY: ", string(contentBytes))
	resp, err2 := client.Post(url, CONTENTTYPE, bytes.NewBuffer(contentBytes))
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func (h *HP_Helper) Invoke(chaincode_name, hp_account, function string, args []string) (Response, error) {
	log.Trace(log.Here(), "Invoke() : calling method -")
	response := Response{}
	url := h.HttpHyperledger + CHAINCODE
	log.Trace(log.Here(), "URL: ", url)
	contentBytes, err1 := Build_invoke_body(chaincode_name, hp_account, function, args)
	if err1 != nil {
		return response, err1
	}
	log.Trace(log.Here(), "BODY: ", string(contentBytes))
	resp, err2 := client.Post(url, CONTENTTYPE, bytes.NewBuffer(contentBytes))
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func (h *HP_Helper) Query(chaincode_name, hp_account, function string, args []string) (Response, error) {
	log.Trace(log.Here(), "Query() : calling method -")
	response := Response{}
	url := h.HttpHyperledger + CHAINCODE
	log.Trace(log.Here(), "URL: ", url)
	contentBytes, err1 := Build_query_body(chaincode_name, hp_account, function, args)
	if err1 != nil {
		return response, err1
	}
	log.Trace(log.Here(), "BODY: ", string(contentBytes))
	resp, err2 := client.Post(url, CONTENTTYPE, bytes.NewBuffer(contentBytes))
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func (h *HP_Helper) Transaction(transaction_uuid string) (Transaction, error) {
	log.Trace(log.Here(), "Transation() : calling method -")
	response := Transaction{}
	url := h.HttpHyperledger + TRANSACTION + "/" + transaction_uuid
	log.Trace(log.Here(), "URL: ", url)
	resp, err2 := client.Get(url)
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	err3 := BuildResponse(&response, resp)
	return response, err3
}

func BuildResponse(response interface{}, resp *http.Response) error {
	log.Trace(log.Here(), "BuildResponse() : calling method -")
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Trace(log.Here(), "RAW RESPONSE: ", string(bytes))
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return err
	}
	responseToString, err := utils.StructToString(response)
	if err == nil {
		log.Trace(log.Here(), "RESPONSE: ", responseToString)
	}
	return nil
}
