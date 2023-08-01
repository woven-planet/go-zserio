package subtypes.schema;


enum uint32 Colour
{
    RED,
    PURPLE,
    BLUE,
    ORANGE,
    CYAN,
};

// Apply two level of indirection for this test.
subtype Colour SubtypedColour;
subtype SubtypedColour OtherSubtypedColour;

struct DummyStruct
{
    Colour colour1;
    SubtypedColour colour2;
    OtherSubtypedColour colour3;

    function bool testColour1()
    {
        return (colour1 == Colour.RED);
    }

    function bool testColour2()
    {
        return (colour2 == SubtypedColour.RED);
    }

    function bool testColor3()
    {
        return (colour3 == OtherSubtypedColour.ORANGE);
    }
};
