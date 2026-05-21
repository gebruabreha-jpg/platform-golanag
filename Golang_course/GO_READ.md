Standard library:-https://pkg.go.dev/std
Effective Go:-https://go.dev/doc/effective_go
Documentation:-https://go.dev/doc/
https://shell.cloud.google.com/?walkthrough_tutorial_url=https%3A%2F%2Fraw.githubusercontent.com%2Fgolang%2Ftour%2Fmaster%2Ftutorial%2Fweb-service-gin.md&pli=1&show=ide&environment_deployment=ide

 go mod init example/hello
 go: creating new go.mod: module example/hello
 hello.go in which to write your code.

go mod tidy keeps your Go module clean and accurate. When run in the module’s root directory, it automatically adds missing dependencies, removes unused ones, and updates the go.sum file to reflect the exact requirements needed to build your code and its dependencies.

The three fundamental building blocks you’ll use every day are:
    Arrays: fixed-size sequences of elements.
    Slices: flexible, dynamic views of arrays.
    Maps: key–value stores implemented as hash tables.
array types: [N]T
slice types: []T
map types: map[K]T 
where T is an arbitrary type. It specifies the element type of a container type. 
Only values of the specified element type can be stored as element values of values of the container type.
K is an arbitrary comparable type,It specifies the key type of a map.