package enum_basic;

// Role is an enum with some values autoset to a value
enum uint8 Role
{
    DEVELOPER = 0,
    TEAM_LEAD = 1,
    CTO       = 2,
    SOMEONE_ELSE,
    NOBODY,
};

enum uint8 RoleTwo
{
    DEVELOPER = 0,
    TEAM_LEAD = 1,
    CTO       = 2,
};

enum bit:3 Color
{
    NONE = 000b,
    RED = 010b,
    BLUE,
    BLACK = 111b,
};
