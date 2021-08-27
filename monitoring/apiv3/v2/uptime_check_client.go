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

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newUptimeCheckClientHook clientHook

// UptimeCheckCallOptions contains the retry settings for each method of UptimeCheckClient.
type UptimeCheckCallOptions struct {
	ListUptimeCheckConfigs  []gax.CallOption
	GetUptimeCheckConfig    []gax.CallOption
	CreateUptimeCheckConfig []gax.CallOption
	UpdateUptimeCheckConfig []gax.CallOption
	DeleteUptimeCheckConfig []gax.CallOption
	ListUptimeCheckIps      []gax.CallOption
}

func defaultUptimeCheckGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultUptimeCheckCallOptions() *UptimeCheckCallOptions {
	return &UptimeCheckCallOptions{
		ListUptimeCheckConfigs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetUptimeCheckConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateUptimeCheckConfig: []gax.CallOption{},
		UpdateUptimeCheckConfig: []gax.CallOption{},
		DeleteUptimeCheckConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListUptimeCheckIps: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalUptimeCheckClient is an interface that defines the methods availaible from Cloud Monitoring API.
type internalUptimeCheckClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListUptimeCheckConfigs(context.Context, *monitoringpb.ListUptimeCheckConfigsRequest, ...gax.CallOption) *UptimeCheckConfigIterator
	GetUptimeCheckConfig(context.Context, *monitoringpb.GetUptimeCheckConfigRequest, ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error)
	CreateUptimeCheckConfig(context.Context, *monitoringpb.CreateUptimeCheckConfigRequest, ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error)
	UpdateUptimeCheckConfig(context.Context, *monitoringpb.UpdateUptimeCheckConfigRequest, ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error)
	DeleteUptimeCheckConfig(context.Context, *monitoringpb.DeleteUptimeCheckConfigRequest, ...gax.CallOption) error
	ListUptimeCheckIps(context.Context, *monitoringpb.ListUptimeCheckIpsRequest, ...gax.CallOption) *UptimeCheckIpIterator
}

// UptimeCheckClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The UptimeCheckService API is used to manage (list, create, delete, edit)
// Uptime check configurations in the Stackdriver Monitoring product. An Uptime
// check is a piece of configuration that determines which resources and
// services to monitor for availability. These configurations can also be
// configured interactively by navigating to the [Cloud Console]
// (http://console.cloud.google.com (at http://console.cloud.google.com)), selecting the appropriate project,
// clicking on “Monitoring” on the left-hand side to navigate to Stackdriver,
// and then clicking on “Uptime”.
type UptimeCheckClient struct {
	// The internal transport-dependent client.
	internalClient internalUptimeCheckClient

	// The call options for this service.
	CallOptions *UptimeCheckCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *UptimeCheckClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *UptimeCheckClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *UptimeCheckClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListUptimeCheckConfigs lists the existing valid Uptime check configurations for the project
// (leaving out any invalid configurations).
func (c *UptimeCheckClient) ListUptimeCheckConfigs(ctx context.Context, req *monitoringpb.ListUptimeCheckConfigsRequest, opts ...gax.CallOption) *UptimeCheckConfigIterator {
	return c.internalClient.ListUptimeCheckConfigs(ctx, req, opts...)
}

// GetUptimeCheckConfig gets a single Uptime check configuration.
func (c *UptimeCheckClient) GetUptimeCheckConfig(ctx context.Context, req *monitoringpb.GetUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	return c.internalClient.GetUptimeCheckConfig(ctx, req, opts...)
}

// CreateUptimeCheckConfig creates a new Uptime check configuration.
func (c *UptimeCheckClient) CreateUptimeCheckConfig(ctx context.Context, req *monitoringpb.CreateUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	return c.internalClient.CreateUptimeCheckConfig(ctx, req, opts...)
}

// UpdateUptimeCheckConfig updates an Uptime check configuration. You can either replace the entire
// configuration with a new one or replace only certain fields in the current
// configuration by specifying the fields to be updated via updateMask.
// Returns the updated configuration.
func (c *UptimeCheckClient) UpdateUptimeCheckConfig(ctx context.Context, req *monitoringpb.UpdateUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	return c.internalClient.UpdateUptimeCheckConfig(ctx, req, opts...)
}

// DeleteUptimeCheckConfig deletes an Uptime check configuration. Note that this method will fail
// if the Uptime check configuration is referenced by an alert policy or
// other dependent configs that would be rendered invalid by the deletion.
func (c *UptimeCheckClient) DeleteUptimeCheckConfig(ctx context.Context, req *monitoringpb.DeleteUptimeCheckConfigRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteUptimeCheckConfig(ctx, req, opts...)
}

// ListUptimeCheckIps returns the list of IP addresses that checkers run from
func (c *UptimeCheckClient) ListUptimeCheckIps(ctx context.Context, req *monitoringpb.ListUptimeCheckIpsRequest, opts ...gax.CallOption) *UptimeCheckIpIterator {
	return c.internalClient.ListUptimeCheckIps(ctx, req, opts...)
}

// uptimeCheckGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type uptimeCheckGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing UptimeCheckClient
	CallOptions **UptimeCheckCallOptions

	// The gRPC API client.
	uptimeCheckClient monitoringpb.UptimeCheckServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewUptimeCheckClient creates a new uptime check service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The UptimeCheckService API is used to manage (list, create, delete, edit)
// Uptime check configurations in the Stackdriver Monitoring product. An Uptime
// check is a piece of configuration that determines which resources and
// services to monitor for availability. These configurations can also be
// configured interactively by navigating to the [Cloud Console]
// (http://console.cloud.google.com (at http://console.cloud.google.com)), selecting the appropriate project,
// clicking on “Monitoring” on the left-hand side to navigate to Stackdriver,
// and then clicking on “Uptime”.
func NewUptimeCheckClient(ctx context.Context, opts ...option.ClientOption) (*UptimeCheckClient, error) {
	clientOpts := defaultUptimeCheckGRPCClientOptions()
	if newUptimeCheckClientHook != nil {
		hookOpts, err := newUptimeCheckClientHook(ctx, clientHookParams{})
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
	client := UptimeCheckClient{CallOptions: defaultUptimeCheckCallOptions()}

	c := &uptimeCheckGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		uptimeCheckClient: monitoringpb.NewUptimeCheckServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *uptimeCheckGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *uptimeCheckGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *uptimeCheckGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *uptimeCheckGRPCClient) ListUptimeCheckConfigs(ctx context.Context, req *monitoringpb.ListUptimeCheckConfigsRequest, opts ...gax.CallOption) *UptimeCheckConfigIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListUptimeCheckConfigs[0:len((*c.CallOptions).ListUptimeCheckConfigs):len((*c.CallOptions).ListUptimeCheckConfigs)], opts...)
	it := &UptimeCheckConfigIterator{}
	req = proto.Clone(req).(*monitoringpb.ListUptimeCheckConfigsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.UptimeCheckConfig, string, error) {
		resp := &monitoringpb.ListUptimeCheckConfigsResponse{}
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
			resp, err = c.uptimeCheckClient.ListUptimeCheckConfigs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetUptimeCheckConfigs(), resp.GetNextPageToken(), nil
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

func (c *uptimeCheckGRPCClient) GetUptimeCheckConfig(ctx context.Context, req *monitoringpb.GetUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetUptimeCheckConfig[0:len((*c.CallOptions).GetUptimeCheckConfig):len((*c.CallOptions).GetUptimeCheckConfig)], opts...)
	var resp *monitoringpb.UptimeCheckConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.uptimeCheckClient.GetUptimeCheckConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *uptimeCheckGRPCClient) CreateUptimeCheckConfig(ctx context.Context, req *monitoringpb.CreateUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateUptimeCheckConfig[0:len((*c.CallOptions).CreateUptimeCheckConfig):len((*c.CallOptions).CreateUptimeCheckConfig)], opts...)
	var resp *monitoringpb.UptimeCheckConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.uptimeCheckClient.CreateUptimeCheckConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *uptimeCheckGRPCClient) UpdateUptimeCheckConfig(ctx context.Context, req *monitoringpb.UpdateUptimeCheckConfigRequest, opts ...gax.CallOption) (*monitoringpb.UptimeCheckConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "uptime_check_config.name", url.QueryEscape(req.GetUptimeCheckConfig().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateUptimeCheckConfig[0:len((*c.CallOptions).UpdateUptimeCheckConfig):len((*c.CallOptions).UpdateUptimeCheckConfig)], opts...)
	var resp *monitoringpb.UptimeCheckConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.uptimeCheckClient.UpdateUptimeCheckConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *uptimeCheckGRPCClient) DeleteUptimeCheckConfig(ctx context.Context, req *monitoringpb.DeleteUptimeCheckConfigRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteUptimeCheckConfig[0:len((*c.CallOptions).DeleteUptimeCheckConfig):len((*c.CallOptions).DeleteUptimeCheckConfig)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.uptimeCheckClient.DeleteUptimeCheckConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *uptimeCheckGRPCClient) ListUptimeCheckIps(ctx context.Context, req *monitoringpb.ListUptimeCheckIpsRequest, opts ...gax.CallOption) *UptimeCheckIpIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListUptimeCheckIps[0:len((*c.CallOptions).ListUptimeCheckIps):len((*c.CallOptions).ListUptimeCheckIps)], opts...)
	it := &UptimeCheckIpIterator{}
	req = proto.Clone(req).(*monitoringpb.ListUptimeCheckIpsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.UptimeCheckIp, string, error) {
		resp := &monitoringpb.ListUptimeCheckIpsResponse{}
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
			resp, err = c.uptimeCheckClient.ListUptimeCheckIps(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetUptimeCheckIps(), resp.GetNextPageToken(), nil
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

// UptimeCheckConfigIterator manages a stream of *monitoringpb.UptimeCheckConfig.
type UptimeCheckConfigIterator struct {
	items    []*monitoringpb.UptimeCheckConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.UptimeCheckConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *UptimeCheckConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *UptimeCheckConfigIterator) Next() (*monitoringpb.UptimeCheckConfig, error) {
	var item *monitoringpb.UptimeCheckConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *UptimeCheckConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *UptimeCheckConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// UptimeCheckIpIterator manages a stream of *monitoringpb.UptimeCheckIp.
type UptimeCheckIpIterator struct {
	items    []*monitoringpb.UptimeCheckIp
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
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.UptimeCheckIp, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *UptimeCheckIpIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *UptimeCheckIpIterator) Next() (*monitoringpb.UptimeCheckIp, error) {
	var item *monitoringpb.UptimeCheckIp
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *UptimeCheckIpIterator) bufLen() int {
	return len(it.items)
}

func (it *UptimeCheckIpIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
