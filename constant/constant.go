package constant

const (
	RpcServerPort = 5001
	HttpServerPort = 5002
)


type DictionaryStoreType int

const (
	FileStore DictionaryStoreType = iota
	MysqlStore
)