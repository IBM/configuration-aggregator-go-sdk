/**
 * (C) Copyright IBM Corp. 2026.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.117.2-f568b7d6-20261016-181021
 */

// Package configurationaggregatorv1 : Operations and models for the ConfigurationAggregatorV1 service
package configurationaggregatorv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/configuration-aggregator-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// ConfigurationAggregatorV1 : Configuration Aggregator
//
// API Version: 1.0.0
// See: https://cloud.ibm.com/docs/app-configuration
type ConfigurationAggregatorV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/provide-here-your-appconfig-instance-uuid"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "configuration_aggregator"

const ParameterizedServiceURL = "https://{region}.apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/{instance_id}"

var defaultUrlVariables = map[string]string{
	"region": "us-south",
	"instance_id": "provide-here-your-appconfig-instance-uuid",
}

// ConfigurationAggregatorV1Options : Service options
type ConfigurationAggregatorV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewConfigurationAggregatorV1UsingExternalConfig : constructs an instance of ConfigurationAggregatorV1 with passed in options and external configuration.
func NewConfigurationAggregatorV1UsingExternalConfig(options *ConfigurationAggregatorV1Options) (configurationAggregator *ConfigurationAggregatorV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	configurationAggregator, err = NewConfigurationAggregatorV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = configurationAggregator.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = configurationAggregator.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewConfigurationAggregatorV1 : constructs an instance of ConfigurationAggregatorV1 with passed in options.
func NewConfigurationAggregatorV1(options *ConfigurationAggregatorV1Options) (service *ConfigurationAggregatorV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &ConfigurationAggregatorV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "configurationAggregator" suitable for processing requests.
func (configurationAggregator *ConfigurationAggregatorV1) Clone() *ConfigurationAggregatorV1 {
	if core.IsNil(configurationAggregator) {
		return nil
	}
	clone := *configurationAggregator
	clone.Service = configurationAggregator.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (configurationAggregator *ConfigurationAggregatorV1) SetServiceURL(url string) error {
	err := configurationAggregator.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (configurationAggregator *ConfigurationAggregatorV1) GetServiceURL() string {
	return configurationAggregator.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (configurationAggregator *ConfigurationAggregatorV1) SetDefaultHeaders(headers http.Header) {
	configurationAggregator.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (configurationAggregator *ConfigurationAggregatorV1) SetEnableGzipCompression(enableGzip bool) {
	configurationAggregator.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (configurationAggregator *ConfigurationAggregatorV1) GetEnableGzipCompression() bool {
	return configurationAggregator.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (configurationAggregator *ConfigurationAggregatorV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	configurationAggregator.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (configurationAggregator *ConfigurationAggregatorV1) DisableRetries() {
	configurationAggregator.Service.DisableRetries()
}

// NlpQueryConfigs : Query configs
// Retrieve config details based on the intent from user's query and generate response using large language model.
func (configurationAggregator *ConfigurationAggregatorV1) NlpQueryConfigs(nlpQueryConfigsOptions *NlpQueryConfigsOptions) (result *NLPQueryResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.NlpQueryConfigsWithContext(context.Background(), nlpQueryConfigsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// NlpQueryConfigsWithContext is an alternate form of the NlpQueryConfigs method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) NlpQueryConfigsWithContext(ctx context.Context, nlpQueryConfigsOptions *NlpQueryConfigsOptions) (result *NLPQueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(nlpQueryConfigsOptions, "nlpQueryConfigsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(nlpQueryConfigsOptions, "nlpQueryConfigsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/query`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "NlpQueryConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range nlpQueryConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if nlpQueryConfigsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*nlpQueryConfigsOptions.Limit))
	}
	if nlpQueryConfigsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*nlpQueryConfigsOptions.Start))
	}

	body := make(map[string]interface{})
	if nlpQueryConfigsOptions.ResourceQuery != nil {
		body["resource_query"] = nlpQueryConfigsOptions.ResourceQuery
	}
	if nlpQueryConfigsOptions.SubAccount != nil {
		body["sub_account"] = nlpQueryConfigsOptions.SubAccount
	}
	if nlpQueryConfigsOptions.HistoryID != nil {
		body["history_id"] = nlpQueryConfigsOptions.HistoryID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "nlp_query_configs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNLPQueryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListNlpChatHistory : List chat history
// Retrieve the chat history for the NLP resource query.
func (configurationAggregator *ConfigurationAggregatorV1) ListNlpChatHistory(listNlpChatHistoryOptions *ListNlpChatHistoryOptions) (result *ListChatHistoryResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ListNlpChatHistoryWithContext(context.Background(), listNlpChatHistoryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListNlpChatHistoryWithContext is an alternate form of the ListNlpChatHistory method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ListNlpChatHistoryWithContext(ctx context.Context, listNlpChatHistoryOptions *ListNlpChatHistoryOptions) (result *ListChatHistoryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNlpChatHistoryOptions, "listNlpChatHistoryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/history`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ListNlpChatHistory")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listNlpChatHistoryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_nlp_chat_history", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListChatHistoryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateNlpChatHistory : Create chat history
// Create a new chat history entry for the NLP resource query.
func (configurationAggregator *ConfigurationAggregatorV1) CreateNlpChatHistory(createNlpChatHistoryOptions *CreateNlpChatHistoryOptions) (result *ChatHistoryResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.CreateNlpChatHistoryWithContext(context.Background(), createNlpChatHistoryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateNlpChatHistoryWithContext is an alternate form of the CreateNlpChatHistory method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) CreateNlpChatHistoryWithContext(ctx context.Context, createNlpChatHistoryOptions *CreateNlpChatHistoryOptions) (result *ChatHistoryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNlpChatHistoryOptions, "createNlpChatHistoryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createNlpChatHistoryOptions, "createNlpChatHistoryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/history`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "CreateNlpChatHistory")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createNlpChatHistoryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createNlpChatHistoryOptions.Title != nil {
		body["title"] = createNlpChatHistoryOptions.Title
	}
	if createNlpChatHistoryOptions.Query != nil {
		body["query"] = createNlpChatHistoryOptions.Query
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_nlp_chat_history", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChatHistoryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateNlpChatHistory : Update chat history
// Update or rename an existing chat history entry for the NLP resource query.
func (configurationAggregator *ConfigurationAggregatorV1) UpdateNlpChatHistory(updateNlpChatHistoryOptions *UpdateNlpChatHistoryOptions) (result *ChatHistoryResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.UpdateNlpChatHistoryWithContext(context.Background(), updateNlpChatHistoryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateNlpChatHistoryWithContext is an alternate form of the UpdateNlpChatHistory method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) UpdateNlpChatHistoryWithContext(ctx context.Context, updateNlpChatHistoryOptions *UpdateNlpChatHistoryOptions) (result *ChatHistoryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNlpChatHistoryOptions, "updateNlpChatHistoryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateNlpChatHistoryOptions, "updateNlpChatHistoryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/history`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "UpdateNlpChatHistory")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateNlpChatHistoryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateNlpChatHistoryOptions.ID != nil {
		body["id"] = updateNlpChatHistoryOptions.ID
	}
	if updateNlpChatHistoryOptions.Title != nil {
		body["title"] = updateNlpChatHistoryOptions.Title
	}
	if updateNlpChatHistoryOptions.Query != nil {
		body["query"] = updateNlpChatHistoryOptions.Query
	}
	if updateNlpChatHistoryOptions.Deleted != nil {
		body["deleted"] = updateNlpChatHistoryOptions.Deleted
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_nlp_chat_history", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChatHistoryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListNlpResourceQueryJobs : List jobs
// Get all jobs for NLP resource query.
func (configurationAggregator *ConfigurationAggregatorV1) ListNlpResourceQueryJobs(listNlpResourceQueryJobsOptions *ListNlpResourceQueryJobsOptions) (result *JobDetailResponseCollection, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ListNlpResourceQueryJobsWithContext(context.Background(), listNlpResourceQueryJobsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListNlpResourceQueryJobsWithContext is an alternate form of the ListNlpResourceQueryJobs method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ListNlpResourceQueryJobsWithContext(ctx context.Context, listNlpResourceQueryJobsOptions *ListNlpResourceQueryJobsOptions) (result *JobDetailResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNlpResourceQueryJobsOptions, "listNlpResourceQueryJobsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/jobs`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ListNlpResourceQueryJobs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listNlpResourceQueryJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listNlpResourceQueryJobsOptions.SubAccount != nil {
		builder.AddQuery("sub_account", fmt.Sprint(*listNlpResourceQueryJobsOptions.SubAccount))
	}
	if listNlpResourceQueryJobsOptions.Status != nil {
		builder.AddQuery("status", fmt.Sprint(*listNlpResourceQueryJobsOptions.Status))
	}
	if listNlpResourceQueryJobsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listNlpResourceQueryJobsOptions.Limit))
	}
	if listNlpResourceQueryJobsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listNlpResourceQueryJobsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_nlp_resource_query_jobs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobDetailResponseCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateNlpResourceQueryJob : Create job
// Create a new job to process NLP resource query asynchronously.
func (configurationAggregator *ConfigurationAggregatorV1) CreateNlpResourceQueryJob(createNlpResourceQueryJobOptions *CreateNlpResourceQueryJobOptions) (result *JobDetailResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.CreateNlpResourceQueryJobWithContext(context.Background(), createNlpResourceQueryJobOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateNlpResourceQueryJobWithContext is an alternate form of the CreateNlpResourceQueryJob method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) CreateNlpResourceQueryJobWithContext(ctx context.Context, createNlpResourceQueryJobOptions *CreateNlpResourceQueryJobOptions) (result *JobDetailResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNlpResourceQueryJobOptions, "createNlpResourceQueryJobOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createNlpResourceQueryJobOptions, "createNlpResourceQueryJobOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/jobs`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "CreateNlpResourceQueryJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createNlpResourceQueryJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createNlpResourceQueryJobOptions.ResourceQuery != nil {
		body["resource_query"] = createNlpResourceQueryJobOptions.ResourceQuery
	}
	if createNlpResourceQueryJobOptions.SubAccount != nil {
		body["sub_account"] = createNlpResourceQueryJobOptions.SubAccount
	}
	if createNlpResourceQueryJobOptions.HistoryID != nil {
		body["history_id"] = createNlpResourceQueryJobOptions.HistoryID
	}
	if createNlpResourceQueryJobOptions.FileType != nil {
		body["file_type"] = createNlpResourceQueryJobOptions.FileType
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_nlp_resource_query_job", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobDetailResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetNlpResourceQueryJob : Get job
// Retrieve the detail of a specific job by its ID.
func (configurationAggregator *ConfigurationAggregatorV1) GetNlpResourceQueryJob(getNlpResourceQueryJobOptions *GetNlpResourceQueryJobOptions) (result *JobDetailResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.GetNlpResourceQueryJobWithContext(context.Background(), getNlpResourceQueryJobOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetNlpResourceQueryJobWithContext is an alternate form of the GetNlpResourceQueryJob method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) GetNlpResourceQueryJobWithContext(ctx context.Context, getNlpResourceQueryJobOptions *GetNlpResourceQueryJobOptions) (result *JobDetailResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getNlpResourceQueryJobOptions, "getNlpResourceQueryJobOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getNlpResourceQueryJobOptions, "getNlpResourceQueryJobOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"job_id": *getNlpResourceQueryJobOptions.JobID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/jobs/{job_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "GetNlpResourceQueryJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getNlpResourceQueryJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_nlp_resource_query_job", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobDetailResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DownloadNlpJobResult : Download job result
// Download the result file of a completed job.
func (configurationAggregator *ConfigurationAggregatorV1) DownloadNlpJobResult(downloadNlpJobResultOptions *DownloadNlpJobResultOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.DownloadNlpJobResultWithContext(context.Background(), downloadNlpJobResultOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DownloadNlpJobResultWithContext is an alternate form of the DownloadNlpJobResult method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) DownloadNlpJobResultWithContext(ctx context.Context, downloadNlpJobResultOptions *DownloadNlpJobResultOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(downloadNlpJobResultOptions, "downloadNlpJobResultOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(downloadNlpJobResultOptions, "downloadNlpJobResultOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"job_id": *downloadNlpJobResultOptions.JobID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/nlp_resource_query/jobs/{job_id}/result`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "DownloadNlpJobResult")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range downloadNlpJobResultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = configurationAggregator.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "download_nlp_job_result", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListConfigs : List configurations
// Retrieve the list of resource configurations collected as part of Configuration Aggregator.
func (configurationAggregator *ConfigurationAggregatorV1) ListConfigs(listConfigsOptions *ListConfigsOptions) (result *ListConfigsResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ListConfigsWithContext(context.Background(), listConfigsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListConfigsWithContext is an alternate form of the ListConfigs method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ListConfigsWithContext(ctx context.Context, listConfigsOptions *ListConfigsOptions) (result *ListConfigsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listConfigsOptions, "listConfigsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/configs`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ListConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listConfigsOptions.ConfigType != nil {
		builder.AddQuery("config_type", fmt.Sprint(*listConfigsOptions.ConfigType))
	}
	if listConfigsOptions.ServiceName != nil {
		builder.AddQuery("service_name", fmt.Sprint(*listConfigsOptions.ServiceName))
	}
	if listConfigsOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*listConfigsOptions.ResourceGroupID))
	}
	if listConfigsOptions.Location != nil {
		builder.AddQuery("location", fmt.Sprint(*listConfigsOptions.Location))
	}
	if listConfigsOptions.ResourceCrn != nil {
		builder.AddQuery("resource_crn", fmt.Sprint(*listConfigsOptions.ResourceCrn))
	}
	if listConfigsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listConfigsOptions.Limit))
	}
	if listConfigsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listConfigsOptions.Start))
	}
	if listConfigsOptions.SubAccount != nil {
		builder.AddQuery("sub_account", fmt.Sprint(*listConfigsOptions.SubAccount))
	}
	if listConfigsOptions.AccessTags != nil {
		builder.AddQuery("access_tags", fmt.Sprint(*listConfigsOptions.AccessTags))
	}
	if listConfigsOptions.UserTags != nil {
		builder.AddQuery("user_tags", fmt.Sprint(*listConfigsOptions.UserTags))
	}
	if listConfigsOptions.ServiceTags != nil {
		builder.AddQuery("service_tags", fmt.Sprint(*listConfigsOptions.ServiceTags))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_configs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListConfigsResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListBatchConfigs : List configurations for set of crn(s)
// List the resource configurations (requested for set of CRNs) collected as part of Configuration Aggregator.
func (configurationAggregator *ConfigurationAggregatorV1) ListBatchConfigs(listBatchConfigsOptions *ListBatchConfigsOptions) (result *ListConfigsQueryResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ListBatchConfigsWithContext(context.Background(), listBatchConfigsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListBatchConfigsWithContext is an alternate form of the ListBatchConfigs method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ListBatchConfigsWithContext(ctx context.Context, listBatchConfigsOptions *ListBatchConfigsOptions) (result *ListConfigsQueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBatchConfigsOptions, "listBatchConfigsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listBatchConfigsOptions, "listBatchConfigsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/configs/query`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ListBatchConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listBatchConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if listBatchConfigsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listBatchConfigsOptions.Limit))
	}
	if listBatchConfigsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listBatchConfigsOptions.Start))
	}

	body := make(map[string]interface{})
	if listBatchConfigsOptions.Configs != nil {
		body["configs"] = listBatchConfigsOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_batch_configs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListConfigsQueryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceSettings : Update settings
// Update resource aggregation settings.
func (configurationAggregator *ConfigurationAggregatorV1) ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions) (result *SettingsResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ReplaceSettingsWithContext(context.Background(), replaceSettingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceSettingsWithContext is an alternate form of the ReplaceSettings method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ReplaceSettingsWithContext(ctx context.Context, replaceSettingsOptions *ReplaceSettingsOptions) (result *SettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceSettingsOptions, "replaceSettingsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceSettingsOptions, "replaceSettingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/settings`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ReplaceSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceSettingsOptions.ResourceCollectionEnabled != nil {
		body["resource_collection_enabled"] = replaceSettingsOptions.ResourceCollectionEnabled
	}
	if replaceSettingsOptions.RealtimeConfigEnabled != nil {
		body["realtime_config_enabled"] = replaceSettingsOptions.RealtimeConfigEnabled
	}
	if replaceSettingsOptions.TrustedProfileID != nil {
		body["trusted_profile_id"] = replaceSettingsOptions.TrustedProfileID
	}
	if replaceSettingsOptions.Regions != nil {
		body["regions"] = replaceSettingsOptions.Regions
	}
	if replaceSettingsOptions.AdditionalScope != nil {
		body["additional_scope"] = replaceSettingsOptions.AdditionalScope
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "replace_settings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettingsResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetSettings : Get settings
// Retrieve settings for resource collection in Configuration Aggregator.
func (configurationAggregator *ConfigurationAggregatorV1) GetSettings(getSettingsOptions *GetSettingsOptions) (result *SettingsResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.GetSettingsWithContext(context.Background(), getSettingsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSettingsWithContext is an alternate form of the GetSettings method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) GetSettingsWithContext(ctx context.Context, getSettingsOptions *GetSettingsOptions) (result *SettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSettingsOptions, "getSettingsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/settings`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "GetSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_settings", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettingsResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetResourceCollectionStatus : Get resource collection status
// Retrieve the status of the resource collection as part of Configuration Aggregator.
func (configurationAggregator *ConfigurationAggregatorV1) GetResourceCollectionStatus(getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions) (result *StatusResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.GetResourceCollectionStatusWithContext(context.Background(), getResourceCollectionStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetResourceCollectionStatusWithContext is an alternate form of the GetResourceCollectionStatus method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) GetResourceCollectionStatusWithContext(ctx context.Context, getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions) (result *StatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getResourceCollectionStatusOptions, "getResourceCollectionStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/resource_collection_status`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "GetResourceCollectionStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getResourceCollectionStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_resource_collection_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatusResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ManualReconcile : Manual reconcile
// Manually trigger the recording of the Configuration items as part of Configuration Aggregator.
func (configurationAggregator *ConfigurationAggregatorV1) ManualReconcile(manualReconcileOptions *ManualReconcileOptions) (result *ManualReconcileResponse, response *core.DetailedResponse, err error) {
	result, response, err = configurationAggregator.ManualReconcileWithContext(context.Background(), manualReconcileOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ManualReconcileWithContext is an alternate form of the ManualReconcile method which supports a Context parameter
func (configurationAggregator *ConfigurationAggregatorV1) ManualReconcileWithContext(ctx context.Context, manualReconcileOptions *ManualReconcileOptions) (result *ManualReconcileResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(manualReconcileOptions, "manualReconcileOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationAggregator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationAggregator.Service.Options.URL, `/reconcile`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("configuration_aggregator", "V1", "ManualReconcile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range manualReconcileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if manualReconcileOptions.ResourceCrns != nil {
		body["resource_crns"] = manualReconcileOptions.ResourceCrns
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationAggregator.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "manual_reconcile", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManualReconcileResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// About : The basic metadata fetched from the query API.
type About struct {
	// The account ID in which the resource exists.
	AccountID *string `json:"account_id" validate:"required"`

	// The type of configuration of the retrieved resource.
	ConfigType *string `json:"config_type" validate:"required"`

	// The unique CRN of the IBM Cloud resource.
	ResourceCrn *string `json:"resource_crn" validate:"required"`

	// The resource group id of the resources.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// The resource group name of the resources.
	ResourceGroupName *string `json:"resource_group_name" validate:"required"`

	// The name of the service to which the resources belongs.
	ServiceName *string `json:"service_name" validate:"required"`

	// User defined name of the resource.
	ResourceName *string `json:"resource_name" validate:"required"`

	// The date and time the configuration was collected, in RFC 3339 format.
	LastConfigRefreshTime *strfmt.DateTime `json:"last_config_refresh_time" validate:"required"`

	// The location or region in which the resources are created.
	Location *string `json:"location" validate:"required"`

	// type of kubernetes cluster.
	Type *string `json:"type,omitempty"`

	// The unique identifier of the resource.
	TypeID *string `json:"type_id,omitempty"`

	// Access tags specified by the user for the resource. For more information, see
	// https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#tag-types.
	AccessTags []string `json:"access_tags" validate:"required"`

	// User tags specified by the user for the resource. For more information, see
	// https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#tag-types.
	UserTags []string `json:"user_tags" validate:"required"`

	// Tags attached to resources or service IDs by an authorized user in the account. For more information, see
	// https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#tag-types.
	ServiceTags []string `json:"service_tags" validate:"required"`

	// Creation time of the resource.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Catalog Tags attached to resources or service IDs by an authorized user in the account. For more information, see
	// https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#tag-types.
	CatalogTags []string `json:"catalog_tags" validate:"required"`
}

// UnmarshalAbout unmarshals an instance of About from the specified map of raw messages.
func UnmarshalAbout(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(About)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		err = core.SDKErrorf(err, "", "account_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "config_type", &obj.ConfigType)
	if err != nil {
		err = core.SDKErrorf(err, "", "config_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_group_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_name", &obj.ResourceGroupName)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_group_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_config_refresh_time", &obj.LastConfigRefreshTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_config_refresh_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		err = core.SDKErrorf(err, "", "location-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type_id", &obj.TypeID)
	if err != nil {
		err = core.SDKErrorf(err, "", "type_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "access_tags", &obj.AccessTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "user_tags", &obj.UserTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "user_tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_tags", &obj.ServiceTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_tags", &obj.CatalogTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AdditionalScope : The additional scope that enables resource collection for Enterprise acccounts.
type AdditionalScope struct {
	// The type of scope. Currently allowed value is Enterprise.
	Type *string `json:"type,omitempty"`

	// The Enterprise ID.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The Profile Template details applied on the enterprise account.
	ProfileTemplate *ProfileTemplate `json:"profile_template,omitempty"`
}

// UnmarshalAdditionalScope unmarshals an instance of AdditionalScope from the specified map of raw messages.
func UnmarshalAdditionalScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AdditionalScope)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		err = core.SDKErrorf(err, "", "enterprise_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "profile_template", &obj.ProfileTemplate, UnmarshalProfileTemplate)
	if err != nil {
		err = core.SDKErrorf(err, "", "profile_template-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChatHistoryResponse : A partial representation of a chat history entry.
type ChatHistoryResponse struct {
	// A unique identifier.
	ID *string `json:"id" validate:"required"`

	// Title for Chat.
	Title *string `json:"title" validate:"required"`

	// The natural language based resource query string.
	Query *string `json:"query" validate:"required"`
}

// UnmarshalChatHistoryResponse unmarshals an instance of ChatHistoryResponse from the specified map of raw messages.
func UnmarshalChatHistoryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChatHistoryResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		err = core.SDKErrorf(err, "", "title-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "query", &obj.Query)
	if err != nil {
		err = core.SDKErrorf(err, "", "query-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Config : Configuration of each individual resource.
type Config struct {
	// The basic metadata fetched from the query API.
	About *About `json:"about" validate:"required"`

	// The configuration of the resource.
	Config *Configuration `json:"config" validate:"required"`

	// The configuration of the resource.
	ConfigV2 *Configuration `json:"config_v2,omitempty"`
}

// UnmarshalConfig unmarshals an instance of Config from the specified map of raw messages.
func UnmarshalConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Config)
	err = core.UnmarshalModel(m, "about", &obj.About, UnmarshalAbout)
	if err != nil {
		err = core.SDKErrorf(err, "", "about-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "config_v2", &obj.ConfigV2, UnmarshalConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "config_v2-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Configuration : The configuration of the resource.
// This type supports additional properties of type interface{}.
type Configuration struct {

	// Allows users to set arbitrary properties of type interface{}.
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of Configuration.
func (o *Configuration) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of Configuration.
func (o *Configuration) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of Configuration.
func (o *Configuration) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of Configuration.
func (o *Configuration) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of Configuration
func (o *Configuration) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	if err != nil {
		err = core.SDKErrorf(err, "", "model-marshal", common.GetComponentInfo())
	}
	return
}

// UnmarshalConfiguration unmarshals an instance of Configuration from the specified map of raw messages.
func UnmarshalConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configuration)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = core.SDKErrorf(e, "", "additional-properties-error", common.GetComponentInfo())
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateNlpChatHistoryOptions : The CreateNlpChatHistory options.
type CreateNlpChatHistoryOptions struct {
	// Title for Chat.
	Title *string `json:"title" validate:"required"`

	// The natural language based resource query string.
	Query *string `json:"query" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateNlpChatHistoryOptions : Instantiate CreateNlpChatHistoryOptions
func (*ConfigurationAggregatorV1) NewCreateNlpChatHistoryOptions(title string, query string) *CreateNlpChatHistoryOptions {
	return &CreateNlpChatHistoryOptions{
		Title: core.StringPtr(title),
		Query: core.StringPtr(query),
	}
}

// SetTitle : Allow user to set Title
func (_options *CreateNlpChatHistoryOptions) SetTitle(title string) *CreateNlpChatHistoryOptions {
	_options.Title = core.StringPtr(title)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *CreateNlpChatHistoryOptions) SetQuery(query string) *CreateNlpChatHistoryOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNlpChatHistoryOptions) SetHeaders(param map[string]string) *CreateNlpChatHistoryOptions {
	options.Headers = param
	return options
}

// CreateNlpResourceQueryJobOptions : The CreateNlpResourceQueryJob options.
type CreateNlpResourceQueryJobOptions struct {
	// The natural language based resource query string.
	ResourceQuery *string `json:"resource_query" validate:"required"`

	// Filter the resource configurations from the specified sub-account in an enterprise hierarchy.
	SubAccount *string `json:"sub_account" validate:"required"`

	// A unique identifier.
	HistoryID *string `json:"history_id" validate:"required"`

	// Contanis the file type for download e.g. json/csv.
	FileType *string `json:"file_type" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateNlpResourceQueryJobOptions.FileType property.
// Contanis the file type for download e.g. json/csv.
const (
	CreateNlpResourceQueryJobOptions_FileType_Csv = "csv"
	CreateNlpResourceQueryJobOptions_FileType_JSON = "json"
)

// NewCreateNlpResourceQueryJobOptions : Instantiate CreateNlpResourceQueryJobOptions
func (*ConfigurationAggregatorV1) NewCreateNlpResourceQueryJobOptions(resourceQuery string, subAccount string, historyID string, fileType string) *CreateNlpResourceQueryJobOptions {
	return &CreateNlpResourceQueryJobOptions{
		ResourceQuery: core.StringPtr(resourceQuery),
		SubAccount: core.StringPtr(subAccount),
		HistoryID: core.StringPtr(historyID),
		FileType: core.StringPtr(fileType),
	}
}

// SetResourceQuery : Allow user to set ResourceQuery
func (_options *CreateNlpResourceQueryJobOptions) SetResourceQuery(resourceQuery string) *CreateNlpResourceQueryJobOptions {
	_options.ResourceQuery = core.StringPtr(resourceQuery)
	return _options
}

// SetSubAccount : Allow user to set SubAccount
func (_options *CreateNlpResourceQueryJobOptions) SetSubAccount(subAccount string) *CreateNlpResourceQueryJobOptions {
	_options.SubAccount = core.StringPtr(subAccount)
	return _options
}

// SetHistoryID : Allow user to set HistoryID
func (_options *CreateNlpResourceQueryJobOptions) SetHistoryID(historyID string) *CreateNlpResourceQueryJobOptions {
	_options.HistoryID = core.StringPtr(historyID)
	return _options
}

// SetFileType : Allow user to set FileType
func (_options *CreateNlpResourceQueryJobOptions) SetFileType(fileType string) *CreateNlpResourceQueryJobOptions {
	_options.FileType = core.StringPtr(fileType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNlpResourceQueryJobOptions) SetHeaders(param map[string]string) *CreateNlpResourceQueryJobOptions {
	options.Headers = param
	return options
}

// DownloadNlpJobResultOptions : The DownloadNlpJobResult options.
type DownloadNlpJobResultOptions struct {
	// The unique identifier of the job.
	JobID *string `json:"job_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDownloadNlpJobResultOptions : Instantiate DownloadNlpJobResultOptions
func (*ConfigurationAggregatorV1) NewDownloadNlpJobResultOptions(jobID string) *DownloadNlpJobResultOptions {
	return &DownloadNlpJobResultOptions{
		JobID: core.StringPtr(jobID),
	}
}

// SetJobID : Allow user to set JobID
func (_options *DownloadNlpJobResultOptions) SetJobID(jobID string) *DownloadNlpJobResultOptions {
	_options.JobID = core.StringPtr(jobID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DownloadNlpJobResultOptions) SetHeaders(param map[string]string) *DownloadNlpJobResultOptions {
	options.Headers = param
	return options
}

// Error : Error of each individual resource not found in aggregator.
type Error struct {
	// The unique CRN of the IBM Cloud resource.
	ResourceCrn *string `json:"resource_crn" validate:"required"`

	// The extended description that describes the solution to the problem.
	Message *string `json:"message" validate:"required"`

	// The code that is used to identify the problem that caused the error.
	ErrorCode *string `json:"error_code" validate:"required"`
}

// Constants associated with the Error.ErrorCode property.
// The code that is used to identify the problem that caused the error.
const (
	Error_ErrorCode_DbTimeout = "db_timeout"
	Error_ErrorCode_ResourceNotFound = "resource_not_found"
)

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "error_code", &obj.ErrorCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "error_code-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetNlpResourceQueryJobOptions : The GetNlpResourceQueryJob options.
type GetNlpResourceQueryJobOptions struct {
	// The unique identifier of the job.
	JobID *string `json:"job_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetNlpResourceQueryJobOptions : Instantiate GetNlpResourceQueryJobOptions
func (*ConfigurationAggregatorV1) NewGetNlpResourceQueryJobOptions(jobID string) *GetNlpResourceQueryJobOptions {
	return &GetNlpResourceQueryJobOptions{
		JobID: core.StringPtr(jobID),
	}
}

// SetJobID : Allow user to set JobID
func (_options *GetNlpResourceQueryJobOptions) SetJobID(jobID string) *GetNlpResourceQueryJobOptions {
	_options.JobID = core.StringPtr(jobID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetNlpResourceQueryJobOptions) SetHeaders(param map[string]string) *GetNlpResourceQueryJobOptions {
	options.Headers = param
	return options
}

// GetResourceCollectionStatusOptions : The GetResourceCollectionStatus options.
type GetResourceCollectionStatusOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetResourceCollectionStatusOptions : Instantiate GetResourceCollectionStatusOptions
func (*ConfigurationAggregatorV1) NewGetResourceCollectionStatusOptions() *GetResourceCollectionStatusOptions {
	return &GetResourceCollectionStatusOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceCollectionStatusOptions) SetHeaders(param map[string]string) *GetResourceCollectionStatusOptions {
	options.Headers = param
	return options
}

// GetSettingsOptions : The GetSettings options.
type GetSettingsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSettingsOptions : Instantiate GetSettingsOptions
func (*ConfigurationAggregatorV1) NewGetSettingsOptions() *GetSettingsOptions {
	return &GetSettingsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSettingsOptions) SetHeaders(param map[string]string) *GetSettingsOptions {
	options.Headers = param
	return options
}

// JobDetailResponse : The response body for job status details.
type JobDetailResponse struct {
	// A unique identifier.
	JobID *string `json:"job_id" validate:"required"`

	// Filter the resource configurations from the specified sub-account in an enterprise hierarchy.
	SubAccount *string `json:"sub_account" validate:"required"`

	// The natural language based resource query string.
	ResourceQuery *string `json:"resource_query" validate:"required"`

	// The current status of the job.
	Status *string `json:"status" validate:"required"`

	// Contanis the file type for download e.g. json/csv.
	FileType *string `json:"file_type,omitempty"`

	// Any error message related to the response.
	ErrorMessage *string `json:"error_message,omitempty"`

	// The timestamp when the job was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The timestamp when the job was updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// The timestamp when the job was completed.
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`
}

// Constants associated with the JobDetailResponse.Status property.
// The current status of the job.
const (
	JobDetailResponse_Status_Completed = "COMPLETED"
	JobDetailResponse_Status_Downloaded = "DOWNLOADED"
	JobDetailResponse_Status_Failed = "FAILED"
	JobDetailResponse_Status_InProgress = "IN_PROGRESS"
	JobDetailResponse_Status_Pending = "PENDING"
	JobDetailResponse_Status_Queued = "QUEUED"
	JobDetailResponse_Status_Unavailable = "UNAVAILABLE"
)

// Constants associated with the JobDetailResponse.FileType property.
// Contanis the file type for download e.g. json/csv.
const (
	JobDetailResponse_FileType_Csv = "csv"
	JobDetailResponse_FileType_JSON = "json"
)

// UnmarshalJobDetailResponse unmarshals an instance of JobDetailResponse from the specified map of raw messages.
func UnmarshalJobDetailResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobDetailResponse)
	err = core.UnmarshalPrimitive(m, "job_id", &obj.JobID)
	if err != nil {
		err = core.SDKErrorf(err, "", "job_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sub_account", &obj.SubAccount)
	if err != nil {
		err = core.SDKErrorf(err, "", "sub_account-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_query", &obj.ResourceQuery)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_query-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "file_type", &obj.FileType)
	if err != nil {
		err = core.SDKErrorf(err, "", "file_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "error_message", &obj.ErrorMessage)
	if err != nil {
		err = core.SDKErrorf(err, "", "error_message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "completed_at-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobDetailResponseCollection : JobDetailResponseCollection struct
type JobDetailResponseCollection struct {
	// The total number of job entries present.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The list of job entries.
	Jobs []JobDetailResponse `json:"jobs" validate:"required"`
}

// UnmarshalJobDetailResponseCollection unmarshals an instance of JobDetailResponseCollection from the specified map of raw messages.
func UnmarshalJobDetailResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobDetailResponseCollection)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "jobs", &obj.Jobs, UnmarshalJobDetailResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "jobs-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBatchConfigsOptions : The ListBatchConfigs options.
type ListBatchConfigsOptions struct {
	// The List of resources requested as part of List Configs Query API.
	Configs []RequestConfigs `json:"configs" validate:"required"`

	// The number of resources to return on the page.
	Limit *int64 `json:"limit,omitempty"`

	// The starting resource for the page, obtained from a previous list request. If unspecified, the first resource will
	// be returned.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListBatchConfigsOptions : Instantiate ListBatchConfigsOptions
func (*ConfigurationAggregatorV1) NewListBatchConfigsOptions(configs []RequestConfigs) *ListBatchConfigsOptions {
	return &ListBatchConfigsOptions{
		Configs: configs,
	}
}

// SetConfigs : Allow user to set Configs
func (_options *ListBatchConfigsOptions) SetConfigs(configs []RequestConfigs) *ListBatchConfigsOptions {
	_options.Configs = configs
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListBatchConfigsOptions) SetLimit(limit int64) *ListBatchConfigsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListBatchConfigsOptions) SetStart(start string) *ListBatchConfigsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBatchConfigsOptions) SetHeaders(param map[string]string) *ListBatchConfigsOptions {
	options.Headers = param
	return options
}

// ListChatHistoryResponse : The response body for listing chat history entries.
type ListChatHistoryResponse struct {
	// The total number of chat history entries present.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The list of chat history entries.
	Chats []ChatHistoryResponse `json:"chats" validate:"required"`
}

// UnmarshalListChatHistoryResponse unmarshals an instance of ListChatHistoryResponse from the specified map of raw messages.
func UnmarshalListChatHistoryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListChatHistoryResponse)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "chats", &obj.Chats, UnmarshalChatHistoryResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "chats-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListConfigsOptions : The ListConfigs options.
type ListConfigsOptions struct {
	// The type of resource configuration that are to be retrieved.
	ConfigType *string `json:"config_type,omitempty"`

	// The name of the IBM Cloud service for which resources are to be retrieved.
	ServiceName *string `json:"service_name,omitempty"`

	// The resource group id of the resources.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The location or region in which the resources are created.
	Location *string `json:"location,omitempty"`

	// The crn of the resource.
	ResourceCrn *string `json:"resource_crn,omitempty"`

	// The number of resources to return on the page.
	Limit *int64 `json:"limit,omitempty"`

	// The starting resource for the page, obtained from a previous list request. If unspecified, the first resource will
	// be returned.
	Start *string `json:"start,omitempty"`

	// Filter the resource configurations from the specified sub-account in an enterprise hierarchy.
	SubAccount *string `json:"sub_account,omitempty"`

	// Filter the resource configurations attached with the specified access tags.
	AccessTags *string `json:"access_tags,omitempty"`

	// Filter the resource configurations attached with the specified user tags.
	UserTags *string `json:"user_tags,omitempty"`

	// Filter the resource configurations attached with the specified service tags.
	ServiceTags *string `json:"service_tags,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListConfigsOptions : Instantiate ListConfigsOptions
func (*ConfigurationAggregatorV1) NewListConfigsOptions() *ListConfigsOptions {
	return &ListConfigsOptions{}
}

// SetConfigType : Allow user to set ConfigType
func (_options *ListConfigsOptions) SetConfigType(configType string) *ListConfigsOptions {
	_options.ConfigType = core.StringPtr(configType)
	return _options
}

// SetServiceName : Allow user to set ServiceName
func (_options *ListConfigsOptions) SetServiceName(serviceName string) *ListConfigsOptions {
	_options.ServiceName = core.StringPtr(serviceName)
	return _options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (_options *ListConfigsOptions) SetResourceGroupID(resourceGroupID string) *ListConfigsOptions {
	_options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return _options
}

// SetLocation : Allow user to set Location
func (_options *ListConfigsOptions) SetLocation(location string) *ListConfigsOptions {
	_options.Location = core.StringPtr(location)
	return _options
}

// SetResourceCrn : Allow user to set ResourceCrn
func (_options *ListConfigsOptions) SetResourceCrn(resourceCrn string) *ListConfigsOptions {
	_options.ResourceCrn = core.StringPtr(resourceCrn)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListConfigsOptions) SetLimit(limit int64) *ListConfigsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListConfigsOptions) SetStart(start string) *ListConfigsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetSubAccount : Allow user to set SubAccount
func (_options *ListConfigsOptions) SetSubAccount(subAccount string) *ListConfigsOptions {
	_options.SubAccount = core.StringPtr(subAccount)
	return _options
}

// SetAccessTags : Allow user to set AccessTags
func (_options *ListConfigsOptions) SetAccessTags(accessTags string) *ListConfigsOptions {
	_options.AccessTags = core.StringPtr(accessTags)
	return _options
}

// SetUserTags : Allow user to set UserTags
func (_options *ListConfigsOptions) SetUserTags(userTags string) *ListConfigsOptions {
	_options.UserTags = core.StringPtr(userTags)
	return _options
}

// SetServiceTags : Allow user to set ServiceTags
func (_options *ListConfigsOptions) SetServiceTags(serviceTags string) *ListConfigsOptions {
	_options.ServiceTags = core.StringPtr(serviceTags)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigsOptions) SetHeaders(param map[string]string) *ListConfigsOptions {
	options.Headers = param
	return options
}

// ListConfigsQueryResponse : List configs batch api response.
type ListConfigsQueryResponse struct {
	// The total number of resources present.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The maximum number of resources retrieved per page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first page of entries.
	First *PaginatedFirst `json:"first,omitempty"`

	// The reference to the previous page of entries. If absent, this is the first page.
	Prev *PaginatedPrevious `json:"prev,omitempty"`

	// The reference to the next page of entries. If absent, this is the last page.
	Next *PaginatedNext `json:"next,omitempty"`

	// Array of resource configurations.
	Configs []Config `json:"configs" validate:"required"`

	// Array of errors for not found resource configurations.
	Errors []Error `json:"errors" validate:"required"`
}

// UnmarshalListConfigsQueryResponse unmarshals an instance of ListConfigsQueryResponse from the specified map of raw messages.
func UnmarshalListConfigsQueryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListConfigsQueryResponse)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedFirst)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "prev", &obj.Prev, UnmarshalPaginatedPrevious)
	if err != nil {
		err = core.SDKErrorf(err, "", "prev-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedNext)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "configs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListConfigsResponse : List configs api response.
type ListConfigsResponse struct {
	// The total number of resources present.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The maximum number of resources retrieved per page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first page of entries.
	First *PaginatedFirst `json:"first,omitempty"`

	// The reference to the previous page of entries. If absent, this is the first page.
	Prev *PaginatedPrevious `json:"prev,omitempty"`

	// The reference to the next page of entries. If absent, this is the last page.
	Next *PaginatedNext `json:"next,omitempty"`

	// Array of resource configurations.
	Configs []Config `json:"configs" validate:"required"`
}

// UnmarshalListConfigsResponse unmarshals an instance of ListConfigsResponse from the specified map of raw messages.
func UnmarshalListConfigsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListConfigsResponse)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedFirst)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "prev", &obj.Prev, UnmarshalPaginatedPrevious)
	if err != nil {
		err = core.SDKErrorf(err, "", "prev-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedNext)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "configs-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListConfigsResponse) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// ListNlpChatHistoryOptions : The ListNlpChatHistory options.
type ListNlpChatHistoryOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListNlpChatHistoryOptions : Instantiate ListNlpChatHistoryOptions
func (*ConfigurationAggregatorV1) NewListNlpChatHistoryOptions() *ListNlpChatHistoryOptions {
	return &ListNlpChatHistoryOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListNlpChatHistoryOptions) SetHeaders(param map[string]string) *ListNlpChatHistoryOptions {
	options.Headers = param
	return options
}

// ListNlpResourceQueryJobsOptions : The ListNlpResourceQueryJobs options.
type ListNlpResourceQueryJobsOptions struct {
	// Filter jobs by the sub-account in an enterprise hierarchy.
	SubAccount *string `json:"sub_account,omitempty"`

	// Filter jobs by their current status.
	Status *string `json:"status,omitempty"`

	// The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different
	// set of records, use `limit` with `offset` to page through the available records.
	Limit *int64 `json:"limit,omitempty"`

	// The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset`
	// value. Use `offset` with `limit` to page through the available records.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the ListNlpResourceQueryJobsOptions.Status property.
// Filter jobs by their current status.
const (
	ListNlpResourceQueryJobsOptions_Status_Completed = "COMPLETED"
	ListNlpResourceQueryJobsOptions_Status_Failed = "FAILED"
	ListNlpResourceQueryJobsOptions_Status_InProgress = "IN_PROGRESS"
	ListNlpResourceQueryJobsOptions_Status_Pending = "PENDING"
	ListNlpResourceQueryJobsOptions_Status_Unavailable = "UNAVAILABLE"
)

// NewListNlpResourceQueryJobsOptions : Instantiate ListNlpResourceQueryJobsOptions
func (*ConfigurationAggregatorV1) NewListNlpResourceQueryJobsOptions() *ListNlpResourceQueryJobsOptions {
	return &ListNlpResourceQueryJobsOptions{}
}

// SetSubAccount : Allow user to set SubAccount
func (_options *ListNlpResourceQueryJobsOptions) SetSubAccount(subAccount string) *ListNlpResourceQueryJobsOptions {
	_options.SubAccount = core.StringPtr(subAccount)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ListNlpResourceQueryJobsOptions) SetStatus(status string) *ListNlpResourceQueryJobsOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListNlpResourceQueryJobsOptions) SetLimit(limit int64) *ListNlpResourceQueryJobsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListNlpResourceQueryJobsOptions) SetOffset(offset int64) *ListNlpResourceQueryJobsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListNlpResourceQueryJobsOptions) SetHeaders(param map[string]string) *ListNlpResourceQueryJobsOptions {
	options.Headers = param
	return options
}

// ManualReconcileOptions : The ManualReconcile options.
type ManualReconcileOptions struct {
	// Specific resource CRNs to reconcile (maximum ~20). The resource CRNs should be the instance CRN of the service.
	// Subresource CRNs (such as a Secrets Manager secret) are not supported.
	ResourceCrns []string `json:"resource_crns,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewManualReconcileOptions : Instantiate ManualReconcileOptions
func (*ConfigurationAggregatorV1) NewManualReconcileOptions() *ManualReconcileOptions {
	return &ManualReconcileOptions{}
}

// SetResourceCrns : Allow user to set ResourceCrns
func (_options *ManualReconcileOptions) SetResourceCrns(resourceCrns []string) *ManualReconcileOptions {
	_options.ResourceCrns = resourceCrns
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ManualReconcileOptions) SetHeaders(param map[string]string) *ManualReconcileOptions {
	options.Headers = param
	return options
}

// ManualReconcileResponse : The reconciliation response.
type ManualReconcileResponse struct {
	// Unique identifier for the reconciliation request.
	ReconcileID *string `json:"reconcile_id,omitempty"`

	// URL to check the status of the reconciliation request.
	Href *string `json:"href,omitempty"`
}

// UnmarshalManualReconcileResponse unmarshals an instance of ManualReconcileResponse from the specified map of raw messages.
func UnmarshalManualReconcileResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManualReconcileResponse)
	err = core.UnmarshalPrimitive(m, "reconcile_id", &obj.ReconcileID)
	if err != nil {
		err = core.SDKErrorf(err, "", "reconcile_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NLPQueryResponse : The resource query response.
type NLPQueryResponse struct {
	// The natural language based resource query string.
	ResourceQuery *string `json:"resource_query" validate:"required"`

	// The response message generated for the resource query.
	ResponseSummary *string `json:"response_summary" validate:"required"`

	// The total number of resources present.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The maximum number of resources retrieved per page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first page of entries.
	First *PaginatedFirst `json:"first,omitempty"`

	// The reference to the previous page of entries. If absent, this is the first page.
	Prev *PaginatedPrevious `json:"prev,omitempty"`

	// The reference to the next page of entries. If absent, this is the last page.
	Next *PaginatedNext `json:"next,omitempty"`

	// The resource details retrieved as a response to the query.
	Configs []map[string]interface{} `json:"configs" validate:"required"`

	// Any additional data related to the response.
	Remarks *string `json:"remarks,omitempty"`

	// The total number of LLM tokens used during the query.
	LlmTokensUsed *int64 `json:"llm_tokens_used,omitempty"`

	// A score between 0.0 and 1.0 indicating the system's confidence in the accuracy or completeness of the resource.
	ResponseScore *float32 `json:"response_score,omitempty"`

	// A unique identifier.
	HistoryID *string `json:"history_id,omitempty"`
}

// UnmarshalNLPQueryResponse unmarshals an instance of NLPQueryResponse from the specified map of raw messages.
func UnmarshalNLPQueryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NLPQueryResponse)
	err = core.UnmarshalPrimitive(m, "resource_query", &obj.ResourceQuery)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_query-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "response_summary", &obj.ResponseSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "response_summary-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedFirst)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "prev", &obj.Prev, UnmarshalPaginatedPrevious)
	if err != nil {
		err = core.SDKErrorf(err, "", "prev-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedNext)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "configs", &obj.Configs)
	if err != nil {
		err = core.SDKErrorf(err, "", "configs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "remarks", &obj.Remarks)
	if err != nil {
		err = core.SDKErrorf(err, "", "remarks-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "llm_tokens_used", &obj.LlmTokensUsed)
	if err != nil {
		err = core.SDKErrorf(err, "", "llm_tokens_used-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "response_score", &obj.ResponseScore)
	if err != nil {
		err = core.SDKErrorf(err, "", "response_score-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "history_id", &obj.HistoryID)
	if err != nil {
		err = core.SDKErrorf(err, "", "history_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NlpQueryConfigsOptions : The NlpQueryConfigs options.
type NlpQueryConfigsOptions struct {
	// The natural language based resource query string.
	ResourceQuery *string `json:"resource_query" validate:"required"`

	// Filter the resource configurations from the specified sub-account in an enterprise hierarchy.
	SubAccount *string `json:"sub_account" validate:"required"`

	// A unique identifier.
	HistoryID *string `json:"history_id,omitempty"`

	// The number of resources to return on the page.
	Limit *int64 `json:"limit,omitempty"`

	// The starting resource for the page, obtained from a previous list request. If unspecified, the first resource will
	// be returned.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewNlpQueryConfigsOptions : Instantiate NlpQueryConfigsOptions
func (*ConfigurationAggregatorV1) NewNlpQueryConfigsOptions(resourceQuery string, subAccount string) *NlpQueryConfigsOptions {
	return &NlpQueryConfigsOptions{
		ResourceQuery: core.StringPtr(resourceQuery),
		SubAccount: core.StringPtr(subAccount),
	}
}

// SetResourceQuery : Allow user to set ResourceQuery
func (_options *NlpQueryConfigsOptions) SetResourceQuery(resourceQuery string) *NlpQueryConfigsOptions {
	_options.ResourceQuery = core.StringPtr(resourceQuery)
	return _options
}

// SetSubAccount : Allow user to set SubAccount
func (_options *NlpQueryConfigsOptions) SetSubAccount(subAccount string) *NlpQueryConfigsOptions {
	_options.SubAccount = core.StringPtr(subAccount)
	return _options
}

// SetHistoryID : Allow user to set HistoryID
func (_options *NlpQueryConfigsOptions) SetHistoryID(historyID string) *NlpQueryConfigsOptions {
	_options.HistoryID = core.StringPtr(historyID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *NlpQueryConfigsOptions) SetLimit(limit int64) *NlpQueryConfigsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *NlpQueryConfigsOptions) SetStart(start string) *NlpQueryConfigsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *NlpQueryConfigsOptions) SetHeaders(param map[string]string) *NlpQueryConfigsOptions {
	options.Headers = param
	return options
}

// PaginatedFirst : The reference to the first page of entries.
type PaginatedFirst struct {
	// The reference to the first page of entries.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPaginatedFirst unmarshals an instance of PaginatedFirst from the specified map of raw messages.
func UnmarshalPaginatedFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginatedFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginatedNext : The reference to the next page of entries. If absent, this is the last page.
type PaginatedNext struct {
	// The reference to the next page of entries.
	Href *string `json:"href,omitempty"`

	// The starting resource for the page, obtained from a previous list request. If unspecified, the first resource will
	// be returned.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginatedNext unmarshals an instance of PaginatedNext from the specified map of raw messages.
func UnmarshalPaginatedNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginatedNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		err = core.SDKErrorf(err, "", "start-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginatedPrevious : The reference to the previous page of entries. If absent, this is the first page.
type PaginatedPrevious struct {
	// The reference to the previous page of entries.
	Href *string `json:"href,omitempty"`

	// The starting resource for the page, obtained from a previous list request. If unspecified, the first resource will
	// be returned.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginatedPrevious unmarshals an instance of PaginatedPrevious from the specified map of raw messages.
func UnmarshalPaginatedPrevious(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginatedPrevious)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		err = core.SDKErrorf(err, "", "start-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileTemplate : The Profile Template details applied on the enterprise account.
type ProfileTemplate struct {
	// The Profile Template ID created in the enterprise account that provides access to App Configuration instance for
	// resource collection.
	ID *string `json:"id,omitempty"`

	// The trusted profile ID that provides access to App Configuration instance to retrieve template information.
	TrustedProfileID *string `json:"trusted_profile_id,omitempty"`
}

// UnmarshalProfileTemplate unmarshals an instance of ProfileTemplate from the specified map of raw messages.
func UnmarshalProfileTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileTemplate)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "trusted_profile_id", &obj.TrustedProfileID)
	if err != nil {
		err = core.SDKErrorf(err, "", "trusted_profile_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReconciliationModeStatus : Status details of a reconciliation mode.
type ReconciliationModeStatus struct {
	// The timestamp at which the configuration was last refreshed by this reconciliation mode.
	LastRefreshTime *strfmt.DateTime `json:"last_refresh_time,omitempty"`
}

// UnmarshalReconciliationModeStatus unmarshals an instance of ReconciliationModeStatus from the specified map of raw messages.
func UnmarshalReconciliationModeStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReconciliationModeStatus)
	err = core.UnmarshalPrimitive(m, "last_refresh_time", &obj.LastRefreshTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_refresh_time-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceSettingsOptions : The ReplaceSettings options.
type ReplaceSettingsOptions struct {
	// The field denoting if the resource collection is enabled.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`

	// The field denoting if the realtime collection is enabled.
	RealtimeConfigEnabled *bool `json:"realtime_config_enabled,omitempty"`

	// The trusted profile id that provides Reader access to the App Configuration instance to collect resource metadata.
	TrustedProfileID *string `json:"trusted_profile_id,omitempty"`

	// The list of regions across which the resource collection is enabled.
	Regions []string `json:"regions,omitempty"`

	// The additional scope that enables resource collection for Enterprise acccounts.
	AdditionalScope []AdditionalScope `json:"additional_scope,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceSettingsOptions : Instantiate ReplaceSettingsOptions
func (*ConfigurationAggregatorV1) NewReplaceSettingsOptions() *ReplaceSettingsOptions {
	return &ReplaceSettingsOptions{}
}

// SetResourceCollectionEnabled : Allow user to set ResourceCollectionEnabled
func (_options *ReplaceSettingsOptions) SetResourceCollectionEnabled(resourceCollectionEnabled bool) *ReplaceSettingsOptions {
	_options.ResourceCollectionEnabled = core.BoolPtr(resourceCollectionEnabled)
	return _options
}

// SetRealtimeConfigEnabled : Allow user to set RealtimeConfigEnabled
func (_options *ReplaceSettingsOptions) SetRealtimeConfigEnabled(realtimeConfigEnabled bool) *ReplaceSettingsOptions {
	_options.RealtimeConfigEnabled = core.BoolPtr(realtimeConfigEnabled)
	return _options
}

// SetTrustedProfileID : Allow user to set TrustedProfileID
func (_options *ReplaceSettingsOptions) SetTrustedProfileID(trustedProfileID string) *ReplaceSettingsOptions {
	_options.TrustedProfileID = core.StringPtr(trustedProfileID)
	return _options
}

// SetRegions : Allow user to set Regions
func (_options *ReplaceSettingsOptions) SetRegions(regions []string) *ReplaceSettingsOptions {
	_options.Regions = regions
	return _options
}

// SetAdditionalScope : Allow user to set AdditionalScope
func (_options *ReplaceSettingsOptions) SetAdditionalScope(additionalScope []AdditionalScope) *ReplaceSettingsOptions {
	_options.AdditionalScope = additionalScope
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceSettingsOptions) SetHeaders(param map[string]string) *ReplaceSettingsOptions {
	options.Headers = param
	return options
}

// RequestConfigs : The details of each resource requested as part of List Configs Query API.
type RequestConfigs struct {
	// The unique CRN of the IBM Cloud resource.
	ResourceCrn *string `json:"resource_crn" validate:"required"`

	// The name of the service to which the resource belongs.
	ServiceName *string `json:"service_name,omitempty"`

	// The list of type for resource configuration that are to be retrieved.
	ConfigType []string `json:"config_type,omitempty"`

	// The unique identifier for each IBM resource.
	TypeID *string `json:"type_id,omitempty"`
}

// NewRequestConfigs : Instantiate RequestConfigs (Generic Model Constructor)
func (*ConfigurationAggregatorV1) NewRequestConfigs(resourceCrn string) (_model *RequestConfigs, err error) {
	_model = &RequestConfigs{
		ResourceCrn: core.StringPtr(resourceCrn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalRequestConfigs unmarshals an instance of RequestConfigs from the specified map of raw messages.
func UnmarshalRequestConfigs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RequestConfigs)
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "config_type", &obj.ConfigType)
	if err != nil {
		err = core.SDKErrorf(err, "", "config_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type_id", &obj.TypeID)
	if err != nil {
		err = core.SDKErrorf(err, "", "type_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SettingsResponse : Settings API response.
type SettingsResponse struct {
	// The field to check if the resource collection is enabled.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`

	// The field to check if the resource collection is enabled.
	RealtimeConfigEnabled *bool `json:"realtime_config_enabled,omitempty"`

	// The trusted profile ID that provides access to App Configuration instance to retrieve resource metadata.
	TrustedProfileID *string `json:"trusted_profile_id,omitempty"`

	// The last time the settings was last updated.
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Regions for which the resource collection is enabled.
	Regions []string `json:"regions" validate:"required"`

	// The additional scope that enables resource collection for Enterprise acccounts.
	AdditionalScope []AdditionalScope `json:"additional_scope" validate:"required"`
}

// UnmarshalSettingsResponse unmarshals an instance of SettingsResponse from the specified map of raw messages.
func UnmarshalSettingsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SettingsResponse)
	err = core.UnmarshalPrimitive(m, "resource_collection_enabled", &obj.ResourceCollectionEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_collection_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "realtime_config_enabled", &obj.RealtimeConfigEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "realtime_config_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "trusted_profile_id", &obj.TrustedProfileID)
	if err != nil {
		err = core.SDKErrorf(err, "", "trusted_profile_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated", &obj.LastUpdated)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "regions", &obj.Regions)
	if err != nil {
		err = core.SDKErrorf(err, "", "regions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "additional_scope", &obj.AdditionalScope, UnmarshalAdditionalScope)
	if err != nil {
		err = core.SDKErrorf(err, "", "additional_scope-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusResponse : The Status response.
type StatusResponse struct {
	// The timestamp at which the configuration was last refreshed.
	LastConfigRefreshTime *strfmt.DateTime `json:"last_config_refresh_time,omitempty"`

	// Latest collection mode.
	LatestReconciliationMode *string `json:"latest_reconciliation_mode,omitempty"`

	// Per reconciliation mode details. Only the modes that have run are present; the object is empty if no reconciliation
	// has completed yet.
	ReconciliationModes *StatusResponseReconciliationModes `json:"reconciliation_modes,omitempty"`

	// Status of the resource collection.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the StatusResponse.LatestReconciliationMode property.
// Latest collection mode.
const (
	StatusResponse_LatestReconciliationMode_Automatic = "automatic"
	StatusResponse_LatestReconciliationMode_Manual = "manual"
	StatusResponse_LatestReconciliationMode_Realtime = "realtime"
)

// Constants associated with the StatusResponse.Status property.
// Status of the resource collection.
const (
	StatusResponse_Status_Complete = "complete"
	StatusResponse_Status_Initialised = "initialised"
	StatusResponse_Status_Inprogress = "inprogress"
)

// UnmarshalStatusResponse unmarshals an instance of StatusResponse from the specified map of raw messages.
func UnmarshalStatusResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusResponse)
	err = core.UnmarshalPrimitive(m, "last_config_refresh_time", &obj.LastConfigRefreshTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_config_refresh_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "latest_reconciliation_mode", &obj.LatestReconciliationMode)
	if err != nil {
		err = core.SDKErrorf(err, "", "latest_reconciliation_mode-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "reconciliation_modes", &obj.ReconciliationModes, UnmarshalStatusResponseReconciliationModes)
	if err != nil {
		err = core.SDKErrorf(err, "", "reconciliation_modes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusResponseReconciliationModes : Per reconciliation mode details. Only the modes that have run are present; the object is empty if no reconciliation
// has completed yet.
type StatusResponseReconciliationModes struct {
	// Status details of a reconciliation mode.
	Automatic *ReconciliationModeStatus `json:"automatic,omitempty"`

	// Status details of a reconciliation mode.
	Manual *ReconciliationModeStatus `json:"manual,omitempty"`

	// Status details of a reconciliation mode.
	Realtime *ReconciliationModeStatus `json:"realtime,omitempty"`
}

// UnmarshalStatusResponseReconciliationModes unmarshals an instance of StatusResponseReconciliationModes from the specified map of raw messages.
func UnmarshalStatusResponseReconciliationModes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusResponseReconciliationModes)
	err = core.UnmarshalModel(m, "automatic", &obj.Automatic, UnmarshalReconciliationModeStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "automatic-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "manual", &obj.Manual, UnmarshalReconciliationModeStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "manual-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "realtime", &obj.Realtime, UnmarshalReconciliationModeStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "realtime-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateNlpChatHistoryOptions : The UpdateNlpChatHistory options.
type UpdateNlpChatHistoryOptions struct {
	// A unique identifier.
	ID *string `json:"id" validate:"required"`

	// Title for Chat.
	Title *string `json:"title,omitempty"`

	// The natural language based resource query string.
	Query *string `json:"query,omitempty"`

	// Indicates whether the chat history entry is deleted.
	Deleted *bool `json:"deleted,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateNlpChatHistoryOptions : Instantiate UpdateNlpChatHistoryOptions
func (*ConfigurationAggregatorV1) NewUpdateNlpChatHistoryOptions(id string) *UpdateNlpChatHistoryOptions {
	return &UpdateNlpChatHistoryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateNlpChatHistoryOptions) SetID(id string) *UpdateNlpChatHistoryOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetTitle : Allow user to set Title
func (_options *UpdateNlpChatHistoryOptions) SetTitle(title string) *UpdateNlpChatHistoryOptions {
	_options.Title = core.StringPtr(title)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *UpdateNlpChatHistoryOptions) SetQuery(query string) *UpdateNlpChatHistoryOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetDeleted : Allow user to set Deleted
func (_options *UpdateNlpChatHistoryOptions) SetDeleted(deleted bool) *UpdateNlpChatHistoryOptions {
	_options.Deleted = core.BoolPtr(deleted)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNlpChatHistoryOptions) SetHeaders(param map[string]string) *UpdateNlpChatHistoryOptions {
	options.Headers = param
	return options
}

//
// ConfigsPager can be used to simplify the use of the "ListConfigs" method.
//
type ConfigsPager struct {
	hasNext bool
	options *ListConfigsOptions
	client  *ConfigurationAggregatorV1
	pageContext struct {
		next *string
	}
}

// NewConfigsPager returns a new ConfigsPager instance.
func (configurationAggregator *ConfigurationAggregatorV1) NewConfigsPager(options *ListConfigsOptions) (pager *ConfigsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	optionsCopy := *options
	pager = &ConfigsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  configurationAggregator,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ConfigsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ConfigsPager) GetNextWithContext(ctx context.Context) (page []Config, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	var result *ListConfigsResponse
	result, _, err = pager.client.ListConfigsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Configs

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ConfigsPager) GetAllWithContext(ctx context.Context) (allItems []Config, err error) {
	for pager.HasNext() {
		var nextPage []Config
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ConfigsPager) GetNext() (page []Config, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ConfigsPager) GetAll() (allItems []Config, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
