//
//  memset.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-14.
//

#ifndef _MEMSET_H
#define _MEMSET_H

#include <stddef.h>

#ifdef memset
#undef memset
#endif	/* memset */

__attribute__ ((noinline)) void *memset(void *b, int c, size_t len);

#endif	/* _MEMSET_H */
