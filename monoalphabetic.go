/****************************************************************************
 * This file contains a program that provides the user with some basic      *  
 * very cryptanalysis tools. It takes as input a text file that contains    *
 * the ciphertext.                                                          *
 *                                                                          *
 * You will need the Go compiler to run this file. Vist http://golang.org   *
 * for more information.                                                    *
 *                                                                          *
 * This program is free software: you can redistribute and/or modify this   *
 * file under the terms of the GNU General Public License as published by   *
 * the Free Software Foundation, either version 3 of the License, or        *
 * (at your option) any later version.                                      *                                             
 *                                                                          *
 * The file distributed in the hope that it will be useful, but WITHOUT     *
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or    *
 * FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License    *
 * for more details.                                                        *
 *                                                                          *
 * Copyright (C) 2011 Akanksha Vyas                                         *
 *                                                                          *
 ****************************************************************************/

package main

import (
   "os"
   "fmt"
   "io"
)

func main () {
   var ciphertext []int

   /* Reads in a string from a file and stores it in var cipher.   */
   if len(os.Args) != 2 {
      fmt.Fprintf( os.Stderr, "decrypt: Enter valid file name \n")
   } else if st, err := os.Stat( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "decrypt: %s\n", err );
   } else if f, err := os.Open( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "dercypt: %s\n", err );
   } else {
      buf := make( []byte, st.Size )
      f.Read( buf )
      ciphertext = []int(string(buf))
      fmt.Println( string( ciphertext ))
   }



   /* Calculates the frequency of each letter in the string and stores it in a map.    */
   singleFrequency := make( map[int] int, 64)
   letterFrequency( ciphertext, singleFrequency )

   /* Prints a letter with its corresponding frequency, for 
   all letters that appear more than once.               */
   for key, value := range singleFrequency {
      if value > 1 {
         fmt.Printf( "%s: %d\n", string( key ), value )
      }
   }

   /* To make the formating pretty. */
   fmt.Println( "\n" )


   /* Calculates the frequency of each sequence of two letters 
   in the string and stores it in a map.                    */
   doubleFrequency := make( map[string] int)
   twoLetterFrequency( string( ciphertext ), doubleFrequency )


   /* Prints a letter with its corresponding frequency, for 
   all letters that appear more than nine times          */
   tripleFrequency := make( map[string] int)
   for key, value := range doubleFrequency {
      if value > 9 {
         fmt.Printf( "%s: %d\n", key, value )
      }
   }

   /* To make the formating pretty. */
   fmt.Println( "\n" )

   /* Calculates the frequency of each sequence of three letters 
   in the string and stores it in a map.                      */
   threeLetterFrequency( string( ciphertext ), tripleFrequency )

   /* Prints a letter with its corresponding frequency, for 
   all letters that appear more than nine times.         */
   for key, value := range tripleFrequency {
      if value > 9 {
         fmt.Printf( "%s: %d\n", key, value )
      }
   }

   fmt.Println( "\n" )

   for {
      fmt.Println(" lalala ")
      if temp, err := io.ReadByte( os.Stdin ); err != nil {
         fmt.Fprintf( os.Stderr, "decrypt: %s\n", err )
      } else if temp == 'q' {
         fmt.Println(string(temp))
         os.Exit(0)
      } /*else if temp[0] == 'r' {
         stringReplace := []int( string( temp ))
         replace( ciphertext, stringReplace[1], stringReplace[2] )
         fmt.Println( string( ciphertext ))
      }*/
   }

  // replace( ciphertext, 'N', 'T')
  // replace( ciphertext, 'Z', 'H')
  // replace( ciphertext, 'Y', 'E')
   fmt.Println( string( ciphertext ))
}

/* Takes a string and a map as input, and maps each letter 
in to the number of times it appears in the string.     */
func letterFrequency( dest []int, frequency map[int] int ) {
   for i := 0; i < len(dest); i++ {
      frequency[dest[i]]++
   }
}

/* Takes a string and a map as input, and maps each sequence of two 
letters in to the number of times they appears in the string.    */
func twoLetterFrequency( dest string, frequency map[string] int ) {
   for i := 0; i < len(dest)-1; i++ {
     frequency[dest[i:i+2]]++
   }
}


/* Takes a string and a map as input, and maps each letter 
in to the number of times it appears in the string.     */
func threeLetterFrequency( dest string, frequency map[string] int ) {
   for i := 0; i < len(dest)-2; i++ {
     frequency[dest[i:i+3]]++
   }
}

/* Takes a byte array and two letters as input, and returns a string with 
every occurrence of the first letter replaced with the second letter.   */
func replace( dest []int, find int, replace int ) {
   changed := dest
   for i := 0; i < len(dest); i++ {
      if dest[i] == find { 
         changed[i] = replace
      }
   }
}
