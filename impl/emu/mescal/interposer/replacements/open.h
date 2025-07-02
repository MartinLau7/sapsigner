//
//  open.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _OPEN_H
#define _OPEN_H

#ifdef open
#undef open
#endif	/* open */

__attribute__ ((noinline)) int open(const char *path, int oflag, ...);

#endif	/* _OPEN_H */
