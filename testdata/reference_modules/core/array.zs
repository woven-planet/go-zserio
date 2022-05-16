package reference_modules.core.array;


// Examples taken from https://github.com/ndsev/zserio/blob/master/doc/ZserioInvisibles.md#packed-arrays-of-integers
struct PackedUint32Array
{
    packed uint32 list[5];
};


struct PackableStructure
{
    uint32 value;
    string text;
};

struct PackedArrayOfCompounds
{
    packed PackableStructure list[5];
};

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
