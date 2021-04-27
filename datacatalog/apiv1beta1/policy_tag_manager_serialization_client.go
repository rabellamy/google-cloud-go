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

package datacatalog

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newPolicyTagManagerSerializationClientHook clientHook

// PolicyTagManagerSerializationCallOptions contains the retry settings for each method of PolicyTagManagerSerializationClient.
type PolicyTagManagerSerializationCallOptions struct {
	ImportTaxonomies []gax.CallOption
	ExportTaxonomies []gax.CallOption
}

func defaultPolicyTagManagerSerializationClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("datacatalog.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("datacatalog.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://datacatalog.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPolicyTagManagerSerializationCallOptions() *PolicyTagManagerSerializationCallOptions {
	return &PolicyTagManagerSerializationCallOptions{
		ImportTaxonomies: []gax.CallOption{},
		ExportTaxonomies: []gax.CallOption{},
	}
}

// PolicyTagManagerSerializationClient is a client for interacting with Google Cloud Data Catalog API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type PolicyTagManagerSerializationClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	policyTagManagerSerializationClient datacatalogpb.PolicyTagManagerSerializationClient

	// The call options for this service.
	CallOptions *PolicyTagManagerSerializationCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPolicyTagManagerSerializationClient creates a new policy tag manager serialization client.
//
// Policy tag manager serialization API service allows clients to manipulate
// their taxonomies and policy tags data with serialized format.
func NewPolicyTagManagerSerializationClient(ctx context.Context, opts ...option.ClientOption) (*PolicyTagManagerSerializationClient, error) {
	clientOpts := defaultPolicyTagManagerSerializationClientOptions()

	if newPolicyTagManagerSerializationClientHook != nil {
		hookOpts, err := newPolicyTagManagerSerializationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &PolicyTagManagerSerializationClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultPolicyTagManagerSerializationCallOptions(),

		policyTagManagerSerializationClient: datacatalogpb.NewPolicyTagManagerSerializationClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PolicyTagManagerSerializationClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PolicyTagManagerSerializationClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PolicyTagManagerSerializationClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ImportTaxonomies imports all taxonomies and their policy tags to a project as new
// taxonomies.
//
// This method provides a bulk taxonomy / policy tag creation using nested
// proto structure.
func (c *PolicyTagManagerSerializationClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ImportTaxonomies[0:len(c.CallOptions.ImportTaxonomies):len(c.CallOptions.ImportTaxonomies)], opts...)
	var resp *datacatalogpb.ImportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerSerializationClient.ImportTaxonomies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExportTaxonomies exports all taxonomies and their policy tags in a project.
//
// This method generates SerializedTaxonomy protos with nested policy tags
// that can be used as an input for future ImportTaxonomies calls.
func (c *PolicyTagManagerSerializationClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExportTaxonomies[0:len(c.CallOptions.ExportTaxonomies):len(c.CallOptions.ExportTaxonomies)], opts...)
	var resp *datacatalogpb.ExportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyTagManagerSerializationClient.ExportTaxonomies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
