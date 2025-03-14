//go:build integration

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
		err                            error
		configurationAggregatorService *configurationaggregatorv1.ConfigurationAggregatorV1
		serviceURL                     string
		config                         map[string]string
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

	Describe(`ListConfigs - List of configurations of the resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) with pagination`, func() {
			listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{
				ConfigType:      core.StringPtr("testString"),
				ServiceName:     core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Location:        core.StringPtr("testString"),
				ResourceCrn:     core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(10)),
				Start:           core.StringPtr("testString"),
				SubAccount:      core.StringPtr("testString"),
				AccessTags:      core.StringPtr("role:admin"),
				UserTags:        core.StringPtr("test"),
				ServiceTags:     core.StringPtr("test:tag"),
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

				listConfigsOptions.Start, err = listConfigsResponse.GetNextStart()
				Expect(err).To(BeNil())

				if listConfigsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) using ConfigsPager`, func() {
			listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{
				ConfigType:      core.StringPtr("testString"),
				ServiceName:     core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Location:        core.StringPtr("testString"),
				ResourceCrn:     core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(10)),
				SubAccount:      core.StringPtr("testString"),
				AccessTags:      core.StringPtr("role:admin"),
				UserTags:        core.StringPtr("test"),
				ServiceTags:     core.StringPtr("test:tag"),
			}

			// Test GetNext().
			pager, err := configurationAggregatorService.NewConfigsPager(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []configurationaggregatorv1.Config
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
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

	Describe(`ReplaceSettings - Replace the settings for Configuration Aggregator`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions)`, func() {
			profileTemplateModel := &configurationaggregatorv1.ProfileTemplate{
				ID:               core.StringPtr("ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"),
				TrustedProfileID: core.StringPtr("Profile-39acf232-8969-4c32-9838-83eb60a037f7"),
			}

			additionalScopeModel := &configurationaggregatorv1.AdditionalScope{
				Type:            core.StringPtr("Enterprise"),
				EnterpriseID:    core.StringPtr("2c99aed413954f93b7cf7ce9fda6de61"),
				ProfileTemplate: profileTemplateModel,
			}

			replaceSettingsOptions := &configurationaggregatorv1.ReplaceSettingsOptions{
				ResourceCollectionEnabled: core.BoolPtr(true),
				TrustedProfileID:          core.StringPtr("Profile-1260aec2-f2fc-44e2-8697-2cc15a447560"),
				Regions:                   []string{"all"},
				AdditionalScope:           []configurationaggregatorv1.AdditionalScope{*additionalScopeModel},
			}

			settingsResponse, response, err := configurationAggregatorService.ReplaceSettings(replaceSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSettings - Retrieve the settings for Configuration Aggregator feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &configurationaggregatorv1.GetSettingsOptions{}

			settingsResponse, response, err := configurationAggregatorService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settingsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetResourceCollectionStatus - Retrieve status for resource collection in Configuration Aggregator`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceCollectionStatus(getResourceCollectionStatusOptions *GetResourceCollectionStatusOptions)`, func() {
			getResourceCollectionStatusOptions := &configurationaggregatorv1.GetResourceCollectionStatusOptions{}

			statusResponse, response, err := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statusResponse).ToNot(BeNil())
		})
	})

	Describe(`ManualReconcile - Manually trigger the recording of the Configuration items as part of Configuration Aggregator`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ManualReconcile(manualReconcileOptions *ManualReconcileOptions)`, func() {
			manualReconcileOptions := &configurationaggregatorv1.ManualReconcileOptions{}

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
