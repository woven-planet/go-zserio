from testdata.reference_modules.core.array.api import PackableStructure
from testdata.reference_modules.core.array.api import PackedArrayOfCompounds


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
