#
#
#
# This makefile assumes a Cygwin environment under Windows.


sample:
	mkdir -p png
	go run ttfsample.go -fontfile FreeSansBold.ttf -verbose

build:
	go build ttfsample.go

win:
	for i in `ls c:/Windows/Fonts/*.TTF`; do ./ttfsample.exe -fontfile $$i ; done
