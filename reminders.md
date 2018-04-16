## reminders


## go

"x" is a string, 'x' is a rune (char)



## data structures

Differences between all the tree DS:
    - Bin-tree = ANY type of tree with degree=2. ORDER DOES NOT MATTER.
    - BST = a sorted tree such that left < root < right
    - Heap = a complete tree (all levels filled except last, no dangly trees), satisfies Min or Max heap property
    - Trie = stores partial keys (aka sequences). SAME THING as a prefix tree.
    - Radix tree = a super compact Trie ( >1 char values per node) he --> llo, y, do -> g, n't, wager, nut  

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

polymorphism = ability for a function to have different behavior based on the type of object passed in

inheritance


mutex = a lock (ensures only one thread can access some shared resource) vs. semaphore (also a lock but supports > 1 thread)


## discrete math

power set = all possible subsets of the originating set including itself and the empty set. # sets in the power set is 2^n where n is the length of the original set. example: {a,b,c} --> Power set is: {{}, {a}, {b}, {c}, {ab}, {bc}, {ac}, {abc}} or 2^3 = len 8 
