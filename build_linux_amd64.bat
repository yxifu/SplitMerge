for /F "tokens=2 delims==" %%I in ('wmic os get LocalDateTime /VALUE') do set datetime=%%I
set datetime1=%datetime:~0,12%

go env -w GOOS=linux
go build -o SplitMerge_linux_amd64_%datetime1%
go env -w GOOS=windows