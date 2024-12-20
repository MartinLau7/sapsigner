//
//  __memset_chk.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef ___MEMSET_CHK_H
#define ___MEMSET_CHK_H

#include <stddef.h>

#ifdef __memset_chk
#undef __memset_chk
#endif	/* __memset_chk */

__attribute__ ((noinline)) void *__memset_chk(void *dest, int val, size_t len, size_t dstlen);

#endif	/* ___MEMSET_CHK_H */
