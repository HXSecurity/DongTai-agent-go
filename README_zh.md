## DongTai-agent-go
---
[English version](README.md)

[![license Apache-2.0](https://img.shields.io/github/license/HXSecurity/DongTai-agent-go)](https://github.com/HXSecurity/DongTai-agent-go/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/HXSecurity/DongTai-agent-go.svg?label=Stars&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)
[![GitHub forks](https://img.shields.io/github/forks/HXSecurity/DongTai-Agent-Go?label=Forks&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)
[![GitHub Contributors](https://img.shields.io/github/contributors-anon/HXSecurity/DongTai-agent-go?label=Contributors&logo=github)](https://github.com/HXSecurity/DongTai-agent-go)


[![Github Version](https://img.shields.io/github/v/release/HXSecurity/DongTai-agent-go?display_name=tag&include_prereleases&sort=semver)](https://github.com/HXSecurity/DongTai-agent-go/releases)
[![Release downloads](https://shields.io/github/downloads/HXSecurity/DongTai-Agent-go/total)](https://github.com/HXSecurity/DongTai-agent-go/releases)



## 项目介绍

DongTai-agent-go 是**洞态IAST** 针对 Go 应用开发的数据采集端。在添加 iast-agent 代理的 Go 应用中，通过改写汇编地址的方式采集所需数据，然后将数据发送至 DongTai-openapi 服务，再由云端引擎处理数据判断是否存在安全漏洞。

DongTai-agent-go 由`core`、`run`、`service`三个主要部分构成，其中：

- `run`用来按需运行需要插装的包的agent
- `core`是核心包，其主要功能是：字节码插桩、数据采集、数据预处理、数据上报、第三方组件管理等。
- `service`用于获取应用发送的请求以及收到的响应，用于数据展示以及请求重放功能。

## 应用场景

- DevOps流程
- 上线前安全测试
- 第三方组件管理
- 代码审计
- 0 Day挖掘


## 快速上手

### 快速使用

请参考：[快速开始](https://doc.dongtai.io)

### 快速开发

1. Fork [DongTai-agent-go](https://github.com/HXSecurity/DongTai-agent-go) 项目到自己的github仓库并 clone 项目：

   ```shell
   git clone https://github.com/<your-username>/DongTai-agent-go
   ```

2. 根据需求编写代码

3. 贡献代码。如果您想要向洞态 IAST 团队贡献代码，请阅读完整的[贡献指南](https://github.com/HXSecurity/DongTai/blob/main/CONTRIBUTING.md)

#### 支持的Go版本及中间件

- Go 1.11+
- Gin,Gorm等主流软件和中间件
