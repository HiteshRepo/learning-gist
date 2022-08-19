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

## Reliability

What makes a software or an application reliable?
1. The application performs the function as the user expects.
2. It can tolerate user making mistakes or software is used in unexpected ways.
3. It performs good enough (in terms of latency and throughput) for required use case, under expected load and data volume.
4. Prevents unauthorized access.
5. In layman terms, 'works correctly' under normal/expected situations and 'continues to work correctly' in abnormal cases.

But things can go wrong and lead to faults. 
Systems that can anticipate 'faults' and have the ability to cope are called 'fault-tolerant' systems.
Thing to keep in mind is we can only tolerate 'certain faults' and not all.

Fault vs Failure
1. Fault is defined as one/some component of the whole system deviating from spec.
2. Failure is when the system as a whole fails to provide required service, everything becomes unavailable.
3. So it is best to design fault tolerant systems that prevent systems to go into failures.
4. Which basically means approach components individually and make them fault tolerant so that minimizes the whole system encountering failure.

Tolerating faults are preferred in some cases than preventing them.
It is 'what to do', 'when things go wrong'. Providing an emphasis on actually things going wrong.
For example, if an attacker has compromised the system, the event has occurred, and it cannot be prevented. Having mechanisms to subsidise the after-effects would be a better approach.
Along with having mechanisms in place to 'cure faults'.

### Hardware faults

Examples of hardware faults:
1. Hard disks crash.
2. RAM becomes faulty.
3. Power grid has a blackout.
4. Unplugging of network cable.

The above events happen very frequently, but we do not usually encounter, thanks to cloud platforms.
For example, Hard disks have an MTTF of 10-50 years, so in a typical data center with 10000 disks, we should expect an average of one disk to die per day.

Prevent hardware faults by redundancy:
1. Disks being setup in RAID configuration.
2. Servers having dual power supplies and hot-swappabble CPUs.
3. Datacenters having batteries and diesel generators as power backups.

So when one component dies, the other takes its place up until the primary component is back in action.

This 'redundancy' method was enough until recent years. 
But with emergence of high data volumes and high compute demanding systems, 'hardware redundancy' is not enough.
Moreover in cloud platforms units/resources are quite ephemeral in nature.

Hardware redundancy in conjunction with software fault tolerant techniques make systems more 'Reliable'.

One of 'software fault tolerant techniques' is 'Rolling Upgrade'. This technique is employed to reduce system downtime to bare minimum while applying system patches, deploying newer versions, etc.

### Software Errors

Usually we assume hardware faults being random and independent. Which is correct.
Also these random hardware faults have lower correlation, which means a CPU malfunctioning because of high temperature in the rack might not affect the backup CPU at some other place.

But there are hardware faults that have higher correlation because of underlying software errors.


Examples of software errors leading to hardware errors:
1. A software bug that causes every instance of an app server to crash when a particular input is provided.
2. A process that uses some shared resource.
3. A service that the system depends on that slows down, becomes unresponsive or gives corrupt responses.
4. Cascading failures, when one fault in a component leads to failures in others.

These types of failures lie for a very long period of time undetected leading to unpleasant surprise.
So it is essential to design components that can make assumptions about their environment when they do not get expected response from the dependent systems.

Steps to do to prevent/avoid above:
1. Carefully thinking about assumptions and interactions in the systems.
2. Thorough testing.
3. Process isolation.
4. Allow process to crash and analyze behavior.
5. Having measuring, monitoring systems in place.

### Human errors

Something that is inevitable but minimizable.

Ways to minimize human errors:
1. Design systems in a way that minimizes the opportunities. Via well-designed abstractions, APis and admin interfaces. This leads to discourage 'wrong thing'.
2. Decouple places where making mistakes are more prevalent. Sandbox environments with real time data allows experimentation and fault detection.
3. Testing: unit, manual, automation (for edge cases), whole-system integration tests.
4. Allow quick and easy recovery from human errors to minimize impact in case of failure. Rolling back configuration, gradual deployments, provide tools to recompute data.
5. Setup detailed monitoring such as performance metrics and error rates (telemetry data).
6. Implement good management practices and training.

### Why is reliability necessary?
1. Even in case of mundane apps, we have a responsibility to our users.
2. Leads to unproductivity to users.
3. Damage to reputation.
4. Loss of revenue of users.

There are places where 'reliability' needs to sacrificed because of operational cost cutting. In those cases being conscious and tracking would help in quick response.

## Scalability

If system is reliable in present times, then it does not guarantee that the system will remain reliable in future time.
This is due to decline in performance because of increase in load, volume of data. Therefore 'Scalability' comes into the picture.

Describing load:
1. Load can be described using few numbers called 'Load parameters'.
2. 'Load parameters' are not generic, they vary on the basis of architecture.
3. A few examples of such parameters are:
   1. Requests per second to a web server.
   2. Ratio of reads to writes in a database.
   3. Number of simultaneous users in a chat room.
   4. Hit ratio on a cache.

### Twitter Anecdote

Two main operations of 'twitter' are:
1. Post Tweet: Average load - 4.6k requests/second, Peak Load - 12k requests/second.
2. Home Timeline: 300k requests/second.

Twitter's scaling challenge was because of the 'fan-out'.

Look twitter has users. Each user has several followers and each user is again followed by a several followees.

There are 2 approaches to achieve the functionality:
1. Posting a tweet simply inserts a new tweet to a global tweet collection and when a user loads his/her timeline a query as below is fired:
   SELECT tweets.*, users.* FROM tweets
      JOIN users ON tweets.sender_id == users.id
      JOIN follows ON follows.followee_id == users.id
      WHERE follows.follower_id == current_user
   
   The results of this tweet is then needed to be sorted by time.
   This is taking place in a relational DB.
   
2. Maintaining of a cache for each user's home timeline. So when an user posts a tweet, the tweet is then fanned out to each of its followee's home timeline cache.
   This results in an expensive write op but an easy read op.
   
Approach 2 makes sense because on an average tweet is delivered to about 75 followers, so 4.6k tweets/second become 345k writes per second to timeline caches.

But some users have over 30million followers. A single tweet from them result in 30million writes to home timelines.

So for twitter, distribution of followers per user can become a key 'Load parameter'.

Twitter user approach1 in its very first version. Since it could not keep up with the load of timeline queries, it switched to approach2.
But Twitter now uses a hybrid approach where approach2 is used for most users while approach1 is used for users with higher count of followers.