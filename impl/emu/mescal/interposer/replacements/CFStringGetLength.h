//
//  CFStringGetLength.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFSTRINGGETLENGTH_H
#define _CFSTRINGGETLENGTH_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFStringGetLength
#undef CFStringGetLength
#endif	/* CFStringGetLength */

__attribute__ ((noinline)) CFIndex CFStringGetLength(CFStringRef theString);

#endif	/* _CFSTRINGGETLENGTH_H */
