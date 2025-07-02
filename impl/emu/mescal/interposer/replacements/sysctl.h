//
//  sysctl.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _SYSCTL_H
#define _SYSCTL_H

#include <stddef.h>

#ifdef sysctl
#undef sysctl
#endif	/* sysctl */

__attribute__ ((noinline)) int sysctl(int *name, u_int namelen, void *oldp, size_t *oldlenp, void *newp, size_t newlen);

#endif	/* _SYSCTL_H */
