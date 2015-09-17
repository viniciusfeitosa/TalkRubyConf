#!/usr/bin/python python
# -*- coding: utf-8 -*-

import tornado.ioloop
import tornado.web


class MainHandler(tornado.web.RequestHandler):
    def get(self, number):
        response = "Python + Tornado fib(" + number + "): "\
            + str(fib(int(number)))
        self.write(response)


def fib(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)


appication = tornado.web.Application([
    (r"/([0-9]+)", MainHandler)
])

if __name__ == "__main__":
    appication.listen(4567)
    tornado.ioloop.IOLoop.current().start()
