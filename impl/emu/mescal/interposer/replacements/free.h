//
//  free.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _FREE_H
#define _FREE_H

#ifdef free
#undef free
#endif	/* free */

__attribute__ ((noinline)) void free(void *ptr);

#endif	/* _FREE_H */
