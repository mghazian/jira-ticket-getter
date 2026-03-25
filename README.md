# JIRA Ticket Getter

This project sources a program to extract JIRA ticket prefix from a kebab-case string or sentence-cased string.

## Installation

1. Clone the repository
2. Run `go build src/main.go`
3. Move and rename the resulting `main` executable to your `/bin` directory
4. Do not forget to `chmod` the executable

## How to Use

Using `jtg` as the executable name,

```
# Extract JIRA ticket in kebab case string
jtg JIRA-11001-copycat-detector
# Output: JIRA-11001

# Extract JIRA ticket in sentence case string
jtg "JIRA-11001 copycat detector feature"
# Output: JIRA-11001
```

## Limitation

Only work if JIRA ticket is prefixed to the string. Will not detect if JIRA ticket is in the middle or even end of the string.