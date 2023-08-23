package bitshift_operator.schema;

struct BitShiftOperator
{
    uint8 u8BitShift;
    float64 f64Value;

    function float64 testBitShift()
    {
        return (1 << u8BitShift) + f64Value;
    }
};
