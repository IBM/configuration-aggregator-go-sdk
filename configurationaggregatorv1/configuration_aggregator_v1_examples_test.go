//go:build examples

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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Configuration Aggregator service.
//
// The following configuration properties are assumed to be defined:
// CONFIGURATION_AGGREGATOR_URL=<service base url>
// CONFIGURATION_AGGREGATOR_AUTH_TYPE=iam
// CONFIGURATION_AGGREGATOR_APIKEY=<IAM apikey>
// CONFIGURATION_AGGREGATOR_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`ConfigurationAggregatorV1 Examples Tests`, func() {

	const externalConfigFile = "../configuration_aggregator_v1.env"

	var (
		configurationAggregatorService *configurationaggregatorv1.ConfigurationAggregatorV1
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configurationaggregatorv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			configurationAggregatorServiceOptions := &configurationaggregatorv1.ConfigurationAggregatorV1Options{}

			configurationAggregatorService, err = configurationaggregatorv1.NewConfigurationAggregatorV1UsingExternalConfig(configurationAggregatorServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(configurationAggregatorService).ToNot(BeNil())
		})
	})

	Describe(`ConfigurationAggregatorV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`NlpQueryConfigs request example`, func() {
			fmt.Println("\nNlpQueryConfigs() result:")
			// begin-nlp_query_configs

			nlpQueryConfigsOptions := configurationAggregatorService.NewNlpQueryConfigsOptions(
				"Give me a summary of my account configurations",
				"e91ff80ae07a4b299a18e1fc695ee40d",
			)

			nlpQueryResponse, response, err := configurationAggregatorService.NlpQueryConfigs(nlpQueryConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(nlpQueryResponse, "", "  ")
			fmt.Println(string(b))

			// end-nlp_query_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(nlpQueryResponse).ToNot(BeNil())
		})
		It(`ListNlpChatHistory request example`, func() {
			fmt.Println("\nListNlpChatHistory() result:")
			// begin-list_nlp_chat_history

			listNlpChatHistoryOptions := configurationAggregatorService.NewListNlpChatHistoryOptions()

			listChatHistoryResponse, response, err := configurationAggregatorService.ListNlpChatHistory(listNlpChatHistoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listChatHistoryResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_nlp_chat_history

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listChatHistoryResponse).ToNot(BeNil())
		})
		It(`CreateNlpChatHistory request example`, func() {
			fmt.Println("\nCreateNlpChatHistory() result:")
			// begin-create_nlp_chat_history

			createNlpChatHistoryOptions := configurationAggregatorService.NewCreateNlpChatHistoryOptions(
				"Chat 3",
				"Show configurations created in the last 30 days",
			)

			chatHistoryResponse, response, err := configurationAggregatorService.CreateNlpChatHistory(createNlpChatHistoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(chatHistoryResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_nlp_chat_history

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(chatHistoryResponse).ToNot(BeNil())
		})
		It(`UpdateNlpChatHistory request example`, func() {
			fmt.Println("\nUpdateNlpChatHistory() result:")
			// begin-update_nlp_chat_history

			updateNlpChatHistoryOptions := configurationAggregatorService.NewUpdateNlpChatHistoryOptions(
				"9e485b9c-6ccd-4623-9a60-6eb4d39af670",
			)
			updateNlpChatHistoryOptions.SetTitle("Chat 3")
			updateNlpChatHistoryOptions.SetQuery("Do I have any VSI? Get me the names, CRNs, and locations of all VSI instances")
			updateNlpChatHistoryOptions.SetDeleted(false)

			chatHistoryResponse, response, err := configurationAggregatorService.UpdateNlpChatHistory(updateNlpChatHistoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(chatHistoryResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_nlp_chat_history

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(chatHistoryResponse).ToNot(BeNil())
		})
		It(`ListNlpResourceQueryJobs request example`, func() {
			fmt.Println("\nListNlpResourceQueryJobs() result:")
			// begin-list_nlp_resource_query_jobs

			listNlpResourceQueryJobsOptions := configurationAggregatorService.NewListNlpResourceQueryJobsOptions()

			jobDetailResponseCollection, response, err := configurationAggregatorService.ListNlpResourceQueryJobs(listNlpResourceQueryJobsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobDetailResponseCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_nlp_resource_query_jobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobDetailResponseCollection).ToNot(BeNil())
		})
		It(`CreateNlpResourceQueryJob request example`, func() {
			fmt.Println("\nCreateNlpResourceQueryJob() result:")
			// begin-create_nlp_resource_query_job

			createNlpResourceQueryJobOptions := configurationAggregatorService.NewCreateNlpResourceQueryJobOptions(
				"Show configurations created in the last 30 days",
				"dc420e985bb849d6b8352f71dac26ab8",
				"beb364aa-9451-4734-a594-c28174573eed",
				"json",
			)

			jobDetailResponse, response, err := configurationAggregatorService.CreateNlpResourceQueryJob(createNlpResourceQueryJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobDetailResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_nlp_resource_query_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(jobDetailResponse).ToNot(BeNil())
		})
		It(`GetNlpResourceQueryJob request example`, func() {
			fmt.Println("\nGetNlpResourceQueryJob() result:")
			// begin-get_nlp_resource_query_job

			getNlpResourceQueryJobOptions := configurationAggregatorService.NewGetNlpResourceQueryJobOptions(
				"89e6c312-9fad-4c5f-98b0-ffce9d171c12",
			)

			jobDetailResponse, response, err := configurationAggregatorService.GetNlpResourceQueryJob(getNlpResourceQueryJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobDetailResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_nlp_resource_query_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobDetailResponse).ToNot(BeNil())
		})
		It(`DownloadNlpJobResult request example`, func() {
			fmt.Println("\nDownloadNlpJobResult() result:")
			// begin-download_nlp_job_result

			downloadNlpJobResultOptions := configurationAggregatorService.NewDownloadNlpJobResultOptions(
				"89e6c312-9fad-4c5f-98b0-ffce9d171c12",
			)

			result, response, err := configurationAggregatorService.DownloadNlpJobResult(downloadNlpJobResultOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, outErr := os.Create("result.out")
				if outErr != nil { panic(outErr) }
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil { panic(err) }
			}

			// end-download_nlp_job_result

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ListConfigs request example`, func() {
			fmt.Println("\nListConfigs() result:")
			// begin-list_configs
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

			pager, err := configurationAggregatorService.NewConfigsPager(listConfigsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []configurationaggregatorv1.Config
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_configs
		})
		It(`ListBatchConfigs request example`, func() {
			fmt.Println("\nListBatchConfigs() result:")
			// begin-list_batch_configs

			requestConfigsModel := &configurationaggregatorv1.RequestConfigs{
				ResourceCrn: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/c1d20fee2fe24c42b8ef6583283d2dcf:fc8de30c-43a5-407b-8c61-2b86dd820922:bucket:cos-pra-1"),
			}

			listBatchConfigsOptions := configurationAggregatorService.NewListBatchConfigsOptions(
				[]configurationaggregatorv1.RequestConfigs{*requestConfigsModel},
			)

			listConfigsQueryResponse, response, err := configurationAggregatorService.ListBatchConfigs(listBatchConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listConfigsQueryResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_batch_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listConfigsQueryResponse).ToNot(BeNil())
		})
		It(`ReplaceSettings request example`, func() {
			fmt.Println("\nReplaceSettings() result:")
			// begin-replace_settings

			profileTemplateModel := &configurationaggregatorv1.ProfileTemplate{
				ID: core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"),
				TrustedProfileID: core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7"),
			}

			additionalScopeModel := &configurationaggregatorv1.AdditionalScope{
				Type: core.StringPtr("Enterprise"),
				EnterpriseID: core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61"),
				ProfileTemplate: profileTemplateModel,
			}

			replaceSettingsOptions := configurationAggregatorService.NewReplaceSettingsOptions()
			replaceSettingsOptions.SetResourceCollectionEnabled(true)
			replaceSettingsOptions.SetRealtimeConfigEnabled(true)
			replaceSettingsOptions.SetTrustedProfileID("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560")
			replaceSettingsOptions.SetRegions([]string{"all"})
			replaceSettingsOptions.SetAdditionalScope([]configurationaggregatorv1.AdditionalScope{*additionalScopeModel})

			settingsResponse, response, err := configurationAggregatorService.ReplaceSettings(replaceSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := configurationAggregatorService.NewGetSettingsOptions()

			settingsResponse, response, err := configurationAggregatorService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
		It(`GetResourceCollectionStatus request example`, func() {
			fmt.Println("\nGetResourceCollectionStatus() result:")
			// begin-get_resource_collection_status

			getResourceCollectionStatusOptions := configurationAggregatorService.NewGetResourceCollectionStatusOptions()

			statusResponse, response, err := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(statusResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_collection_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statusResponse).ToNot(BeNil())
		})
		It(`ManualReconcile request example`, func() {
			fmt.Println("\nManualReconcile() result:")
			// begin-manual_reconcile

			manualReconcileOptions := configurationAggregatorService.NewManualReconcileOptions()

			manualReconcileResponse, response, err := configurationAggregatorService.ManualReconcile(manualReconcileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(manualReconcileResponse, "", "  ")
			fmt.Println(string(b))

			// end-manual_reconcile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(manualReconcileResponse).ToNot(BeNil())
		})
	})
})
