package provider

//Provider interface
type Provider interface {
	GetProxys() ([]string, error)
}
