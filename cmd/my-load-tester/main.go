package main

import (
    "github.com/interchainio/tm-load-test/pkg/loadtest"
    "github.com/you/my-load-tester/pkg/myabciapp"
    // "fmt"
)

func main() {
    err := loadtest.RegisterClientFactory("my-abci-app-name", &myabciapp.MyABCIAppClientFactory{}); //err != nil {
        // panic(err)
    // }
    // if  err not equal null then this below panic will excute.
    if err!=nil{
        panic(err)
    }
    // fmt.Println(err)
    // The loadtest.Run method will handle CLI argument parsing, errors, 
    // configuration, instantiating the load test and/or master/slave 
    // operations, etc. All it needs is to know which client factory to use for
    // its load testing.
    loadtest.Run(&loadtest.CLIConfig{
        AppName:              "my-load-tester",
        AppShortDesc:         "Load testing application for My ABCI App (TM)",
        AppLongDesc:          "Some long description on how to use the tool",
        DefaultClientFactory: "my-abci-app-name",
    })
}
