package asa

import (
	"context"
	"testing"
)

func TestGetCreativeAppAssets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &MediaCreativeSetDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.GetCreativeAppAssets(ctx, 1, &MediaCreativeSetRequest{})
	})
}

func TestGetAppPreviewDeviceSizes(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewDevicesMappingResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.GetAppPreviewDeviceSizes(ctx)
	})
}

func TestCreateAdGroupCreativeSets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupCreativeSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.CreateAdGroupCreativeSets(ctx, 1, 99, &CreateAdGroupCreativeSetRequest{})
	})
}

func TestFindAdGroupCreativeSets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupCreativeSetListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.FindAdGroupCreativeSets(ctx, 1, &FindAdGroupCreativeSetRequest{})
	})
}

func TestUpdateAdGroupCreativeSets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupCreativeSetResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.UpdateAdGroupCreativeSets(ctx, 1, 99, 10001, &AdGroupCreativeSetUpdate{})
	})
}

func TestDeleteAdGroupCreativeSets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &IntegerResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.DeleteAdGroupCreativeSets(ctx, 1, 99, []int64{10001})
	})
}
