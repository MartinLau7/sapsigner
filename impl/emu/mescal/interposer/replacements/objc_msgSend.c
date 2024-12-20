// #include "strcmp.h"

#include "objc_msgSend.h"

id objc_msgSend(id self, SEL op, ...) {
    const char *sel_name = (const char *) op;
    if (
        false
        // || !strcmp((const char *) op, "objectForKey:")
        || (sel_name[0] == 'o' && sel_name[1] == 'b' && sel_name[2] == 'j' && sel_name[3] == 'e' && sel_name[4] == 'c' && sel_name[5] == 't' && sel_name[6] == 'F' && sel_name[7] == 'o' && sel_name[8] == 'r' && sel_name[9] == 'K' && sel_name[10] == 'e' && sel_name[11] == 'y' && sel_name[12] == ':' && sel_name[13] == '\0')
    ) {
        return (void *) ~0;
    }

    return nil;

    (void) self;
    (void) op;
}
