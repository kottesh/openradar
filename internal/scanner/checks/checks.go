package checks

type CheckFunc func(src string) bool

var Checks []CheckFunc
