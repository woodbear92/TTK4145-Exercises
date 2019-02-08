# Exercise 2

### What is an atomic operation?
Atomic operations in concurrent programming are program operations that run completely independently of any other processes.

### What is a semaphore?
An integer flag is always equal to or larger than zero. Contains the operations initialize, wait() (decrements) and signal() (incremens). It is used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.(WIKI)

### What is a mutex?
A mutex is a special case of binary semaphore with the concept of ownership. When a process lock a shared resource it owns it, and no one can access it until the process unlocks it.

### What is the difference between a mutex and a binary semaphore?
Binary semaphores don't have the concepts of ownership

### What is a critical section?
In concurrent programming, concurrent accesses to shared resources can lead to unexpected or erroneous behavior, so parts of the program where the shared resource is accessed are protected. This protected section is the critical section. (WIKI)

### What is the difference between race conditions and data races?
Race conditions is the behavior of an software or other system where the output is dependent on the sequence or timing of other uncontrollable events. It becomes a bug when events do not happen in the order the programmer intended. (WIKI)

A data race occurs when two instructions from different threads access the same memory location, at least one of these accesses is a write and there is no synchronization that is mandating any particular order among these accesses. (STACKOVERFLOW)

### List some advantages of using message passing over lock-based synchronization primitives.
* The state of Mutable/Shared objects are harder to reason about in a context where multiple threads run concurrently. 
* Synchronizing on a Shared Objects would lead to algorithms that are inherently non-wait free or non-lock free.
* In a multiprocessor system, A shared object can be duplicated across processor caches. Even with the use of Compare and swap based algorithms that doesn't require synchronization, it is possible that a lot of processor cycles will be spent sending cache coherence messages to each of the processors.
* A system built of Message passing semantics is inherently more scalable. Since message passing implies that messages are sent asynchronously, the sender is not required to block until the receiver acts on the message. 

### List some advantages of using lock-based synchronization primitives over message passing.
* Some algorithms tend to be much simpler.
* A message passing system that requires resources to be locked will eventually degenerate into a shared object systems. This is sometimes apparent in Erlang when programmers start using ets tables etc. to store shared state.
* If algorithms are wait-free, you will see improved performance and reduced memory footprint as there is much less object allocation in the form of new messages.




