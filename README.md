# sc-interns-2022

The technical take home for SC interns of 2022.

## Compile instruction

Requires `Go` >= `1.19`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
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
```

## Instructions

- This technical assessment consist of 2 components.
- Component 1:
  - within `folders.go`. 
    - We would like you to read through the code and write some comments around what you think the code does.
    - After looking through the code, suggest some improvements we could make to the `GetAllFolders` method.
    - Write up some unit tests in `folders_test.go` for the `GetAllFolders` method

- Component 2:
  - Extend the existing code to facilitate pagination. You can copy over the existing methods into `folders_pagination.go` to get started.
  - Pagination helps break down a large dataset into smaller chunks.
  - The small data chunk will then be served to the client side usually accompanied a token that points to the next chunk.
  - Write a short explanation on why you choosen the solution you implemented.
  - The end result should look like this:
```
  original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
  
  The current result will look like this:
  { data: [1, 2, 3, ..., 10] }
  
  With pagination implementation, the result should look like this: 
  { data: [1, 2], token: "Mw==" }

  The token can then be used to fetch more result:
  
  { data : [3, 4], token: "NQ==" }

  And more results util there's no data left:
  
  { data: [9, 10], token: null }
```

## Submission

Create a repo in your chosen git repository (make sure it is public so we can access it) and reply with the link to your code. We recommend using Github.
