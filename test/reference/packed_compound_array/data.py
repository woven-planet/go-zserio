from testdata.packed_compound_array.schema.api import (PackableStructure, PackedArrayOfCompounds)

def new():
    array = PackedArrayOfCompounds(
        [
            PackableStructure(15, "fünfzehn"),
            PackableStructure(14, "fourteen"),
            PackableStructure(13, "dertien"),
            PackableStructure(16, "seksten"),
            PackableStructure(0, "nul"),
        ]
    )
    return array
