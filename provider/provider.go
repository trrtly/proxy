package provider

//Provider interface
type Provider interface {
	GetProxys() (res []string, err error)
}
