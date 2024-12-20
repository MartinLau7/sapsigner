//
//  kCFAllocatorDefault.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _KCFALLOCATORDEFAULT_H
#define _KCFALLOCATORDEFAULT_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef kCFAllocatorDefault
#undef kCFAllocatorDefault
#endif	/* kCFAllocatorDefault */

const CFAllocatorRef kCFAllocatorDefault;

#endif	/* _KCFALLOCATORDEFAULT_H */
