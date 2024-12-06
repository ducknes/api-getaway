package settings

type Config struct {
	Port    int     `json:"port"`
	Cluster Cluster `json:"cluster"`
}

type Cluster struct {
	UserServiceUrl    string `json:"user_service_url"`
	StorageServiceUrl string `json:"storage_service_url"`
	AuthServiceUrl    string `json:"auth_service_url"`
}
