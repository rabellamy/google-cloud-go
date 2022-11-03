// Copyright 2022 Google LLC
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

package dataproc

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	dataprocpb "cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newBatchControllerClientHook clientHook

// BatchControllerCallOptions contains the retry settings for each method of BatchControllerClient.
type BatchControllerCallOptions struct {
	CreateBatch []gax.CallOption
	GetBatch    []gax.CallOption
	ListBatches []gax.CallOption
	DeleteBatch []gax.CallOption
}

func defaultBatchControllerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataproc.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataproc.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataproc.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBatchControllerCallOptions() *BatchControllerCallOptions {
	return &BatchControllerCallOptions{
		CreateBatch: []gax.CallOption{},
		GetBatch:    []gax.CallOption{},
		ListBatches: []gax.CallOption{},
		DeleteBatch: []gax.CallOption{},
	}
}

// internalBatchControllerClient is an interface that defines the methods available from Cloud Dataproc API.
type internalBatchControllerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateBatch(context.Context, *dataprocpb.CreateBatchRequest, ...gax.CallOption) (*CreateBatchOperation, error)
	CreateBatchOperation(name string) *CreateBatchOperation
	GetBatch(context.Context, *dataprocpb.GetBatchRequest, ...gax.CallOption) (*dataprocpb.Batch, error)
	ListBatches(context.Context, *dataprocpb.ListBatchesRequest, ...gax.CallOption) *BatchIterator
	DeleteBatch(context.Context, *dataprocpb.DeleteBatchRequest, ...gax.CallOption) error
}

// BatchControllerClient is a client for interacting with Cloud Dataproc API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The BatchController provides methods to manage batch workloads.
type BatchControllerClient struct {
	// The internal transport-dependent client.
	internalClient internalBatchControllerClient

	// The call options for this service.
	CallOptions *BatchControllerCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BatchControllerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BatchControllerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BatchControllerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateBatch creates a batch workload that executes asynchronously.
func (c *BatchControllerClient) CreateBatch(ctx context.Context, req *dataprocpb.CreateBatchRequest, opts ...gax.CallOption) (*CreateBatchOperation, error) {
	return c.internalClient.CreateBatch(ctx, req, opts...)
}

// CreateBatchOperation returns a new CreateBatchOperation from a given name.
// The name must be that of a previously created CreateBatchOperation, possibly from a different process.
func (c *BatchControllerClient) CreateBatchOperation(name string) *CreateBatchOperation {
	return c.internalClient.CreateBatchOperation(name)
}

// GetBatch gets the batch workload resource representation.
func (c *BatchControllerClient) GetBatch(ctx context.Context, req *dataprocpb.GetBatchRequest, opts ...gax.CallOption) (*dataprocpb.Batch, error) {
	return c.internalClient.GetBatch(ctx, req, opts...)
}

// ListBatches lists batch workloads.
func (c *BatchControllerClient) ListBatches(ctx context.Context, req *dataprocpb.ListBatchesRequest, opts ...gax.CallOption) *BatchIterator {
	return c.internalClient.ListBatches(ctx, req, opts...)
}

// DeleteBatch deletes the batch workload resource. If the batch is not in terminal state,
// the delete fails and the response returns FAILED_PRECONDITION.
func (c *BatchControllerClient) DeleteBatch(ctx context.Context, req *dataprocpb.DeleteBatchRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteBatch(ctx, req, opts...)
}

// batchControllerGRPCClient is a client for interacting with Cloud Dataproc API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type batchControllerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BatchControllerClient
	CallOptions **BatchControllerCallOptions

	// The gRPC API client.
	batchControllerClient dataprocpb.BatchControllerClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBatchControllerClient creates a new batch controller client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The BatchController provides methods to manage batch workloads.
func NewBatchControllerClient(ctx context.Context, opts ...option.ClientOption) (*BatchControllerClient, error) {
	clientOpts := defaultBatchControllerGRPCClientOptions()
	if newBatchControllerClientHook != nil {
		hookOpts, err := newBatchControllerClientHook(ctx, clientHookParams{})
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
	client := BatchControllerClient{CallOptions: defaultBatchControllerCallOptions()}

	c := &batchControllerGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		batchControllerClient: dataprocpb.NewBatchControllerClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *batchControllerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *batchControllerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *batchControllerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *batchControllerGRPCClient) CreateBatch(ctx context.Context, req *dataprocpb.CreateBatchRequest, opts ...gax.CallOption) (*CreateBatchOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateBatch[0:len((*c.CallOptions).CreateBatch):len((*c.CallOptions).CreateBatch)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchControllerClient.CreateBatch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateBatchOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *batchControllerGRPCClient) GetBatch(ctx context.Context, req *dataprocpb.GetBatchRequest, opts ...gax.CallOption) (*dataprocpb.Batch, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBatch[0:len((*c.CallOptions).GetBatch):len((*c.CallOptions).GetBatch)], opts...)
	var resp *dataprocpb.Batch
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.batchControllerClient.GetBatch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *batchControllerGRPCClient) ListBatches(ctx context.Context, req *dataprocpb.ListBatchesRequest, opts ...gax.CallOption) *BatchIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListBatches[0:len((*c.CallOptions).ListBatches):len((*c.CallOptions).ListBatches)], opts...)
	it := &BatchIterator{}
	req = proto.Clone(req).(*dataprocpb.ListBatchesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataprocpb.Batch, string, error) {
		resp := &dataprocpb.ListBatchesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.batchControllerClient.ListBatches(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetBatches(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *batchControllerGRPCClient) DeleteBatch(ctx context.Context, req *dataprocpb.DeleteBatchRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteBatch[0:len((*c.CallOptions).DeleteBatch):len((*c.CallOptions).DeleteBatch)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.batchControllerClient.DeleteBatch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateBatchOperation manages a long-running operation from CreateBatch.
type CreateBatchOperation struct {
	lro *longrunning.Operation
}

// CreateBatchOperation returns a new CreateBatchOperation from a given name.
// The name must be that of a previously created CreateBatchOperation, possibly from a different process.
func (c *batchControllerGRPCClient) CreateBatchOperation(name string) *CreateBatchOperation {
	return &CreateBatchOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateBatchOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Batch, error) {
	var resp dataprocpb.Batch
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateBatchOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Batch, error) {
	var resp dataprocpb.Batch
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateBatchOperation) Metadata() (*dataprocpb.BatchOperationMetadata, error) {
	var meta dataprocpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateBatchOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateBatchOperation) Name() string {
	return op.lro.Name()
}

// BatchIterator manages a stream of *dataprocpb.Batch.
type BatchIterator struct {
	items    []*dataprocpb.Batch
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataprocpb.Batch, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BatchIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BatchIterator) Next() (*dataprocpb.Batch, error) {
	var item *dataprocpb.Batch
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BatchIterator) bufLen() int {
	return len(it.items)
}

func (it *BatchIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
