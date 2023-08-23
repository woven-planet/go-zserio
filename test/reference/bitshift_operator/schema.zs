package bitshift_operator.schema;

struct BitShiftOperator
{
    uint8 u8BitShift;
    float64 f64Value;

    function float64 testBitShift()
    {
        // This test case validates the correct generation of bitshift expressions.
        // See https://github.com/woven-planet/go-zserio/pull/123.
        return (1 << u8BitShift) + f64Value;
    }
};
