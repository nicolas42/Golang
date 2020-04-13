

#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>


#include <ctype.h>
// includes isspace and isdigit and others
// isalpha(c) non-zero if c is alphabetic, 0 if not
// isupper(c) non-zero if c is upper case, 0 if not
// islower(c) non-zero if c is lower case, 0 if not
// isdigit(c) non-zero if c is digit, 0 if not
// isalnum(c) non-zero if isalpha(c) or isdigit(c), 0 if not
// isspace(c) non-zero if c is blank, tab, newline, return, formfeed, vertical tab toupper(c) return c converted to upper case
// tolower(c) return c converted to lower case


int main() {

	char *str = "123*(2+3)^4";
	char *cp = 0;
	
	printf("%s %d\n", str+(int)cp, (int)cp);
	
//	printf("%lu", sizeof(str));
	
	for (int k=0;k<sizeof(str);k+=(int)i) {

		char *i = str;
		double ret;
		ret = strtod(str+(int)i, &i);
		printf("%d \"%s\"\n", (int)(ret), i);
		i++;		
	}



	return 0;
}

