# Ready Set (Go)

This is a simple, naive and non-production ready Set implementation for Golang.

## Purpose

For me to learn about interface{} in Golang.

## What's missing

A lot, for example:

* Set difference
* Set iteration
* Other cool set operations

## Usage

Should you decide to use this please consider the following:

*DON'T*

You still here? Use as follows:

    package main

    import (
      "fmt"
      "github.com/bulters/readyset"
    )

    func main() {
      s := set.New(1,2,3)
      fmt.Println(s.String())

      s2 := set.New(2,3,4)

      i := s.Intersect(s2)
      fmt.Println("i: " + i.String())

      u1 := s.Union(s2)
      fmt.Println("u1: " + u1.String())

      s2.Add(5)

      u2 := s.Union(s2)
      fmt.Println("u2: " + u2.String())
    }

## Feedback

Please send me feedback on my usage of the Go language, or basically anything you can think of (sans personal trait evaluations, I have my wife for that). Contact me via jeroen@bulte dot rs.

