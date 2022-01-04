## DongTai-agent-go
---
[中文文档](README_zh.md)

[![license Apache-2.0](https://img.shields.io/github/license/HXSecurity/DongTai-agent-go)](https://github.com/HXSecurity/DongTai-agent-go/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/HXSecurity/DongTai-agent-go.svg?label=Stars&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)
[![GitHub forks](https://img.shields.io/github/forks/HXSecurity/DongTai-Agent-Go?label=Forks&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)
[![GitHub Contributors](https://img.shields.io/github/contributors-anon/HXSecurity/DongTai-agent-go?label=Contributors&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)

[![Github Version](https://img.shields.io/github/v/release/HXSecurity/DongTai-agent-go?display_name=tag&include_prereleases&sort=semver)](https://github.com/HXSecurity/DongTai-agent-go/releases)
[![Release downloads](https://shields.io/github/downloads/HXSecurity/DongTai-Agent-go/total)](https://github.com/HXSecurity/DongTai-agent-go/releases)




## Project Introduction

DongTai-agent-go is the data collection terminal developed by **Dongtai IAST** for Go applications. In the Go application with the iast-agent agent, the required data is collected by rewriting the assembly address, and then the data is sent to the DongTai-openapi service, and the cloud engine processes the data to determine whether there are security vulnerabilities.

DongTai-agent-go is composed of three main parts: `core`, `run`, and `service`, among which:

`run` is used to run the agent of the package that needs to be instrumented on demand

`core` is the core package, and its main functions are: bytecode instrumentation, data collection, data preprocessing, data reporting, third-party component management, etc.

`service` is used to obtain the request sent by the application and the response received, for data display and request replay function.

## Application scenario

DevOps process

Safety test before going live

Third-party component management

Code audit

0 Day mining

## Get started quickly

### Quick use

Please refer to: [Quick Start](https://doc.dongtai.io)

### Rapid development

1. Fork [DongTai-agent-go](https://github.com/HXSecurity/DongTai-agent-go) project to your github repository and clone the project:

   ```shell
   git clone https://github.com/<your-username>/DongTai-agent-go
   ```

2. Write code according to requirements

3. Contribute code. If you want to contribute code to the Dongtai IAST team, please read the complete [Contribution Guide](https://github.com/HXSecurity/DongTai/blob/main/CONTRIBUTING.md)

#### Supported Go versions and middleware

Go 1.11+

Gin, Gorm and other mainstream software and middleware 
