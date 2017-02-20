# cf-multi-instance
Plugin for Cloud Foundry CLI to obtain status from multiple instances, e.g., several Bluemix Public regions

The development of this plugin was motivated by an interest in how cf CLI plugins work and their potential. See these blog posts for some related reading:
* Extend the Bluemix CLI Through Plugins: http://blog.4loeser.net/2016/10/extend-bluemix-cli-through-plugins.html
* Extend the Bluemix Command Line Interface Through Plugins: https://www.ibm.com/blogs/bluemix/2016/11/extend-bluemix-command-line-interface-plugins/
* Write Your Own CLI Plugins for Bluemix Cloud Foundry: http://blog.4loeser.net/2017/02/write-your-own-cli-plugins-for-bluemix.html

### Installation & Uninstallation
I created a releas with binaries for Linux64, OSX and Win64. Download the binary for your platform. They are named after the supported platform, e.g. `mi.linux64`. Then, in a command shell with the Cloud Foundry cf CLI present, change to the directory whith the binary.
The following command with install (register) the plugin. Replace `linux64` with the name for your platform:
 ```
 cf install-plugin mi.linux64
 ```   

If you want to remove the plugin, use the following command:
 ```
cf uninstall-plugin multi-instance
 ```   
 
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
