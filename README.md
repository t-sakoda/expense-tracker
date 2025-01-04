# expense-tracker

Build a simple expense tracker to manage your finances.

## Build

```sh
go build
```

## Usage

### Add an Expense

To add an expense, use the `add` command with the `--description` and `--amount` flags:

```sh
expense-tracker add --description "Lunch" --amount 20
```

### List All Expenses

To list all expenses, use the `list` command:

```sh
expense-tracker list
```

### Delete an Expense

To delete an expense, use the `delete` command with the `--id` flag:

```sh
expense-tracker delete --id 1
```

### Summary of All Expenses

To get a summary of all expenses, use the `summary` command:

```sh
expense-tracker summary
```

### Summary of Expenses for a Specific Month

To get a summary of expenses for a specific month, use the `summary` command with the`--month` flag:

```sh
expense-tracker summary --month 8
```

## Project Roadmap

For more details on the project roadmap, visit: <https://roadmap.sh/projects/expense-tracker>
