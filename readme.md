# Go Foundations

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoint
- Discussion & Code

## Repository
- https://github.com/tkmagesh/Synechron-Go-Dec-2023

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go extension for VSCode (https://marketplace.visualstudio.com/items?itemName=golang.Go)

### Verification
> go version

## Who Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (No private/public/protected etc)
    - No classes (only structs)
    - No Inheritance (only composition)
    - No Reference types (everything is a value)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try..catch..finally
    - No implicit type conversion
- Performance
    - Equivalent to C++
- Close to hardware
    - Application builds targeting each OS
    - No VM (JVM, CLR)
    - No JIT
    - Tooling support for cross compilation
- Managed Concurrency
    - Concurrency is NOT parallelism
    - Ability to have more than one execution path
    - Concurrent Operations are represented as Goroutines
    - Goroutines are cheap (4 KB / goroutine)
    - Builtin scheduler to schedule & manage the goroutines
    - Support for concurrency built in the language
        - go keyword, channel (data type), channel operator ( <- ), range & select-case constructs
    - API support
        - "sync" packages
        - "sync/atomic" packages
### Concurrency Model (Using OS Threads)
![image concurrency-os-thread](./images/concurrency-os-thread.png)
### Concurrency Model (Using Goroutines)
![image concurrency-goroutine](./images//concurrency-goroutines.png)