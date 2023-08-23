from testdata.extern_bytes_array.schema.api import ExternBytesArray
import zserio

def new():
    array = ExternBytesArray(
        extern_array_=[zserio.bitbuffer.BitBuffer(buffer=bytes([0xff, i, 0xe0]),bitsize=19) for i in range(3)],
        bytes_array_=[bytearray([0xff, i]) for i in range(3)],
    )
    return array
