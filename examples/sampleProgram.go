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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var configAggregatorInstance *configurationaggregatorv1.ConfigurationAggregatorV1

func initAndReturnSingletonInstanceWithAPIKey(authToken string, guid string, region string) *configurationaggregatorv1.ConfigurationAggregatorV1 {
	instanceURL := "https://" + region + ".apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/" + guid
	var once sync.Once
	if configAggregatorInstance == nil {
		once.Do(func() {
			if configAggregatorInstance == nil {
				authenticator := &core.IamAuthenticator{
					ApiKey: authToken,
					URL:    "https://iam.cloud.ibm.com",
				}
				options := &configurationaggregatorv1.ConfigurationAggregatorV1Options{
					Authenticator: authenticator,
					URL:           instanceURL,
				}
				var error error
				configAggregatorInstance, error = configurationaggregatorv1.NewConfigurationAggregatorV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}

			}
		})
	}
	return configAggregatorInstance
}

func initAndReturnSingletonInstanceWithBearertoken(authToken string, guid string, region string) *configurationaggregatorv1.ConfigurationAggregatorV1 {
	instanceURL := "https://" + region + ".apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/" + guid
	var once sync.Once
	if configAggregatorInstance == nil {
		once.Do(func() {
			if configAggregatorInstance == nil {
				authenticator := &core.BearerTokenAuthenticator{
					BearerToken: authToken,
				}
				options := &configurationaggregatorv1.ConfigurationAggregatorV1Options{
					Authenticator: authenticator,
					URL:           instanceURL,
				}
				var error error
				configAggregatorInstance, error = configurationaggregatorv1.NewConfigurationAggregatorV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}

			}
		})
	}

	return configAggregatorInstance
}

func listConfigs() {
	fmt.Println("Executing the List Configs Function")
	listConfigsOptions := configAggregatorInstance.NewListConfigsOptions()
	listConfigsResponse, listconfigsRespnseCode, err := configAggregatorInstance.ListConfigs(listConfigsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	listConfigs, _ := json.MarshalIndent(listConfigsResponse, "", "  ")
	fmt.Println(string(listConfigs))
	fmt.Println(listconfigsRespnseCode.StatusCode)

}

func replaceSettings(trustedProfileID string, resourceCollectionEnabled bool, regions []string) {
	fmt.Println("Executing the Replace Settings function")
	replaceSettingsOptions := configAggregatorInstance.NewReplaceSettingsOptions()
	replaceSettingsOptions.TrustedProfileID = core.StringPtr(trustedProfileID)
	replaceSettingsOptions.ResourceCollectionEnabled = core.BoolPtr(resourceCollectionEnabled)
	replaceSettingsOptions.Regions = regions
	replaceSettingsResponse, replaceSettingsResponseCode, err := configAggregatorInstance.ReplaceSettings(replaceSettingsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	replaceSettings, _ := json.MarshalIndent(replaceSettingsResponse, "", "  ")
	fmt.Println(string(replaceSettings))
	fmt.Println(replaceSettingsResponseCode.StatusCode)

}

func getSettings() {
	fmt.Println("Executing the List Settings Function")
	getSettingsOptions := configAggregatorInstance.NewGetSettingsOptions()
	getSettingsResponse, getSettingsResponsecode, err := configAggregatorInstance.GetSettings(getSettingsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	getSettings, _ := json.MarshalIndent(getSettingsResponse, "", "  ")
	fmt.Println(string(getSettings))
	fmt.Println(getSettingsResponsecode.StatusCode)
}

func getResourceCollectionStatus() {
	fmt.Println("Executing the Get Resource Collection Status Function")
	getResourceCollectionStatusOptions := configAggregatorInstance.NewGetResourceCollectionStatusOptions()
	getResourceCollectionStatusResponse, getResourceCollectionStatusResponseCode, err := configAggregatorInstance.GetResourceCollectionStatus(getResourceCollectionStatusOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	getResourceCollectionStatus, _ := json.MarshalIndent(getResourceCollectionStatusResponse, "", "  ")
	fmt.Println(string(getResourceCollectionStatus))
	fmt.Println(getResourceCollectionStatusResponseCode.StatusCode)
}

func ManualConfigReconcile() {
	fmt.Println("Executing the Manual Configuration Reconcile Function")
	manualConfigReconcileOptions := configAggregatorInstance.NewManualReconcileOptions()
	manualConfigReconcileResponse, manualConfigReconcileResponseCode, err := configAggregatorInstance.ManualReconcile(manualConfigReconcileOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	ManualConfigReconcile, _ := json.MarshalIndent(manualConfigReconcileResponse, "", "  ")
	fmt.Println(string(ManualConfigReconcile))
	fmt.Println(manualConfigReconcileResponseCode.StatusCode)
}

func main() {

	authToken := "<authToken>"
	guid := "<guid>"
	region := "<region>"
	initAndReturnSingletonInstanceWithAPIKey(authToken, guid, region)
	replaceSettings("your_trusted_profile_id", true, []string{"all"})
	getSettings()
	getResourceCollectionStatus()
	ManualConfigReconcile()
	listConfigs()
}
