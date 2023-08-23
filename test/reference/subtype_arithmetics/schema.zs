package subtype_arithmetics.schema;


enum uint32 Colour
{
    RED,
    PURPLE,
    BLUE,
    ORANGE,
    CYAN,
};

// Apply two level of indirection for this test.
subtype int32 SubtypedInt32;
subtype varuint64 SubtypedVarUint64;
subtype float64 SubtypedFloat64;
subtype SubtypedFloat64 DoubleSubtypedFloat64;

struct SubTypeArithmetics
{
    SubtypedInt32 value1;
    SubtypedVarUint64 value2;
    SubtypedFloat64 value3;
    int64 value4;
    DoubleSubtypedFloat64 value5;
    Colour colourValue;



    function float64 retunSum()
    {
        return value1 + value2 + value3 + (1 << 4) + value4 + value5 + valueof(colourValue);
    }

};
