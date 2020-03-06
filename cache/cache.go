package cache

//Cache interface
type Cache interface {
	Get(key string) ([]byte, error)
	SAdd(key string, values ...interface{}) (int, error)
	SRem(key string, values ...interface{}) (int, error)
	SMembers(key string) ([]string, error)
	SRandMember(key string) (string, error)
	SRandMembers(key string, count int) ([]string, error)
}
