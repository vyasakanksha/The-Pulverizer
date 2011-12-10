/****************************************************************************
 * This file contains a program that provides the user with simple          *  
 * decrypting tools for a Vigenere Cipher. It takes a text file containing  *
 * the ciphertext as input. After reading the input it asks the user for    *
 * the key. The program assumes that only letters are encrypted, and leaves *
 * numbers and puntuation as they were.                                     *
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

/* Things to do :: Figure out how to account for multiple input files. 
                   Figure out how to seperate a random size key      
*/


package main

import (
   "os"
   "fmt"
   "strings"
   )

func main() {
   var ciphertext []int
   var keylen []int

   /* Reads the ciphertext from input file */
   if len(os.Args) < 2 {
      fmt.Fprintf( os.Stderr, "vigerene: Enter valid filename \n" );
   } else if st, err := os.Stat( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "vigenere: %s\n", err );
   } else if f, err := os.Open( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "vigenere: %s\n", err);
   } else {
      buf := make( []byte, st.Size )
      f.Read( buf )
      // Our substitution function can not deal with lower case characters.
      temp := strings.ToLower(string(buf))
      ciphertext = []int(temp)
      fmt.Println( string( ciphertext ))
   }

   for {
      fmt.Printf( "Enter your keylen followed by your key in this format: <keylength key>")
      fmt.Println("Eg: for key ABTD")
      fmt.Println{"4 ABTD")
      key := make( []int, keylen[0] )
      fmt.Println( "Enter your key with each number in a new line. Eg:" )
      fmt.Println( "23" )
      fmt.Println( "14" )
      fmt.Println( "2" )
      fmt.Println( "" )

      for i := 0; i < keylen[0]; i++ {
         fmt.Scanf( "%d", key[1] )
      }

      substitute( ciphertext, key )
      fmt.Println( string( ciphertext ))
   } else {
      fmt.Println( "Your key is empty" )
      }

}

/*Takes a messege and a key as input and runs the vigenere cipher.*/
func substitute( ciphertext, key []int ) {
   j := 0
   for i :=0; i < len(ciphertext); i++ {
      // Checking to ensure only lower case letters get replaced
      if ciphertext[i] > 96 && ciphertext[i] < 123 {
         temp := ciphertext[i] - key[j]
         // Addition mod 26
         if temp < 97 {
            temp += 26
         }
         ciphertext[i] = temp
         j++
      }

      if j == len(key) {
         j = 0
      }
   }
}

