//
//  kCFAllocatorNull.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-07-04.
//

#ifndef _KCFALLOCATORNULL_H
#define _KCFALLOCATORNULL_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef kCFAllocatorNull
#undef kCFAllocatorNull
#endif	/* kCFAllocatorNull */

const CFAllocatorRef kCFAllocatorNull;

#endif	/* _KCFALLOCATORNULL_H */
