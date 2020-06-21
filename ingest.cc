#include <iostream>
#include <vector>
#include <list>
#include <string>

using namespace std;

template<typename C, typename P> //requires Sequence<C> && Callable<P, value_type<P>>
int count(const C& c, P pred)
{
	int cnt = 0;
	for(const auto& x: c)
		if(pred(x))
			++cnt;
	return cnt;
}


void f(const vector<int>& vec, const list<string>& lst, int x, const string& s)
{
	cout << "number of values less than " << x
	<< ": " << count(vec, [&](int a) { return a < x; })
	<< '\n';

	cout << "number of values less than " << s
	<< ": " << count(lst, [&](const string& a) { return a < s; })
	<< '\n';
}

int main() {

	f(vector{1,2,200}, list<string>{"a"s, "b"s}, 10, "b"s);

}