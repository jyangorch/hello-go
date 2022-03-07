// Domain Layer code
//
// Business logic expressed in code.
//
// This layer should be rich in business logic (representations of business states, business rules, etc).
//
// It should have no awareness of the concrete technology we use to invoke the interface encapsulating
// of the business logic. For example, it exposes technology-netural APIs and do not assume the code
// is invoked by a Service Activity from a particular service or protocol (HTTP/gPRC), or invoked by
// Queue Consumer Job, or just a plain CLI.
//
// It should have no awareness of the concrete technology we use to support the implementation of
// the business logic. For example, it interacts with a repostiory interface to access data, but it
// doesn't include which storage we actually use and the specifics to operate that storage.
package domain
