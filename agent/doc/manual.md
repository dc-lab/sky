## Sky Agent User Manual

<!-- vim-markdown-toc GFM -->

* [Installation](#installation)
    * [System Requirements](#system-requirements)
    * [Downloading Agent](#downloading-agent)
    * [Upgrading Agent](#upgrading-agent)
    * [Downloading agent external modules](#downloading-agent-external-modules)
* [Configuration](#configuration)
* [Quick Start](#quick-start)
* [Running Agent](#running-agent)

<!-- vim-markdown-toc -->

### Installation

#### System Requirements

Everest agent is a Go application supporting Linux and Docker Images.

The system should have `go` installed.
Supported `go` version: 1.13+ 

#### Downloading Agent

The recommended way to download the agent by cloning sky repository:
```
git clone https://github.com/dc-lab/sky.git
```
Agent source folder is `sky/agent/src`

#### Upgrading Agent

* Run `git pull` inside the directory where repository was downloaded.

#### Downloading agent external modules

Sky agent use several external Golang modules from github repositories. \
To download them all you need to run following command in source directory:
```
go mod tidy
```
It will download all modules from `go.mod` file.

### Configuration

The agent configuration is stored in a plain text file using a JSON format.

Path to the configuration file must be specified via a command line option `-c (--config)` when starting the agent.

### Quick Start
* Create and open your configuration file in a text editor. You can find an example in source directory.

* You can configure the following parameters:
    * *ResourceManagerAddress* – IPv4/IPv6 address of running resource manager
    * *AgentDirectory* – path to directory  in which the agent will execute the launched tasks for which a separate subfolder will be created
    * *TokenPath* – path to the file with the registration token. You must register your computing resource in resource manager. You will get registration token which must be written to file with this path
    * *LogsDirectory* – path to the directory with agent logs
    * *MaxCacheSize* -- bytes restriction on local cache folder capacity.

### Running Agent

To start the agent run the following command in agent sources directory:
`go run main.go`

The following optional command-line parameters are supported:

| Option             | Description |
|--------------------|----------------------------------------------|
| `-c`, `--config`   |  Custom path to the agent configuration file.        |
