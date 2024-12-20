//
//  malloc.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _MALLOC_H
#define _MALLOC_H

#include <stddef.h>

#ifdef malloc
#undef malloc
#endif	/* malloc */

__attribute__ ((noinline)) void *malloc(size_t size);

#endif	/* _MALLOC_H */
