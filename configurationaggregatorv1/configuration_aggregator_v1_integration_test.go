//go:build integration

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

package configurationaggregatorv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the configurationaggregatorv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ConfigurationAggregatorV1 Integration Tests`, func() {
	const externalConfigFile = "../configuration_aggregator_v1.env"

	var (
		err          error
		configurationAggregatorService *configurationaggregatorv1.ConfigurationAggregatorV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configurationaggregatorv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			configurationAggregatorServiceOptions := &configurationaggregatorv1.ConfigurationAggregatorV1Options{}

			configurationAggregatorService, err = configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(configurationAggregatorServiceOptions)
			Expect(err).To(BeNil())
			Expect(configurationAggregatorService).ToNot(BeNil())
			Expect(configurationAggregatorService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			configurationAggregatorService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`NlpQueryConfigs - Query configs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`NlpQueryConfigs(nlpQueryConfigsOptions *NlpQueryConfigsOptions)`, func() {
			nlpQueryConfigsOptions := &configurationaggregatorv1.NlpQueryConfigsOptions{
				ResourceQuery: core.StringPtr("Give me a summary of my account configurations"),
				SubAccount: core.StringPtr("e91ff80ae07a4b299a18e1fc695ee40d"),
				HistoryID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
				Start: core.StringPtr("testString"),
			}

			nlpQueryResponse, response, err := configurationAggregatorService.NlpQueryConfigs(nlpQueryConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(nlpQueryResponse).ToNot(BeNil())
		})
	})

	Describe(`ListNlpChatHistory - List chat history`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNlpChatHistory(listNlpChatHistoryOptions *ListNlpChatHistoryOptions)`, func() {
			listNlpChatHistoryOptions := &configurationaggregatorv1.ListNlpChatHistoryOptions{
			}

			listChatHistoryResponse, response, err := configurationAggregatorService.ListNlpChatHistory(listNlpChatHistoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listChatHistoryResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateNlpChatHistory - Create chat history`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNlpChatHistory(createNlpChatHistoryOptions *CreateNlpChatHistoryOptions)`, func() {
			createNlpChatHistoryOptions := &configurationaggregatorv1.CreateNlpChatHistoryOptions{
				Title: core.StringPtr("Chat 3"),
				Query: core.StringPtr("Show configurations created in the last 30 days"),
			}

			chatHistoryResponse, response, err := configurationAggregatorService.CreateNlpChatHistory(createNlpChatHistoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(chatHistoryResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateNlpChatHistory - Update chat history`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNlpChatHistory(updateNlpChatHistoryOptions *UpdateNlpChatHistoryOptions)`, func() {
			updateNlpChatHistoryOptions := &configurationaggregatorv1.UpdateNlpChatHistoryOptions{
				ID: core.StringPtr("9e485b9c-6ccd-4623-9a60-6eb4d39af670"),
				Title: core.StringPtr("Chat 3"),
				Query: core.StringPtr("Do I have any VSI? Get me the names, CRNs, and locations of all VSI instances"),
				Deleted: core.BoolPtr(false),
			}

			chatHistoryResponse, response, err := configurationAggregatorService.UpdateNlpChatHistory(updateNlpChatHistoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(chatHistoryResponse).ToNot(BeNil())
		})
	})

	Describe(`ListNlpResourceQueryJobs - List jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNlpResourceQueryJobs(listNlpResourceQueryJobsOptions *ListNlpResourceQueryJobsOptions)`, func() {
			listNlpResourceQueryJobsOptions := &configurationaggregatorv1.ListNlpResourceQueryJobsOptions{
				SubAccount: core.StringPtr("testString"),
				Status: core.StringPtr("PENDING"),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
			}

			jobDetailResponseCollection, response, err := configurationAggregatorService.ListNlpResourceQueryJobs(listNlpResourceQueryJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobDetailResponseCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateNlpResourceQueryJob - Create job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNlpResourceQueryJob(createNlpResourceQueryJobOptions *CreateNlpResourceQueryJobOptions)`, func() {
			createNlpResourceQueryJobOptions := &configurationaggregatorv1.CreateNlpResourceQueryJobOptions{
				ResourceQuery: core.StringPtr("Show configurations created in the last 30 days"),
				SubAccount: core.StringPtr("dc420e985bb849d6b8352f71dac26ab8"),
				HistoryID: core.StringPtr("beb364aa-9451-4734-a594-c28174573eed"),
				FileType: core.StringPtr("json"),
			}

			jobDetailResponse, response, err := configurationAggregatorService.CreateNlpResourceQueryJob(createNlpResourceQueryJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(jobDetailResponse).ToNot(BeNil())
		})
	})

	Describe(`GetNlpResourceQueryJob - Get job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNlpResourceQueryJob(getNlpResourceQueryJobOptions *GetNlpResourceQueryJobOptions)`, func() {
			getNlpResourceQueryJobOptions := &configurationaggregatorv1.GetNlpResourceQueryJobOptions{
				JobID: core.StringPtr("89e6c312-9fad-4c5f-98b0-ffce9d171c12"),
			}

			jobDetailResponse, response, err := configurationAggregatorService.GetNlpResourceQueryJob(getNlpResourceQueryJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobDetailResponse).ToNot(BeNil())
		})
	})

	Describe(`DownloadNlpJobResult - Download job result`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DownloadNlpJobResult(downloadNlpJobResultOptions *DownloadNlpJobResultOptions)`, func() {
			downloadNlpJobResultOptions := &configurationaggregatorv1.DownloadNlpJobResultOptions{
				JobID: core.StringPtr("89e6c312-9fad-4c5f-98b0-ffce9d171c12"),
			}

			result, response, err := configurationAggregatorService.DownloadNlpJobResult(downloadNlpJobResultOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ListConfigs - List configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) with pagination`, func(){
			listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{
				ConfigType: core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				ResourceCrn: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
				SubAccount: core.StringPtr("testString"),
				AccessTags: core.StringPtr("role:admin"),
				UserTags: core.StringPtr("test"),
				ServiceTags: core.StringPtr("test:tag"),
			}

			listConfigsOptions.Start = nil
			listConfigsOptions.Limit = core.Int64Ptr(1)

			var allResults []configurationaggregatorv1.Config
			for {
				listConfigsResponse, response, err := configurationAggregatorService.ListConfigs(listConfigsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listConfigsResponse).ToNot(BeNil())
				allResults = append(allResults, listConfigsResponse.Configs...)

				nextToken, nextErr := listConfigsResponse.GetNextStart()
				Expect(nextErr).To(BeNil())
				listConfigsOptions.Start = nextToken

				if listConfigsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) using ConfigsPager`, func(){
			listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{
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

			// Test GetNext().
			pager, err := configurationAggregatorService.NewConfigsPager(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []configurationaggregatorv1.Config
			for pager.HasNext() {
				nextPage, nextPageErr := pager.GetNext()
				Expect(nextPageErr).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = configurationAggregatorService.NewConfigsPager(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListConfigs() returned a total of %d item(s) using ConfigsPager.\n", len(allResults))
		})
	})

	Describe(`ListBatchConfigs - List configurations for set of crn(s)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBatchConfigs(listBatchConfigsOptions *ListBatchConfigsOptions)`, func() {
			requestConfigsModel := &configurationaggregatorv1.RequestConfigs{
				ResourceCrn: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/c1d20fee2fe24c42b8ef6583283d2dcf:fc8de30c-43a5-407b-8c61-2b86dd820922:bucket:cos-pra-1"),
				ServiceName: core.StringPtr("testString"),
				ConfigType: []string{"testString"},
				TypeID: core.StringPtr("testString"),
			}

			listBatchConfigsOptions := &configurationaggregatorv1.ListBatchConfigsOptions{
				Configs: []configurationaggregatorv1.RequestConfigs{*requestConfigsModel},
				Limit: core.Int64Ptr(int64(1)),
				Start: core.StringPtr("testString"),
			}

			listConfigsQueryResponse, response, err := configurationAggregatorService.ListBatchConfigs(listBatchConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listConfigsQueryResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSettings - Update settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions)`, func() {
			profileTemplateModel := &configurationaggregatorv1.ProfileTemplate{
				ID: core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"),
				TrustedProfileID: core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7"),
			}

			additionalScopeModel := &configurationaggregatorv1.AdditionalScope{
				Type: core.StringPtr("Enterprise"),
				EnterpriseID: core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61"),
				ProfileTemplate: profileTemplateModel,
			}

			replaceSettingsOptions := &configurationaggregatorv1.ReplaceSettingsOptions{
				ResourceCollectionEnabled: core.BoolPtr(true),
				RealtimeConfigEnabled: core.BoolPtr(true),
				TrustedProfileID: core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560"),
				Regions: []string{"all"},
				AdditionalScope: []configurationaggregatorv1.AdditionalScope{*additionalScopeModel},
			}

			settingsResponse, response, err := configurationAggregatorService.ReplaceSettings(replaceSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSettings - Get settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &configurationaggregatorv1.GetSettingsOptions{
			}

			settingsResponse, response, err := configurationAggregatorService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetResourceCollectionStatus - Get resource collection status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceCollectionStatus(getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions)`, func() {
			getResourceCollectionStatusOptions := &configurationaggregatorv1.GetResourceCollectionStatusOptions{
			}

			statusResponse, response, err := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statusResponse).ToNot(BeNil())
		})
	})

	Describe(`ManualReconcile - Manual reconcile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ManualReconcile(manualReconcileOptions *ManualReconcileOptions)`, func() {
			manualReconcileOptions := &configurationaggregatorv1.ManualReconcileOptions{
				ResourceCrns: []string{"testString"},
			}

			manualReconcileResponse, response, err := configurationAggregatorService.ManualReconcile(manualReconcileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(manualReconcileResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
