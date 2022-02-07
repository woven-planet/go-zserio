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
