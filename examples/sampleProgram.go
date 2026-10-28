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
	"io"
	"os"
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

func replaceSettings(trustedProfileID string, resourceCollectionEnabled bool, realtimeConfigEnabled bool, regions []string) {
	fmt.Println("Executing the Replace Settings function")
	replaceSettingsOptions := configAggregatorInstance.NewReplaceSettingsOptions()
	replaceSettingsOptions.TrustedProfileID = core.StringPtr(trustedProfileID)
	replaceSettingsOptions.ResourceCollectionEnabled = core.BoolPtr(resourceCollectionEnabled)
	replaceSettingsOptions.RealtimeConfigEnabled = core.BoolPtr(realtimeConfigEnabled)
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

func listBatchConfigs(configs []configurationaggregatorv1.RequestConfigs, limit int64, start string) {
	fmt.Println("Executing the List Batch Configs Function")

	// Create options with the array of configs
	listBatchConfigsOptions := configAggregatorInstance.NewListBatchConfigsOptions(configs)
	listBatchConfigsOptions.SetLimit(limit)
	listBatchConfigsOptions.SetStart(start)

	// Call the function
	listBatchConfigsResponse, listBatchConfigsResponseCode, err := configAggregatorInstance.ListBatchConfigs(listBatchConfigsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	// Print results
	listBatchConfigs, _ := json.MarshalIndent(listBatchConfigsResponse, "", "  ")
	fmt.Println(string(listBatchConfigs))
	fmt.Println(listBatchConfigsResponseCode.StatusCode)
}

func nlpQueryConfigs(resourceQuery string, subAccount string, historyID string, limit int64) {
	fmt.Println("Executing the NLP Query Configs Function")
	nlpQueryConfigsOptions := configAggregatorInstance.NewNlpQueryConfigsOptions(resourceQuery, subAccount)
	if historyID != "" {
		nlpQueryConfigsOptions.SetHistoryID(historyID)
	}
	nlpQueryConfigsOptions.SetLimit(limit)
	nlpQueryResponse, nlpQueryResponseCode, err := configAggregatorInstance.NlpQueryConfigs(nlpQueryConfigsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	nlpQuery, _ := json.MarshalIndent(nlpQueryResponse, "", "  ")
	fmt.Println(string(nlpQuery))
	fmt.Println(nlpQueryResponseCode.StatusCode)
}

func listNlpChatHistory() {
	fmt.Println("Executing the List NLP Chat History Function")
	listNlpChatHistoryOptions := configAggregatorInstance.NewListNlpChatHistoryOptions()
	listChatHistoryResponse, listChatHistoryResponseCode, err := configAggregatorInstance.ListNlpChatHistory(listNlpChatHistoryOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	listChatHistory, _ := json.MarshalIndent(listChatHistoryResponse, "", "  ")
	fmt.Println(string(listChatHistory))
	fmt.Println(listChatHistoryResponseCode.StatusCode)
}

func createNlpChatHistory(title string, query string) *string {
	fmt.Println("Executing the Create NLP Chat History Function")
	createNlpChatHistoryOptions := configAggregatorInstance.NewCreateNlpChatHistoryOptions(title, query)
	chatHistoryResponse, chatHistoryResponseCode, err := configAggregatorInstance.CreateNlpChatHistory(createNlpChatHistoryOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	}
	chatHistory, _ := json.MarshalIndent(chatHistoryResponse, "", "  ")
	fmt.Println(string(chatHistory))
	fmt.Println(chatHistoryResponseCode.StatusCode)
	return chatHistoryResponse.ID
}

func updateNlpChatHistory(id string, title string, query string, deleted bool) {
	fmt.Println("Executing the Update NLP Chat History Function")
	updateNlpChatHistoryOptions := configAggregatorInstance.NewUpdateNlpChatHistoryOptions(id)
	updateNlpChatHistoryOptions.SetTitle(title)
	updateNlpChatHistoryOptions.SetQuery(query)
	updateNlpChatHistoryOptions.SetDeleted(deleted)
	chatHistoryResponse, chatHistoryResponseCode, err := configAggregatorInstance.UpdateNlpChatHistory(updateNlpChatHistoryOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	chatHistory, _ := json.MarshalIndent(chatHistoryResponse, "", "  ")
	fmt.Println(string(chatHistory))
	fmt.Println(chatHistoryResponseCode.StatusCode)
}

func listNlpResourceQueryJobs(subAccount string, status string, limit int64, offset int64) {
	fmt.Println("Executing the List NLP Resource Query Jobs Function")
	listNlpResourceQueryJobsOptions := configAggregatorInstance.NewListNlpResourceQueryJobsOptions()
	listNlpResourceQueryJobsOptions.SetSubAccount(subAccount)
	listNlpResourceQueryJobsOptions.SetStatus(status)
	listNlpResourceQueryJobsOptions.SetLimit(limit)
	listNlpResourceQueryJobsOptions.SetOffset(offset)
	jobsResponse, jobsResponseCode, err := configAggregatorInstance.ListNlpResourceQueryJobs(listNlpResourceQueryJobsOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	jobs, _ := json.MarshalIndent(jobsResponse, "", "  ")
	fmt.Println(string(jobs))
	fmt.Println(jobsResponseCode.StatusCode)
}

func createNlpResourceQueryJob(resourceQuery string, subAccount string, historyID string, fileType string) *string {
	fmt.Println("Executing the Create NLP Resource Query Job Function")
	createNlpResourceQueryJobOptions := configAggregatorInstance.NewCreateNlpResourceQueryJobOptions(resourceQuery, subAccount, historyID, fileType)
	jobResponse, jobResponseCode, err := configAggregatorInstance.CreateNlpResourceQueryJob(createNlpResourceQueryJobOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	}
	job, _ := json.MarshalIndent(jobResponse, "", "  ")
	fmt.Println(string(job))
	fmt.Println(jobResponseCode.StatusCode)
	return jobResponse.JobID
}

func getNlpResourceQueryJob(jobID string) {
	fmt.Println("Executing the Get NLP Resource Query Job Function")
	getNlpResourceQueryJobOptions := configAggregatorInstance.NewGetNlpResourceQueryJobOptions(jobID)
	jobResponse, jobResponseCode, err := configAggregatorInstance.GetNlpResourceQueryJob(getNlpResourceQueryJobOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	job, _ := json.MarshalIndent(jobResponse, "", "  ")
	fmt.Println(string(job))
	fmt.Println(jobResponseCode.StatusCode)
}

func downloadNlpJobResult(jobID string, outputFile string) {
	fmt.Println("Executing the Download NLP Job Result Function")
	jobResponse, _, err := configAggregatorInstance.GetNlpResourceQueryJob(configAggregatorInstance.NewGetNlpResourceQueryJobOptions(jobID))
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	if jobResponse.Status == nil || *jobResponse.Status != configurationaggregatorv1.JobDetailResponse_Status_Completed {
		status := "<nil>"
		if jobResponse.Status != nil {
			status = *jobResponse.Status
		}
		fmt.Println("Job is not ready for download. Current status: " + status)
		return
	}
	downloadNlpJobResultOptions := configAggregatorInstance.NewDownloadNlpJobResultOptions(jobID)
	result, downloadResponseCode, err := configAggregatorInstance.DownloadNlpJobResult(downloadNlpJobResultOptions)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	if result == nil {
		return
	}
	defer result.Close()
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	defer outFile.Close()
	if _, err = io.Copy(outFile, result); err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println("Job result written to " + outputFile)
	fmt.Println(downloadResponseCode.StatusCode)
}

func main() {

	authToken := "<authToken>"
	guid := "<guid>"
	region := "<region>"
	resourceCrn := "<resourceCrn>"
	subAccount := "<accountid>"
	initAndReturnSingletonInstanceWithAPIKey(authToken, guid, region)
	replaceSettings("your_trusted_profile_id", true, false, []string{"all"})
	replaceSettings("your_trusted_profile_id", true, true, []string{"all"})
	getSettings()
	getResourceCollectionStatus()
	ManualConfigReconcile()
	listConfigs()
	listBatchConfigs([]configurationaggregatorv1.RequestConfigs{{ResourceCrn: core.StringPtr(resourceCrn)}}, 100, "page token")

	nlpQuery := "Show me all resources"
	nlpQueryConfigs(nlpQuery, subAccount, "", 10)
	chatID := createNlpChatHistory("Chat 1", nlpQuery)
	listNlpChatHistory()
	if chatID != nil {
		updateNlpChatHistory(*chatID, "Chat 1 updated", nlpQuery, false)

		jobID := createNlpResourceQueryJob(nlpQuery, subAccount, *chatID, "json")
		listNlpResourceQueryJobs(subAccount, "COMPLETED", 10, 0)
		if jobID != nil {
			getNlpResourceQueryJob(*jobID)
			downloadNlpJobResult(*jobID, "nlp_job_result.json")
		}
	}
}
