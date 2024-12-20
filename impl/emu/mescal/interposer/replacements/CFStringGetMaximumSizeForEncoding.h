//
//  CFStringGetMaximumSizeForEncoding.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFSTRINGGETMAXIMUMSIZEFORENCODING_H
#define _CFSTRINGGETMAXIMUMSIZEFORENCODING_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFStringGetMaximumSizeForEncoding
#undef CFStringGetMaximumSizeForEncoding
#endif	/* CFStringGetMaximumSizeForEncoding */

__attribute__ ((noinline)) CFIndex CFStringGetMaximumSizeForEncoding(CFIndex length, CFStringEncoding encoding);

#endif	/* _CFSTRINGGETMAXIMUMSIZEFORENCODING_H */
