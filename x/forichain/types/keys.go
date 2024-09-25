package types

const (
	// ModuleName defines the module name
	ModuleName = "forichain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_forichain"
)

var (
	ParamsKey = []byte("p_forichain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
