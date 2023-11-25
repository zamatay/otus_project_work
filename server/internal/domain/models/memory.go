package models

type Memory struct {
	Total, Used, Buffers, Cached, Free, Available, Active, Inactive, SwapTotal, SwapUsed,
	SwapCached, SwapFree, Mapped, Shmem, Slab, PageTables, Committed, VmallocUsed uint64
	MemAvailableEnabled bool
}
