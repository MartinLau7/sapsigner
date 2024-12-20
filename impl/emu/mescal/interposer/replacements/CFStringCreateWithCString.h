//
//  CFStringCreateWithCString.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _CFSTRINGCREATEWITHCSTRING_H
#define _CFSTRINGCREATEWITHCSTRING_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFStringCreateWithCString
#undef CFStringCreateWithCString
#endif	/* CFStringCreateWithCString */

__attribute__ ((noinline)) CFStringRef CFStringCreateWithCString(CFAllocatorRef alloc, const char *cStr, CFStringEncoding encoding);

#endif	/* _CFSTRINGCREATEWITHCSTRING_H */
