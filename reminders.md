## reminders


## go

"x" is a string, 'x' is a rune (char)



## data structures

Differences between all the tree DS:
    - Bin-tree =
    - BST =
    - Heap = complete (all levels filled except last, no dangly trees), satisfies Min or Max heap property

Priority queue = a cheating queue (each item has a "weight" aka priority). USE HEAPS FOR PRIORITY QUEUES.



## API design

API is a part of a server designed to send and receive requests. It's the interface that allows one piece of software
to communicate with others.

REST like a postcard (representational state transfer)
    -- REST = uses HTTP GET/PUT/POST/DELETE. GET reqs should be idempotent
    -- aka REST is STATELESS

SOAP like an envelope = simple object access protocol / XML. soap is more secure, and has retry logic, better for messaging

RPC (Remote procedure call)


## computer science

idempotent (doing it once = doing it forever).
    - abs(x) is idempotent because abs(x) = abs(abs(x)).

polymorphism

inheritance
