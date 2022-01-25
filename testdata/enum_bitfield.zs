package bitfield_static;

enum bit:3 Color
{
    NONE = 000b,
    RED = 010b,
    BLUE,
    BLACK = 111b
};

struct Palette
{
    Color primary;
    Color secondary;
    Color background;
};

