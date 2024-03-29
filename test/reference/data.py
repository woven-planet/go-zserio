from testdata.reference_modules.core.instantiations.instantiated_template_struct import (
    InstantiatedTemplateStruct,
)
from testdata.reference_modules.core.instantiations.instantiated_template_choice import (
    InstantiatedTemplateChoice,
)
from testdata.reference_modules.testobject1.testobject.test_object import TestObject
from testdata.reference_modules.core.types.value_wrapper import ValueWrapper
from testdata.reference_modules.core.types.basic_choice import BasicChoice
from testdata.reference_modules.core.types.color import Color
from testdata.reference_modules.core.types.city_attributes import CityAttributes
from testdata.reference_modules.core.types.some_enum import SomeEnum
from zserio.bitbuffer import BitBuffer


def new():
    """Create a dummy object with some random values."""
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

    # d.option_choice1 is deliberately not set
    d.parameter3 = 12
    d.option_choice2 = InstantiatedTemplateChoice(d.parameter3)
    d.option_choice2.default_field = 707

    # The next test is for the correct lookup of choice selector types. The
    # choice below is using an enum, whose values also exist in a different enum
    d.choice_selector = SomeEnum.HAS_A
    d.basic_choice = BasicChoice(d.choice_selector)
    d.basic_choice.has_a = 5

    # The next test checks that float types are correctly stored in packed and
    # nonpacked arrays
    d.float_member = 23.5
    for i in range(100):
        d.float_array.append(float(i))

    d.data = BitBuffer(b"ABC", 13)
    d.foo = 42
    return d
