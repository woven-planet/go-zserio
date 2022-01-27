from re import S
import os
import zserio
from testdata.reference_modules.core.instantiations.instantiated_template_struct import (
    InstantiatedTemplateStruct,
)
from testdata.reference_modules.testobject1.testobject.test_object import TestObject
from testdata.reference_modules.core.types.value_wrapper import ValueWrapper
from testdata.reference_modules.core.types.color import Color
from testdata.reference_modules.core.types.city_attributes import CityAttributes

if __name__ == "__main__":
    # Create a dummy object with some random values
    test_data = TestObject()
    test_data.parameter1 = 7

    struct1_field = ValueWrapper(test_data.parameter1)
    struct1_field.value = 72
    struct1_field.other_value = 121
    struct1_field.enum_value = Color.BLUE
    struct1_field.description = "test data"

    struct1 = InstantiatedTemplateStruct(test_data.parameter1)
    struct1.field = struct1_field
    test_data.struct1 = struct1
    test_data.color1 = Color.RED

    test_data.parameter2 = 12

    for i in range(22):
        current_struct_field = ValueWrapper(test_data.parameter2)
        current_struct_field.value = i + 100
        current_struct_field.other_value = i
        current_struct_field.description = "some dummy description"
        current_struct_field.enum_value = Color.BLACK
        current_struct = InstantiatedTemplateStruct(test_data.parameter2)
        current_struct.field = current_struct_field
        test_data.struct_array.append(current_struct)

    for i in range(100):
        test_data.bitmask_array.append(CityAttributes.from_value(2))

    # write the testdata
    os.mkdir("bin")
    byte_data = zserio.serialize_to_bytes(test_data)
    handle = open("bin/testdata.bin", "wb")
    handle.write(byte_data)
    handle.close()
