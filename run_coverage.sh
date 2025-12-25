#!/bin/bash

FRAMEWORK=$1
EXECUTABLE=""
REPORT_DIR=""

case $FRAMEWORK in
    "boost")
        EXECUTABLE="./BoostTest"
        REPORT_DIR="boost_report"
        ;;
    "gtest")
        EXECUTABLE="./GoogleTest"
        REPORT_DIR="gtest_report"
        ;;
    "catch2")
        EXECUTABLE="./Catch2Test"
        REPORT_DIR="catch2_report"
        ;;
esac

$EXECUTABLE

cd CMakeFiles/${EXECUTABLE:2}.dir/tests || exit

lcov --ignore-errors mismatch -c -d . -o main.info
lcov --ignore-errors mismatch -c -d ../src/Array -o array.info
lcov --ignore-errors mismatch -c -d ../src/Stack -o stack.info
lcov --ignore-errors mismatch -c -d ../src/Queue -o queue.info
lcov --ignore-errors mismatch -c -d ../src/SLL -o sll.info
lcov --ignore-errors mismatch -c -d ../src/DLL -o dll.info
lcov --ignore-errors mismatch -c -d ../src/FBT -o fbt.info
lcov --ignore-errors mismatch -c -d ../src/ChainingHash -o ch.info
lcov --ignore-errors mismatch -c -d ../src/OpenAddrHash -o oah.info

lcov -a main.info -a array.info -a stack.info -a queue.info -a sll.info \
      -a dll.info -a fbt.info -a ch.info -a oah.info \
      -o total.info

lcov -r total.info \
     '/usr/*' \
     '*/tests/*' \
     '*/CMakeFiles/*' \
     '*/External Libraries/*' \
     '*/Scratches and Consoles/*' \
     '*/cmake-build-*/*' \
     '*/reports/*' \
     '*/main.cpp' \
     -o total.info

genhtml total.info --output-directory $REPORT_DIR
