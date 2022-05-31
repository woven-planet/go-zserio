package packed_compound_array.schema;

struct PackableStructure
{
    uint32 value;
    string text;
};

struct PackedArrayOfCompounds
{
    packed PackableStructure list[5];
};
