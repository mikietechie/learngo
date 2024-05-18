package main

import "testing"

func Test1(t *testing.T) {
	if !IsPrime(1) {
		t.Errorf("1 is a prime number")
	}
}

func Test2(t *testing.T) {
	if !IsPrime(2) {
		t.Errorf("2 is a prime number")
	}
}

func Test10(t *testing.T) {
	if IsPrime(10) {
		t.Errorf("10 is a not prime number")
	}
}
