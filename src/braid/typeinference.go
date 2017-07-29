package main

import (
    "fmt"
)

func infer(mod Module) map[string]interface{} {
    fmt.Println(mod.Name)
    types := make(map[string]interface{})
    return types
}