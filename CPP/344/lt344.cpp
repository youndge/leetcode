#include <iostream>
#include <vector>
#include "344.h"
using namespace std;

int test344() {
	vector<char> s;
	s.push_back('H');
	s.push_back('e');
	s.push_back('l');
	s.push_back('l');
	s.push_back('o');
	s.push_back('!');


	Solution sl;
	for (int i = 0; i < s.size(); i++) {
		cout << s[i] << " ";
	}
	cout << endl;
	sl.reverseString(s);
	for (int i = 0; i < s.size(); i++) {
		cout << s[i] << " ";
	}
	cout << endl;
	return 0;
}