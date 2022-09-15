# Description

Implement simple sitemap (https://www.sitemaps.org) generator as command line tool.
Please implement this test task in the same way as you would do it for production
code, which means pay attention to edge cases and details.

## It should

1. Accept start url as argument.
2. Recursively navigate by site pages in parallel.
3. Should not use any external dependencies, only standard golang library.
4. Extract page urls only from <a> elements and take in account <base> element if declared.
5. Should be well tested (automated testing).
6. Save the urls to a Postgres database.
7. Use TDD.
8. Use Clean Architecture.
9. Use Docker.
10. Upload to GitHub.
11. Use Git.

### Suggested program options

-parallel= number of parallel workers to navigate through site.
-output-file= output file path.
-max-depth=  max depth of url navigation recursion.

## Webcraller Analysis

1. Recursivity
   1. Given depth of recursion, how many pages will be crawled.
2. Parallelism
   1. How many workers will be needed
   2. Check the max number of workers that can be used and use them.
3. Regex
   1. <a>
   2. <base>
4. Testing
5. UI (cli)
6. Use only standard golang library
7. Create and save XML File
   1. XML file directory
8. TDD
9. Clean Architecture



# Commands Examples:

```sh
-u http://localhost:8081/index.html -p 4 -of sitemap.xml -md 0
```

```sh
-u http://www.example.com -p 4 -of sitemap.xml -md 0
```

```sh
-u http://localhost:8081/index.html -p 4 -of ./file/sitemap.xml -md 2
```

```sh
-u http://localhost:8081/index.html -p 1 -of ./file/sitemap.xml -md 5
```

```sh
-u http://localhost:8081/index.html -p 2 -of ./file/sitemap.xml -md 4
```

```sh
-u http://localhost:8081/index.html -p 1 -of ./sitemap.xml -md 10
```

```sh
-u http://localhost:8081/index.html -p 0 -of ./file.xml -md 5
```

## Testing webserver 

For testing purposes I have created a web server in localhost:8081

it can be launched with the following command:

```sh
go run test-webserver/server.go
```