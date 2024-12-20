//
//  CFDictionaryGetValue.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFDICTIONARYGETVALUE_H
#define _CFDICTIONARYGETVALUE_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFDictionaryGetValue
#undef CFDictionaryGetValue
#endif	/* CFDictionaryGetValue */

__attribute__ ((noinline)) const void *CFDictionaryGetValue(CFDictionaryRef theDict, const void *key);

#endif	/* _CFDICTIONARYGETVALUE_H */
