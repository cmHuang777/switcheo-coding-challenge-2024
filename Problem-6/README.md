### Design Overview

To design the transaction broadcaster service, we need to focus on building a robust and scalable system that can handle broadcasting transactions to an EVM-compatible blockchain network reliably. The system should handle the signing of transactions, broadcasting them to the blockchain network, handling failures, retries, and providing visibility into the status of transactions.

                      +-------------------+
                      |    API Layer      |
                      +---------+---------+
                                |
                                v
                      +-------------------+
                      | Transaction       |
                      |     Manager       |
                      +---------+---------+
                                |
                                v
                      +-------------------+
                      |Blockchain Node RPC|
                      |      Client       |
                      +---------+---------+
                                |
                                v
                      +-------------------+
                      |     Database      |
                      +---------+---------+
                                |
                                v
                      +-------------------+
                      |   Admin Interface |
                      +-------------------+


### Architecture

#### Components:

1. **API Layer**: Exposes the internal API for broadcasting transactions.
2. **Transaction Manager**: Responsible for signing transactions, managing broadcasting, handling failures, and retries.
3. **Blockchain Node RPC Client**: Makes RPC requests to the blockchain node for broadcasting transactions.
4. **Database**: Stores transaction data and status for tracking and retrying.
5. **Admin Interface**: Provides an interface for admins to retry failed broadcasts and view transaction status.

#### Workflow:

1. When a POST request is made to `/transaction/broadcast`, the API Layer receives the request.
2. The API Layer validates the request and forwards it to the Transaction Manager.
3. The Transaction Manager signs the transaction data and stores it in the database with a status of "pending".
4. It then initiates the broadcast process by sending the signed transaction to the Blockchain Node RPC Client.
5. The Blockchain Node RPC Client sends the transaction to the blockchain node for broadcasting.
6. Depending on the response from the blockchain node:
   - If successful, the Transaction Manager updates the transaction status to "success" in the database.
   - If failed, the Transaction Manager updates the status to "failed" and schedules a retry.
7. The Admin Interface allows admins to view the list of transactions, their statuses, and manually retry failed broadcasts.

### Detailed Components

#### API Layer:

- Exposes `/transaction/broadcast` endpoint.
- Validates incoming requests.
- Forwards valid requests to the Transaction Manager.
- Returns appropriate HTTP status codes (`200`, `4xx`, `5xx`).

#### Transaction Manager:

- Receives broadcast requests from the API Layer.
- Signs transaction data using a private key.
- Stores transactions in the database with a status of "pending".
- Initiates broadcast to the blockchain node using the Blockchain Node RPC Client.
- Handles responses from the blockchain node.
- Updates transaction statuses in the database accordingly.
- Implements retry logic for failed transactions.
- Provides an interface for querying transaction status.

#### Blockchain Node RPC Client:

- Makes RPC requests to the blockchain node for broadcasting transactions.
- Handles timeouts and failures gracefully.
- Provides asynchronous communication for non-blocking requests.
- Reports success or failure responses to the Transaction Manager.

#### Database:

- Stores transaction data including message type, data, status, and timestamps.
- Provides CRUD operations for transaction management.
- Supports indexing for efficient querying.
- Ensures data consistency and durability.

#### Admin Interface:

- Allows admins to view the list of transactions.
- Displays transaction details including status and timestamps.
- Provides functionality to manually retry failed transactions.
- Requires authentication for admin access.

### Scalability and Robustness

- **Horizontal Scalability**: The system can be horizontally scaled by deploying multiple instances of each component behind a load balancer to handle increased traffic.
- **Fault Tolerance**: The system is fault-tolerant, with retry mechanisms for failed transactions and handling unexpected restarts gracefully.
- **Monitoring and Alerting**: Implement monitoring and alerting systems to detect issues such as failures in broadcasting transactions or performance degradation.
- **Transaction Queue**: Implement a queue system to handle a large number of transaction requests efficiently, ensuring that transactions are processed in the order they are received and preventing overload on the blockchain node.
- **Rate Limiting**: Apply rate limiting to prevent abuse or overload of the system, ensuring fair access to resources.
- **Logging and Auditing**: Log all transactions and system activities for auditing and debugging purposes.
- **Security**: Implement security measures such as encryption for sensitive data, access controls, and regular security audits to protect against potential threats.
