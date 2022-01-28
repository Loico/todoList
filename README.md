Todo List
=========

A simple CLI todo list app written in Go and unsing redis.

How to build
------------

1. `git clone https://github.com/Loico/todoList.git todoList`
2. `cd todoList`
3. `go build -o todo`
4. `docker pull redis`
5. `docker run -d --name todoList_redis -p 6379:6379 redis`
