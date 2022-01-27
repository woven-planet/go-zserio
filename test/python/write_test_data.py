import os
import sys

import zserio

from testdata.reference_modules.core.instantiations.instantiated_template_struct import (
    InstantiatedTemplateStruct,
)
from testdata.reference_modules.testobject1.testobject.test_object import TestObject
from testdata.reference_modules.core.types.value_wrapper import ValueWrapper
from testdata.reference_modules.core.types.color import Color
from testdata.reference_modules.core.types.city_attributes import CityAttributes

def _new_test_object():
    """Create a dummy object with some random values.
    """
    d = TestObject()
    d.parameter1 = 7

    struct1_field = ValueWrapper(d.parameter1)
    struct1_field.value = 72
    struct1_field.other_value = 121
    struct1_field.enum_value = Color.BLUE
    struct1_field.description = "test data"

    struct1 = InstantiatedTemplateStruct(d.parameter1)
    struct1.field = struct1_field
    d.struct1 = struct1
    d.color1 = Color.RED

    d.parameter2 = 12

    for i in range(22):
        current_struct_field = ValueWrapper(d.parameter2)
        current_struct_field.value = i + 100
        current_struct_field.other_value = i
        current_struct_field.description = "some dummy description"
        current_struct_field.enum_value = Color.BLACK
        current_struct = InstantiatedTemplateStruct(d.parameter2)
        current_struct.field = current_struct_field
        d.struct_array.append(current_struct)

    for i in range(100):
        d.bitmask_array.append(CityAttributes.from_value(2))

    return d

def main():
    out = sys.argv[1]
    # write the testdata
    byte_data = zserio.serialize_to_bytes(_new_test_object())
    with open(out, "wb") as f:
        f.write(byte_data)

if __name__ == "__main__":
    main()
