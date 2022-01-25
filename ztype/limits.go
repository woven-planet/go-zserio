package ztype

const (
	INT8_MAX  = (1 << 7) - 1
	INT8_MIN  = -1 << 7
	INT16_MAX = (1 << 15) - 1
	INT16_MIN = -1 << 15
	INT32_MAX = (1 << 31) - 1
	INT32_MIN = -1 << 31
	INT64_MAX = (1 << 63) - 1
	INT64_MIN = -1 << 63

	UINT8_MAX  = (1 << 8) - 1
	UINT8_MIN  = 0
	UINT16_MAX = (1 << 16) - 1
	UINT16_MIN = 0
	UINT32_MAX = (1 << 32) - 1
	UINT32_MIN = 0
	UINT64_MAX = (1 << 64) - 1
	UINT64_MIN = 0

	VARINT16_MAX = (1 << (6 + 8)) - 1
	VARINT16_MIN = -VARINT16_MAX
	VARINT32_MAX = (1 << (6 + 7 + 7 + 8)) - 1
	VARINT32_MIN = -VARINT32_MAX
	VARINT64_MAX = (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1
	VARINT64_MIN = -VARINT64_MAX
	VARINT_MAX   = INT64_MAX
	VARINT_MIN   = INT64_MIN

	VARUINT16_MAX = (1 << (7 + 8)) - 1
	VARUINT16_MIN = 0
	VARUINT32_MIN = 0
	VARUINT32_MAX = (1 << (7 + 7 + 7 + 8)) - 1
	VARUINT64_MIN = 0
	VARUINT64_MAX = (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1
	VARUINT_MIN   = 0
	VARUINT_MAX   = UINT64_MAX
	VARSIZE_MIN   = 0
	VARSIZE_MAX   = (1 << (2 + 7 + 7 + 7 + 8)) - 1
)
