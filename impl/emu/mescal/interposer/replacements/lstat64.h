//
//  lstat64.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-04.
//

#ifndef _LSTAT64_H
#define _LSTAT64_H

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

#include <sys/stat.h>

#ifdef lstat64
#undef lstat64
#endif	/* lstat64 */

__attribute__ ((noinline)) int lstat64(const char *restrict path, struct stat64 *restrict buf) __asm__ ("_lstat$INODE64");

#endif	/* _LSTAT64_H */
