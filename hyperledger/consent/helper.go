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

package consent

import (
	"github.com/pascallimeux/ocms/hyperledger"
	"github.com/pascallimeux/ocms/utils/log"
)

type Consent_Helper struct {
	HP_helper hyperledger.HP_Helper
}

func (c *Consent_Helper) CreateConsent() {
	log.Trace(log.Here(), "CreateConsent() : calling method -")
}

func (c *Consent_Helper) GetConsent() {
	log.Trace(log.Here(), "GetConsent() : calling method -")
}

func (c *Consent_Helper) GetAllConsents() {
	log.Trace(log.Here(), "GetAllConsents() : calling method -")
}

func (c *Consent_Helper) GetTRConsent() {
	log.Trace(log.Here(), "GetTRConsent() : calling method -")
}
