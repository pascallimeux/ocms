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
	"encoding/json"
	"github.com/pascallimeux/ocms/hyperledger"
	"github.com/pascallimeux/ocms/utils/log"
	"net/http"
)

//HTTP Post - /ocms/v1/api/consent
func (a *AppContext) processConsent(w http.ResponseWriter, r *http.Request) {
	log.Trace(log.Here(), "processConsent() : calling method -")
	var bytes []byte
	var consent Consent
	err := json.NewDecoder(r.Body).Decode(&consent)
	if err != nil {
		sendError(log.Here(), w, err)
		return
	}
	switch action := consent.Action; action {
	case "create":
		bytes, err = a.createConsent(consent)
	case "list":
		bytes, err = a.listConsents(consent.Appid)
	case "get":
		bytes, err = a.getConsent(consent.Appid, consent.Consentid)
	case "remove":
		bytes, err = a.unactivateConsent(consent.Appid, consent.Consentid)
	case "list4owner":
		bytes, err = a.getConsents4Consumer(consent.Appid, consent.Ownerid)
	case "list4consumer":
		bytes, err = a.getConsents4Owner(consent.Appid, consent.Consumerid)
	case "isconsent":
		bytes, err = a.isConsent(consent)
	default:
		log.Error(log.Here(), "bad action request")
		sendError(log.Here(), w, err)
	}
	if err != nil {
		sendError(log.Here(), w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

//HTTP Get - /ocms/v1/api/hyperledger/consenttr
func (a *AppContext) processConsentTR(w http.ResponseWriter, r *http.Request) {
	log.Trace(log.Here(), "processConsentTR() : calling method -")
}

func (a *AppContext) createConsent(consent Consent) ([]byte, error) {
	log.Trace(log.Here(), "createConsent() : calling method -")
	_, err := a.Consent_helper.CreateConsent(a.Configuration.ApplicationID, consent.Ownerid, consent.Consumerid, consent.Datatype, consent.Dataaccess, consent.Dt_begin, consent.Dt_end)
	if err != nil {
		return nil, err
	}
	return consent2Bytes(consent)
}

func (a *AppContext) listConsents(applicationID string) ([]byte, error) {
	log.Trace(log.Here(), "listConsents() : calling method -")
	consents, err := a.Consent_helper.GetACtivesConsents(a.Configuration.ApplicationID)
	if err != nil {
		return nil, err
	}
	return HPconsents2ConsentsBytes(consents)
}

func (a *AppContext) getConsent(applicationID, consentID string) ([]byte, error) {
	log.Trace(log.Here(), "getConsent() : calling method -")
	consent, err := a.Consent_helper.GetConsent(a.Configuration.ApplicationID, consentID)
	if err != nil {
		return nil, err
	}
	return HPconsent2ConsentBytes(consent)
}

func (a *AppContext) unactivateConsent(applicationID, consentID string) ([]byte, error) {
	log.Trace(log.Here(), "unactivateConsent() : calling method -")
	_, err := a.Consent_helper.UnactivateConsent(a.Configuration.ApplicationID, consentID)
	if err != nil {
		return nil, err
	}
	consent, err := a.Consent_helper.GetConsent(a.Configuration.ApplicationID, consentID)
	if err != nil {
		return nil, err
	}
	return HPconsent2ConsentBytes(consent)
}

func (a *AppContext) getConsents4Consumer(applicationID, consumerID string) ([]byte, error) {
	log.Trace(log.Here(), "getConsents4Consumer() : calling method -")
	consents, err := a.Consent_helper.GetConsents4Consumer(a.Configuration.ApplicationID, consumerID)
	if err != nil {
		return nil, err
	}
	return HPconsents2ConsentsBytes(consents)
}

func (a *AppContext) getConsents4Owner(applicationID, ownerID string) ([]byte, error) {
	log.Trace(log.Here(), "getConsents4Owner() : calling method -")
	consents, err := a.Consent_helper.GetConsents4Owner(a.Configuration.ApplicationID, ownerID)
	if err != nil {
		return nil, err
	}
	return HPconsents2ConsentsBytes(consents)
}

func (a *AppContext) isConsent(consent Consent) ([]byte, error) {
	log.Trace(log.Here(), "isConsent() : calling method -")
	isconsent, err := a.Consent_helper.IsConsent(a.Configuration.ApplicationID, consent.Ownerid, consent.Consumerid, consent.Datatype, consent.Dataaccess)
	if err != nil {
		return nil, err
	}
	if isconsent {
		return []byte("True"), nil
	} else {
		return []byte("False"), nil
	}
}

func convertHPConsents2APIConsents(HPconsents []hyperledger.Consent) []Consent {
	var consents []Consent
	for i, HPconsent := range HPconsents {
		consents[i] = convertHPConsent2APIConsent(HPconsent)
	}
	return consents
}

func convertHPConsent2APIConsent(HPconsent hyperledger.Consent) Consent {
	var consent Consent
	consent.Consentid = HPconsent.ConsentID
	consent.Ownerid = HPconsent.OwnerID
	consent.Consumerid = HPconsent.ConsumerID
	consent.Dataaccess = HPconsent.Dataaccess
	consent.Datatype = HPconsent.Datatype
	consent.Dt_begin = HPconsent.Dt_begin
	consent.Dt_end = HPconsent.Dt_end
	return consent
}

func HPconsents2ConsentsBytes(HPconsents []hyperledger.Consent) ([]byte, error) {
	consents := convertHPConsents2APIConsents(HPconsents)
	j, err := json.Marshal(consents)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func HPconsent2ConsentBytes(HPconsent hyperledger.Consent) ([]byte, error) {
	consent := convertHPConsent2APIConsent(HPconsent)
	return consent2Bytes(consent)
}

func consent2Bytes(consent Consent) ([]byte, error) {
	j, err := json.Marshal(consent)
	if err != nil {
		return nil, err
	}
	return j, nil
}
