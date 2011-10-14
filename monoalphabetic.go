package main

import (
   "os"
   "fmt"
)

func main () {
   var cipher string

/* Reads in a string from a file and stores it in var cipher.   */
   if len(os.Args) != 2 {
      fmt.Fprintf( os.Stderr, "decrypt: Enter valid file name \n")
   } else if st, err := os.Stat( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "decrypt: %s\n", err );
   } else if f, err := os.Open( os.Args[1] ); err != nil {
      fmt.Fprintf( os.Stderr, "dercypt: %s\n", err);
   } else {
      buf := make( []byte, st.Size )
      f.Read( buf )
      cipher = string(buf)
      fmt.Println( cipher )
   }


/* Calculates the frequency of each letter in the string and stores it in a map.    */
   frequency := make(map[string] int, len(cipher))

   for i := 0; i < len(cipher); i++ {
      frequency[string(cipher[i])] = letterfrequency(cipher, string(cipher[i]))
   }

/* Prints each letter with its corrosponding frequency.              */
   for key, value := range frequency {
      fmt.Printf("%s: %d\n", key, value)
   }

/* Replaces every occurence of "Y" with "E" in the string cipher.    */
   cipher = replace( []byte(cipher), "Y", "E" )
   fmt.Println( cipher )


}

/* Takes a string and a letter as input, and returns the count
of the number of times the letter appears in the string.    */
func letterfrequency( dest string, find string ) ( int ) {
   count := 0
   for i := 0; i < len(dest); i++ {
      if dest[i] == find[0] { 
         count++
      }
   }
   return count
}

/* Takes a byte array and two letters as input, and returns a string with 
every occurence of the first letter replaced with the second letter.   */
func replace( dest []byte, find string, replace string ) ( string ) {
   changed := dest
   for i := 0; i < len(dest); i++ {
      if dest[i] == find[0] { 
         changed[i] = replace[0]
      }
   }
   return string(changed)
}



