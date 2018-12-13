// C++ Builtins

#include <iostream>
#include <string>

using namespace std;


// Nothing Class
class Nothing {
	
};

// Base Class
class Base {
public:
	string val;
	Nothing PRINT(void) {
		cout << val << endl;
	}
};

const string True = "true";
const string False = "false";

// Bool Class
class Bool: public Base {
public:
	Bool(string x) {
		val = x;
	}

	Bool And(Bool x) {
		if (val == False || x.val == False) {
			return Bool(False);
		}
		return Bool(False);
	}

	Bool Or(Bool x) {
		if (val == True || x.val == True) {
			return Bool(True);
		}
		return Bool(False);
	}
};

// String Class
class String: public Base {
public: 
	String(string x) {
		val = x;
	}
	String PLUS(String str) {
		return String(val + str.val);
	}

	Bool EQ(String str) {
		if (val == str.val) {
			return Bool(True);
		} else {
			return Bool(False);
		}
	}
};

// Int Class
class Int: public Base {
public:
	int valInt;
	Int(int x) {
		val = to_string(x);
		valInt = x;
	}

	Int PLUS(Int num) {
		return Int(valInt + num.valInt);
	}

	Int MINUS(Int num) {
		return Int(valInt - num.valInt);
	}

	Int TIMES(Int num) {
		return Int(valInt * num.valInt);
	}

	Bool GT(Int num) {
		if (valInt < num.valInt) {
			return Bool(False);
		} else {
			return Bool(True);
		}
	}

	Bool EQ(Int num) {
		if (valInt == num.valInt) {
			return Bool(True);
		}
		return Bool(False);
	}

	String Stringify() {
		return String(val);
	}
};