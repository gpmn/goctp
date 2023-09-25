
install:

	go install -v -x -a -buildmode=shared runtime sync/atomic
	go install -v -x -a -buildmode=archive runtime sync/atomic
#	go install -v -x -work -buildmode=shared -linkshared

build:
#	go build -v -x -linkshared -o ctp ./demo/*.go
	go build -v -x -work -o ctp ./demo/*.go

gentmp:
	swig -go -intgosize 32   -c++  goctp.swigcxx
