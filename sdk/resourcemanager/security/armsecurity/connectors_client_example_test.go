//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/GetSecurityConnectorsSubscription_example.json
func ExampleConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/GetSecurityConnectorsResourceGroup_example.json
func ExampleConnectorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("exampleResourceGroup", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
func ExampleConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/PutSecurityConnector_example.json
func ExampleConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", armsecurity.Connector{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.ConnectorProperties{
			EnvironmentData: &armsecurity.AwsEnvironmentData{
				EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
			},
			EnvironmentName:     to.Ptr(armsecurity.CloudNameAWS),
			HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
			Offerings: []armsecurity.CloudOfferingClassification{
				&armsecurity.CspmMonitorAwsOffering{
					OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
					NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
						CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/PatchSecurityConnector_example.json
func ExampleConnectorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", armsecurity.Connector{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.ConnectorProperties{
			EnvironmentData: &armsecurity.AwsEnvironmentData{
				EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
			},
			EnvironmentName:     to.Ptr(armsecurity.CloudNameAWS),
			HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
			Offerings: []armsecurity.CloudOfferingClassification{
				&armsecurity.CspmMonitorAwsOffering{
					OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
					NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
						CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/DeleteSecurityConnector_example.json
func ExampleConnectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "myRg", "mySecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
