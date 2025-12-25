#include "OutputRedirect.h"

OutputRedirect::OutputRedirect() {
    oldBuffer = cout.rdbuf();
    cout.rdbuf(newBuffer.rdbuf());
}

OutputRedirect::~OutputRedirect() {
    cout.rdbuf(oldBuffer);
}

string OutputRedirect::getOutput() const {
    return newBuffer.str();
}

void OutputRedirect::clear() {
    newBuffer.str("");
    newBuffer.clear();
}



