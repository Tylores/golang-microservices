#!/bin/bash

go run metadata/cmd/main.go &
go run rating/cmd/main.go &
go run movie/cmd/main.go
