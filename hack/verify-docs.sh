#!/usr/bin/env bash
#
# Verify that generated Markdown docs are up-to-date.
#

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

pushd "${PROJECT_ROOT}"
trap popd EXIT

tmpdir=$(mktemp -d)
go run cmd/help/main.go --dir "$tmpdir"
diff -Naur -x 'testing.md' "$tmpdir" docs/
