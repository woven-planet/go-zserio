package struct_optional_field

enum uint8 State
{
    ERROR = 0,
    SUCCESS = 1,
};

struct Result {
    State status;
    string message if status == State.ERROR;
};
