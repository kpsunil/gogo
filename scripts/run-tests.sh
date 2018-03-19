#!/usr/bin/env bash
# Script for running tests.

set -euo pipefail

srcDir=$(dirname "$0")/..
binDir="$srcDir/bin"
testDir="$srcDir/test"
executable="parser"
testName=

checkBuildStatus() {
    if [ ! -e "$binDir/$executable" ]; then
	# http://mywiki.wooledge.org/BashFAQ/105
	# http://fvue.nl/wiki/Bash:_Error_handling
	( cd "$srcDir" && make )
    fi
}

# runIRTests generates MIPS assembly files of the form:
#	"test(i)ir.asm", i ∈ Z
# from corresponding test files (in $testdir) of the form:
#	"test(i).ir", i ∈ Z
runIRTests() {
    for f in "$testDir/ir"/*.ir; do
	# Remove everything after and including the last '.'
	testName=$(echo "$f" | sed -E 's/(.*)\.(.*)/\1/')
	rm -f "$testName.asm"
	"$binDir/$executable" "$f" > "$testName.asm"
    done
}

runParserTests() {
    for f in "$testDir/parser"/*.go; do
    	echo "$f"
	# Remove everything after and including the last '.'
	testName=$(echo "$f" | sed -E 's/(.*)\.(.*)/\1/')
	rm -f "$testName.html"
	"$binDir/$executable" "$f" > "$testName.html"
    done
}

checkBuildStatus
runParserTests
