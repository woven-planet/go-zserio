package reference_modules.core.types;

enum bit:3 Color
{
    NONE = 000b,
    RED = 010b,
    BLUE,
    BLACK = 111b
};

bitmask uint32 CityAttributes
{
    HAS_TOWNHALL,
    HAS_RAILWAY_STATION,
    HAS_UNIVERSITY,
    HAS_HOKKIEN_MEE_STORE = 0x100,
    HAS_SUBWAY,
    HAS_HIGHSCHOOL
};

struct ValueWrapper(int32 parameter)
{
    int32 value;
    int8 otherValue;
    align(16):
    
    Color enumValue if parameter == 7;
    align(8):
    string description;
    
    function int32 getValue()
    {
        return value + parameter;
    }
};


// SomeEnum defines some values
enum int32 SomeEnum
{
    ATRR_A,
    ATTR_B,
    ATTR_C,
    HAS_A
};

// SomeOtherEnum defines other values, but note that one entry is already defined
// in another enum
enum int32 SomeOtherEnum
{
    HAS_A,
    HAS_B,
    HAS_C,
    HAS_D
};

choice BasicChoice(SomeEnum type) on type
{
    case ATRR_A: int32 fieldA;
    case ATTR_B: int64 fieldB;
    case SomeEnum.ATTR_C: int8 fieldC;

    // the purpose of this structure is this field - HAS_A must be resolved from SomeOtherEnum
    case HAS_A: int8 hasA;
};
