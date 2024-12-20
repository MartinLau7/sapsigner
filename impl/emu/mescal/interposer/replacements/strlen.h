//
//  strlen.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _STRLEN_H
#define _STRLEN_H

#include <stddef.h>

#ifdef strlen
#undef strlen
#endif	/* strlen */

__attribute__ ((noinline)) size_t strlen(const char *s);

#endif	/* _STRLEN_H */
