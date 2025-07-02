//
//  __memcpy_chk.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-01.
//

#ifndef ___MEMCPY_CHK_H
#define ___MEMCPY_CHK_H

#include <stddef.h>

#ifdef __memcpy_chk
#undef __memcpy_chk
#endif	/* __memcpy_chk */

__attribute__ ((noinline)) char *__memcpy_chk(void *restrict dst, const void *restrict src, size_t len, size_t dstlen);

#endif	/* ___MEMCPY_CHK_H */
