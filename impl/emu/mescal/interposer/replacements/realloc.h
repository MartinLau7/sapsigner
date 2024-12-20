//
//  realloc.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _REALLOC_H
#define _REALLOC_H

#include <stddef.h>

#ifdef realloc
#undef realloc
#endif	/* realloc */

__attribute__ ((noinline)) void *realloc(void *old_ptr, size_t new_size);

#endif	/* _REALLOC_H */
