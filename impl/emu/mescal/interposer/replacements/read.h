//
//  read.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _READ_H
#define _READ_H

#include <stddef.h>

#ifdef read
#undef read
#endif	/* read */

__attribute__ ((noinline)) ssize_t read(int fildes, void *buf, size_t nbyte);

#endif	/* _READ_H */
