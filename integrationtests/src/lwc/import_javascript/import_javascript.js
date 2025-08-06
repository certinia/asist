/* SPECIFYING JS (all below should be reported) */

import { LoadingMixin } from 'c/application.js';
import './bin/test/jest-matchers/matchers.js';

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import
import defaultExport from "module-name.js";
import * as name from "module-name.js";
import { export1 } from "module-name.js";
import { export1 as alias1 } from "module-name.js";
import { default as alias } from "module-name.js";
import { export1, export2 } from "module-name.js";
import { export1, export2 as alias2, /*  */ } from "module-name.js";
import { "string name" as alias } from "module-name.js";
import defaultExport, { export1, /*  */ } from "module-name.js";
import defaultExport, * as name from "module-name.js";``
import "module-name.js";

// Single quotes
import defaultExport from 'module-name.js';
import * as name from 'module-name.js';
import { export1 } from 'module-name.js';
import { export1 as alias1 } from 'module-name.js';
import { default as alias } from 'module-name.js';
import { export1, export2 } from 'module-name.js';
import { export1, export2 as alias2, /*  */ } from 'module-name.js';
import { 'string name' as alias } from 'module-name.js';
import defaultExport, { export1, /*  */ } from 'module-name.js';
import defaultExport, * as name from 'module-name.js';
import 'module-name.js';

// Funky but valid variable names
import defaultExport from "module-name.js";
import * as name from "module-name.js";
import { _export1 } from "module-name.js";
import { _export1 as alias-1 } from "module-name.js";
import { default as alias- } from "module-name.js";
import { _export1, $export2 } from "module-name.js";
import { _export1, $export2 as alias-2, /*  */ } from "module-name.js";
import { "string name" as alias- } from "module-name.js";
import defaultExport, { _export1, /*  */ } from "module-name.js";
import defaultExport, * as name from "module-name.js";
import "module-name.js";

// More spaces
 import  defaultExport  from  "module-name.js"  ;
 import  *  as  name  from  "module-name.js"  ;
 import  {  export1  }  from  "module-name.js"  ;
 import  {  export1  as  alias1  }  from  "module-name.js"  ;
 import  {  default  as  alias  }  from  "module-name.js"  ;
 import  {  export1,  export2  }  from  "module-name.js"  ;
 import  {  export1,  export2  as  alias2,  /*  */  }  from  "module-name.js"  ;
 import  {  "string  name"  as  alias  }  from  "module-name.js"  ;
 import  defaultExport,  {  export1,  /*  */  }  from  "module-name.js"  ;
 import  defaultExport,  *  as  name  from  "module-name.js"  ;
 import  "module-name.js"  ;

/* NOT SPECIFYING JS (none below should be reported) */

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import
import defaultExport from "module-name";
import * as name from "module-name";
import { export1 } from "module-name";
import { export1 as alias1 } from "module-name";
import { default as alias } from "module-name";
import { export1, export2 } from "module-name";
import { export1, export2 as alias2, /*  */ } from "module-name";
import { "string name" as alias } from "module-name";
import defaultExport, { export1, /*  */ } from "module-name";
import defaultExport, * as name from "module-name";
import "module-name";

// Single quotes
import defaultExport from 'module-name';
import * as name from 'module-name';
import { export1 } from 'module-name';
import { export1 as alias1 } from 'module-name';
import { default as alias } from 'module-name';
import { export1, export2 } from 'module-name';
import { export1, export2 as alias2, /*  */ } from 'module-name';
import { 'string name' as alias } from 'module-name';
import defaultExport, { export1, /*  */ } from 'module-name';
import defaultExport, * as name from 'module-name';
import 'module-name';

// Funky but valid variable names
import defaultExport from "module-name";
import * as name from "module-name";
import { _export1 } from "module-name";
import { _export1 as alias-1 } from "module-name";
import { default as alias- } from "module-name";
import { _export1, $export2 } from "module-name";
import { _export1, $export2 as alias-2, /*  */ } from "module-name";
import { "string name" as alias- } from "module-name";
import defaultExport, { _export1, /*  */ } from "module-name";
import defaultExport, * as name from "module-name";
import "module-name";

// More spaces
 import  defaultExport  from  "module-name"  ;
 import  *  as  name  from  "module-name"  ;
 import  {  export1  }  from  "module-name"  ;
 import  {  export1  as  alias1  }  from  "module-name"  ;
 import  {  default  as  alias  }  from  "module-name"  ;
 import  {  export1,  export2  }  from  "module-name"  ;
 import  {  export1,  export2  as  alias2,  /*  */  }  from  "module-name"  ;
 import  {  "string  name"  as  alias  }  from  "module-name"  ;
 import  defaultExport,  {  export1,  /*  */  }  from  "module-name"  ;
 import  defaultExport,  *  as  name  from  "module-name"  ;
 import  "module-name"  ;

// Double-quoted strings (shouldn't be reported)
a="import defaultExport from 'module-name';"
a="import * as name from 'module-name';"
a="import { export1 } from 'module-name';"
a="import { export1 as alias1 } from 'module-name';"
a="import { default as alias } from 'module-name';"
a="import { export1, export2 } from 'module-name';"
a="import { export1, export2 as alias2, /*  */ } from 'module-name';"
a="import { 'string name' as alias } from 'module-name';"
a="import defaultExport, { export1, /*  */ } from 'module-name';"
a="import defaultExport, * as name from 'module-name';"
a="import 'module-name';"

// Single-quoted strings (shouldn't be reported)
a='import defaultExport from "module-name";'
a='import * as name from "module-name";'
a='import { export1 } from "module-name";'
a='import { export1 as alias1 } from "module-name";'
a='import { default as alias } from "module-name";'
a='import { export1, export2 } from "module-name";'
a='import { export1, export2 as alias2, /*  */ } from "module-name";'
a='import { "string name" as alias } from "module-name";'
a='import defaultExport, { export1, /*  */ } from "module-name";'
a='import defaultExport, * as name from "module-name";'
a='import "module-name";'

// Single-line comments (shouldn't be reported)
// import defaultExport from "module-name";
// import * as name from "module-name";
// import { export1 } from "module-name";
// import { export1 as alias1 } from "module-name";
// import { default as alias } from "module-name";
// import { export1, export2 } from "module-name";
// import { export1, export2 as alias2, /*  */ } from "module-name";
// import { "string name" as alias } from "module-name";
// import defaultExport, { export1, /*  */ } from "module-name";
// import defaultExport, * as name from "module-name";
// import "module-name";

// Multi-line comments (shouldn't be reported)
/*
 import defaultExport from "module-name";
*/
/* import * as name from "module-name"; */
/* import { export1 } from "module-name"; */
/* import { export1 as alias1 } from "module-name"; */
/* import { default as alias } from "module-name"; */
/* import { export1, export2 } from "module-name"; */
/* import { export1, export2 as alias2 } from "module-name"; */
/* import { "string name" as alias } from "module-name"; */
/* import defaultExport, { export1 } from "module-name"; */
/* import defaultExport, * as name from "module-name"; */
/* import "module-name"; */