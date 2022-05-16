from testdata.reference_modules.core.array.api import PackableNestedStructure
from testdata.reference_modules.core.array.api import PackedNestedArray
from testdata.reference_modules.core.array.api import InnerStructure


def new():
    array = PackedNestedArray(
        [
            PackableNestedStructure(15, "fünfzehn", InnerStructure(15**2, 15 * 2)),
            PackableNestedStructure(14, "fourteen", InnerStructure(14**2, 14 * 2)),
            PackableNestedStructure(13, "dertien", InnerStructure(13**2, 13 * 2)),
            PackableNestedStructure(16, "seksten", InnerStructure(16**2, 16 * 2)),
            PackableNestedStructure(0, "nul", InnerStructure(0**2, 0 * 2)),
        ]
    )
    return array
