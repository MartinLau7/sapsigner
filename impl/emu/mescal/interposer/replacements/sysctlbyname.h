//
//  sysctlbyname.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _SYSCTLBYNAME_H
#define _SYSCTLBYNAME_H

#include <stddef.h>

#ifdef sysctlbyname
#undef sysctlbyname
#endif	/* sysctlbyname */

__attribute__ ((noinline)) int sysctlbyname(const char *name, void *oldp, size_t *oldlenp, void *newp, size_t newlen);

#endif	/* _SYSCTLBYNAME_H */
