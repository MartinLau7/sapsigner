//
//  CFUUIDCreateString.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFUUIDCREATESTRING_H
#define _CFUUIDCREATESTRING_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFUUIDCreateString
#undef CFUUIDCreateString
#endif	/* CFUUIDCreateString */

__attribute__ ((noinline)) CFStringRef CFUUIDCreateString(CFAllocatorRef alloc, CFUUIDRef uuid);

#endif	/* _CFUUIDCREATESTRING_H */
