// Application Layer code
//
// From client & interfacing perspective, Application Layer defines the API facade of the licensing business logic,
// organized by business *use cases*. The facade is exposed to and invoked by "Driving Layer" in any form
// (CLI, RPC Service Activity, Queue Consumer Handler, Temporal Activity, etc.) to execute the business logic.
// This promotes resuability of the same business logic regardless where they are executed.
//
// From responsibility & implementation perspective, Application Layer contain the implementation of the
// "business logic" (e.g., plumbing/coordinating domain entities and repositories to implement business use cases,
// handling cross-cutting concern like security, emiting application events)", but actually delegate most
// business logic to Domain Layer (entities, aggregates, repositories) as much as possible.
package application
