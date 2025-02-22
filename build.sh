#!/bin/bash
go mod tidy
go build -o health_check main.go
chmod +x health_check 