// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package phishingprotection_test

import (
	"context"

	phishingprotection "cloud.google.com/go/phishingprotection/apiv1beta1"
	phishingprotectionpb "google.golang.org/genproto/googleapis/cloud/phishingprotection/v1beta1"
)

func ExampleNewPhishingProtectionServiceV1Beta1Client() {
	ctx := context.Background()
	c, err := phishingprotection.NewPhishingProtectionServiceV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExamplePhishingProtectionServiceV1Beta1Client_ReportPhishing() {
	// import phishingprotectionpb "google.golang.org/genproto/googleapis/cloud/phishingprotection/v1beta1"

	ctx := context.Background()
	c, err := phishingprotection.NewPhishingProtectionServiceV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &phishingprotectionpb.ReportPhishingRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportPhishing(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
