package main

import (
   "os"
   "fmt"
   //"strings"
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
func replace( dest []int, find int, replace int ) ( []int ) {
   changed := dest
   for i := 0; i < len(dest); i++ {
      if dest[i] == find { 
         changed[i] = replace
      }
   }
   return changed
}
