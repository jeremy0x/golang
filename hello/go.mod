module example.com/hello

go 1.18

require example.com/greetings v0.0.0-00010101000000-000000000000

replace example.com/greetings => ../greetings

// notes: --> love how go handles modules, unused modules are removed when you tidy up the project. being able to have a module locally and also have it in a remote repository is a great feature.

// multiline comments are also supported in go, which is great for writing notes in the code.
// q: how do you write a multiline comment in go?
// a: use the /* */ syntax
