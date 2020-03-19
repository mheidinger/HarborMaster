package models

type RegistryInfo struct {
	URL             string
	RepositoryInfos []*RepositoryInfo
}

type RepositoryInfo struct {
	Name string
	Tags []string
}
