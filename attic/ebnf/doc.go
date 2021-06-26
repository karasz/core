/*
Package ebmf implements an ISO/IEC 14977
Extended Backus–Naur Form parser, verifiers,
and additional related helpers for AsciiGoat

A syntax highlighter for vim and a copy of the final draft of the standard
are included in the doc/ directory. The official standard can be downloaded from
http://standards.iso.org/ittf/PubliclyAvailableStandards/s026153_ISO_IEC_14977_1996(E).zip

An uberly simplified version of the EBNF grammar looks like:

  letter = "A" | "B" | "C" | "D" | "E" | "F" | "G"
         | "H" | "I" | "J" | "K" | "L" | "M" | "N"
         | "O" | "P" | "Q" | "R" | "S" | "T" | "U"
         | "V" | "W" | "X" | "Y" | "Z" ;
  digit = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;
  symbol = "[" | "]" | "{" | "}" | "(" | ")" | "<" | ">"
         | "'" | '"' | "=" | "|" | "." | "," | ";" ;
  character = letter | digit | symbol | "_" ;

  identifier = letter , { letter | digit | "_" } ;
  terminal = "'" , character , { character } , "'"
           | '"' , character , { character } , '"' ;

  lhs = identifier ;
  rhs = identifier
      | terminal
      | "[" , rhs , "]"
      | "{" , rhs , "}"
      | "(" , rhs , ")"
      | rhs , "|" , rhs
      | rhs , "," , rhs ;

  rule = lhs , "=" , rhs , ";" ;
  grammar = { rule } ;
*/
package ebnf
