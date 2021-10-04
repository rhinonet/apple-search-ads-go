/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asa

import (
	"context"
	"fmt"
)

// BudgetService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/campaigns
type BudgetService service

// BudgetOrderInfoResponse is a container for the budget order response body
//
// https://developer.apple.com/documentation/apple_search_ads/budgetorderinforesponse
type BudgetOrderInfoResponse struct {
	Data       *BudgetOrderInfo   `json:"data"`
	Error      *ErrorResponseBody `json:"error"`
	Pagination PageDetail         `json:"pagination"`
}

// BudgetOrderInfo is the response to a request for specific details of a budget order
//
// https://developer.apple.com/documentation/apple_search_ads/budgetorderinfo
type BudgetOrderInfo struct {
	Bo *BudgetOrder `json:"bo"`
}

type BudgetOrderStatus string

const (
	BudgetOrderStatusActive    BudgetOrderStatus = "ACTIVE"
	BudgetOrderStatusCanceled  BudgetOrderStatus = "CANCELED"
	BudgetOrderStatusExhausted BudgetOrderStatus = "EXHAUSTED"
	BudgetOrderStatusInActive  BudgetOrderStatus = "INACTIVE"
	BudgetOrderStatusComplete  BudgetOrderStatus = "COMPLETED"
)

type BudgetOrderSupplySources string

const (
	BudgetOrderSupplySourcesAppStoreSearchResults BudgetOrderSupplySources = "APPSTORE_SEARCH_RESULTS"
	BudgetOrderSupplySourcesAppStoreSearchTab     BudgetOrderSupplySources = "APPSTORE_SEARCH_TAB"
)

// BudgetOrder is the response to requests for budget order details
//
// https://developer.apple.com/documentation/apple_search_ads/budgetorder
type BudgetOrder struct {
	BillingEmail      string                   `json:"billingEmail,omitempty"`
	Budget            *Money                   `json:"budget,omitempty"`
	ClientName        string                   `json:"clientName,omitempty"`
	EndDate           DateTime                 `json:"endDate,omitempty"`
	Id                int64                    `json:"id,omitempty"`
	Name              string                   `json:"name,omitempty"`
	OrderNumber       string                   `json:"orderNumber,omitempty"`
	ParentOrgId       int64                    `json:"parentOrgId,omitempty"`
	PrimaryBuyerEmail string                   `json:"primaryBuyerEmail"`
	PrimaryBuyerName  string                   `json:"primaryBuyerName"`
	StartDate         DateTime                 `json:"startDate"`
	Status            BudgetOrderStatus        `json:"status"`
	SupplySources     BudgetOrderSupplySources `json:"supplySources"`
}

// GetBudgetOrder Fetches a specific budget order using a budget order identifier
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_budget_order
func (s *BudgetService) GetBudgetOrder(ctx context.Context, boId int64) (*BudgetOrderInfoResponse, *Response, error) {
	url := fmt.Sprintf("budgetorders/%d", boId)
	res := new(BudgetOrderInfoResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}

// GetAllBudgetOrdersQuery defines query parameter for GetAllBudgetOrders endpoint.
type GetAllBudgetOrdersQuery struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type BudgetOrderInfoListResponse struct {
	Data       []*BudgetOrderInfo `json:"data"`
	Error      *ErrorResponseBody `json:"error"`
	Pagination PageDetail         `json:"pagination"`
}

// GetAllBudgetOrders Fetches all assigned budget orders for an organization
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_budget_orders
func (s *BudgetService) GetAllBudgetOrders(ctx context.Context, params *GetAllBudgetOrdersQuery) (*BudgetOrderInfoListResponse, *Response, error) {
	url := fmt.Sprintf("budgetorders")
	res := new(BudgetOrderInfoListResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
