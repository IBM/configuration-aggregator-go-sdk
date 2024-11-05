Build Status
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Configuration Aggregator Go SDK Version 0.0.2
Go client library to interact with the various IBM Cloud Configuration Aggregator API SDK APIs.

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.
## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Configuration Aggregator Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
Configuration Aggregator | configurationaggregatorv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.20 or above.

## Installation
The current version of this SDK: 0.0.2

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and replace your application's
`go.mod` file.  

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/configuration-aggregator-go-sdk
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
Examples

Construct a service client and use it to create, retrieve and manage resources from your App Configuration instance.

Here's an example ```main.go``` file:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
)

func main() {
	authenticator := &core.IamAuthenticator{
		ApiKey: "<IBM_CLOUD_API_KEY>",
		URL:    "https://iam.cloud.ibm.com",
	}
	instanceID := "provide_your_instance_id"
	region:="region_of_the_url"
	configurationAggregatorServiceOptions := &configurationaggregatorv1.ConfigurationAggregatorV1Options{
		Authenticator: authenticator,
		URL:           "https://" + region + ".apprapp.cloud.ibm.com/apprapp/config_aggregator/v1/instances/" + instanceID,
	}
	configurationAggregatorService, err := configurationaggregatorv1.NewConfigurationAggregatorV1(configurationAggregatorServiceOptions)

	if err != nil {
		panic(err)
	}

	replaceSettingsOptions := configurationAggregatorService.NewReplaceSettingsOptions()
	replaceSettingsOptions.TrustedProfileID = core.StringPtr("your_trusted_profile_id")
	replaceSettingsOptions.ResourceCollectionEnabled = core.BoolPtr(true)
	replaceSettingsOptions.Regions = []string{"all"}
	replaceSettingsResponse, _, err := configurationAggregatorService.ReplaceSettings(replaceSettingsOptions)
	if err != nil {
		panic(err)
	}
	replaceSettings, _ := json.MarshalIndent(replaceSettingsResponse, "", "  ")
	fmt.Println(string(replaceSettings))
	getSettingsOptions := configurationAggregatorService.NewGetSettingsOptions()
	getSettingsResponse, _, err := configurationAggregatorService.GetSettings(getSettingsOptions)
	if err != nil {
		panic(err)
	}
	getSettings, _ := json.MarshalIndent(getSettingsResponse, "", "  ")
	fmt.Println(string(getSettings))
	getResourceCollectionStatusOptions := configurationAggregatorService.NewGetResourceCollectionStatusOptions()
    getResourceCollectionStatusResponse, _, err := configurationAggregatorService.GetResourceCollectionStatus(getResourceCollectionStatusOptions)
    if err != nil {
        panic(err)
    }
    getResourceCollectionStatus, _ := json.MarshalIndent(getResourceCollectionStatusResponse, "", "  ")
    fmt.Println(string(getResourceCollectionStatus))
	listConfigsOptions := configurationAggregatorService.NewListConfigsOptions()

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
	listConfigs, _ := json.MarshalIndent(listConfigsResponse, "", "  ")
	fmt.Println(string(listConfigs))

}



```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](<github-repo-url>/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
