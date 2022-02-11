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

    // Parameter for optional choices
    int32 parameter3;

    // The first optional choice, expected to not be present in the data
    optional InstantiatedTemplateChoice(parameter3) optionChoice1;

    // the second optional choice, expected to be present in the data
    optional InstantiatedTemplateChoice(parameter3) optionChoice2;

    // This tests the correct type lookup if the choice selector type is enum
    SomeEnum choiceSelector;
    BasicChoice(choiceSelector) basicChoice;

    // foo is just there to have something after an optional type
    varint32 foo;
    
};
