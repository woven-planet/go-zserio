from testdata.isset_operator.schema.api import IssetOperator, Permission


def new() -> IssetOperator:
    return IssetOperator(
        u8_value_=Permission.Values.READABLE,
    )
