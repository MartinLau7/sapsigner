//
//  malloc_size.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _MALLOC_SIZE_H
#define _MALLOC_SIZE_H

#include <stddef.h>

#ifdef malloc_size
#undef malloc_size
#endif	/* malloc_size */

__attribute__ ((noinline)) size_t malloc_size(const void *ptr);

#endif	/* _MALLOC_SIZE_H */
