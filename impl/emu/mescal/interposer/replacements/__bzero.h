//
//  __bzero.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef ___BZERO_H
#define ___BZERO_H

#include <stddef.h>

#ifdef __bzero
#undef __bzero
#endif	/* __bzero */

__attribute__ ((noinline)) void __bzero(void *s, size_t n);

#endif	/* ___BZERO_H */
