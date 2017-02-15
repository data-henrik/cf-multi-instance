# cf-multi-instance
Plugin for Cloud Foundry CLI to obtain status from multiple instances, e.g., several Bluemix Public regions


### Installation
At this time you need to download and build the plugin on your own (no help offered). I am in the process of building plugin binaries that could be directly installed using the `cf` plugin-related commands. 

Right now, you can download the binary for linux86 and install it with `cf install-plugin pathToBinary/multi-instance`.

### Configuration
A file `miconfig.yml` needs to be placed in the default configuration path for the cf CLI environment (typically `~/.cf`). The content is based on YAML and has the following structure:   
```
Name: someName
Identifier: 1
APIs:
  - https://api.your-cf-instance1.com
  - https://api.your-cf-instance2.com
```   

`Name` is used to identify the configuration (not yet used). The `Identifier` determines which of the provided API endpoints is used to match the configuration against the current cf CLI environment. If you are logged in to a different Cloud Foundry instance then the configured one, the plugin obviously is not going to work. The URLs specified under `APIs` are those queried by the plugin. 

An example for two public regions of [IBM Bluemix](http://www.ibm.com/cloud-computing/bluemix/) is provided in the file [`miconfig.yml.sample`](miconfig.yml.sample).

### Usage
The installed plugin is identified as `multi-instance`. It can be used with either the `multi-instance` command or the short version `mi`:   
```
cf mi [apps | orgs]
```   
If no parameter is used, the configured API endpoints (instances) are printed. The parameter `apps` lets the plugin to return status information for all apps visible to the user across the configured the instances. With the parameter `orgs` the available organisations for each of the instances are shown:   
```
[henrik@machine]$ cf mi orgs
Endpoint:  https://api.your-cf-instance1.com
Org 0: data-henrik
Org 1: BluemixSamples
Org 2: PluginDev
Endpoint:  https://api.your-cf-instance2.com
Org 0: data-henrik
Org 1: TestOrg
```   
# Useful Links
Here are some useful links for developing plugins on your own:
* Using cf CLI Plugins: http://docs.cloudfoundry.org/cf-cli/use-cli-plugins.html
* Developing cf CLI Plugins: http://docs.cloudfoundry.org/cf-cli/develop-cli-plugins.html
* Cloud Foundry plugin registry: https://plugins.cloudfoundry.org/
* IBM Bluemix plugin registry: http://plugins.ng.bluemix.net/ui/repository.html
* Repository for cf CLI, including plugin API: https://github.com/cloudfoundry/cli
* Go cf client (used by this plugin): https://github.com/cloudfoundry-community/go-cfclient

# License
See [LICENSE](LICENSE) for license information.

# Contact Information
If you have found errors or some instructions are not working anymore, then please open an GitHub issue or, better, create a pull request with your desired changes.

You can find more tutorials and sample code at:
https://ibm-bluemix.github.io/
