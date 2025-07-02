//
//  memmove.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _MEMMOVE_H
#define _MEMMOVE_H

#include <stddef.h>

#ifdef memmove
#undef memmove
#endif	/* memmove */

__attribute__ ((noinline)) void *memmove(void *restrict dst, const void *restrict src, size_t n);

#endif	/* _MEMMOVE_H */
