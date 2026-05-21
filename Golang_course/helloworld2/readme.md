 go run .\main.go
 output:-
.\main.go:6:12: undefined: helperFunction


go run .\main.go .\helper.go
output:-
Result: 10



or use go build 
go build   // now module.exe is generated
./module.exe
output:-
Result: 10


or use go run . 
output:-
Result: 10
