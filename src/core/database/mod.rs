//! Database module containing core database abstractions and implementations.

// Traits
pub mod database_manager;

// Implementations
pub mod inmemory_database_manager;

// Re-export for cleaner imports
pub use database_manager::DatabaseManager;
pub use inmemory_database_manager::InMemoryDatabaseManager;
