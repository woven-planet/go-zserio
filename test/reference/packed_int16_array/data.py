from testdata.packed_int16_array.schema.api import PackedInt16Array


def new():
    array = PackedInt16Array([-15, -14, -13, -16, -11])
    return array
