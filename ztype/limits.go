package ztype

const (
	MaxInt8  = (1 << 7) - 1
	MinInt8  = -1 << 7
	MaxInt16 = (1 << 15) - 1
	MinInt16 = -1 << 15
	MaxInt32 = (1 << 31) - 1
	MinInt32 = -1 << 31
	MaxInt64 = (1 << 63) - 1
	MinInt64 = -1 << 63

	MaxUint8  = (1 << 8) - 1
	MinUint8  = 0
	MaxUint16 = (1 << 16) - 1
	MinUint16 = 0
	MaxUint32 = (1 << 32) - 1
	MinUint32 = 0
	MaxUint64 = (1 << 64) - 1
	MinUint64 = 0

	MaxVarint16 = (1 << (6 + 8)) - 1
	MinVarint16 = -MaxVarint16
	MaxVarint32 = (1 << (6 + 7 + 7 + 8)) - 1
	MinVarint32 = -MaxVarint32
	MaxVarint64 = (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1
	MinVarint64 = -MaxVarint64
	MaxVarint   = MaxInt64
	MinVarint   = MinInt64

	MaxVaruint16 = (1 << (7 + 8)) - 1
	MinVaruint16 = 0
	MaxVaruint32 = (1 << (7 + 7 + 7 + 8)) - 1
	MinVaruint32 = 0
	MaxVaruint64 = (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1
	MinVaruint64 = 0
	MaxVaruint   = MaxUint64
	MinVaruint   = 0
	MaxVarsize   = (1 << (2 + 7 + 7 + 7 + 8)) - 1
	MinVarsize   = 0
)
