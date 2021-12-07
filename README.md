# talk2frog

## About this plugin
This plugin brings the power of natural language processing to the JFrog command line. It transforms descriptions of command line tasks in English to their jfrog CLI syntax

## Installation with JFrog CLI
### Environment installation
Make sure you have python 3 installed on the machine and accessible. Install the nlc2cmd requirements:

`pip install -r nlc2cmd/requirements.txt`

Train the model as described in the nlc2cmd [readme](https://github.com/cohmoti/Magnum-NLC2CMD/blob/feature/jfrog-cli/README.md). You can also download a pretrained [model](https://drive.google.com/file/d/1P-59TwoBIWc-nNvCbEgudc4ZSZ4HrTvN/view?usp=sharing) and place in nlc2cmd/src/model/run.

Define and environment variable TALK2FROG_MODEL_HOME to point to the nlc2cmd path (jfrog-cli-nlc/nlc2cmd).

### Plugin installation
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
  Translating to command ...
  Result: jfrog xr ago --watches watch1
  Would you like to execute it now? (y/n) [n]?

  ```

## Additional info
This plugin uses a learning module named nlc2cmd. We forked the project's repository and added the things that are relevant for learning jfrog cli specific commands. This module is based on the transformer architecture for sequence to sequence translation. 

## Release Notes
The release notes are available [here](RELEASE.md).
