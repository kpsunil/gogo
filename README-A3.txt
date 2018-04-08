CS335: Assignment 2
~~~~~~~~~~~~~~~~~~~

Instructions
~~~~~~~~~~~~
Place the directory "asgn3" inside $GOPATH.

The following should generate relevant binaries inside the directory `bin` -

        make gogo

Alternatively, if "gocc" is installed under $GOPATH then the dependencies can be built separately via -

        make deps

Tests
~~~~~
The test cases are available under "test/codegen" directory, and all the relevant tests are of the form *.go. The IR for a single test of the form *.go can be generated via -

        bin/gogo test/codegen/test1.go
