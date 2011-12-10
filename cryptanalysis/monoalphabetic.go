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

/* Things to do: max function for maps */


package main

import (
   "os"
   "fmt"
   //"io"
)

type IntMap map[int] int
type StringMap map[string] int

type MaxMap interface {
   Ma() int
}

func (s IntMap ) Len() int                      { return len(s) }
func (s IntMap ) Less( i, j int ) bool          { return s[i] < s[j] }
func (s IntMap ) Swap( i, j int )               { s[i], s[j] = s[j], s[i] }

func (s StringMap ) Len() int                   { return len(s) }
func (s StringMap ) Less( i, j string ) bool    { return s[i] < s[j] }
func (s StringMap ) Swap( i, j string )         { s[i], s[j] = s[j], s[i] }
*/

func main () {
   var ciphertext []int

   /* Reads in a string from a file and stores it in var cipher.   */
   if len(os.Args) < 2 {
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
   singleFrequency := make( IntMap )
   letterFrequency( ciphertext, singleFrequency )
   //sort( singleFrequency )

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
   doubleFrequency := make( StringMap )
   twoLetterFrequency( string( ciphertext ), doubleFrequency )
   //sort( doubleFrequency )

   /* Prints a letter with its corresponding frequency, for 
   all letters that appear more than nine times          */
   for key, value := range doubleFrequency {
      if value > 9 {
         fmt.Printf( "%s: %d\n", key, value )
      }
   }

   /* To make the formating pretty. */
   fmt.Println( "\n" )

   /* Calculates the frequency of each sequence of three letters 
   in the string and stores it in a map.                      */
   tripleFrequency := make( StringMap )
   threeLetterFrequency( string( ciphertext ), tripleFrequency )
   //sort( tripleFrequency )

   /* Prints a letter with its corresponding frequency, for 
   all letters that appear more than nine times.         */
   for key, value := range tripleFrequency {
      if value > 7 {
         fmt.Printf( "%s: %d\n", key, value )
      }
   }


   changed := make( []int, len( ciphertext ) )
   copy( changed, ciphertext )

   fmt.Println( "\n" )
   var input string


/* Function to replace letters from the command line */

   fmt.Println("You can use this prompt to replace letters in your ciphertext. Remember to use upper case letter and ensure that there are no spaces between entries")
   fmt.Println("To replace  R:[Dest][Source]" )
   fmt.Println("To Quit     Q" )
   fmt.Println("")

   for {
      fmt.Fscanln( os.Stdin, &input )

      if input == "help" || input == "Help" {
         fmt.Println(" To Replace letters type R followed by the replacement without spaces: R:[Dest][Source]" )
         fmt.Println(" To Quit: Q" )
         } else if input[0] == 'q' || input[0] == 'Q' {
            break;
      } else if input[0] == 'r' || input[0] == 'R' {
         temp := []int(input)
         replace( changed, ciphertext, temp[1], temp[2] )
      } else {
         fmt.Println( "Enter a valid input. Type Help for more information" )
      }
   }


/* replace( changed, ciphertext, 'N', 'T')
   replace( changed, ciphertext, 'Z', 'H')
   replace( changed, ciphertext, 'Y', 'E')
   replace( changed, ciphertext, 'R', 'S')
   replace( changed, ciphertext, 'W', 'I')
   replace( changed, ciphertext, 'K', 'O')
   replace( changed, ciphertext, 'D', 'N')
   replace( changed, ciphertext, 'J', 'G')
   replace( changed, ciphertext, 'C', 'M')
   replace( changed, ciphertext, 'H', 'W')
   replace( changed, ciphertext, 'S', 'L')
   replace( changed, ciphertext, 'L', 'C')
   replace( changed, ciphertext, 'G', 'F')
   replace( changed, ciphertext, 'X', 'R')
   replace( changed, ciphertext, 'F', 'Y')
   replace( changed, ciphertext, 'E', 'P')
   replace( changed, ciphertext, 'Q', 'U')
   replace( changed, ciphertext, 'P', 'X')
   replace( changed, ciphertext, 'I', 'B')
   replace( changed, ciphertext, 'M', 'V')
   replace( changed, ciphertext, 'B', 'D')
   replace( changed, ciphertext, 'T', 'K')
*/
   fmt.Println( string( changed ))
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
func replace( changed, dest []int, find, replace int ) {
   for i := 0; i < len(dest); i++ {
      if dest[i] == find {
         changed[i] = replace
      }
   }
}

func mapMax(k MapKey, m MapMax ) {
   for i := 1; i < a.Len(); i++ {
      for j := i; j > 0 && m.Less(j, j-1); j-- {
         data.Swap(j, j-1)
         }
      }
}
