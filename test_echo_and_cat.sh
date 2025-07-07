#!/usr/bin/env bash
./echo/echo "the quick brown fox" >fox.txt
./echo/echo "jumps over the lazy dog" >dog.txt
./cat/cat fox.txt dog.txt
rm fox.txt
rm dog.txt
