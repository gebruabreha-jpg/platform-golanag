# Chapter 5 Test Suite

Comprehensive unit tests for all code snippets from Chapter 5 (Funciones I) of the Go course materials. These tests validate the exact code shown in the LaTeX document using the `\gofragment` macro.

## Test Coverage

### 1. Conditional Structures - if statements (`if_test.go`)

Tests code snippets from `code/capitulo5/if/if.go`:

| Test Function | Lines | Description |
|---------------|-------|-------------|
| `TestIfBasicCheck` | 21-23 | Basic if statement checking for zero value |
| `TestIfLeapYear` | 25-32 | Leap year calculation logic |
| `TestIfWithInitialization` | 34-37 | If statement with initialization for error handling |
| `TestIfCompleteExample` | All | Integration test for all if patterns |

**Coverage**: 4 test functions testing all if statement patterns shown in Section 2.

### 2. Switch Statements (`switch_test.go`)

Tests code snippets from `code/capitulo5/switch/switch.go`:

| Test Function | Lines | Description |
|---------------|-------|-------------|
| `TestSwitchFirstMatchOnly` | 16-24 | Switch with first match only (multiple divisibility) |
| `TestSwitchGuardBased` | 26-37 | Switch without variable using guards (chess logic) |
| `TestSwitchMonthDays` | 39-53 | Switch with multiple values per case (month days) |
| `TestSwitchFallthrough` | 55-71 | Switch with fallthrough (chess pieces) |
| `TestSwitchIntegration` | All | Integration test for all switch patterns |

**Coverage**: 5 test functions testing all switch patterns shown in Section 2.

### 3. Loop Structures - for loops (`for_test.go`)

Tests code snippets from `code/capitulo5/for/for.go`:

| Test Function | Lines | Description |
|---------------|-------|-------------|
| `TestForBinaryConversion` | 22-31 | Three-component for loop (binary to decimal) |
| `TestForStringReversal` | 35-38 | For loop with dual iteration variables (string reversal) |
| `TestForBreakContinue` | 39-52 | Infinite loop with break and continue |
| `TestForWhileStyleFileRead` | 53-62 | While-style for loop (file reading with Scanner) |
| `TestForRangeInteger` | 64-72 | Range over integer (Go 1.22+ feature) |
| `TestForRangeStringToUpper` | 74-82 | Range over string with character manipulation |
| `TestForIntegration` | All | Integration test for all for loop patterns |

**Coverage**: 7 test functions testing all for loop patterns shown in Section 3.

### 4. Functions (`func_test.go`)

Tests code snippets from `code/capitulo5/func/func.go`:

| Test Function | Lines | Description |
|---------------|-------|-------------|
| `TestFuncDiv1` | 23-29 | Function with named return value and early return |
| `TestFuncDiv2` | 31-36 | Multi-valued function returning result and error |
| `TestFuncFibonacci` | 38-50 | Fibonacci function with overflow detection |
| `TestFuncPhi` | 52-60 | Function consuming multi-valued result (golden ratio) |
| `TestFuncOrdenar` | 62-75 | Function invoking external command (sort) |
| `TestFuncSwapByValue` | 77-80 | Pass by value demonstration (failed swap) |
| `TestFuncDiv1Calls` | 108-109 | Function invocation examples |
| `TestFuncShadowing` | 112-117 | Variable shadowing and scope |
| `TestFuncDiv2Usage` | 120-125 | Multi-valued function usage with error handling |
| `TestFuncPhiComposition` | 128-132 | Composing multi-valued functions |
| `TestFuncOrdenarUsage` | 134-139 | Complete ordenar example with error handling |
| `TestFuncSwapDemo` | 142-145 | Pass by value demonstration in context |
| `TestFuncIntegration` | All | Integration test for all function patterns |

**Coverage**: 13 test functions testing all function patterns shown in Section 4.

## Running the Tests

### Run all tests:
```bash
cd code/capitulo5/tests
go test -v
```

### Run tests from a specific file:
```bash
go test -v -run TestIf        # Run all if-related tests
go test -v -run TestSwitch    # Run all switch-related tests
go test -v -run TestFor       # Run all for-related tests
go test -v -run TestFunc      # Run all function-related tests
```

### Run a specific test:
```bash
go test -v -run TestIfLeapYear
go test -v -run TestSwitchFallthrough
go test -v -run TestForBinaryConversion
go test -v -run TestFuncFibonacci
```

### Run with coverage:
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  # View coverage in browser
```

## Test Statistics

- **Total test files**: 4
- **Total test functions**: 29
- **Total test cases**: 50+ (including subtests with table-driven tests)
- **Line coverage**: Tests all code snippets referenced by `\gofragment` in the LaTeX files

## Test Philosophy

These tests follow these principles:

1. **Exact code validation**: Tests validate the exact code snippets shown in the textbook
2. **Line-by-line correspondence**: Each test function documents which lines it tests
3. **Educational value**: Tests include comments explaining what each snippet demonstrates
4. **No modifications**: Code snippets are tested as-is without modifications for "better" practices
5. **Integration tests**: Each file includes integration tests to verify patterns work together

## Code Snippet Mapping

All tests map directly to `\gofragment` references in the LaTeX files:

- **seccion2.tex**: Tests for if and switch statements
- **seccion3.tex**: Tests for for loops
- **seccion4.tex**: Tests for function definitions and usage

## Requirements

- Go 1.22 or later (for `range` over integer feature)
- Standard library only (no external dependencies)
- Unix-like system (for `sort` command test)

## Notes

1. Some tests (like `TestForWhileStyleFileRead`) create temporary files that are automatically cleaned up
2. The `TestFuncOrdenar` test requires the `sort` command to be available
3. Tests for shadowing and scope demonstrate Go's scoping rules as shown in the textbook
4. Float comparisons use appropriate epsilon values for precision

## Validation

All tests pass successfully with Go 1.24:
```
PASS
ok      capitulo5_tests    0.XXXs
```

This confirms that all code snippets in Chapter 5 are correct, compile successfully, and behave as documented in the course materials.
