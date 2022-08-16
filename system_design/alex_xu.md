## Database isolation levels
1. Purpose of database isolation levels: to allow DB transactions to execute as if there are no other concurrent running transactions.

2. Types of DB locks:
    1. Shared Lock (S Lock): if transactionA locks some data, then transactionB can only read the data and not modify it.
    2. Exclusive Lock (X Lock): if this is applied, no other transaction can neither read/modify the data.

3. Types of Isolation levels:
    1. Serializable: The highest isolation level. Concurrent transactions are guaranteed to be executed in sequence.
    2. Repeatable Read: Data read during the transaction remains the same as the transaction starts.
    3. Read Committed: Data modification can only be read after the transaction is committed.

    4. Multi-Version Consistency Control (MVCC) for Repeatable Read:
        - There are 2 hidden columns for each row - transaction_id and roll_pointer
        - when transactionA starts, a new read view with tx_id=201 is created
        - shortly after transactionB starts, a new read view with tx_id=202 is created
        - when transactionA modifies data, a new row of the log is created and the roll pointer points to the old row
        - before transactionA commits, transactionB reads the balance data, but transactionA is not commited so it reads the previous commited record via roll pointer
        - even when transactionA commits, transactionB reads from the read view created during start of transactionB
    
| dirty_read      | non_repeatable | phantom_read | isolation_level | read | write |
| ----------- | ----------- | ----------- | ----------- | ----------- | ----------- |
| Impossible      | Impossible       | Impossible | Serializable | S Lock | X Lock |
| Impossible      | Impossible       | Probably | Repeatable Read | MVCC (beginning) | X Lock |
| Impossible      | Probably       | Probably | Read Committed | MVCC (last commit) | X Lock |
| Probably      | Probably       | Probably | Read Uncommitted | No Lock | X Lock |

## IAAS/PAAS/SAAS
1. A traditional IT manages below computing services:
    - applications
    - data
    - runtime
    - middleware
    - os
    - virtualization
    - servers
    - storage
    - networking
    
2. IAAS manages: os, virtualization, servers, storage, networking only.
3. PAAS manages: runtime, middleware, os, virtualization, servers, storage, networking only.
4. SAAS manages: all traditional IT computing services.

