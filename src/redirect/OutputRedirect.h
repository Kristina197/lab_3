#pragma once

#include <iostream>
#include <sstream>
#include <streambuf>

using namespace std;

class OutputRedirect {
    streambuf* oldBuffer;
    stringstream newBuffer;

public:
    OutputRedirect();
    ~OutputRedirect();
    string getOutput() const;
    void clear();
};