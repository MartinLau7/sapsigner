//
//  CFStringGetCString.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFSTRINGGETCSTRING_H
#define _CFSTRINGGETCSTRING_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFStringGetCString
#undef CFStringGetCString
#endif	/* CFStringGetCString */

__attribute__ ((noinline)) Boolean CFStringGetCString(CFStringRef theString, char *buffer, CFIndex bufferSize, CFStringEncoding encoding);

#endif	/* _CFSTRINGGETCSTRING_H */
