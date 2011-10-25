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
package main

import (
   "os"
   "fmt"
   "strings"
   )

func main() {
   var ciphertext []int

   /* Reads the ciphertext from input file */
   if len(os.Args) != 3 {
      fmt.Fprintf( os.Stderr, "vigerene: Enter valid filename \n" );
   } else if st, err := os.Stat( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "vigenere: %s\n", err );
   } else if f, err := os.Open( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "vigenere: %s\n", err);
   } else {
      buf := make( []byte, st.Size );
      f.Read( buf )
      // Our substitution function can not deal with lower case characters.
      temp := strings.ToLower(string(buf))
      ciphertext = []int(temp)
      fmt.Println( string( ciphertext ))
   }


   /*fmt.Println( "Enter the length of your key" )
   type KeyLength int
   fmt.Fscan( os.Stdin, keylength ) */

   key := make( []int, 3 )
   key[0] = 15
   key[1] = 14
   key[2] = 4

   substitute( ciphertext, key )
   fmt.Println( string( ciphertext ))
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

