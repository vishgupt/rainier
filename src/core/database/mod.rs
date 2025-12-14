//! Database module containing core database abstractions and implementations.

// Traits
pub mod manager;

// Implementations
pub mod inmemory_manager;

// Re-export for cleaner imports
pub use manager::DatabaseManager;
pub use inmemory_manager::InMemoryDatabaseManager;
