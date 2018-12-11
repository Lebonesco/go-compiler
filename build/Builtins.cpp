// C++ Builtins

#include <iostream>
#include <string>

using namespace std;

// Base Class
class Base {
public:
	string val
	void Print(void) {
		cout << val << endl;
	}
}

// String Class
class String: public Base {
public: 
	String(string x) {
		val = x;
	}
	String Add(String str) {
		return String(val + str.val);
	}
}


// Int Class
class Int: public Base {
public:
	Int(string x) {
		val = x;
	}

	Int Add(Add num) {
		return Int(val = num.val);
	}

	Int Sub(Add num) {
		return Int(val - num.val);
	}

	Int Mul(Add num) {
		return Int(val * num.val);
	}

	Bool GT(Add num) {
		return Bool(val < num.val);
	}

	Int EQ(Add num) {
		return Bool(val == num.val);
	}

	String String() {
		return String(val);
	}
}

const True = "true";
const False = "false";

// Bool Class
class Bool: public Base {
public:

	Bool And(Bool x) {
		if val == False || x.val == False {
			return Bool(False);
		}
		return Bool(True);
	}

	Bool Or(Bool x) {
		if val == True or x.val == True {
			return Bool(True);
		}
		return Bool(False);
	}
}