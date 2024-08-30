//go:build examples

/**
 * (C) Copyright IBM Corp. 2024.
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
	"os"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
var _ = Describe(`ConfigurationAggregatorV1 Examples Tests`, func() {

	const externalConfigFile = "../configuration_aggregator_v1.env"

	var (
		configurationAggregatorService *configurationaggregatorv1.ConfigurationAggregatorV1
		config                         map[string]string
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
		It(`ListConfigs request example`, func() {
			fmt.Println("\nListConfigs() result:")
			// begin-list_configs
			listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{
				ConfigType:      core.StringPtr("testString"),
				ServiceName:     core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Location:        core.StringPtr("testString"),
				ResourceCrn:     core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(10)),
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
		It(`ReplaceSettings request example`, func() {
			fmt.Println("\nReplaceSettings() result:")
			// begin-replace_settings

			replaceSettingsOptions := configurationAggregatorService.NewReplaceSettingsOptions()

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
	})
})
