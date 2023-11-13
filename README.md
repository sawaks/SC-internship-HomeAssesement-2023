# SC-internship-HomeAssesement-2023

## Description
This is the completed version of the technical take home for SC interns of 2023.
We need to follow instructions wich has below and commit to merge the repository.

## Usage

Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
```

To run the code for the test 
```
  go test -v ./folders
```

## Folder structure

```
| go.mod
| README.md
| sample.json
| main.go
| folders
    | folders.go
    | folders_test.go
    | static.go
    | folders_pagination.go
    | folders_pagination_helper.go
    | types.go
```

## Instructions

- This technical assessment consist of 2 components.
- Component 1:
  - within `folders.go`. 
    - We would like you to read through the code and run the code.
    - Write some comments on what you think the code does.
    - suggest some improvement that can be made to the code.
    - Implement the suggested improvement.
    - Write up some unit tests in `folders_test.go` for your new `GetAllFolders` method

- Component 2:
  - Extend your improved code to now facilitate pagination. 
  - You can copy over the existing methods into `folders_pagination.go` to get started.
  - Write a short explanation on why you choosen the solution you implemented.

    I've research between the different types such as offset-based, cursor-based, keyset and seek method and considering the Requirements I decided to go with the cursor-based.
    I believe it is a good fit for this assignment because it is efficient for the server as there is no counting of rows and can handles the potential of new data well.


## Technologies
go

## GitHub Repository Link
https://github.com/sawaks/SC-internship-HomeAssesement-2023.git
