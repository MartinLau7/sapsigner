//
//  reallocf.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _REALLOCF_H
#define _REALLOCF_H

#include <stddef.h>

#ifdef reallocf
#undef reallocf
#endif	/* reallocf */

__attribute__ ((noinline)) void *reallocf(void *ptr, size_t size);

#endif	/* _REALLOCF_H */
