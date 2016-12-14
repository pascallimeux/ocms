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
	"github.com/pascallimeux/ocms/api"
	"github.com/pascallimeux/ocms/utils"
	"github.com/pascallimeux/ocms/utils/log"
	"gopkg.in/mgo.v2"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup(true)
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var applicationJSON string = "application/json"
var AppContext api.AppContext
var httpServerTest *httptest.Server
var logfile *os.File
var MOCK_HR string

func DropDB(session *mgo.Session, dbname string) {
	err := session.DB(dbname).DropDatabase()
	if err != nil {
		log.Fatal(log.Here(), "error:", err.Error())
	}
}

func setup(isDropDB bool) {

	// Read configuration file
	configuration, err := utils.Readconf("../config/configtest.json")
	if err != nil {
		log.Fatal(log.Here(), "error:", err.Error())
	}

	// Init logger
	logfile = log.Init_log(configuration.LogFileName, configuration.Logger)

	// Init application context
	AppContext = api.AppContext{}

	// Init http server
	router := AppContext.CreateRoutes()
	httpServerTest = httptest.NewServer(router)

}

func shutdown() {
	log.Trace(log.Here(), "End of tests..")
}
