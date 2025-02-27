//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ReportsClient contains the methods for the Reports group.
// Don't use this type directly, use NewReportsClient() instead.
type ReportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewReportsClient creates a new instance of ReportsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReportsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReportsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetLatencyScorecards - Gets a Latency Scorecard for a given Experiment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// aggregationInterval - The aggregation interval of the Latency Scorecard
// options - ReportsClientGetLatencyScorecardsOptions contains the optional parameters for the ReportsClient.GetLatencyScorecards
// method.
func (client *ReportsClient) GetLatencyScorecards(ctx context.Context, resourceGroupName string, profileName string, experimentName string, aggregationInterval LatencyScorecardAggregationInterval, options *ReportsClientGetLatencyScorecardsOptions) (ReportsClientGetLatencyScorecardsResponse, error) {
	req, err := client.getLatencyScorecardsCreateRequest(ctx, resourceGroupName, profileName, experimentName, aggregationInterval, options)
	if err != nil {
		return ReportsClientGetLatencyScorecardsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReportsClientGetLatencyScorecardsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReportsClientGetLatencyScorecardsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLatencyScorecardsHandleResponse(resp)
}

// getLatencyScorecardsCreateRequest creates the GetLatencyScorecards request.
func (client *ReportsClient) getLatencyScorecardsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, aggregationInterval LatencyScorecardAggregationInterval, options *ReportsClientGetLatencyScorecardsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/LatencyScorecard"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	if options != nil && options.EndDateTimeUTC != nil {
		reqQP.Set("endDateTimeUTC", *options.EndDateTimeUTC)
	}
	if options != nil && options.Country != nil {
		reqQP.Set("country", *options.Country)
	}
	reqQP.Set("aggregationInterval", string(aggregationInterval))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLatencyScorecardsHandleResponse handles the GetLatencyScorecards response.
func (client *ReportsClient) getLatencyScorecardsHandleResponse(resp *http.Response) (ReportsClientGetLatencyScorecardsResponse, error) {
	result := ReportsClientGetLatencyScorecardsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LatencyScorecard); err != nil {
		return ReportsClientGetLatencyScorecardsResponse{}, err
	}
	return result, nil
}

// GetTimeseries - Gets a Timeseries for a given Experiment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// experimentName - The Experiment identifier associated with the Experiment
// startDateTimeUTC - The start DateTime of the Timeseries in UTC
// endDateTimeUTC - The end DateTime of the Timeseries in UTC
// aggregationInterval - The aggregation interval of the Timeseries
// timeseriesType - The type of Timeseries
// options - ReportsClientGetTimeseriesOptions contains the optional parameters for the ReportsClient.GetTimeseries method.
func (client *ReportsClient) GetTimeseries(ctx context.Context, resourceGroupName string, profileName string, experimentName string, startDateTimeUTC time.Time, endDateTimeUTC time.Time, aggregationInterval TimeseriesAggregationInterval, timeseriesType TimeseriesType, options *ReportsClientGetTimeseriesOptions) (ReportsClientGetTimeseriesResponse, error) {
	req, err := client.getTimeseriesCreateRequest(ctx, resourceGroupName, profileName, experimentName, startDateTimeUTC, endDateTimeUTC, aggregationInterval, timeseriesType, options)
	if err != nil {
		return ReportsClientGetTimeseriesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReportsClientGetTimeseriesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReportsClientGetTimeseriesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getTimeseriesHandleResponse(resp)
}

// getTimeseriesCreateRequest creates the GetTimeseries request.
func (client *ReportsClient) getTimeseriesCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, startDateTimeUTC time.Time, endDateTimeUTC time.Time, aggregationInterval TimeseriesAggregationInterval, timeseriesType TimeseriesType, options *ReportsClientGetTimeseriesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/Timeseries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	reqQP.Set("startDateTimeUTC", startDateTimeUTC.Format(time.RFC3339Nano))
	reqQP.Set("endDateTimeUTC", endDateTimeUTC.Format(time.RFC3339Nano))
	reqQP.Set("aggregationInterval", string(aggregationInterval))
	reqQP.Set("timeseriesType", string(timeseriesType))
	if options != nil && options.Endpoint != nil {
		reqQP.Set("endpoint", *options.Endpoint)
	}
	if options != nil && options.Country != nil {
		reqQP.Set("country", *options.Country)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTimeseriesHandleResponse handles the GetTimeseries response.
func (client *ReportsClient) getTimeseriesHandleResponse(resp *http.Response) (ReportsClientGetTimeseriesResponse, error) {
	result := ReportsClientGetTimeseriesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Timeseries); err != nil {
		return ReportsClientGetTimeseriesResponse{}, err
	}
	return result, nil
}
