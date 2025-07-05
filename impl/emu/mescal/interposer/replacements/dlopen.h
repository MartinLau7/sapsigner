//
//  dlopen.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-02.
//

#ifndef _DLOPEN_H
#define _DLOPEN_H

#ifdef dlopen
#undef dlopen
#endif	/* dlopen */

__attribute__ ((noinline)) void *dlopen(const char* path, int mode);

#endif	/* _DLOPEN_H */
