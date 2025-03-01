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

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ExampleNewGlobalNetworkEndpointGroupsRESTClient() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleGlobalNetworkEndpointGroupsClient_AttachNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest.
	}
	resp, err := c.AttachNetworkEndpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalNetworkEndpointGroupsClient_Delete() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteGlobalNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#DeleteGlobalNetworkEndpointGroupRequest.
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalNetworkEndpointGroupsClient_DetachNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest.
	}
	resp, err := c.DetachNetworkEndpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalNetworkEndpointGroupsClient_Get() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetGlobalNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#GetGlobalNetworkEndpointGroupRequest.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalNetworkEndpointGroupsClient_Insert() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertGlobalNetworkEndpointGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#InsertGlobalNetworkEndpointGroupRequest.
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalNetworkEndpointGroupsClient_List() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListGlobalNetworkEndpointGroupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#ListGlobalNetworkEndpointGroupsRequest.
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleGlobalNetworkEndpointGroupsClient_ListNetworkEndpoints() {
	ctx := context.Background()
	c, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest.
	}
	it := c.ListNetworkEndpoints(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
