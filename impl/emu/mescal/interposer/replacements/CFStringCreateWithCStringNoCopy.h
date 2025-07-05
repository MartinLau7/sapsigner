//
//  CFStringCreateWithCStringNoCopy.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-04.
//

#ifndef _CFSTRINGCREATEWITHCSTRINGNOCOPY_H
#define _CFSTRINGCREATEWITHCSTRINGNOCOPY_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFStringCreateWithCStringNoCopy
#undef CFStringCreateWithCStringNoCopy
#endif	/* CFStringCreateWithCStringNoCopy */

__attribute__ ((noinline)) CFStringRef CFStringCreateWithCStringNoCopy(CFAllocatorRef alloc, const char *cStr, CFStringEncoding encoding, CFAllocatorRef contentsDeallocator);

#endif	/* _CFSTRINGCREATEWITHCSTRINGNOCOPY_H */
