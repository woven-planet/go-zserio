package array_basic;

struct FixedStringArray {
    string fixedStringArray[5];
};

struct AutoStringArray {
    string stringArray[];
};


struct PackedFixedVarintArray {
    packed varint fixedVarintArray[5];
};

struct AutoVarintArray {
    varint autoVarintArray[];
};

struct PackedAutoVarintArray {
    packed varint autoVarintArray[];
};

struct AutoFloat16Array {
    float16 array[];
};

struct AutoFloat32Array {
    float32 array[];
};

struct AutoFloat64Array {
    float64 array[];
};

struct AutoUint32Array {
    uint32 array[];
};

struct PackedAutoUint32Array {
    packed uint32 array[];
};

struct PackedFixedUint32Array {
    packed uint32 array[5];
};

struct AutoVarInt32Array {
    varint32 array[];
};

struct PackedAutoVarInt32Array {
    packed varint32 array[];
};

struct PackedFixedVarInt32Array {
    packed varint32 array[6];
};

subtype uint32 BlockIndex;

struct PackedAutoSubtypeUInt32Array {
    packed BlockIndex blockIndexArray[];
};
