//
//  malloc_good_size.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _MALLOC_GOOD_SIZE_H
#define _MALLOC_GOOD_SIZE_H

#include <stddef.h>

#ifdef malloc_good_size
#undef malloc_good_size
#endif	/* malloc_good_size */

__attribute__ ((noinline)) size_t malloc_good_size(size_t size);

#endif	/* _MALLOC_GOOD_SIZE_H */
