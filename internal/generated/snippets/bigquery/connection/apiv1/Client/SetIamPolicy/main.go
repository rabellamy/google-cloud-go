// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START bigqueryconnection_generated_bigquery_connection_apiv1_Client_SetIamPolicy]

package main

import (
	"context"

	connection "cloud.google.com/go/bigquery/connection/apiv1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func main() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := connection.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END bigqueryconnection_generated_bigquery_connection_apiv1_Client_SetIamPolicy]
