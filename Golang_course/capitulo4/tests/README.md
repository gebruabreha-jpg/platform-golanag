# Test Suite for Chapter 4: Tipos Básicos (Basic Types)

This directory contains a comprehensive test suite for all examples and concepts presented in Chapter 4 of the Go course.

## Overview

The test suite validates all code examples from Chapter 4, covering:
- Constants declarations (typed and untyped, iota usage)
- Variable declarations and initialization
- Boolean types and operators
- Numeric types (integers, floats, complex numbers)
- Pointer operations
- String handling and UTF-8 support

## Test Files

### `constants_test.go` - Section 2: Constantes (8 tests)
Tests for constant declarations and usage:
- Explicit and implicit constant declarations
- Factorized constants
- Constant value reuse
- iota usage (basic and with expressions)
- Month constants pattern (starting at 1)

**Key Examples Tested:**
- Chess piece values with constant reuse (`peon`, `caballo`, `alfil`, `torre`, `dama`)
- Days of the week with iota (`Lunes`...`Domingo`)
- Bitwise operations with iota

### `variables_test.go` - Section 3: Variables (6 tests)
Tests for variable declarations and properties:
- Explicit variable declarations
- Factorized declarations
- Variable initialization
- Short declaration syntax (`:=`)
- Variable mutability
- Default values for all basic types

### `boolean_test.go` - Section 4: Tipo Booleano (9 tests)
Tests for boolean type and operators:
- Boolean constants (`true`, `false`)
- Boolean not related to integers or nil
- Unary NOT operator (`!`)
- Logical AND (`&&`) and OR (`||`)
- Short-circuit evaluation

**Critical Tests:**
- Short-circuit evaluation prevents panics (division by zero, empty string indexing)

### `numeric_test.go` - Section 5: Tipos Numéricos (18 tests)
Tests for numeric types and operations:
- Integer types (int8, int16, int32, int64, int)
- Unsigned integer types (uint8, uint16, uint32, uint64, uint)
- Type aliases (byte, rune)
- Math package constants
- Arithmetic operators
- Bitwise operators
- Shift operators
- Floating-point types (float32, float64)
- Complex number types (complex64, complex128)
- **NEW**: `min` and `max` built-in functions (Go 1.21+)

**Key Examples Tested:**
- Modulo with negative numbers (`34%13=8`, `-34%13=-8`)
- IMC (BMI) calculation
- Scientific notation
- NaN handling
- Variadic min/max functions

### `pointers_test.go` - Section 6: Punteros (9 tests)
Tests for pointer operations:
- Pointer basics
- Referencing operator (`&`)
- Dereferencing operator (`*`)
- Pointer chaining
- Nil pointer default value
- Pointer comparison
- Pointer to pointer

### `strings_test.go` - Section 7: Cadenas de Caracteres (16 tests)
Tests for string handling:
- String immutability
- String indexing and slicing
- String concatenation
- UTF-8 support
- `len()` vs `utf8.RuneCountInString()`
- Rune iteration with `range`
- Raw strings
- Escape sequences (hex, unicode)
- String comparison
- Byte slice conversion

**Key Examples Tested:**
- Greek letters (θ, π) multi-byte handling
- String formula: "sin(θ)=cos(π/2-θ)"
- Raw string regex patterns

## Running the Tests

### Run all tests:
```bash
cd code/capitulo4/tests
go test -v
```

### Run specific test file:
```bash
go test -v -run TestConstants    # Section 2: Constants
go test -v -run TestVariables     # Section 3: Variables
go test -v -run TestBoolean       # Section 4: Booleans
go test -v -run TestNumeric       # Section 5: Numeric types
go test -v -run TestPointer       # Section 6: Pointers
go test -v -run TestString        # Section 7: Strings
```

### Run with coverage:
```bash
go test -v -cover
```

### Run specific test function:
```bash
go test -v -run TestIotaBasic
go test -v -run TestShortCircuitAnd
go test -v -run TestMinMaxBuiltins
```

## Test Statistics

- **Total Test Functions**: 66
- **Test Files**: 6
- **Lines of Test Code**: ~1,000 lines
- **Coverage**: All Chapter 4 examples validated

## Test Results

All tests validate that:
1. Code examples from the chapter compile correctly
2. Examples produce expected results
3. Edge cases are handled properly
4. Go 1.21+ features work as documented

Example output:
```
PASS
ok  github.com/go-uc3m/intro-go/capitulo4/tests0.003s
```

## Requirements

- Go 1.21 or later (required for `min`/`max` built-in functions)
- Standard library packages: `math`, `testing`, `unicode/utf8`, `strings`

## Test Coverage Summary

| Section | Topic | Tests | Key Features |
|---------|-------|-------|--------------|
| 2 | Constantes | 8 | iota, factorized, reuse |
| 3 | Variables | 6 | declarations, defaults, mutability |
| 4 | Tipo Booleano | 9 | operators, short-circuit |
| 5 | Tipos Numéricos | 18 | int, float, complex, min/max |
| 6 | Punteros | 9 | &, *, nil, chaining |
| 7 | Cadenas | 16 | UTF-8, runes, immutability |
| **Total** | | **66** | |

## Key Test Features

### Comprehensive Coverage
- Every code example from Chapter 4 is tested
- All 7 sections of the chapter covered
- Edge cases and error conditions tested

### Safety Tests
- Short-circuit evaluation prevents panics
- Nil pointer handling
- String immutability verification
- Division by zero prevention

### Modern Go Features
- `min()` built-in function (Go 1.21+)
- `max()` built-in function (Go 1.21+)
- Variadic function support

### UTF-8 Validation
- Multi-byte character handling (θ, π)
- Byte vs rune count differences
- Range iteration over runes
- Proper string indexing

## Example Test Cases

### Constants with iota
```go
const (
    Lunes = iota      // 0
    Martes            // 1
    Miercoles         // 2
    Jueves            // 3
    Viernes           // 4
    Sabado            // 5
    Domingo           // 6
)
```

### Short-circuit evaluation (prevents panic)
```go
y := 0
result := y > 0 && x/y > 10  // Right side not evaluated, no panic
```

### UTF-8 string handling
```go
formula := "sin(θ)=cos(π/2-θ)"
len(formula)                          // 20 bytes
utf8.RuneCountInString(formula)       // 17 runes
```

### New Go 1.21+ features
```go
min(5, 10)              // 5
max(5, 10)              // 10
min(3.14, 2.71, 1.41)   // 1.41 (variadic)
```

## Notes

- Tests use table-driven test patterns where appropriate
- Critical safety features (short-circuit evaluation) are explicitly tested
- UTF-8 tests use actual Greek letters (θ, π) to match course examples
- All numeric ranges and constants verified against Go 1.24 specifications
- No external dependencies - standard library only

## Purpose

This test suite serves multiple purposes:

1. **Validation**: Ensures all code examples in Chapter 4 are correct and compile
2. **Documentation**: Provides executable examples of concepts taught
3. **Regression Testing**: Prevents future changes from breaking examples
4. **Learning Tool**: Students can run tests to see concepts in action
5. **Quality Assurance**: Verifies compatibility with modern Go versions

## Contributing

When adding new examples to Chapter 4, please:
1. Add corresponding test cases to the appropriate test file
2. Ensure tests pass with `go test -v`
3. Update this README if adding new concepts
4. Follow existing test naming conventions

## Support

For questions or issues with the test suite, please refer to the main repository documentation or open an issue on GitHub.
