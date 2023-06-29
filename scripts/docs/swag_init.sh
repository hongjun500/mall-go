#!/bin/bash
# 生成 mall_admin 的 swagger 文档
swag init --instanceName mall_admin -g ./cmd/main.go --exclude ./internal/routers/r_mall_search/ -o docs/mall_admin

# 生成 mall_search 的 swagger 文档

swag init --instanceName mall_search -g ./cmd/main.go --exclude ./internal/routers/r_mall_admin/ -o docs/mall_search

# Path: scripts\sql-script\swag_init.sh