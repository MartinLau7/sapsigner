//
//  CFRelease.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFRELEASE_H
#define _CFRELEASE_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFRelease
#undef CFRelease
#endif	/* CFRelease */

__attribute__ ((noinline)) void CFRelease(CFTypeRef cf);

#endif	/* _CFRELEASE_H */
