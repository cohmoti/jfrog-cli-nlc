# talk2frog

## About this plugin
This plugin brings the power of natural language processing to the JFrog command line. It transforms descriptions of command line tasks in English to their jfrog CLI syntax

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install talk2frog`

Installing a specific version:

`$ jfrog plugin install talk2frog@version`

Uninstalling a plugin

`$ jfrog plugin uninstall talk2frog`

## Usage
### Commands
* do
    - Argument:
        Quoted English command description
    - Example:
    ```
  $ jfrog talk2frog do "Audit the Go project at the current directory using the watch1 watch defined in Xray"
  
  jfrog xr ago --watches "watch1"
  ```

## Additional info
This plugin uses a learning module named nlc2cmd. We forked the project's repository and added the things that are relevant for learning jfrog cli specific commands. This module is based on the transformer architecture for sequence to sequence translation. 

## Release Notes
The release notes are available [here](RELEASE.md).
