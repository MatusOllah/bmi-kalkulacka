@echo off

go mod tidy

:: windows
set GOOS=windows
go build -v -o bin/windows/bmi-kalkulacka.exe

:: mac
set GOOS=darwin
go build -v -o bin/mac/bmi-kalkulacka

:: linux
set GOOS=linux
go build -v -o bin/linux/bmi-kalkulacka

set GOOS=windows

echo.
echo hotovo