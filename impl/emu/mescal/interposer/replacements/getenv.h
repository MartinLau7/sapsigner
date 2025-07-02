//
//  getenv.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _GETENV_H
#define _GETENV_H

#ifdef getenv
#undef getenv
#endif	/* getenv */

__attribute__ ((noinline)) char *getenv(const char *name);

#endif	/* _GETENV_H */
