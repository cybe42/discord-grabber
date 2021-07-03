set GOARCH=386
set GOOS=windows
garble -tiny -literals -seed glnagmlawaeretostpodbasefstlttne -debugdir=out build grabber.go
cd out/command-line-arguments
go build -ldflags "-s -w" -ldflags -H=windowsgui grabber.go
cd ../../
upx -9 out/command-line-arguments/grabber.exe
cd out/command-line-arguments
move grabber.exe ../../
cd ../../
rmdir /Q /S out
pause