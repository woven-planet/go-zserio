package reference_modules.core.templates;

/** This templated struct is using a type parameter on a template field.  */
struct TestTemplatedStructWithTypeParamters<T, V>(T parameter)
{
  /** The templated field */
  V(parameter) field;
  
  // A templated function to test expression evaluation
  function T getValue()
  {
    return (parameter > 10) ? field.getValue() << 1 : field.getValue() >> 1;
  }
};

/** This is a templated choice */
choice TemplatedChoice<T, V>(uint8 parameter) on parameter
{
    case 8: T fieldT;
    case 16: V(parameter) fieldV;
    default: int32 defaultField;
};