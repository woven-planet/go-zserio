from testdata.subtype_arithmetics.schema.api import SubTypeArithmetics, Colour, SubtypedInt32, SubtypedVarUint64, SubtypedFloat64, DoubleSubtypedFloat64


def new() -> SubTypeArithmetics:
    return SubTypeArithmetics(
        value1_=-55543,
		value2_=43634,
		value3_=22.0,
		value4_=64565,
		value5_=-24.0,
        colour_value_=Colour.PURPLE,
    )
