# minutil - 一个通用的 Go 工具包

`minutil` 是一个通用的 Go 工具包，提供了一些常用的功能，包括 Gin 框架的统一响应封装、GORM 数据库操作的泛型封装等。

## 功能概述

- **Gin 统一响应封装**: 提供了 `OK` 和 `Err` 函数，用于生成统一的 JSON 响应。
- **GORM 数据库操作封装**: 提供了泛型的数据库操作函数，如 `GetOne`、`GetAll`、`Create`、`Update`、`Delete`、`Like` 和 `Search`。

## 安装

使用 `go get` 命令安装 `minutil` 包：

```bash
go get github.com/yowaimono/min-util
```