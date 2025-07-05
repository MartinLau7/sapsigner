//
//  dlsym.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-02.
//

#ifndef _DLSYM_H
#define _DLSYM_H

#ifdef dlsym
#undef dlsym
#endif	/* dlsym */

__attribute__ ((noinline)) void *dlsym(void* handle, const char* symbol);

#endif	/* _DLSYM_H */
