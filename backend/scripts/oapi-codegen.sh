#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/.."

oapi-codegen --package api --generate server -o "$PROJECT_DIR/api/server_gen.go" "$PROJECT_DIR/docs/api.yaml"
oapi-codegen --package api --generate spec -o "$PROJECT_DIR/api/spec_gen.go" "$PROJECT_DIR/docs/api.yaml"
oapi-codegen --package api --generate types -o "$PROJECT_DIR/api/types_gen.go" "$PROJECT_DIR/docs/api.yaml"