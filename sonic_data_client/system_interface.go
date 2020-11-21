package client

import "io/ioutil"

// ImplIoutilReadFile points to the implementation of ioutil.ReadFile
var ImplIoutilReadFile func(string) ([]byte, error) = ioutil.ReadFile
