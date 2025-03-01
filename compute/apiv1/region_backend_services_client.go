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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionBackendServicesClientHook clientHook

// RegionBackendServicesCallOptions contains the retry settings for each method of RegionBackendServicesClient.
type RegionBackendServicesCallOptions struct {
	Delete    []gax.CallOption
	Get       []gax.CallOption
	GetHealth []gax.CallOption
	Insert    []gax.CallOption
	List      []gax.CallOption
	Patch     []gax.CallOption
	Update    []gax.CallOption
}

// internalRegionBackendServicesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalRegionBackendServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteRegionBackendServiceRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetRegionBackendServiceRequest, ...gax.CallOption) (*computepb.BackendService, error)
	GetHealth(context.Context, *computepb.GetHealthRegionBackendServiceRequest, ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error)
	Insert(context.Context, *computepb.InsertRegionBackendServiceRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListRegionBackendServicesRequest, ...gax.CallOption) *BackendServiceIterator
	Patch(context.Context, *computepb.PatchRegionBackendServiceRequest, ...gax.CallOption) (*Operation, error)
	Update(context.Context, *computepb.UpdateRegionBackendServiceRequest, ...gax.CallOption) (*Operation, error)
}

// RegionBackendServicesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionBackendServices API.
type RegionBackendServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionBackendServicesClient

	// The call options for this service.
	CallOptions *RegionBackendServicesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionBackendServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionBackendServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RegionBackendServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified regional BackendService resource.
func (c *RegionBackendServicesClient) Delete(ctx context.Context, req *computepb.DeleteRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified regional BackendService resource.
func (c *RegionBackendServicesClient) Get(ctx context.Context, req *computepb.GetRegionBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetHealth gets the most recent health check results for this regional BackendService.
func (c *RegionBackendServicesClient) GetHealth(ctx context.Context, req *computepb.GetHealthRegionBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	return c.internalClient.GetHealth(ctx, req, opts...)
}

// Insert creates a regional BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview.
func (c *RegionBackendServicesClient) Insert(ctx context.Context, req *computepb.InsertRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of regional BackendService resources available to the specified project in the given region.
func (c *RegionBackendServicesClient) List(ctx context.Context, req *computepb.ListRegionBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified regional BackendService resource with the data included in the request. For more information, see Understanding backend services This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *RegionBackendServicesClient) Patch(ctx context.Context, req *computepb.PatchRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Update updates the specified regional BackendService resource with the data included in the request. For more information, see Backend services overview .
func (c *RegionBackendServicesClient) Update(ctx context.Context, req *computepb.UpdateRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionBackendServicesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRegionBackendServicesRESTClient creates a new region backend services rest client.
//
// The RegionBackendServices API.
func NewRegionBackendServicesRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionBackendServicesClient, error) {
	clientOpts := append(defaultRegionBackendServicesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &regionBackendServicesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &RegionBackendServicesClient{internalClient: c, CallOptions: &RegionBackendServicesCallOptions{}}, nil
}

func defaultRegionBackendServicesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionBackendServicesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionBackendServicesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *regionBackendServicesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified regional BackendService resource.
func (c *regionBackendServicesRESTClient) Delete(ctx context.Context, req *computepb.DeleteRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices/%v", req.GetProject(), req.GetRegion(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified regional BackendService resource.
func (c *regionBackendServicesRESTClient) Get(ctx context.Context, req *computepb.GetRegionBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices/%v", req.GetProject(), req.GetRegion(), req.GetBackendService())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendService{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// GetHealth gets the most recent health check results for this regional BackendService.
func (c *regionBackendServicesRESTClient) GetHealth(ctx context.Context, req *computepb.GetHealthRegionBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetResourceGroupReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices/%v/getHealth", req.GetProject(), req.GetRegion(), req.GetBackendService())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendServiceGroupHealth{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// Insert creates a regional BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview.
func (c *regionBackendServicesRESTClient) Insert(ctx context.Context, req *computepb.InsertRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of regional BackendService resources available to the specified project in the given region.
func (c *regionBackendServicesRESTClient) List(ctx context.Context, req *computepb.ListRegionBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	it := &BackendServiceIterator{}
	req = proto.Clone(req).(*computepb.ListRegionBackendServicesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.BackendService, string, error) {
		resp := &computepb.BackendServiceList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch updates the specified regional BackendService resource with the data included in the request. For more information, see Understanding backend services This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *regionBackendServicesRESTClient) Patch(ctx context.Context, req *computepb.PatchRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices/%v", req.GetProject(), req.GetRegion(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Update updates the specified regional BackendService resource with the data included in the request. For more information, see Backend services overview .
func (c *regionBackendServicesRESTClient) Update(ctx context.Context, req *computepb.UpdateRegionBackendServiceRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/backendServices/%v", req.GetProject(), req.GetRegion(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}
