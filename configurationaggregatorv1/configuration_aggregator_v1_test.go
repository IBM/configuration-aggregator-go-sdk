/**
 * (C) Copyright IBM Corp. 2025.
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

package configurationaggregatorv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ConfigurationAggregatorV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(configurationAggregatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(configurationAggregatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				URL: "https://configurationaggregatorv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(configurationAggregatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_AGGREGATOR_URL": "https://configurationaggregatorv1/api",
				"CONFIGURATION_AGGREGATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				})
				Expect(configurationAggregatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := configurationAggregatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationAggregatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationAggregatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationAggregatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL: "https://testService/api",
				})
				Expect(configurationAggregatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationAggregatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationAggregatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationAggregatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationAggregatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				})
				err := configurationAggregatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationAggregatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationAggregatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationAggregatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationAggregatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_AGGREGATOR_URL": "https://configurationaggregatorv1/api",
				"CONFIGURATION_AGGREGATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(configurationAggregatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_AGGREGATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(configurationAggregatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = configurationaggregatorv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := configurationaggregatorv1.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/provide-here-your-appconfig-instance-uuid"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := configurationaggregatorv1.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions) - Operation response error`, func() {
		listConfigsPath := "/configs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["config_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_tags"]).To(Equal([]string{"role:admin"}))
					Expect(req.URL.Query()["user_tags"]).To(Equal([]string{"test"}))
					Expect(req.URL.Query()["service_tags"]).To(Equal([]string{"test:tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigs with error: Operation response processing error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(configurationaggregatorv1.ListConfigsOptions)
				listConfigsOptionsModel.ConfigType = core.StringPtr("testString")
				listConfigsOptionsModel.ServiceName = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listConfigsOptionsModel.Location = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceCrn = core.StringPtr("testString")
				listConfigsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConfigsOptionsModel.Start = core.StringPtr("testString")
				listConfigsOptionsModel.SubAccount = core.StringPtr("testString")
				listConfigsOptionsModel.AccessTags = core.StringPtr("role:admin")
				listConfigsOptionsModel.UserTags = core.StringPtr("test")
				listConfigsOptionsModel.ServiceTags = core.StringPtr("test:tag")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationAggregatorService.EnableRetries(0, 0)
				result, response, operationErr = configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions)`, func() {
		listConfigsPath := "/configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["config_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_tags"]).To(Equal([]string{"role:admin"}))
					Expect(req.URL.Query()["user_tags"]).To(Equal([]string{"test"}))
					Expect(req.URL.Query()["service_tags"]).To(Equal([]string{"test:tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "prev": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "configs": [{"about": {"account_id": "AccountID", "config_type": "ConfigType", "resource_crn": "ResourceCrn", "resource_group_id": "ResourceGroupID", "service_name": "ServiceName", "resource_name": "ResourceName", "last_config_refresh_time": "2019-01-01T12:00:00.000Z", "location": "Location", "type_id": "TypeID", "access_tags": ["role:admin"], "user_tags": ["UserTags"], "service_tags": ["ServiceTags"], "created_at": "2021-05-12T23:20:50.520Z", "catalog_tags": "tag"}, "config": {}, "config_v2": {}}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully with retries`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				configurationAggregatorService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(configurationaggregatorv1.ListConfigsOptions)
				listConfigsOptionsModel.ConfigType = core.StringPtr("testString")
				listConfigsOptionsModel.ServiceName = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listConfigsOptionsModel.Location = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceCrn = core.StringPtr("testString")
				listConfigsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConfigsOptionsModel.Start = core.StringPtr("testString")
				listConfigsOptionsModel.SubAccount = core.StringPtr("testString")
				listConfigsOptionsModel.AccessTags = core.StringPtr("role:admin")
				listConfigsOptionsModel.UserTags = core.StringPtr("test")
				listConfigsOptionsModel.ServiceTags = core.StringPtr("test:tag")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationAggregatorService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationAggregatorService.DisableRetries()
				result, response, operationErr := configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationAggregatorService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["config_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_tags"]).To(Equal([]string{"role:admin"}))
					Expect(req.URL.Query()["user_tags"]).To(Equal([]string{"test"}))
					Expect(req.URL.Query()["service_tags"]).To(Equal([]string{"test:tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "prev": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "configs": [{"about": {"account_id": "AccountID", "config_type": "ConfigType", "resource_crn": "ResourceCrn", "resource_group_id": "ResourceGroupID", "service_name": "ServiceName", "resource_name": "ResourceName", "last_config_refresh_time": "2019-01-01T12:00:00.000Z", "location": "Location", "type_id": "TypeID", "access_tags": ["role:admin"], "user_tags": ["UserTags"], "service_tags": ["ServiceTags"], "created_at": "2021-05-12T23:20:50.520Z", "catalog_tags": "tag"}, "config": {}, "config_v2": {}}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationAggregatorService.ListConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(configurationaggregatorv1.ListConfigsOptions)
				listConfigsOptionsModel.ConfigType = core.StringPtr("testString")
				listConfigsOptionsModel.ServiceName = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listConfigsOptionsModel.Location = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceCrn = core.StringPtr("testString")
				listConfigsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConfigsOptionsModel.Start = core.StringPtr("testString")
				listConfigsOptionsModel.SubAccount = core.StringPtr("testString")
				listConfigsOptionsModel.AccessTags = core.StringPtr("role:admin")
				listConfigsOptionsModel.UserTags = core.StringPtr("test")
				listConfigsOptionsModel.ServiceTags = core.StringPtr("test:tag")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigs with error: Operation request error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(configurationaggregatorv1.ListConfigsOptions)
				listConfigsOptionsModel.ConfigType = core.StringPtr("testString")
				listConfigsOptionsModel.ServiceName = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listConfigsOptionsModel.Location = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceCrn = core.StringPtr("testString")
				listConfigsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConfigsOptionsModel.Start = core.StringPtr("testString")
				listConfigsOptionsModel.SubAccount = core.StringPtr("testString")
				listConfigsOptionsModel.AccessTags = core.StringPtr("role:admin")
				listConfigsOptionsModel.UserTags = core.StringPtr("test")
				listConfigsOptionsModel.ServiceTags = core.StringPtr("test:tag")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationAggregatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListConfigs successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(configurationaggregatorv1.ListConfigsOptions)
				listConfigsOptionsModel.ConfigType = core.StringPtr("testString")
				listConfigsOptionsModel.ServiceName = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listConfigsOptionsModel.Location = core.StringPtr("testString")
				listConfigsOptionsModel.ResourceCrn = core.StringPtr("testString")
				listConfigsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listConfigsOptionsModel.Start = core.StringPtr("testString")
				listConfigsOptionsModel.SubAccount = core.StringPtr("testString")
				listConfigsOptionsModel.AccessTags = core.StringPtr("role:admin")
				listConfigsOptionsModel.UserTags = core.StringPtr("test")
				listConfigsOptionsModel.ServiceTags = core.StringPtr("test:tag")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationAggregatorService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(configurationaggregatorv1.ListConfigsResponse)
				nextObject := new(configurationaggregatorv1.PaginatedNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(configurationaggregatorv1.ListConfigsResponse)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"configs":[{"about":{"account_id":"AccountID","config_type":"ConfigType","resource_crn":"ResourceCrn","resource_group_id":"ResourceGroupID","service_name":"ServiceName","resource_name":"ResourceName","last_config_refresh_time":"2019-01-01T12:00:00.000Z","location":"Location","type_id":"TypeID","access_tags":["role:admin"],"user_tags":["UserTags"],"service_tags":["ServiceTags"],"created_at":"2021-05-12T23:20:50.520Z","catalog_tags":"tag"},"config":{},"config_v2":{}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"configs":[{"about":{"account_id":"AccountID","config_type":"ConfigType","resource_crn":"ResourceCrn","resource_group_id":"ResourceGroupID","service_name":"ServiceName","resource_name":"ResourceName","last_config_refresh_time":"2019-01-01T12:00:00.000Z","location":"Location","type_id":"TypeID","access_tags":["role:admin"],"user_tags":["UserTags"],"service_tags":["ServiceTags"],"created_at":"2021-05-12T23:20:50.520Z","catalog_tags":"tag"},"config":{},"config_v2":{}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ConfigsPager.GetNext successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				listConfigsOptionsModel := &configurationaggregatorv1.ListConfigsOptions{
					ConfigType: core.StringPtr("testString"),
					ServiceName: core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					Location: core.StringPtr("testString"),
					ResourceCrn: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					SubAccount: core.StringPtr("testString"),
					AccessTags: core.StringPtr("role:admin"),
					UserTags: core.StringPtr("test"),
					ServiceTags: core.StringPtr("test:tag"),
				}

				pager, err := configurationAggregatorService.NewConfigsPager(listConfigsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []configurationaggregatorv1.Config
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ConfigsPager.GetAll successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				listConfigsOptionsModel := &configurationaggregatorv1.ListConfigsOptions{
					ConfigType: core.StringPtr("testString"),
					ServiceName: core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					Location: core.StringPtr("testString"),
					ResourceCrn: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					SubAccount: core.StringPtr("testString"),
					AccessTags: core.StringPtr("role:admin"),
					UserTags: core.StringPtr("test"),
					ServiceTags: core.StringPtr("test:tag"),
				}

				pager, err := configurationAggregatorService.NewConfigsPager(listConfigsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions) - Operation response error`, func() {
		replaceSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceSettings with error: Operation response processing error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := new(configurationaggregatorv1.ReplaceSettingsOptions)
				replaceSettingsOptionsModel.ResourceCollectionEnabled = core.BoolPtr(true)
				replaceSettingsOptionsModel.TrustedProfileID = core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.Regions = []string{"all"}
				replaceSettingsOptionsModel.AdditionalScope = []configurationaggregatorv1.AdditionalScope{*additionalScopeModel}
				replaceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationAggregatorService.EnableRetries(0, 0)
				result, response, operationErr = configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions)`, func() {
		replaceSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSettingsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_collection_enabled": false, "trusted_profile_id": "TrustedProfileID", "last_updated": "2019-01-01T12:00:00.000Z", "regions": ["us-south"], "additional_scope": [{"type": "Enterprise", "enterprise_id": "EnterpriseID", "profile_template": {"id": "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57", "trusted_profile_id": "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"}}]}`)
				}))
			})
			It(`Invoke ReplaceSettings successfully with retries`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				configurationAggregatorService.EnableRetries(0, 0)

				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := new(configurationaggregatorv1.ReplaceSettingsOptions)
				replaceSettingsOptionsModel.ResourceCollectionEnabled = core.BoolPtr(true)
				replaceSettingsOptionsModel.TrustedProfileID = core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.Regions = []string{"all"}
				replaceSettingsOptionsModel.AdditionalScope = []configurationaggregatorv1.AdditionalScope{*additionalScopeModel}
				replaceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationAggregatorService.ReplaceSettingsWithContext(ctx, replaceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationAggregatorService.DisableRetries()
				result, response, operationErr := configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationAggregatorService.ReplaceSettingsWithContext(ctx, replaceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSettingsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_collection_enabled": false, "trusted_profile_id": "TrustedProfileID", "last_updated": "2019-01-01T12:00:00.000Z", "regions": ["us-south"], "additional_scope": [{"type": "Enterprise", "enterprise_id": "EnterpriseID", "profile_template": {"id": "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57", "trusted_profile_id": "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"}}]}`)
				}))
			})
			It(`Invoke ReplaceSettings successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationAggregatorService.ReplaceSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := new(configurationaggregatorv1.ReplaceSettingsOptions)
				replaceSettingsOptionsModel.ResourceCollectionEnabled = core.BoolPtr(true)
				replaceSettingsOptionsModel.TrustedProfileID = core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.Regions = []string{"all"}
				replaceSettingsOptionsModel.AdditionalScope = []configurationaggregatorv1.AdditionalScope{*additionalScopeModel}
				replaceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceSettings with error: Operation request error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := new(configurationaggregatorv1.ReplaceSettingsOptions)
				replaceSettingsOptionsModel.ResourceCollectionEnabled = core.BoolPtr(true)
				replaceSettingsOptionsModel.TrustedProfileID = core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.Regions = []string{"all"}
				replaceSettingsOptionsModel.AdditionalScope = []configurationaggregatorv1.AdditionalScope{*additionalScopeModel}
				replaceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationAggregatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceSettings successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := new(configurationaggregatorv1.ReplaceSettingsOptions)
				replaceSettingsOptionsModel.ResourceCollectionEnabled = core.BoolPtr(true)
				replaceSettingsOptionsModel.TrustedProfileID = core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.Regions = []string{"all"}
				replaceSettingsOptionsModel.AdditionalScope = []configurationaggregatorv1.AdditionalScope{*additionalScopeModel}
				replaceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationAggregatorService.ReplaceSettings(replaceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(configurationaggregatorv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationAggregatorService.EnableRetries(0, 0)
				result, response, operationErr = configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_collection_enabled": false, "trusted_profile_id": "TrustedProfileID", "last_updated": "2019-01-01T12:00:00.000Z", "regions": ["us-south"], "additional_scope": [{"type": "Enterprise", "enterprise_id": "EnterpriseID", "profile_template": {"id": "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57", "trusted_profile_id": "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"}}]}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				configurationAggregatorService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(configurationaggregatorv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationAggregatorService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationAggregatorService.DisableRetries()
				result, response, operationErr := configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationAggregatorService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_collection_enabled": false, "trusted_profile_id": "TrustedProfileID", "last_updated": "2019-01-01T12:00:00.000Z", "regions": ["us-south"], "additional_scope": [{"type": "Enterprise", "enterprise_id": "EnterpriseID", "profile_template": {"id": "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57", "trusted_profile_id": "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"}}]}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationAggregatorService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(configurationaggregatorv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(configurationaggregatorv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationAggregatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(configurationaggregatorv1.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationAggregatorService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceCollectionStatus(getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions) - Operation response error`, func() {
		getResourceCollectionStatusPath := "/resource_collection_status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceCollectionStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceCollectionStatus with error: Operation response processing error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := new(configurationaggregatorv1.GetResourceCollectionStatusOptions)
				getResourceCollectionStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationAggregatorService.EnableRetries(0, 0)
				result, response, operationErr = configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceCollectionStatus(getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions)`, func() {
		getResourceCollectionStatusPath := "/resource_collection_status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceCollectionStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_config_refresh_time": "2019-01-01T12:00:00.000Z", "status": "initiated"}`)
				}))
			})
			It(`Invoke GetResourceCollectionStatus successfully with retries`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				configurationAggregatorService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := new(configurationaggregatorv1.GetResourceCollectionStatusOptions)
				getResourceCollectionStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationAggregatorService.GetResourceCollectionStatusWithContext(ctx, getResourceCollectionStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationAggregatorService.DisableRetries()
				result, response, operationErr := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationAggregatorService.GetResourceCollectionStatusWithContext(ctx, getResourceCollectionStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceCollectionStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_config_refresh_time": "2019-01-01T12:00:00.000Z", "status": "initiated"}`)
				}))
			})
			It(`Invoke GetResourceCollectionStatus successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationAggregatorService.GetResourceCollectionStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := new(configurationaggregatorv1.GetResourceCollectionStatusOptions)
				getResourceCollectionStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceCollectionStatus with error: Operation request error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := new(configurationaggregatorv1.GetResourceCollectionStatusOptions)
				getResourceCollectionStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationAggregatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetResourceCollectionStatus successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := new(configurationaggregatorv1.GetResourceCollectionStatusOptions)
				getResourceCollectionStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ManualReconcile(manualReconcileOptions *ManualReconcileOptions) - Operation response error`, func() {
		manualReconcilePath := "/reconcile"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(manualReconcilePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ManualReconcile with error: Operation response processing error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := new(configurationaggregatorv1.ManualReconcileOptions)
				manualReconcileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationAggregatorService.EnableRetries(0, 0)
				result, response, operationErr = configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ManualReconcile(manualReconcileOptions *ManualReconcileOptions)`, func() {
		manualReconcilePath := "/reconcile"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(manualReconcilePath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke ManualReconcile successfully with retries`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())
				configurationAggregatorService.EnableRetries(0, 0)

				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := new(configurationaggregatorv1.ManualReconcileOptions)
				manualReconcileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationAggregatorService.ManualReconcileWithContext(ctx, manualReconcileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationAggregatorService.DisableRetries()
				result, response, operationErr := configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationAggregatorService.ManualReconcileWithContext(ctx, manualReconcileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(manualReconcilePath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke ManualReconcile successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationAggregatorService.ManualReconcile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := new(configurationaggregatorv1.ManualReconcileOptions)
				manualReconcileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ManualReconcile with error: Operation request error`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := new(configurationaggregatorv1.ManualReconcileOptions)
				manualReconcileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationAggregatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ManualReconcile successfully`, func() {
				configurationAggregatorService, serviceErr := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationAggregatorService).ToNot(BeNil())

				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := new(configurationaggregatorv1.ManualReconcileOptions)
				manualReconcileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationAggregatorService.ManualReconcile(manualReconcileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			configurationAggregatorService, _ := configurationaggregatorv1.NewConfigurationAggregatorV1(&configurationaggregatorv1.ConfigurationAggregatorV1Options{
				URL:           "http://configurationaggregatorv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetResourceCollectionStatusOptions successfully`, func() {
				// Construct an instance of the GetResourceCollectionStatusOptions model
				getResourceCollectionStatusOptionsModel := configurationAggregatorService.NewGetResourceCollectionStatusOptions()
				getResourceCollectionStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceCollectionStatusOptionsModel).ToNot(BeNil())
				Expect(getResourceCollectionStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := configurationAggregatorService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigsOptions successfully`, func() {
				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := configurationAggregatorService.NewListConfigsOptions()
				listConfigsOptionsModel.SetConfigType("testString")
				listConfigsOptionsModel.SetServiceName("testString")
				listConfigsOptionsModel.SetResourceGroupID("testString")
				listConfigsOptionsModel.SetLocation("testString")
				listConfigsOptionsModel.SetResourceCrn("testString")
				listConfigsOptionsModel.SetLimit(int64(10))
				listConfigsOptionsModel.SetStart("testString")
				listConfigsOptionsModel.SetSubAccount("testString")
				listConfigsOptionsModel.SetAccessTags("role:admin")
				listConfigsOptionsModel.SetUserTags("test")
				listConfigsOptionsModel.SetServiceTags("test:tag")
				listConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigsOptionsModel).ToNot(BeNil())
				Expect(listConfigsOptionsModel.ConfigType).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.ResourceCrn).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listConfigsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.SubAccount).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.AccessTags).To(Equal(core.StringPtr("role:admin")))
				Expect(listConfigsOptionsModel.UserTags).To(Equal(core.StringPtr("test")))
				Expect(listConfigsOptionsModel.ServiceTags).To(Equal(core.StringPtr("test:tag")))
				Expect(listConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewManualReconcileOptions successfully`, func() {
				// Construct an instance of the ManualReconcileOptions model
				manualReconcileOptionsModel := configurationAggregatorService.NewManualReconcileOptions()
				manualReconcileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(manualReconcileOptionsModel).ToNot(BeNil())
				Expect(manualReconcileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceSettingsOptions successfully`, func() {
				// Construct an instance of the ProfileTemplate model
				profileTemplateModel := new(configurationaggregatorv1.ProfileTemplate)
				Expect(profileTemplateModel).ToNot(BeNil())
				profileTemplateModel.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
				profileTemplateModel.TrustedProfileID = core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")
				Expect(profileTemplateModel.ID).To(Equal(core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")))
				Expect(profileTemplateModel.TrustedProfileID).To(Equal(core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7")))

				// Construct an instance of the AdditionalScope model
				additionalScopeModel := new(configurationaggregatorv1.AdditionalScope)
				Expect(additionalScopeModel).ToNot(BeNil())
				additionalScopeModel.Type = core.StringPtr("Enterprise")
				additionalScopeModel.EnterpriseID = core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")
				additionalScopeModel.ProfileTemplate = profileTemplateModel
				Expect(additionalScopeModel.Type).To(Equal(core.StringPtr("Enterprise")))
				Expect(additionalScopeModel.EnterpriseID).To(Equal(core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61")))
				Expect(additionalScopeModel.ProfileTemplate).To(Equal(profileTemplateModel))

				// Construct an instance of the ReplaceSettingsOptions model
				replaceSettingsOptionsModel := configurationAggregatorService.NewReplaceSettingsOptions()
				replaceSettingsOptionsModel.SetResourceCollectionEnabled(true)
				replaceSettingsOptionsModel.SetTrustedProfileID("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
				replaceSettingsOptionsModel.SetRegions([]string{"all"})
				replaceSettingsOptionsModel.SetAdditionalScope([]configurationaggregatorv1.AdditionalScope{*additionalScopeModel})
				replaceSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceSettingsOptionsModel).ToNot(BeNil())
				Expect(replaceSettingsOptionsModel.ResourceCollectionEnabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceSettingsOptionsModel.TrustedProfileID).To(Equal(core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")))
				Expect(replaceSettingsOptionsModel.Regions).To(Equal([]string{"all"}))
				Expect(replaceSettingsOptionsModel.AdditionalScope).To(Equal([]configurationaggregatorv1.AdditionalScope{*additionalScopeModel}))
				Expect(replaceSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAdditionalScope successfully`, func() {
			// Construct an instance of the model.
			model := new(configurationaggregatorv1.AdditionalScope)
			model.Type = core.StringPtr("Enterprise")
			model.EnterpriseID = core.StringPtr("testString")
			model.ProfileTemplate = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *configurationaggregatorv1.AdditionalScope
			err = configurationaggregatorv1.UnmarshalAdditionalScope(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalProfileTemplate successfully`, func() {
			// Construct an instance of the model.
			model := new(configurationaggregatorv1.ProfileTemplate)
			model.ID = core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57")
			model.TrustedProfileID = core.StringPtr("Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *configurationaggregatorv1.ProfileTemplate
			err = configurationaggregatorv1.UnmarshalProfileTemplate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
