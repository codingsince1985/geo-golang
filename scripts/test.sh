#! /usr/bin/env bash
set -e

curdir="$( builtin cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
projectdir="${curdir}/.."

# DVR_ARGS default don't use them
DVR_ARGS=""
DVR_FILE="testdata/{{.Name}}.dvr"

# Record if specified (e.g. we need to run this when we update API servers or tests)
if [[ "$DVR_MODE" = "record" ]]; then
	DVR_ARGS="-dvr.record -dvr.file=${DVR_FILE}"
# Else if explicitly playback mode or if the API keys aren't available
elif [[ "$DVR_MODE" = "replay" || \
						"$GOOGLE_API_KEY" = ""  ||
						"$HERE_APP_ID" = ""  || \
						"$HERE_APP_CODE" = ""  || \
						"$MAPBOX_API_KEY" = "" || \
						"$MAPQUEST_NOMINATUM_KEY" = "" || \
						"$MAPQUEST_OPEN_KEY" = "" \
				]]; then
	DVR_ARGS="-dvr.replay -dvr.file=${DVR_FILE}"
fi

# Uses go list to inspect all packages and dependencies to run full code coverage across all packages in project
go list -f '"go test '"${DVR_ARGS}"' -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg={{range $i, $f := .Deps}}{{if eq (printf "%.37s" $f) "github.com/codingsince1985/geo-golang" }}{{$f}},{{end}}{{end}}{{.ImportPath}} {{.ImportPath}}"' ./... | xargs -I {} bash -c {}
