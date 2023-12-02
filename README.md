# watchgod
A simple watchdog context written in Go

## What does it do?
Unlike the `context.WithTimeout` function in Go, which automatically cancels after a fixed time, Watchgod's timer, created using `WithWatchdog`, can be reset. This feature grants users the ability to extend the context and prevent cancellation as needed, offering more fine-grained control over the context's lifecycle.

## Go Get watchgod
Run this command to try it out!
```bash
 go get github.com/stefanbildl/watchgod@v1.0.2
```

## WithWatchdog
The WithWatchdog function returns a context that is derived from the provided parent context, along with a ResetFn function to reset the watchdog timer.

## Behavior
The function creates a child context from the provided parent context.
A separate goroutine manages a timer that triggers a cancellation of the context after the specified timeout duration.
The ResetFn function allows resetting the timer to prevent the context from being cancelled.
If the timeout elapses without the watchdog being reset, the context is cancelled, and any operations depending on it should handle the cancellation appropriately.


## Example

```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/stefanbildl/watchgod"
)

func main() {
	parentContext := context.Background()
	timeoutDuration := 5 * time.Second

	ctx, reset := watchgod.WithWatchdog(parentContext, timeoutDuration)

	// Use the context and reset function as needed
	// ...

	// Reset the watchdog timer
	reset()

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled due to timeout.")
		// Handle cancellation logic
		// ...
	}
}
```


# Contributing

We welcome contributions from the community to improve this project! Please follow these guidelines to contribute effectively:

## How to Contribute

1. Fork the repository to your GitHub account.
2. Clone the forked repository to your local machine:
    ```bash
    git clone https://github.com/your-username/repository.git
    ```
3. Create a new branch for your changes:
    ```bash
    git checkout -b feature/YourFeature
    ```
4. Make your changes, following the project's coding style and guidelines.
5. Test your changes thoroughly.
6. Commit your changes:
    ```bash
    git commit -am 'feat: Description of your changes'
    ```
7. Push to your forked repository:
    ```bash
    git push origin feature/YourFeature
    ```
8. Open a Pull Request (PR) from your forked repository to the main repository's `main` branch.
9. Provide a detailed description of your changes in the PR, including the motivation and context for the changes.

## Code Style and Guidelines

- Follow the existing code style and conventions.
- Write descriptive commit messages in the present tense.
- Ensure any new code includes relevant comments and documentation.
- Test your changes thoroughly and ensure there are no linting errors or warnings.

## Issue Reporting

If you encounter bugs, have feature requests, or want to propose enhancements, please open an issue in the repository. Ensure to include relevant details and steps to reproduce the issue.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). Please adhere to this code by treating all contributors and users with respect.

## Questions and Support

For questions or support related to the project, please open an issue in the repository or reach out to the maintainers.

We appreciate and welcome your contributions to this project!
