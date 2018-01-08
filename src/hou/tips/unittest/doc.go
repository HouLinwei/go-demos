package unittest

/*
Boolean values are comparable. Two boolean values are equal if they are either both true or both false.
Integer values are comparable and ordered, in the usual way.
Floating point values are comparable and ordered, as defined by the IEEE-754 standard.
Complex values are comparable. Two complex values u and v are equal if both real(u) == real(v) and imag(u) == imag(v).
String values are comparable and ordered, lexically byte-wise.
Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value nil. Pointers to distinct zero-size variables may or may not be equal.
Channel values are comparable. Two channel values are equal if they were created by the same call to make or if both have value nil.
Interface values are comparable. Two interface values are equal if they have identical dynamic types and equal dynamic values or if both have value nil.
A value x of non-interface type X and a value t of interface type T are comparable when values of type X are comparable and X implements T. They are equal if t's dynamic type is identical to X and t's dynamic value is equal to x.
Struct values are comparable if all their fields are comparable. Two struct values are equal if their corresponding non-blank fields are equal.
Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.
*/
