# go-lucky

Go wrapper for the [LuckPerms REST API](https://luckperms.net/wiki/Standalone-and-REST-API)

## Example Usage

```go
// Creating a client
config := golucky.Config{
    RestIp:  "http://127.0.0.1:8080",
}
client := golucky.New(config)

// Creating a client with authentication
config := golucky.Config{
    RestIp:  "http://127.0.0.1:8080",
    AuthKey: "myverysecurekey",
}
client := golucky.New(config)


// Permission check
permCheck, err := client.UserHasPermission("7bd5b459-1e6b-4753-8274-1fbd2fe9a4d5", "cool.permission")
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(permCheck.Result) // One of [ true, false, undefined ]


// Add user to group
nodes, err := client.AddUserNode("7bd5b459-1e6b-4753-8274-1fbd2fe9a4d5", golucky.NewNode{
    Key:   "group.cool",
    Value: true,
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(nodes)
```