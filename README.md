# Zula, The Reading Tribe

Zula is on a mission to amplify diversity and inclusion in children's libraries and bookshelves around the globe!

We are aiming to develop an open-source + premium subscription online library for inclusive childrens books.

Find out more on our [wiki](https://github.com/reading-tribe/zula/wiki)

## Project Setup

- Backend written in Golang
- Frontend written in TypeScript using React

### Folder Structure

```
src/
  api/
    cmd/
      api/ -- contains API Server start command
    pkg/
      api/ -- contains API Server code
        test/ -- contains Go tests
  go2ts/ -- contains script to generate TS interfaces
  webapp/
    public/  -- contains frontend public assets
    src/  -- contains React frontend code
```

## Running the project locally

The recommended way to run the program locally is to use docker.

    $ docker-compose up --build

You can then visit the homepage at `localhost:8080` or use the API at `localhost:1323`
