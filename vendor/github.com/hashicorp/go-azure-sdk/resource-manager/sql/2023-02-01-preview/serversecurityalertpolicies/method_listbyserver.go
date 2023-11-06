package serversecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

<<<<<<< HEAD
=======
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
>>>>>>> 5e957238fca9519400c2479c7d1f73e3d1b0871c
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerSecurityAlertPolicy
}

type ListByServerCompleteResult struct {
	Items []ServerSecurityAlertPolicy
}

// ListByServer ...
<<<<<<< HEAD
func (c ServerSecurityAlertPoliciesClient) ListByServer(ctx context.Context, id ServerId) (result ListByServerOperationResponse, err error) {
=======
func (c ServerSecurityAlertPoliciesClient) ListByServer(ctx context.Context, id commonids.SqlServerId) (result ListByServerOperationResponse, err error) {
>>>>>>> 5e957238fca9519400c2479c7d1f73e3d1b0871c
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/securityAlertPolicies", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ServerSecurityAlertPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServerComplete retrieves all the results into a single object
<<<<<<< HEAD
func (c ServerSecurityAlertPoliciesClient) ListByServerComplete(ctx context.Context, id ServerId) (ListByServerCompleteResult, error) {
=======
func (c ServerSecurityAlertPoliciesClient) ListByServerComplete(ctx context.Context, id commonids.SqlServerId) (ListByServerCompleteResult, error) {
>>>>>>> 5e957238fca9519400c2479c7d1f73e3d1b0871c
	return c.ListByServerCompleteMatchingPredicate(ctx, id, ServerSecurityAlertPolicyOperationPredicate{})
}

// ListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
<<<<<<< HEAD
func (c ServerSecurityAlertPoliciesClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id ServerId, predicate ServerSecurityAlertPolicyOperationPredicate) (result ListByServerCompleteResult, err error) {
=======
func (c ServerSecurityAlertPoliciesClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id commonids.SqlServerId, predicate ServerSecurityAlertPolicyOperationPredicate) (result ListByServerCompleteResult, err error) {
>>>>>>> 5e957238fca9519400c2479c7d1f73e3d1b0871c
	items := make([]ServerSecurityAlertPolicy, 0)

	resp, err := c.ListByServer(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByServerCompleteResult{
		Items: items,
	}
	return
}
