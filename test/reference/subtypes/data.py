from testdata.subtypes.schema.api import DummyStruct, Colour, SubtypedColour, OtherSubtypedColour


def new() -> DummyStruct:
    return DummyStruct(
        colour1_=Colour.PURPLE,
        colour2_=SubtypedColour.BLUE,
        colour3_=OtherSubtypedColour.ORANGE,
    )
