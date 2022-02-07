package enum_expressions;

const uint32 cu32_random_constant = 3253;

enum bit:3 Color
{
    NONE = 000b,
    RED = 010b,
    BLUE,
    BLACK = 111b,
};

enum uint32 CrazyEnum
{
    HEX_EXPR = 0xdeadbeef,
    DEC_EXPR = 835,
    ADD_EXPR = 2 + 6,
    SUB_EXPR = 10 - 5,
    MULT_EXPR = 1 * 7,
    DIV_EXPR = 10 / 4,
    STRANGE = (10 + (-2)) * 5,
    LSHIFT = 5 << 2,
    RSHIFT = 3276 >> 2,
    TEST_TERT = false ? 5: 2,
    BIN_AND = 0x11011 & 0x10010,
    BIN_OR = 0x11011 | 0x10111,
    BIN_XOR = 0x11011 ^ 0x10110,
    NUMBITS = numbits(14235235),
    CONSTANT = cu32_random_constant,
    ENUM_VALUE = Color.BLACK,
};
