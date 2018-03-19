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
The test cases are available under "tests/" directory, and all the relevant tests are of the form *.ir. The assembly equivalent of these tests can be generated via -

        make test

The MIPS assembly for a single test of the form *.ir can be generated via -

        bin/gogo test/test1.ir
