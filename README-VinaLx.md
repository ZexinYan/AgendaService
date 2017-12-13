# AgendaService

- **Student No.**: 15331344
- **Name**: 薛明淇
- **Forked Repo**: [https://github.com/VinaLx/AgendaService](https://github.com/VinaLx/AgendaService)

## Responsibility

My tasks in this project includes:

- Development of backend (service)
- Docker Integration
- Project Management

## Commit History

**All my commits can be found [here](https://github.com/VinaLx/AgendaService/commits?author=VinaLx)**

### 1. 7e4f19cbefe13d9f1a275a1238bcccf3ebab9e5d

> initialize service

Initializing the service project structure, creating empty source files.

### 2. 3b353365dca0033579ce1b42146f2022e4effc3d

> initialize service 2

Initializing basic interfaces for command line argument parsing, database, router and server

### 3. 827cdbfd1f4a16e5c81bbebd3cd25e9744290b31

> db stuff

Removing the abstraction of DAO to plain functions. Basic implementation of sqlite database involving manipulation of `User`

### 4. 04efd6e70ee2868a643a1bc0eb0be64edd7ec53

> db & model

Adding database access to the login token. Implementing basic business logics in `model`. (Functionalities of Agenda)

### 5. ed8bf4d02baa3a585710a6cee021b8c8d411e551

> database debugged, a huge workaround

Debugging database, manage concurrent access to sqlite database, but involving a workaround due to the weakness of `go` and `sql` library structure.

### 6. 6068b7fb9d0eda02d2dc648242c4080beabbf510

> user and token tested

Completing the tests for `User` and `Token` for database

### 7. 58de51e95d02f2cd3c4ced1926b872cd3393452e

> refactor test util functions

Moving utilities of testing into a new package

### 8. ffa355d7458c8b1fcc6d77b08013a24dbc5bbdc4

> model tested

Finishing the tests for business logic

### 9. e9178d7b2dd2c00224623e9978e73567ae82f352

> update to adapt http api

Updating various interfaces and tests accordingly to adapting the updated http api

### 10. df7f089a123b07f1e629863bfdb649b6d9f5f6d7

> update model a bit

Slight adjustment on the `RemoveUser` interface, adding `GetUser` logic and its tests.

### 11. 848b57723e190db5d96b72edc38e97c55b751f64

> finish router

Finishing the router according to http api, managing json serialization of data.

### 12. ed9a88c7ad41e9d296abed159f31a939fa59323e

> finish server

Basic debugging of the whole server functionaility and make slight adjustment.


### All commits after `12`

- Minor debugging of code
- Adding tests to .travis.yml
- Add Dockerfile
- Update Readme.md

## Project Summary

In this project I learn or review the following:

- Basic abstraction techniques of code
- Writing non-trivial tests for logics
- Some Sqlites features
- Docker integration and automated build
- Web backend programming to the API
- Collaboration with team
- Being patient

To be honest, I don't think the whole process a experience of a lot of fun, working and engineering may not be a very amusing process after all :)

The logic part of the program being straightforward, 70% of the work is for engineering discipline, tests, standardized api and travis and docker etc. They are not hard to do at all, but still involves a lot of work. We talk about them point by point.

### Something about Implementations

I talked a little about basic abstraction and modular method in [the first homework](https://github.com/VinaLx/service-computing-homework/tree/master/selpg), and the homework following involve nothing more in principle. But go is a fantastic language that complicate nearly all things about abstraction. Even with a somewhat acceptable implementation in mind, I find it hard to express them in `go`. Although the coroutine and concurrent utlities are very flexible, but that doesn't really cover the weakness of abstraction capability of `go`.

And the thing is, I want to decouple the logic of "retrieving a record in database" to "manipulating the record", But since we know nothing about the row until the process of concrete implementation. So I fail to decouple the thing generically, my first attempt is to create a generator of `sql.Row` when "retrieving records" and let caller do its job extracting the data. But it fail to run because of the mutability of `sql.Row`, before we process the row, the row move to the next record. So I have to pass callbacks into the retrieving method to make the code works, but it's interface is just disgusting. There are several method I comming out to solve this problem perfectly but `go` and `database.sql` just support none of them, of course reflection is not an option.

Leaving aside all disgusting works, the whole structure of the code is clean, involving the database part of retrieval of data from primitive storage, the model part that implements the business logic based on the database, and finally the server and router calling the model and handling http data exchange. What's inside of the package includes maybe two things, business logic, and calling external libraries.

### Tests

With a language that fail to provide abstraction on language level, I find testing very important in the development. But even complete testing for logic is so important in stacking business logic, the time doesn't permit me to test the completely edge cases in logic and only relatively basic correct samples and error handling are included in the testing.

The development process in this homework includes tests, but is not test driven, since it involves extra work and thinking and clearly we don't have much time for that. The testing in this homework is all about rushing, But I can feel the power of strict testing in a serious engineering project.

### Docker

I learned docker before, about half a year ago, so the job may be easier for me. But it doesn't mean that the whole job being easy. First I find it hard to pull the golang image down to my laptop, but it mysteriously solves itself after I wait for about an hour or two.

The second problem is the real thing, how do we download the package of google inside of the image without the support of socks proxy of docker engine. It's a lot of work finding the http proxying method for my local proxy but fail, and the final solution is a little workaround, based on the fact that go cli tool support socks5 proxy, so I configure the `http_proxy` environment variable inside of the image connecting the proxy outside of it, it sounds really simple, but all my time spends on finding a cleaner solution once and for all. It's acceptable though, comparing to the `sql.Row` thing.

## Last

This homework did successfully takes a lot of my time without paying me a single bit XD. And, looking forward to the final homework.