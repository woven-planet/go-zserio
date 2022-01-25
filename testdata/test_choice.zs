package test_choice;



/*!

## Core Property Type

`CorePropertyType` defines attribute properties.

**Definitions**

!*/

enum varuint16 CorePropertyType
{
  /** Reason why an attribute is there. */
  ATTRIBUTE_REASON,

  /**
    * Explicitly use metric unit. The attribute property is only used if units
    * are defined to be imperial in the layer metadata defaults and need to be
    * overwritten for a single attribute.
    */
  METRIC_UNIT,

  /**
    * Explicitly use imperial unit. The attribute property is only used if units
    * are defined to be metric in the layer metadata defaults and need to be
    * overwritten for a single attribute.
    */
  IMPERIAL_UNIT,
};


choice CorePropertyValue(CorePropertyType type) on type
{
  case ATTRIBUTE_REASON:
          uint32 reason;
  case CorePropertyType.METRIC_UNIT:
          ;
  case IMPERIAL_UNIT:
          ;
};

enum uint32 AreaType {
    COUNTRY,
    STATE,
    CITY,
    MAP,
    ROAD,
    OTHER,
};

choice AreaAttributes(AreaType type) on type
{
    case AreaType.COUNTRY: // prefix "AreaType." is optional
    case STATE:
    case CITY:
        int32 regionAttr;
    case MAP:
        /* empty */ ;
    case ROAD:
        varint roadAttr;
    default:
        uint8 defaultAttr;
};


/** Type of geometries stored in the layer. */
enum uint8 GeometryLayerType
{
  /** List of 2-dimensional positions. */
  POSITION_2D,

  /** List of 3-dimensional positions. */
  POSITION_3D,

  /** 2-dimensional lines using non-indexed storage. */
  LINE_2D,

  /** 3-dimensional lines using non-indexed storage. */
  LINE_3D,

  /** Triangulated 2-dimensional polygons using indexed storage. */
  POLYGON_2D,

  /** Triangulated 3-dimensional polygons using indexed storage. */
  POLYGON_3D,

  /** Triangulated 3-dimensional meshes using indexed storage. */
  MESH_3D,
};

/** Number of bits to shift coordinate values to reduce precision. */
subtype bit:5 CoordShift;

/**
  * Absolute 2-dimensional geographic position using NDS coordinates.
  * 2-dimensional positions are encoded with variable size.
  * Coordinate shift is applied.
  */
struct Position2D(CoordShift shift)
{
  /** Longitude. */
  int<(31-shift) + 1> longitude;

  /** Latitude. */
  int<(31-shift) + 1> latitude;

  /** Shift longitude value back to the left. */
  function int32 lon()
  {
    return (longitude << shift);
  }

  /** Shift latitude value back to the left. */
  function int32 lat()
  {
    return (latitude << shift);
  }
};

/** Buffer for 2-dimensional geographic positions. */
struct Position2DBuffers(CoordShift shift, varsize numElements)
{
  /** Base position. */
  Position2D(shift) basePosition;

  /** Buffer of 2-dimensional offset positions relative to the base position. */
  PositionBuffer2D(basePosition, shift) positions : positions.numPositions == numElements;
};

/** Buffer for 2-dimensional positions. */
struct PositionBuffer2D(Position2D base, CoordShift shift)
{
  /** Number of geographic positions stored in the buffer. */
  varsize numPositions;
};


choice Buffers(GeometryLayerType type, CoordShift shiftXY, CoordShift shiftZ, varsize numElements) on type
{
  /** List of 2-dimensional positions. */
  case POSITION_2D:
          Position2DBuffers(shiftXY, numElements) positions2D;

  /** List of 3-dimensional positions. */
  case POSITION_3D:
          Position2DBuffers(shiftXY, numElements) positions3D;

};