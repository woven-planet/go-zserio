package reference_modules.core.instantiations;

import reference_modules.core.templates.*;
import reference_modules.core.types.*;

/** Test instantiation with int32 and dummy struct */
instantiate TestTemplatedStructWithTypeParamters<int32, ValueWrapper> InstantiatedTemplateStruct;

/** Test instantiation of choice */
instantiate TemplatedChoice<int32, ValueWrapper> InstantiatedTemplateChoice;
