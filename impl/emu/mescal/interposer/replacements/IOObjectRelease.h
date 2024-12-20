//
//  IOObjectRelease.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOOBJECTRELEASE_H
#define _IOOBJECTRELEASE_H

#include <IOKit/IOKitLib.h>

#ifdef IOObjectRelease
#undef IOObjectRelease
#endif	/* IOObjectRelease */

__attribute__ ((noinline)) kern_return_t IOObjectRelease(io_object_t object);

#endif	/* _IOOBJECTRELEASE_H */
