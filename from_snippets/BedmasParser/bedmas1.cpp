/*
Chapter: Own Project
Title: BEDMAS text equation to result calculator
Author: SRH
Date: 27/02/16
From: own pseudocode


*/

#include<iostream>
#include<string>
#include<sstream>
#include<vector>
#include<algorithm>
#include<math.h>

using namespace std;

union Entry
{
	char c;
	double d;
};

bool operator==(const Entry& lhs, const char rhs)
{
	return lhs.c == rhs;
}

bool operator!=(const Entry& lhs, const char rhs)
{
	return !(lhs == rhs);
}

bool isdp(char c) // check if character is a decimal point from the ASCII table
{
	int a = c;
	if (a == 46)
	{
		return true;
	}
	else
	{
		return false;
	}
}

bool isnum(char c) // check if character is a digit 0-9 from the ASCII table
{
	int a = c;
	if (a >= 48 && a <= 57)
	{
		return true;
	}
	else
	{
		return false;
	}
}

void groomInput(string& input) // remove all whitespace from input string
{
	input.erase(remove(input.begin(), input.end(), ' '), input.end());
}

double computePair(Entry e1, Entry e2,  Entry e3) // e1 (number) e2 (operation) e3 (number)
{
	switch (e2.c)
	{
		case '+':
			return e1.d + e3.d;
		case '-':
			return e1.d - e3.d;
		case '*':
			return e1.d * e3.d;
		case '/':
			return e1.d / e3.d;
		case '^':
			return pow(e1.d, e3.d);
	}
}

void formExpression(string input, vector<Entry>& pieces)
{
	groomInput(input); // remove all whitespace

	stringstream out;

	// create a vectors of the pieces of the equation, doubles and chars
	for (string::iterator it = input.begin(); it < input.end(); it++)
	{
		if (isnum(*it) || isdp(*it))
		{
			out << *it;
		}
		if ((!isnum(*it) && !isdp(*it)) || it == input.end() - 1)
		{
			if (!out.str().empty())
			{
				Entry entry1;
				entry1.d = stod(out.str());
				pieces.push_back(entry1);
				//cout << entry1.d;
			}

			if ((it != input.end()-1) || *it == ')')
			{
				Entry entry2;
				entry2.c = *it;
				pieces.push_back(entry2);
				//cout << entry2.c;
			}

			out.str("");
		}
	}
}




double Evaluate(string input) 
{
	vector<Entry> express;

	formExpression(input, express);

	while (express.size() > 1)
	{
		vector<Entry> subexpress;

		// * what does it do when there are no brackets? code pertinent to this case marked with a *

		int idx1 = 0;
		int idx2 = 0;

		if (find(express.begin(), express.end(), ')') != express.end())
		{
			vector<Entry>::iterator it1 = find(express.begin(), express.end(), ')');
			vector<Entry>::iterator it2 = it1;
			while (*it1 != '(') { it1--; }
			idx1 = it1 - express.begin();
			idx2 = it2 - express.begin();
			vector<Entry> temp(it1+1,it2);
			subexpress = temp;
		}
		else
		{
			subexpress = express;
		}

		while (subexpress.size() > 1)
		{
			// EDMAS - order of operations
			char op[5] = { '^', '/', '*', '+', '-' };

			for (int i = 0; i < 5; i++)
			{
				while (find(subexpress.begin(), subexpress.end(), op[i]) != subexpress.end())
				{
					vector<Entry>::iterator it = find(subexpress.begin(), subexpress.end(), op[i]);

					Entry entry;
					entry.d = computePair(*(it - 1), *it, *(it + 1));
					subexpress.insert(it + 2, entry); // this insert invalidates the previous iterator

					// use new iterator to perform erase on old section
					vector<Entry>::iterator itt = find(subexpress.begin(), subexpress.end(), op[i]);
					subexpress.erase(itt - 1, itt + 2);		
				}
			}
		}

		
		if (find(express.begin(), express.end(), ')') != express.end())
		{
			express.insert(express.begin() + idx2 + 1, subexpress[0]);
			express.erase(express.begin() + idx1, express.begin() + idx2 + 1);
		}
		else
		{
			express = subexpress;
		}		

	}

	return express[0].d;
}


int main()
{
	cout << "Allowed inputs: +,-,*,/,^,(,),0,1,2,3,4,5,6,7,8,9" << endl;
	cout << "Disallowed: starting equation with - to indicate negative number" << endl;
	
	string input = "(1*(2+3)^3)^2 + (3*((2)+3))";

	cout << Evaluate(input);

	cin.get();
	return 0;
}
