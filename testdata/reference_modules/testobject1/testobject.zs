package reference_modules.testobject1.testobject;

import reference_modules.core.instantiations.*;
import reference_modules.core.types.*;


struct TestObject
{
    int32 parameter1;
    InstantiatedTemplateStruct(parameter1) struct1;
    // Test an enum
    Color color1;
    align(8):
    // Some array for testing
    int32 parameter2;
    packed InstantiatedTemplateStruct(parameter2) structArray[];
    align(5):
    CityAttributes bitmaskArray[100];
};
