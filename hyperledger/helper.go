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
	"fmt"
	"github.com/pascallimeux/ocms/utils"
	"github.com/pascallimeux/ocms/utils/log"
	"io/ioutil"
	"net/http"
)

type Result struct {
	status  string
	message string
}

type Response struct {
	Jsonrpc string
	Result  Result
	Id      int
}

type HP_Helper struct {
	HttpHyperledger string
	ChainCodePath   string
	EnrollID        string
	EnrollSecret    string
}

type Invoke struct {
	Jsonrpc string
	Method  string
	Params  Params2
	Id      int
}

type Query struct {
	Jsonrpc string
	Method  string
	Params  Params2
	Id      int
}

type Deploy struct {
	Jsonrpc string
	Method  string
	Params  Params
	Id      int
}

type ChaincodeID struct {
	Path string
}

type ChaincodeID2 struct {
	Name string
}

type CtorMsg struct {
	Function string
	Args     []string
}

type Params struct {
	Type          int
	ChaincodeID   ChaincodeID
	SecureContext string
	CtorMsg       CtorMsg
}

type Params2 struct {
	Type          int
	ChaincodeID   ChaincodeID2
	SecureContext string
	CtorMsg       CtorMsg
}

const (
	JSONRPC     = "2.0"
	SC_PATH     = "github.com/pascallimeux/blockchain/consent"
	HP_ACCOUNT  = "pascal"
	CHAINCODE   = "/chaincode"
	CONTENTTYPE = "application/json"
)

func Build_query_body(chaincode_name, hp_account, function string, args []string) ([]byte, error) {
	query := &Query{Jsonrpc: JSONRPC, Method: "query", Id: 1, Params: Params2{Type: 1, ChaincodeID: ChaincodeID2{Name: chaincode_name}, SecureContext: hp_account, CtorMsg: CtorMsg{Function: function, Args: args}}}
	bytes, err := json.Marshal(query)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

func Build_invoke_body(chaincode_name, hp_account, function string, args []string) ([]byte, error) {
	invoke := &Invoke{Jsonrpc: JSONRPC, Method: "invoke", Id: 1, Params: Params2{Type: 1, ChaincodeID: ChaincodeID2{Name: chaincode_name}, SecureContext: hp_account, CtorMsg: CtorMsg{Function: function, Args: args}}}
	bytes, err := json.Marshal(invoke)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

func Build_deploy_body(smartcontract_path, hp_account, function string, args []string) ([]byte, error) {
	deploy := &Deploy{Jsonrpc: JSONRPC, Method: "deploy", Id: 1, Params: Params{Type: 1, ChaincodeID: ChaincodeID{Path: smartcontract_path}, SecureContext: hp_account, CtorMsg: CtorMsg{Function: function, Args: args}}}
	bytes, err := json.Marshal(deploy)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

func Display_json(jsonbytes []byte) {
	var out bytes.Buffer
	json.Indent(&out, jsonbytes, "", "  ")
	fmt.Println("Json object: ", out.String())
}

func (h *HP_Helper) DeployChainCode(smartcontract_path, function string, args []string) (Response, error) {
	log.Trace(log.Here(), "deployChainCode() : calling method -")
	response := Response{}
	url := h.HttpHyperledger + CHAINCODE
	log.Trace(log.Here(), "URL: ", url)
	contentBytes, err1 := Build_deploy_body(SC_PATH, HP_ACCOUNT, function, args)
	if err1 != nil {
		return response, err1
	}
	log.Trace(log.Here(), "BODY: ", string(contentBytes))
	resp, err2 := http.Post(url, CONTENTTYPE, bytes.NewBuffer(contentBytes))
	if err2 != nil {
		return response, err2
	}
	defer resp.Body.Close()
	bytes, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return response, err3
	}
	err4 := json.Unmarshal(bytes, &response)
	if err4 != nil {
		return response, err4
	}
	responseToString, err5 := utils.StructToString(response)
	if err5 == nil {
		log.Trace(log.Here(), "RESPONSE: ", responseToString)
	}
	return response, nil
}

func (h *HP_Helper) invoke() {
	log.Trace(log.Here(), "invoke() : calling method -")
}

func (h *HP_Helper) query() {
	log.Trace(log.Here(), "query() : calling method -")
}

func main() {

	chaincodeName := "1234567890"
	emptyargs := make([]string, 0)
	args := []string{"000", "111", "222", "BP", "R", "2016-09-27", "2017-10-19"}
	deploy, _ := Build_deploy_body(SC_PATH, HP_ACCOUNT, "init", emptyargs)
	invoke, _ := Build_invoke_body(chaincodeName, HP_ACCOUNT, "PostConsent", args)
	query, _ := Build_query_body(chaincodeName, HP_ACCOUNT, "GetVersion", emptyargs)
	Display_json(deploy)
	Display_json(invoke)
	Display_json(query)
}
