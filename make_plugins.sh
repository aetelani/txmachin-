#!/bin/sh
ROOTDIR=~/src/gocode/src/txmachinae
WORKERS=$ROOTDIR/workers
cd ${WORKERS} || exit 1
build() {
	echo building plugins..
	cd ${WORKERS} || exit 1
	go build -buildmode=plugin -o ${1}.so ${1}.go && echo build $1 || echo failed to build $1
	echo done
}

clean() {
	echo cleaning $WORKERS directory
	rm -v ${WORKERS}/*.so
}

clean
build "downloader"
cd ${ROOTDIR}
