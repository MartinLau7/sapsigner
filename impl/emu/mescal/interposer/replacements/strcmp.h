//
//  strcmp.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _STRCMP_H
#define _STRCMP_H

#ifdef strcmp
#undef strcmp
#endif	/* strcmp */

__attribute__ ((noinline)) int strcmp(const char *s1, const char *s2);

#endif	/* _STRCMP_H */
