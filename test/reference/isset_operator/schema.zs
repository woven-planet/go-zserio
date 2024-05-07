package isset_operator.schema;


bitmask uint8 Permission
{
    EXECUTABLE = 0x01,
    READABLE = 0x02,
    WRITEABLE = 0x04,

};

struct IssetOperator
{
    Permission u8Value;

    function bool isExecutable()
    {
        return isset(u8Value, Permission.EXECUTABLE);
    }
    function bool isReadable()
    {
        return isset(u8Value, Permission.READABLE);
    }
    function bool isWriteable()
    {
        return isset(u8Value, Permission.WRITEABLE);
    }

};
