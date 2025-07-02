//
//  memcmp.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _MEMCMP_H
#define _MEMCMP_H

#include <stddef.h>

#ifdef memcmp
#undef memcmp
#endif	/* memcmp */

__attribute__ ((noinline)) int memcmp(const void *s1, const void *s2, size_t n);

#endif	/* _MEMCMP_H */
