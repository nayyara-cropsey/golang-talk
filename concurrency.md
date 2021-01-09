---
marp: true
theme: uncover
paginate: true
---


# Concurrency in Go Lang
 

---

## Concurrency 

* Making progress on more than one task simultaneously is known as concurrency.
* Concurrency is especially important when designing systems (beyond simple tools).
* Golang is a high-level programming language with high-level concurrency constructs.

---

## Concepts

* Event-driven systems
* Asynchronous systems
* Concurrency*
* Parallelism

---

## Goal

* Become better at writing ***idiomatic Go***
* Learn Go lang concurrency constructs
* Understand and use Go lang concurrency patterns effectively

---


## Idiomatic Go


1. Orthogonality

2. Simplicity

3. Readability


---

## Orthogonality

*pieces are independent from each other; changes to one part, a type, package, program, etc, have minimal to no effect on other parts*

---

[![width:550px](images/proverbs.png)](https://go-proverbs.github.io/) 

---

## Concurrency vs Parallelism

* Concurrency is a programming concept
    * code designed to do multiple tasks and coordinate them 
* Parallelism is an execution concept 
    * a program with concurrent design utilizing multiple cores

[![](images/code.png)](https://play.golang.org/p/3FS7RDIKaQ1)

---

## Event-Driven & Asynchronous

* Event-driven: a programming paradigm in which the flow of the program is determined by events. 
    * Such as user interactions, sensor outputs, or messages from other programs or threads.
* Asynchronous programming is synonymous - events can happen asynchronously. 

---

## Go Concurrency

* Goroutines
* Channels
* Mutexes* (serializing concurrent access)

```markdown
go func() {
  for {
    select {
      case <- events_1:
        // handle event 1
      case <- events_2:
        // handle event 2
      case <- events_3:
       // handle event 3 
    }
  }
}

```

---


# Goroutines

* Like threads conceptually
* Can resource leaks if not terminated properly

[![](images/code.png)](https://play.golang.org/p/th7SCxdbr4Q)

![](images/goroutines.png)

---

# Channels

* Blocking (always!)
* Synchronization construct
* `select` with multiple channels

---

# Synchronization Events

***Cancellation***

![](images/cancellation.png)

--- 

# Synchronization Events

***Cancellation Acknowledgement***

![](images/finished.png)

--- 

## Concurrency is not parallelism

---


## Channels orchestrate; mutexes serialize

---

## Make the zero value useful

---


# Resources

[Go Proverbs](https://go-proverbs.github.io/)
[Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
[Concurrency is Not Parallellism](https://www.youtube.com/watch?v=oV9rvDllKEg) 
[Idiomatic Go](https://about.sourcegraph.com/go/idiomatic-go/)


![width:100px](images/scholar.png) 


