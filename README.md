# Using JetBrains GoLand for Test-Driven Development

## Table of Contents
1. [Introduction](#introduction)
2. [Setting Up GoLand](#setting-up-goland)
3. [Running Tests in GoLand](#running-tests-in-goland)
4. [TDD Workflow in GoLand](#tdd-workflow-in-goland)
5. [Productivity Features](#productivity-features)
6. [Debugging Tests](#debugging-tests)
7. [Code Coverage](#code-coverage)
8. [Live Templates for Testing](#live-templates-for-testing)
9. [Keyboard Shortcuts](#keyboard-shortcuts)
10. [Tips and Best Practices](#tips-and-best-practices)

---

## Introduction

JetBrains GoLand is a powerful IDE specifically designed for Go development. It has excellent built-in support for testing and TDD workflows, making it an ideal choice for practicing Test-Driven Development.

### Why Use GoLand for TDD?

- **Integrated test runner** with visual feedback
- **Code generation** for tests
- **Live code coverage** visualization
- **Powerful refactoring tools**
- **Smart code completion** in tests
- **Quick navigation** between tests and implementation
- **Built-in debugger** for tests

---

## Setting Up GoLand

### Installation

1. **Download GoLand:**
    - Visit [jetbrains.com/go](https://www.jetbrains.com/go/)
    - Download for your OS (Windows, macOS, Linux)
    - 30-day free trial available

2. **Install Go SDK:**
    - GoLand will detect Go installation automatically
    - If not: `File → Settings → Go → GOROOT`
    - Verify: Check Go version in status bar

3. **Open Your Project:**
    - `File → Open` → Select your project directory
    - GoLand will automatically detect `go.mod`
    - Wait for indexing to complete

### Essential Settings for TDD

**Enable Auto-Import:**
```
File → Settings → Editor → General → Auto Import
☑ Show import popup
☑ Add unambiguous imports on the fly
```

**Configure Test Runner:**
```
File → Settings → Go → Test Runner
☑ Enable test runner
☑ Auto-scroll to source
Test output: Show passed tests
```

**Enable File Watchers (Optional):**
```
File → Settings → Tools → File Watchers
Add → Go fmt (formats code on save)
```

---

## Running Tests in GoLand

### Method 1: Run Icon (Gutter)

**Look for the green play icon** next to:
- Individual test functions
- Test files
- Packages

**Click the icon to:**
- **Run** (green play) - Run tests
- **Debug** (green bug) - Debug tests
- **Run with Coverage** (play with shield)

```go
func TestAdd(t *testing.T) {  // ▶️ Green icon appears here
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### Method 2: Right-Click Menu

**Right-click anywhere in the file:**
- `Run 'TestAdd'` - Run single test
- `Run 'calculator_test.go'` - Run all tests in file
- `Run 'All Tests'` - Run all tests in project

### Method 3: Keyboard Shortcuts

**Run tests:**
- `Ctrl+Shift+F10` (Windows/Linux) or `Ctrl+Shift+R` (macOS)
- Runs the test at cursor position

**Rerun last test:**
- `Shift+F10` (Windows/Linux) or `Ctrl+R` (macOS)

**Run with coverage:**
- `Ctrl+Shift+F10` then select "Run with Coverage"

### Method 4: Run Configurations

**Create a permanent run configuration:**

1. `Run → Edit Configurations`
2. Click `+` → `Go Test`
3. Configure:
   ```
   Name: All Tests
   Test kind: Directory
   Directory: ./...
   Pattern: .*
   ```
4. Click `OK`

Now you can run from the toolbar dropdown!

---

## TDD Workflow in GoLand

### The Red-Green-Refactor Cycle in GoLand

#### 1. RED Phase: Write Failing Test

**Step 1:** Create test file
```
Right-click package → New → Go File
Name: calculator_test.go
```

**Step 2:** Write test
```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)  // ⚠️ Red underline - Add() doesn't exist
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

**Step 3:** Run test (it fails)
- Click green play icon → Red X appears
- Test output shows: `undefined: Add`

#### 2. GREEN Phase: Make It Pass

**Quick Fix to Generate Function:**

1. Place cursor on `Add(2, 3)`
2. Press `Alt+Enter` (Windows/Linux) or `Option+Enter` (macOS)
3. Select `Create function 'Add'`
4. GoLand generates:
```go
func Add(i int, i2 int) int {
    //TODO implement me
    panic("implement me")
}
```

**Implement the function:**
```go
func Add(a, b int) int {
    return a + b
}
```

**Run test again:**
- Click green play icon → Green checkmark!
- Test output shows: `PASS`

#### 3. REFACTOR Phase: Improve Code

**Use GoLand's refactoring tools:**

1. **Rename** - `Shift+F6`
    - Rename variables, functions, parameters
    - All references updated automatically

2. **Extract Method** - `Ctrl+Alt+M` (Windows/Linux) or `Cmd+Alt+M` (macOS)
    - Select code → Extract to new function

3. **Inline** - `Ctrl+Alt+N` (Windows/Linux) or `Cmd+Alt+N` (macOS)
    - Inline function or variable

**Run tests after each refactor to ensure they still pass!**

---

## Productivity Features

### 1. Generate Test Functions

**Automatically generate test stubs:**

1. Place cursor on function name
2. Press `Ctrl+Shift+T` (Windows/Linux) or `Cmd+Shift+T` (macOS)
3. Select `Create New Test`
4. Choose:
    - Test file name
    - Test function name
    - Testing framework

**Example:** For function `Divide`, GoLand generates:
```go
func TestDivide(t *testing.T) {
    type args struct {
        a int
        b int
    }
    tests := []struct {
        name string
        args args
        want float64
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.args.a, tt.args.b)
            if (err != nil) != tt.wantErr {
                t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Divide() got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 2. Navigate Between Test and Implementation

**Quick navigation:**

- `Ctrl+Shift+T` (Windows/Linux) or `Cmd+Shift+T` (macOS)
- Toggles between `calculator.go` ↔ `calculator_test.go`

**Navigate to related symbol:**
- `Ctrl+Alt+Home` - Go to related files
- Shows all files related to current file

### 3. Smart Code Completion

**IntelliSense in tests:**

```go
func TestAccount(t *testing.T) {
    account := NewAccount("John", 100.0)
    account.  // ← Type dot, get all methods
    // GoLand shows: Deposit, Withdraw, GetBalance, Transfer
}
```

**Completion for testing package:**
```go
t.  // ← Shows all testing.T methods
// Error, Errorf, Fatal, Fatalf, Helper, Run, etc.
```

### 4. Live Templates

**Built-in test templates:**

Type abbreviation + `Tab`:

**`test`** - Basic test function:
```go
func TestName(t *testing.T) {
    
}
```

**`testt`** - Table-driven test:
```go
func TestName(t *testing.T) {
    tests := []struct {
        name string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            
        })
    }
}
```

**`bench`** - Benchmark test:
```go
func BenchmarkName(b *testing.B) {
    for i := 0; i < b.N; i++ {
        
    }
}
```

### 5. Quick Fixes and Intentions

**Alt+Enter magic:**

```go
result := Add(2, 3)  // ← Alt+Enter on undefined function
// Options:
// - Create function 'Add'
// - Create method 'Add'
// - Import package

if err != nil {  // ← Alt+Enter here
    // Options:
    // - Add return statement
    // - Add log statement
}
```

---

## Debugging Tests

### Setting Breakpoints

1. **Click left gutter** next to line number
2. **Red dot appears** - breakpoint set
3. **Right-click breakpoint** for conditions

### Debug a Test

**Start debugging:**
1. Click debug icon (green bug) next to test
2. Or: `Ctrl+Shift+F9` (Windows/Linux) or `Ctrl+Shift+D` (macOS)

**Debug panel shows:**
- **Variables** - Current values
- **Watches** - Custom expressions
- **Call Stack** - Function call hierarchy
- **Console** - Test output

### Debug Controls

```
F8          - Step Over
F7          - Step Into
Shift+F8    - Step Out
F9          - Resume Program
Ctrl+F8     - Toggle Breakpoint
```

### Example Debug Session

```go
func TestWithdraw(t *testing.T) {
    account := NewAccount("John", 100.0)  // ← Set breakpoint here
    err := account.Withdraw(50.0)
    
    if err != nil {                        // ← Set breakpoint here
        t.Errorf("Unexpected error: %v", err)
    }
}
```

**During debug:**
- Inspect `account.Balance` value
- Check `err` value
- Step through `Withdraw` function

---

## Code Coverage

### Running Tests with Coverage

**Method 1: Run icon**
- Right-click green play icon
- Select `Run 'TestName' with Coverage`

**Method 2: Keyboard**
- `Ctrl+Shift+F10` then select coverage option

**Method 3: Menu**
- `Run → Run 'TestName' with Coverage`

### Coverage Visualization

**Green/Red highlighting:**
- **Green** - Code is covered by tests
- **Red** - Code is NOT covered by tests
- **Yellow** - Code is partially covered

**Example:**
```go
func Divide(a, b int) (float64, error) {
    if b == 0 {              // ✓ Green - tested
        return 0, fmt.Errorf("division by zero")
    }
    return float64(a) / float64(b), nil  // ✗ Red - not tested
}
```

### Coverage Tool Window

**View detailed coverage:**
- `Run → Show Code Coverage Data`
- Shows percentage per file/package
- Click to navigate to uncovered code

**Coverage settings:**
```
Run → Edit Configurations → Coverage tab
☑ Track per test coverage
☑ Enable coverage in test runner
```

### Export Coverage Report

```
Run → Generate Coverage Report
Choose format: HTML, XML, or Plain Text
```

---

## Live Templates for Testing

### Create Custom Templates

**Settings → Editor → Live Templates → Go**

**Example: Create "terr" template**

1. Click `+` → Live Template
2. Abbreviation: `terr`
3. Template text:
```go
if err != nil {
    t.Errorf("Unexpected error: %v", err)
}
```
4. Define applicable contexts: Go → Statement

**Now type `terr` + Tab in any test!**

### Useful Custom Templates

**Table-driven test with error:**
```go
// Abbreviation: testte
tests := []struct {
    name        string
    input       $INPUT_TYPE$
    expected    $EXPECTED_TYPE$
    expectError bool
}{
    {$FIRST_CASE$},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result, err := $FUNC$(tt.input)
        
        if tt.expectError && err == nil {
            t.Error("Expected error but got none")
        }
        
        if !tt.expectError && err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
        
        if result != tt.expected {
            t.Errorf("got %v; want %v", result, tt.expected)
        }
    })
}
```

**Mock struct:**
```go
// Abbreviation: mock
type Mock$NAME$ struct {
    $FIELDS$
}

func (m *Mock$NAME$) $METHOD$($PARAMS$) $RETURNS$ {
    $END$
}
```

---

## Keyboard Shortcuts

### Essential TDD Shortcuts

| Action | Windows/Linux | macOS |
|--------|--------------|-------|
| **Running Tests** |
| Run test at cursor | `Ctrl+Shift+F10` | `Ctrl+Shift+R` |
| Rerun last test | `Shift+F10` | `Ctrl+R` |
| Run with coverage | `Ctrl+Shift+F10` → Coverage | Same |
| Stop running test | `Ctrl+F2` | `Cmd+F2` |
| **Navigation** |
| Go to test/implementation | `Ctrl+Shift+T` | `Cmd+Shift+T` |
| Go to declaration | `Ctrl+B` | `Cmd+B` |
| Find usages | `Alt+F7` | `Option+F7` |
| Recent files | `Ctrl+E` | `Cmd+E` |
| **Editing** |
| Generate code | `Alt+Insert` | `Cmd+N` |
| Quick fix | `Alt+Enter` | `Option+Enter` |
| Complete code | `Ctrl+Space` | `Ctrl+Space` |
| **Refactoring** |
| Rename | `Shift+F6` | `Shift+F6` |
| Extract method | `Ctrl+Alt+M` | `Cmd+Alt+M` |
| Inline | `Ctrl+Alt+N` | `Cmd+Alt+N` |
| **Debugging** |
| Debug test | `Ctrl+Shift+F9` | `Ctrl+Shift+D` |
| Toggle breakpoint | `Ctrl+F8` | `Cmd+F8` |
| Step over | `F8` | `F8` |
| Step into | `F7` | `F7` |

### Custom Keymap

**Create your own shortcuts:**

1. `File → Settings → Keymap`
2. Search for action
3. Right-click → Add Keyboard Shortcut
4. Press your desired combination

---

## Tips and Best Practices

### 1. Use Test Runner Window

**Enable persistent test runner:**
```
View → Tool Windows → Run
Pin the window (pin icon)
```

**Benefits:**
- See all test results at a glance
- Quickly rerun failed tests
- Filter tests by status (passed/failed)

### 2. Organize Test Files

**Project structure view:**
- `Alt+1` - Toggle Project view
- Tests appear under their packages
- Group by test kind

### 3. Use TODO Comments

```go
func TestNewFeature(t *testing.T) {
    // TODO: Add more test cases
    // TODO: Test edge cases
}
```

**View all TODOs:**
- `View → Tool Windows → TODO`
- Click to navigate

### 4. Quick Documentation

**View docs without leaving editor:**
- `Ctrl+Q` (Windows/Linux) or `F1` (macOS)
- Shows function documentation in popup

### 5. Split Editor for TDD

**View test and implementation side-by-side:**

1. Right-click editor tab
2. Select `Split Right` or `Split Down`
3. Navigate to related file in other pane
4. Run tests while viewing implementation

### 6. Use Scratches for Experiments

**Create scratch files:**
- `Ctrl+Alt+Shift+Insert` → Go Scratch
- Test ideas without creating files
- Not saved to project

### 7. Git Integration

**Commit tests with implementation:**
- `Ctrl+K` - Commit
- Review changes in diff view
- Ensure tests pass before commit

### 8. Run Tests on Save

**File Watchers (Advanced):**
```
File → Settings → Tools → File Watchers
Add custom watcher:
  File type: Go files
  Program: go
  Arguments: test $FileDir$
```

---

## Common Workflows

### Workflow 1: Adding New Feature with TDD

1. **Create test file** (if needed)
    - Right-click package → New → Go File

2. **Write failing test**
    - Use `test` live template
    - Run test (should fail)

3. **Generate function stub**
    - `Alt+Enter` on undefined function
    - Select "Create function"

4. **Implement function**
    - Write minimal code
    - Run test (should pass)

5. **Refactor**
    - Use refactoring shortcuts
    - Run tests after each change

6. **Check coverage**
    - Run with coverage
    - Add tests for uncovered code

### Workflow 2: Fixing Bug with TDD

1. **Write test that reproduces bug**
    - Test should fail

2. **Debug test**
    - Set breakpoints
    - Inspect variables
    - Find root cause

3. **Fix implementation**
    - Modify code
    - Run test (should pass)

4. **Run all tests**
    - Ensure no regressions

### Workflow 3: Refactoring with Confidence

1. **Ensure all tests pass**
    - Run all tests first

2. **Make refactoring**
    - Use IDE refactoring tools
    - Rename, extract, inline, etc.

3. **Run tests continuously**
    - After each small change
    - If tests fail, undo and try differently

4. **Check coverage**
    - Ensure coverage hasn't decreased

---

## Troubleshooting

### Tests Not Running

**Check GOROOT:**
```
File → Settings → Go → GOROOT
Verify correct Go SDK is selected
```

**Re-index project:**
```
File → Invalidate Caches / Restart
```

### Tests Not Found

**Check test file naming:**
- Must end with `_test.go`
- Must be in same package

**Check function naming:**
- Must start with `Test`
- Must have signature `func TestXxx(t *testing.T)`

### Coverage Not Showing

**Enable coverage:**
```
Run → Edit Configurations → Coverage tab
☑ Enable coverage
```

**Clear coverage:**
```
Run → Clear All Code Coverage Data
```

### Slow Test Execution

**Run specific tests:**
- Don't run all tests every time
- Use focused test runs

**Disable file watchers:**
```
File → Settings → Tools → File Watchers
Uncheck watchers if too slow
```

---

## Advanced Features

### 1. Remote Development

**Run tests on remote machine:**
```
Tools → Deployment → Configuration
Add remote server
Map local to remote paths
```

### 2. Database Testing

**Built-in database tools:**
```
View → Tool Windows → Database
Connect to test database
Verify test data
```

### 3. HTTP Client Testing

**Built-in HTTP client:**
```
Tools → HTTP Client → Test RESTful Web Service
Test API endpoints
Save requests for reuse
```

### 4. Benchmarking

**Run benchmarks:**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Right-click → Run with Benchmarks

---

## Summary

GoLand makes TDD in Go incredibly productive with:

✅ **Integrated test runner** - Instant feedback
✅ **Smart code generation** - Less typing
✅ **Visual coverage** - See what's tested
✅ **Powerful refactoring** - Change with confidence
✅ **Debugging tools** - Fix issues quickly

**Master these three things:**
1. Run tests frequently (keyboard shortcuts!)
2. Use quick fixes for code generation
3. Refactor fearlessly with IDE tools

Happy TDD with GoLand! 🚀    