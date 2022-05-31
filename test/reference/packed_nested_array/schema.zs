package packed_nested_array.schema;

struct InnerStructure
{
    uint64 value64;
    uint16 value16;
};

struct PackableNestedStructure
{
    uint32 value32;
    string text;
    InnerStructure innerStructure;
};

struct PackedNestedArray
{
    packed PackableNestedStructure list[5];
};
