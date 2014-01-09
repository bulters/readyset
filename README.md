# Ready Set (Go)

This is a simple, naive and non-production ready Set implementation for Golang.

## Purpose

For me to learn about interface{} in Golang.

## Usage

Should you decide to use this please consider the following:

*DON'T*

You still here? Use as follows:

    package main

    import (
      "github.com/bulters/readyset"
    )

    func main() {
      s1 := set.New(1,2,3)
      s2 := set.New(2,3,4)
      i := s1.Intersect(s2)
      u1 := s1.Union(s2)
      s2.Add(5)
      u2 := s1.Union(s2)
    }

## Feedback

Please send me feedback on my usage of the Go language, or basically anything you can think of (sans personal trait evaluations, I have my wife for that). Contact me via jeroen@bulte dot rs.

