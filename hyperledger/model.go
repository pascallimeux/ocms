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

type Result struct {
	status  string
	message string
}

type Response struct {
	Jsonrpc string
	Result  Result
	Id      int
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
