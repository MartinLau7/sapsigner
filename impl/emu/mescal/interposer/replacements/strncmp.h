//
//  strncmp.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _STRNCMP_H
#define _STRNCMP_H

#ifdef strncmp
#undef strncmp
#endif	/* strncmp */

__attribute__ ((noinline)) int strncmp(const char *s1, const char *s2, size_t n);

#endif	/* _STRNCMP_H */
