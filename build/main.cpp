#include <string>
#include <iostream>
#include "Builtins.cpp"

int main() {
Int tmp_1 = Int(0);
Int x = tmp_1;
if ("true" == Bool("true").val) {
Int tmp_2 = Int(5);
x = tmp_2;
} else {
Int tmp_3 = Int(6);
x = tmp_3;
}

return 0;
}