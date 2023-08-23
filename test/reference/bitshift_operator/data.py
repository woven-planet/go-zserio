from testdata.bitshift_operator.schema.api import BitShiftOperator


def new() -> BitShiftOperator:
    return BitShiftOperator(
        u8_bit_shift_=2,
        f64_value_=35.55,
    )
