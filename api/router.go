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

package api

import (
	"github.com/gorilla/mux"
	"github.com/pascallimeux/ocms/common"
	"github.com/pascallimeux/ocms/hyperledger/consent"
	"github.com/pascallimeux/ocms/utils/log"
	"net/http"
)

const (
	VERSIONURI   = "/ocms/v1/version"
	CONSENTAPI   = "/ocms/v1/api/consent/"
	CONSENTTRAPI = "/ocms/v1/api/hyperledger/consenttr"
)

type AppContext struct {
	HttpServer     *http.Server
	Consent_helper consent.Consent_Helper
	Configuration  common.Configuration
}

// Initialize API
func (appContext *AppContext) CreateRoutes() *mux.Router {
	log.Trace(log.Here(), "createRoutes() : calling method -")
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc(VERSIONURI, appContext.getVersion).Methods("GET")
	router.HandleFunc(CONSENTAPI, appContext.processConsent).Methods("POST")
	router.HandleFunc(CONSENTTRAPI+"/{truuid}", appContext.processConsentTR).Methods("GET")
	return router
}
