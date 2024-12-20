//
//  memcpy.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _MEMCPY_H
#define _MEMCPY_H

#include <stddef.h>

#ifdef memcpy
#undef memcpy
#endif	/* memcpy */

__attribute__ ((noinline)) void *memcpy(void *restrict dst, const void *restrict src, size_t n);

#endif	/* _MEMCPY_H */
