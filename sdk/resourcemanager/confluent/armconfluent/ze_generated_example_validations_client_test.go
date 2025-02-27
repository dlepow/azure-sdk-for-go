//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Validations_ValidateOrganizations.json
func ExampleValidationsClient_ValidateOrganization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconfluent.NewValidationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ValidateOrganization(ctx,
		"myResourceGroup",
		"myOrganization",
		armconfluent.OrganizationResource{
			Location: to.Ptr("West US"),
			Properties: &armconfluent.OrganizationResourceProperties{
				OfferDetail: &armconfluent.OfferDetail{
					ID:          to.Ptr("string"),
					PlanID:      to.Ptr("string"),
					PlanName:    to.Ptr("string"),
					PublisherID: to.Ptr("string"),
					TermUnit:    to.Ptr("string"),
				},
				UserDetail: &armconfluent.UserDetail{
					EmailAddress: to.Ptr("abc@microsoft.com"),
					FirstName:    to.Ptr("string"),
					LastName:     to.Ptr("string"),
				},
			},
			Tags: map[string]*string{
				"Environment": to.Ptr("Dev"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
