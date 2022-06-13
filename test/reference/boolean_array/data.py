from testdata.boolean_array.schema.api import BooleanArray


def new():
    array = BooleanArray([True, False, True, True, False])
    return array
