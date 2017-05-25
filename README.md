A Tour of Go is an introduction to the Go programming language.

To install the tour from source, first 
[set up a workspace](https://golang.org/doc/code.html) and then run:

```
  $ go get golang.org/x/tour/gotour
  $ cd $GOPATH/src/golang.org/x/tour
  $ git remote add vaskoz https://github.com/vaskoz/tour.git
  $ git fetch --all
  $ git checkout -b extended vaskoz/master
  $ cd gotour
  $ go install
  $ gotour # this will start this extended 'gotour' application
```

Unless otherwise noted, the go-tour source files are distributed
under the BSD-style license found in the LICENSE file.

Contributions should follow the same procedure as for the Go project:
http://golang.org/doc/contribute.html
