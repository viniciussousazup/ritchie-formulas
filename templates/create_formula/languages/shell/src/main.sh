#!/bin/bash

# shellcheck source=/dev/null
. ./hello/hello.sh --source-only
#In sh for receive inputs of CLI use: $SAMPLE_TEXT, $SAMPLE_LIST and $SAMPLE_BOOL for this example

runFormula "$SAMPLE_TEXT" "$SAMPLE_LIST" "$SAMPLE_BOOL"
