#
#
#
# This makefile assumes a Cygwin environment under Windows.

build:
	go build 

win:
	for i in c:/Windows/Fonts/*.TTF; do ./ttfsample.exe -fontfile $$i ; done
