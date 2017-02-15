// Plugin "multi-instance" to obtain status from multiple Cloud Foundry
// instances, e.g., from several Bluemix Public regions.
//
// (C) 2017 IBM - Licensed under Apache License, Version 2.0
// Author: data-henrik (Henrik Loeser)
//
//


package main

import (
	"fmt"
	"io/ioutil"
  "path/filepath"
  "gopkg.in/yaml.v2"
	"code.cloudfoundry.org/cli/plugin"
  cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry/cli/cf/configuration/confighelpers"
)

type instanceConfig struct {
	 Name       string   `yaml:"Name"`
	 Identifier int      `yaml:"Identifier"`
	 APIs       []string `yaml:"APIs"`
}

// MultiInstance holds info for this plugin
type MultiInstance struct{
	iConfig        instanceConfig
	clients        [5]*cfclient.Client
}




// Run an invocation of this plugin
func (c *MultiInstance) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "multi-instance" || args[0] == "mi" {
    cfConfigFile, err := confighelpers.DefaultFilePath()
		if err != nil {
		   panic(err)
		}
		// We assume a config "miconfig.yml" in the default cf cli config path
		filename := filepath.Join(filepath.Dir(cfConfigFile),"miconfig.yml")
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			 fmt.Println("Error: Expected miconfig.yml in default cf configuration path!")
		   return
		}

		err = yaml.Unmarshal(yamlFile, &c.iConfig)
    if err != nil {
        panic(err)
    }

		li, err := cliConnection.IsLoggedIn()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			if li != true {
				fmt.Println("You are not logged in to a Cloud Foundry instance!")
				fmt.Println("Please log in and retry.")
				return
			}
		}
		endpoint, err := cliConnection.ApiEndpoint()
		if err != nil {
			 panic(err)
		}

		// Check that the configured and current endpoint match
		if c.iConfig.APIs[c.iConfig.Identifier-1]!=endpoint{
			fmt.Println("Expected endpoint does not match current endpoint!")
			fmt.Println("Configured: ",c.iConfig.APIs[c.iConfig.Identifier-1])
			fmt.Println("Current:    ", endpoint)
			return
		}

		un, err := cliConnection.Username()
		if err != nil {
			 panic(err)
		}
		//fmt.Println("You are ", un)

		// Get and prepare the access token
		token, err := cliConnection.AccessToken()
		if err != nil {
			 panic(err)
		}
		// Cut of the "bearer" prefix
		stoken := token[7:len(token)]

		// Now initialize the clients for each configured API endpoint
		for i:=0; i<len(c.iConfig.APIs);i++{
			configTemp := &cfclient.Config{
		 	  ApiAddress:   c.iConfig.APIs[i],
		 		Username:     un,
		 		Token:        stoken,
		 		ClientSecret: "",
		 		ClientID:     "cf"}
				if c.clients[i], err = cfclient.NewClient(configTemp); err != nil {
	  		 	fmt.Println(err)
	  		 	panic(err)
	  		}
			}

      // Individual actions based on Command
			// - they should be moved to functions
			// - and fetching results from each endpoint parallized
			// - could sort results
			switch {
			case len(args) > 1 && args[1] == "orgs":
				for i:=0; i<len(c.iConfig.APIs); i++ {
					orgs, err := c.clients[i].ListOrgs()
					if err != nil {
						panic(err)
					}
					fmt.Println("Endpoint: ",c.iConfig.APIs[i])
					for j := 0; j < len(orgs); j++ {
						fmt.Printf("Org %d: %v\n", j, orgs[j].Name)
					}
				}
			case len(args) > 1 && args[1] == "apps":
				for i:=0; i<len(c.iConfig.APIs); i++ {
					apps, err := c.clients[i].ListApps()
					if err != nil {
						panic(err)
					}
					fmt.Println("Endpoint: ",c.iConfig.APIs[i])
					for j := 0; j < len(apps); j++ {
						fmt.Printf("App %d: %v %v\n", j, apps[j].State, apps[j].Name)
					}
				}

				default:
				fmt.Println("You have the following endpoints configured:")
				for i:=0; i<len(c.iConfig.APIs); i++ {
					fmt.Println("Endpoint: ",c.iConfig.APIs[i])
				}
			}


	}
}

// GetMetadata provides info on our plugin
func (c *MultiInstance) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "multi-instance",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 2,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "multi-instance",
				HelpText: "Get status from multiple instances",
				UsageDetails: plugin.Usage{
					Usage: "multi-instance \n   cf multi-instance [apps | orgs]",
				},
			},
			{
					Name:     "mi",
					HelpText: "Get status from multiple instances",
					UsageDetails: plugin.Usage{
						Usage: "mi - multi-instance\n   cf mi [apps | orgs]",
					},
			},
		},
	}
}

// Initialize this plugin
func main() {
	plugin.Start(new(MultiInstance))
}
