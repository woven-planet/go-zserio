from testdata.packed_uint32_array.schema.api import PackedUint32Array


def new():
    array = PackedUint32Array([15, 14, 13, 16, 0])
    return array
