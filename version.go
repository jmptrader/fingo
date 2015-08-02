package main

import (
  "fmt"
)

const (

  VERSION = "0.0.1"

)

func version() string {
  return fmt.SPrintf( "VERSION %s" , VERSION )
}
