#!/bin/bash

# shellcheck source=/dev/null
. ./repo/repo.sh --source-only

run "$PROJECT_NAME" "$PROJECT_DESCRIPTION" "$PRIVATE $USERNAME" "$TOKEN"
