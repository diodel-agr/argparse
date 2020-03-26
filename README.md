# argparse
Command line arguments parser for Go.

This package is used to ease the reading of command line arguments and the variable instantiation.
Let's suppose that your program needs to read several variables from the command line: an integer, a string and a slice of floating point numbers.
Usually you should read the first argument, parse it to integer and check whether the parsing succeeded. For the slice, you should instantiate an empty slice, and then read each parameter, converting it to float, checking that the paring succeeded, then appending it to the slice.
To ease this step, the ```argparse package``` takes care of the reading, converting and checking of command line arguments.

#### Description
This package can deal with all the basic types of variables, pointers or slices found in Go:
- int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
- rune, byte.
- float32, float64.
- complex64, complex128.
- bool, string.

In order for ```argparse``` to parse your arguments, a ```format``` has to be specified before all arguments, specifying the expected types.
For the previous example (an integer, a string and a slice of floating point numbers), the format is: ```is[]5f32``` where:
- ```i``` specifies that the first argument is an integer.
- ```s``` specifies that the second argument is a string.
- ```[]5f32``` specifies that the last ```5``` arguments is a slice of 5 variables of float32 type.

#### Usage
In your program, the first line should be: ```variableList := argparse.ParseArgList()```.
The result of this function is a ```*list.List``` object with all the variables.
If an error occurs, the error message will be printed and the result is ```nil```.

The list can be printed like this:
```
if listOfVariables != nil {
    fmt.Println("The result: ")
    for i := 0; i < listOfVariables.Len(); i++ {
        el := argparse.GetVariableAt(listOfVariables, i)
        fmt.Println(i, ": ", el, "(", reflect.TypeOf(el), ")")
    }
}
```
#### Result
For this input: ```[]3i[]4s[]2f32 10 20 30 hello how are you? 3.14259 2.71``` the output is this:
```
The result:
0 :  [10 20 30] ( []int )
1 :  [hello, how are you?] ( []string )
2 :  [3.14259 2.71] ( []float32 )
```

#### Specifying a format
A ```format``` is composed out of 1 or more ```specifiers```. A specifier describes the type of a variable.
A ```specifier``` has the following general structure: ```[] n * t x```, where:
- ```[]``` specifies if the variable is a slice.
- ```n``` a number, specifies the slice size. ```n``` has to be equal to the number of arguments of the specified type passed to the command line. Otherwise an error will occur. If ```n``` is not specified, an empty slice will be created.
- ```*``` specifies that the variable is a pointer to the specified type.
- ```t``` specifies the type.
- ```x``` a number, specifies the bit size of the specified type.

The type ```t``` may have one of the following values:
- ```i``` for ```int``` or ```ui``` for ```unsigned int```. Valid bit sizes {8, 16, 32, 64}.
- ```r``` for ```rune```, ```by``` for ```byte```.
- ```f``` for ```float```. Valid bit sizes {32, 64}.
- ```c``` for ```complex```. Valid bit sizes {64, 128}.
- ```b``` for ```bool``` and ```s``` for ```string```.
