from testdata.packed_position2d_array.schema.api import PackedPosition2DArray, Position2D


def new():
    positions: list[Position2D] = [
        Position2D(shift_=0, x_=-3112, y_=-12),
        Position2D(shift_=0, x_=-3113, y_=-11),
        Position2D(shift_=0, x_=-3114, y_=-12),
        Position2D(shift_=0, x_=-3115, y_=-11),
        Position2D(shift_=0, x_=-3116, y_=-12),
        Position2D(shift_=0, x_=-3117, y_=-11),
    ]
    array = PackedPosition2DArray(positions)
    return array
