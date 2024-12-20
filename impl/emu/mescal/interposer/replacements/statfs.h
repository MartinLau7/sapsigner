//
//  statfs.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _STATFS_H
#define _STATFS_H

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

#ifdef statfs
#undef statfs
#endif	/* statfs */

__attribute__ ((noinline)) int statfs(const char *path, struct statfs *buf) __asm__ ("_statfs");

#endif	/* _STATFS_H */
