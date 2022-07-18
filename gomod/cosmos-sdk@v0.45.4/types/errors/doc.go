 div, nil
   }

The first time an error instance is wrapped a stacktrace is attached as well.
Stacktrace information can be printed using %+v and %v formats.

  %s  is just the error message
  %+v is the full stack trace
  %v  appends a compressed [filename:line] where the error was created

*/
package errors
