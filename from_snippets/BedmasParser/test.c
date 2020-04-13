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

struct Entry {
	char c;
	double d;
};

int main() {

//	arg = "2234*((1+2)^2 + (3-5)*2)"

	int debug = 0;

	char *str = "123*(2+3)^4";
	
//	printf("%lu", sizeof(str));
	double result = 0.0;
	
	// Make Entries from the input string

	struct Entry entries[100];

	int k;	
	for (k=0;k<4;k+=1) {

		char *i = str;
		double ret;
		ret = strtod(str+i, &i);
		printf("%d \"%s\"\n", (int)(ret), i);

		char symbols[7] = {'*', '/', '+', '-', '^', '(', ')'};
		
		for( int j=0; j<sizeof(symbols); j+=1 ) {
			if (str[i] == symbols[j]) {
				printf("%c", str[i]);
				i+=1;
				break;
			}
		}
		
	}



	return 0;
}

