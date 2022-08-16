# Reliable, Scalable and Maintainable applications

Standard building blocks of application development:
1. Data storage - databases
2. Remember the result of an expensive operation to speed up reads - caches
3. Search data by keywords/filter - search indexes
4. Send message to another process to be handled asynchronously - stream processing
5. Periodically crunch a large amount of accumulated data - batch processing

Thinking about data systems
1. Even though data systems like DBs, Caches, Message buses, etc have superficial similarities like storing data but they have different access patterns and performance - so why put them under same umbrella called 'Data Systems'?
2. The boundaries between data systems are getting blurred. For example:
    1. Datastores are also used as message queues - redis.
    2. Message queues are giving DB like durability guarantees - kafka
3. Because many applications have requirements that a single tool cannot fulfill, the apps are broken down into smaller tasks and each task is fulfilled by a tool and the application developers stitch these tools together and make things work.

As a result, you are not just an app developer but a data system designer.
And as a data system designer you are answerable for:
1. How do you ensure data remains correct and complete even when things go wrong internally? - ensure consistency
2. How do you provide consistently good performance to clients, even when parts of your system has degraded? - ensure high availability
3. How do you scale to handle loads? - scalability

Some factors that influence application design:
1. Skills and experience of people involved.
2. Organization tolerance to risks.
3. Regulatory constraints

Three concerns of software systems:
1. Reliability
2. Scalability
3. Maintainability