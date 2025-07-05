//
//  __error.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-04.
//

#ifndef ___ERROR_H
#define ___ERROR_H

#ifdef __error
#undef __error
#endif	/* __error */

__attribute__ ((noinline)) int *__error(void);

#endif	/* ___ERROR_H */
