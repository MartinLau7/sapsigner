//
//  calloc.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _CALLOC_H
#define _CALLOC_H

#include <stddef.h>

#ifdef calloc
#undef calloc
#endif	/* calloc */

__attribute__ ((noinline)) void *calloc(size_t count, size_t size);

#endif	/* _CALLOC_H */
