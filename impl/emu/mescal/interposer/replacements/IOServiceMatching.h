//
//  IOServiceMatching.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOSERVICEMATCHING_H
#define _IOSERVICEMATCHING_H

#include <IOKit/IOKitLib.h>

#ifdef IOServiceMatching
#undef IOServiceMatching
#endif	/* IOServiceMatching */

__attribute__ ((noinline)) CFMutableDictionaryRef IOServiceMatching(const char *name);

#endif	/* _IOSERVICEMATCHING_H */
