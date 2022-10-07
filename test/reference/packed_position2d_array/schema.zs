package packed_position2d_array.schema;

struct Position2D(int shift)
{
  int<(31-shift) + 1> x;
  int<(31-shift) + 1> y;
};

struct PackedPosition2DArray
{
    packed Position2D(0) positions[6];
};
