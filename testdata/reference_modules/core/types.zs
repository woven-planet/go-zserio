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
    float64 f64Value;
    float32 f32Value;
    float16 f16Value;

    varint16 vari16Value;
    varint32 vari32Value;
    varint64 vari64Value;
    varint variValue;
    varuint16 varu16Value;
    varuint32 varu32Value;
    varuint64 varu64Value;
    varuint varuValue;
    varsize varSizeValue;
    int32 i32ValueArray[5];


    function int32 getValue()
    {
        return value + parameter;
    }

    function float64 getSum()
    {
        // Just some random additions and divisions.
        return 12.0 + f16Value + f32Value + f64Value / (value << 2);
    }

    function float32 testTypeCasts()
    {
        // Test mixing different integer types, and make
        // sure the generated code handles type casts correctly.
        return value + vari16Value + vari32Value + vari64Value + variValue + 
            varu16Value + varu32Value + varu64Value + varuValue + varSizeValue +
            i32ValueArray[0] + lengthof(i32ValueArray);   
    }
};


// SomeEnum defines some values
enum int32 SomeEnum
{
    ATTR_A,
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
    case ATTR_A: int32 fieldA;
    case ATTR_B: int64 fieldB;
    case SomeEnum.ATTR_C: int8 fieldC;

    // the purpose of this structure is this field - HAS_A must be resolved from SomeOtherEnum
    case HAS_A: int8 hasA;
};
