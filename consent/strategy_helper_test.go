/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package consent_test

import (
	"net/http"
	"testing"

	"github.com/ory/x/pointerx"
	"github.com/tyaps/hydra/sdk/go/hydra/client/admin"
	"github.com/tyaps/hydra/sdk/go/hydra/models"

	"github.com/stretchr/testify/require"

	"github.com/tyaps/hydra/sdk/go/hydra/client"
)

var passAuthentication = func(apiClient *client.OryHydra, remember bool) func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			v, err := apiClient.Admin.AcceptLoginRequest(admin.NewAcceptLoginRequestParams().WithLoginChallenge(r.URL.Query().Get("login_challenge")).WithBody(&models.HandledLoginRequest{
				Subject:     pointerx.String("user"),
				Remember:    remember,
				RememberFor: 0,
				ACR:         "1",
			}))
			require.NoError(t, err)
			require.NotEmpty(t, v.Payload.RedirectTo)
			http.Redirect(w, r, v.Payload.RedirectTo, http.StatusFound)
		}
	}
}

var passAuthorization = func(apiClient *client.OryHydra, remember bool) func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			v, err := apiClient.Admin.AcceptConsentRequest(admin.NewAcceptConsentRequestParams().WithConsentChallenge(r.URL.Query().Get("consent_challenge")).WithBody(&models.HandledConsentRequest{
				GrantedScope: []string{"scope-a"},
				Remember:     remember,
				RememberFor:  0,
				Session: &models.ConsentRequestSessionData{
					AccessToken: map[string]interface{}{"foo": "bar"},
					IDToken:     map[string]interface{}{"bar": "baz"},
				},
			}))
			require.NoError(t, err)
			require.NotEmpty(t, v.Payload.RedirectTo)
			http.Redirect(w, r, v.Payload.RedirectTo, http.StatusFound)
		}
	}
}
