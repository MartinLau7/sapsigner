//
//  statfs64.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-15.
//

#ifndef _STATFS64_H
#define _STATFS64_H

#ifdef __DARWIN_64_BIT_INO_T
#undef __DARWIN_64_BIT_INO_T
#endif	/* __DARWIN_64_BIT_INO_T */
#define __DARWIN_64_BIT_INO_T 0

#ifdef __DARWIN_ONLY_64_BIT_INO_T
#undef __DARWIN_ONLY_64_BIT_INO_T
#endif	/* __DARWIN_ONLY_64_BIT_INO_T */
#define __DARWIN_ONLY_64_BIT_INO_T 0

#ifdef __DARWIN_INODE64
#undef __DARWIN_INODE64
#endif	/* __DARWIN_INODE64 */
#define __DARWIN_INODE64(...)

#include <sys/mount.h>

#ifdef statfs64
#undef statfs64
#endif	/* statfs64 */

__attribute__ ((noinline)) int statfs64(const char *path, struct statfs64 *buf) __asm__ ("_statfs$INODE64");

#endif	/* _STATFS64_H */
