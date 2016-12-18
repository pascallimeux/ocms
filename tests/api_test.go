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

package tests

import (
	"bytes"
	"encoding/json"
	"github.com/pascallimeux/ocms/api"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestCreateConsentFromAPINominal(t *testing.T) {
	ownerID := "1111"
	consumerID := "2222"
	dt_begin := time.Now().Format("2006-01-02")
	//dt_end := dt_begin.Add(time.Hour * 24 )
	consent := api.Consent{Action: "create", Appid: AppContext.Configuration.ApplicationID, Ownerid: ownerID, Consumerid: consumerID, Dt_begin: dt_begin}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(consent)

	res, err := http.Post(httpServerTest.URL+api.CONSENTAPI, "application/json", b)

	rec_bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	body := string(rec_bytes)
	if res.StatusCode != http.StatusOK {
		t.Fatal("Non-expected status code: %v\n\tbody: %v, data:%s\n", http.StatusCreated, res.StatusCode, body)
	}
}
