package model

type Deployment struct {
	Name      string
	Namespace string
	Image     string
	Replicas  int
	Port      []int32
	Env       map[string]string
}
